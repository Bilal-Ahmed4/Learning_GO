package storage

type storage interface {
	 createStudent(name string , email string , age int)(int64,error)
}