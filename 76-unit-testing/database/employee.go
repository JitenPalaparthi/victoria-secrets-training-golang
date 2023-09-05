package database

//go:generate mockgen -source employee.go -destination=emoployee_mock.go -package=database
type EmployeeDB struct {
	DB any
}

type Employee struct {
	ID, Name string
}

type IEmployee interface {
	InsertEmployee(emp *Employee) (*Employee, error)
}

func (e *EmployeeDB) InsertEmployee(emp *Employee) (*Employee, error) {
	// you write some code hare that code
	// inserts data into mongodb and return employee object

	// Use gomock here becase you are using mongodb .. which is a separate component
	// and it is third party go client code well
	return new(Employee), nil
}
