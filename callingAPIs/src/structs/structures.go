package structs

type User struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
