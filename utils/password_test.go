package utils

import (
	"fmt"
	"testing"
)

func TestMakePassword(t *testing.T) {
	fmt.Println(GeneratePassword("rootme@123"))
}
