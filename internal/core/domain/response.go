package domain

// type interface{} means any data type
type Response struct {
	ErrorCode int
	Message string
	Data    interface{}
}
