package utils

import (
	"strings"

	"github.com/google/uuid"
)

func PKID() string {
	ret := uuid.New().String()
	return strings.ReplaceAll(ret, "-", "")
}
