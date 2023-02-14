package storage

import(
	"fmt"
)

type Storage interface {
	Connect()
	Get(int Uid) (string Response, error)
	Set(int Uid) (string Response, error)
	Delete(int Uid) (string Response, error)
}
