package testgomod

import (
	"fmt"
	"errors"
)

// Say hello
func Hello(who, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hello, %s!", who), nil
	cse "cn":
		return fmt.Sprintf("你好, %s!", who), nil
	default:
		return "", errors.New("unknown language")
	}
}
