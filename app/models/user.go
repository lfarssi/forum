package models

type User struct  {
	ID   int    `json:"id"`
	UserName string `json:"username"`
	Email string `json:"email"`

}