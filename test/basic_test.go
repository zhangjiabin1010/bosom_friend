package test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBasic(t *testing.T) {
	a := "2"
	b := 1
	d, _ := strconv.Atoi(a)
	c := d - b
	fmt.Println(c)
}
