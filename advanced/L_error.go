package main

import (
	"fmt"
	"runtime"
	"time"
)

type MyError struct {
	time time.Time
	msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v, %s",
		e.time, e.msg)
}

func run() error {
	return &MyError{
		time: time.Now(),
		msg:  "服务器出差错了"}
}

func main() {
	num := runtime.NumCPU()
	fmt.Println(num)
	//Go 程序使用 error 值来表示错误状态
	if err := run(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("没有任何异常")
	}
}
