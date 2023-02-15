package storage

import(
	"fmt"
	"sync"
	"time"
	"errors"
)

type Cache struct{
	sync.RWMutex
	Data map[string] Item 
}

type Item struct {
	Value 		string
	Created    	int64
	Expiration 	int64
}

func (c *Cache) Connect(){
	c.Data=make(map[string]Item)

	fmt.Println("Connect")
}

func (c *Cache) Get(key string) (value []byte, flags int, err error){
	c.RLock()
    defer c.RUnlock()
	
	if c.Data[key].Expiration > 0 {
        if time.Now().Unix() > (c.Data[key].Created+c.Data[key].Expiration ) {
			
			delete(c.Data,key)

            return nil, 0, nil
        }
    }

	return []byte( c.Data[key].Value), 0, nil}

func (c *Cache) Set(key string, value []byte, flags int, exptime int64) (err error){
	c.Lock()
	defer c.Unlock()

	c.Data[key]=Item{string(value), int64(time.Now().Unix()), exptime}
	
	return nil
}

func (c *Cache) Delete(key string) (err error) {
	c.Lock()
    defer c.Unlock()

	if _, found := c.Data[key]; !found {
        return errors.New("Key not found")
    }

	delete(c.Data, key);

	return nil
}
