package myrpc

import (
	"fmt"
	"log"
	"time"
)

type Hello struct{}

func (*Hello) Haha(a string, result *string) error {
	time.Sleep(time.Second * 10)
	*result = fmt.Sprintf("hello %s", a)
	return nil
}

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	log.Println("got req", request)
	return nil
}
