package myrpc

import "fmt"

type Hello struct {
}

func (Hello) Haha(a string, result *string) error {
	*result = fmt.Sprintf("hello %s", a)
	return nil
}
