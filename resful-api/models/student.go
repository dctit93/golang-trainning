package models

type Student struct {
	Id      int    `json:id`
	Name    string `json:name`
	Address string `json:address`
}
