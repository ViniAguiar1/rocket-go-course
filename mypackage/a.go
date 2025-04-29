package mypackage

import (
	"fmt"

	"github.com/ViniAguiar1/rocket-go-course/mypackage/internal/foo"
)

var Age int = 10

func GetMyName() string {
	fmt.Println(foo.MyName)
	return foo.MyName
}
