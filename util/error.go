package util

import "fmt"

type CustomError struct {
	ErrorCode int16
	ErrorMsg  string
	ErrorR    string
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("Error Code : %v \n Error Msg : %v \n Error : %v \n", c.ErrorCode, c.ErrorMsg, c.ErrorR)
}
