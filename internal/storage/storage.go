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
}
