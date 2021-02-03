package utils

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func Pk() string {
	return UUID()
}

func UUID() string {
	u1 := uuid.New()
	return strings.Replace(fmt.Sprintf("%s", u1), "-", "", -1)
}
