package storage

type Storage interface {
	// methods for student storage
	CreateStudent(
		name string,
		email string,
		age int,
	) (
		int64, error,
	)
}
