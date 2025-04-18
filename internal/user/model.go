package user

type CreateUserInput struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
