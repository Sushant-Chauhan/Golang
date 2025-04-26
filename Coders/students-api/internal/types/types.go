package types

type Student struct {
	Id    int
	Name  string `validate:"required,min=3,max=100"`
	Email string `validate:"required,email"`
	Age   int    `validate:"gte=0,lte=130"`
}
