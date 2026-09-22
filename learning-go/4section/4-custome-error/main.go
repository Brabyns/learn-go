package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrDivisionByZero = errors.New("Division by Zero")
var ErrNumTooLarge = errors.New("Number too Large")

type OpError struct{
	Op string
	Code int
	Message string
	Time time.Time
}

func NewOpError(op string, code int, message string, t time.Time) *OpError{
	return  &OpError{
		Op: op,
		Code: code,
		Message: message,
		Time: t,
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}

	if a > 1000 {
		return 0, ErrNumTooLarge
	}

	return a / b, nil
}

func main() {

	value, err := divide(10,5)
	if err != nil{
		if errors.Is(err, ErrDivisionByZero){
			fmt.Println("Divide by Zeror")
		}else if errors.Is(err, ErrNumTooLarge){
			fmt.Println("Number too Large")
		}
		
	}

	fmt.Println(value)

}