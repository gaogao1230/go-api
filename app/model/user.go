package model

type User struct {
	Id   int    `json:"-"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
