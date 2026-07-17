package types

type Student struct {
	Id    int64  `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string  `validate:"required"`
	Age   int    `validate:"required"`
}
