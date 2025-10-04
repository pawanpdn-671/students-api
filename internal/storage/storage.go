package storage

type Storage interface {
	CreateStudent(name string, email string, age uint) (int64, error)
}