package model

import (
	"fmt"
	"testing"
)

func TestNewStudent2(t *testing.T) {
	stu := NewStudent2(WithName("kaixin2"), WithAge(5))
	fmt.Println(stu.Name, stu.Age)
	WithAge(8)(stu)
	fmt.Println(stu.Age)
}
