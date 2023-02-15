package storage

import(
	"fmt"
)

type Cache struct{
	Data map[string] string 
}

func (c *Cache) Connect(){
	c.Data=make(map[string]string)

	fmt.Println("Connect")
}

func (c *Cache) Get(key string) (value []byte, flags int, err error){
	return []byte( c.Data[key]), 0, nil}

func (c *Cache) Set(key string, value []byte, flags int, exptime int64) (err error){
	c.Data[key]=string(value)
	return nil
}

func (c *Cache) Delete(key string) (err error) {
	delete(c.Data, key);

	return nil
}
