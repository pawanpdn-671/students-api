package storage

import "github.com/pawanpdn-671/students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age uint) (int64, error)
	GetStudentById(id uint64) (types.Student, error)
}