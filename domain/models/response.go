package models

type Response struct {
	ErrorCode int
	Message   string
	Data      interface{}
}
