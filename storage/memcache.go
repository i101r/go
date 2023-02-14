package storage

import(
	"fmt"
	"net"
	"bufio"
	"strconv"
	"strings"
	"errors"
)

type Memcached struct {
	conn net.Conn
}

var (
	ConnectionError = errors.New("memcache: not connected")
	ReadError       = errors.New("memcache: read error")
	DeleteError     = errors.New("memcache: delete error")
	FlushAllError   = errors.New("memcache: flush_all error")
	NotFoundError   = errors.New("memcache: not found")
)


func (m *Memcached) Connect(){
	conn, err := net.Dial("tcp", "localhost:11211" )
	
	if err != nil {
		return
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	status, err := bufio.NewReader(conn).ReadString('\n')
	
	fmt.Fprintf(status)

	m.conn=conn
}


func (m *Memcache) Get(key string) (value []byte, flags int, err error) {
	command := "get " + key + "\r\n"

	_, err = m.conn.Write([]uint8(command))

	if err != nil {
		return
	}

	response := bufio.NewReader(m.conn)

	return m.readResponse(response,key)
}




func (m *Memcached) readResponse( response *bufio.Reader, key string) (value []byte, flags int, err error) {
	
	row, err1 := response.ReadString('\n')

	if err1 != nil {
		err = err1
		return
	}
	
	a := strings.Split(strings.TrimSpace(row), " ")
	
	if len(a) != 4 || a[0] != "VALUE" {
		
		if row == "END\r\n" {
			err = NotFoundError
		} else {
			err = ReadError
		}
		return
	}

	flags, _ = strconv.Atoi(a[2])
	
	l, _ := strconv.Atoi(a[3])
	
	value = make([]byte, l)

	n := 0

	for {
		i, err1 := response.Read(value[n:])

		if i == 0 && err == io.EOF {
			break
		}

		if err1 != nil {
			err = err1
			return
		}

		n += i

		if n >= l {
			break
		}
	}

	if n != l {
		err = ReadError
		return
	}
	
	row, err = response.ReadString('\n')
	
	if err != nil {
		return
	}

	if row != "\r\n" {
		err = ReadError
		return
	}

	return 

}

func (m *Memcache) Set(key string, value []byte, flags int, exptime int64) (err error) {
	return m.store("set", key, value, flags, exptime)
}

func (m *Memcache) Delete(key string) (err error) {
	if m == nil || m.conn == nil {
		return ConnectionError
	}

	command := "delete " + key + "\r\n"

	_, err1 := m.conn.Write([]uint8(command))
	
	if err1 != nil {
		err = err1
		return err
	}

	response := bufio.NewReader(m.conn)
	
	row, err1 := response.ReadString('\n')

	if err1 != nil {
		err = err1
		return err
	}

	if row != "DELETED\r\n" {
		return DeleteError
	}
	return nil
}


func (m *Memcache) store(command string, key string, value []byte, flags int, exptime int64) (err error) {
	if m == nil || m.conn == nil {
		return ConnectionError
	}

	l := len(value)
	s := command + " " + key + " " + strconv.Itoa(flags) + " " + strconv.FormatInt(exptime, 10) + " " + strconv.Itoa(l) + "\r\n"
	
	writer := bufio.NewWriter(memc.conn)
	
	_, err = writer.WriteString(s)
	if err != nil {
		return err
	}
	_, err = writer.Write(value)
	if err != nil {
		return err
	}
	_, err = writer.WriteString("\r\n")
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}

	response := bufio.NewReader(m.conn)
	row, err1 := response.ReadString('\n')
	
	if err1 != nil {
		err = err1
		return err
	}
	
	if row != "STORED\r\n" {
		WriteError := errors.New("memcache: " + strings.TrimSpace(line))
		return WriteError
	}
	
	return nil
}
