package storage

import(
	"fmt"
)

type Storage interface {
	Connect()
	Get(int Uid) (string Response, error)
	Get(int Uid) (string Response, error)
	Get(int Uid) (string Response, error)
}

func (s *Storage) connect(){

}

func (s *Storage) Get(int Uid) (string Response, error){

}

func (s *Storage) Set(int Uid) (string Response, error){

}

func (s *Storage) Delete(int Uid) (string Response, error){

}