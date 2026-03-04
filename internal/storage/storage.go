package storage

import "github.com/hasnathahmedtamim/students-api/internal/types"

type Storage interface {
	// methods for student storage
	CreateStudent(
		name string,
		email string,
		age int,
	) (
		int64, error,
	)

	GetStudentById(id int64) (types.Student, error)
	GetAllStudents() ([]types.Student, error)
	UpdateStudent(id int64, name string, email string, age int) error
	DeleteStudent(id int64) error
}
