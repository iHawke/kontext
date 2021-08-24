package pkg

import (
	"fmt"
	"os"
	"time"
)

var RED = "\033[31m"
var GREEN = "\033[32m"
var YELLOW = "\033[33m   "
var CLEAR = "\033[0m"

var IS_DEBUG = true

func Fatal(v ...interface{}) {
	fmt.Print(RED)
	fmt.Println(fmt.Sprint(v...), CLEAR)
	os.Exit(1)
}

func Error(v ...interface{}) {
	fmt.Print(RED)
	fmt.Println(fmt.Sprint(v...), CLEAR)
}

func Warn(v ...interface{}) {
	fmt.Print(YELLOW)
	fmt.Println(fmt.Sprint(v...), CLEAR)
}

func Info(v ...interface{}) {
	fmt.Print(GREEN)
	fmt.Println(fmt.Sprint(v...), CLEAR)
	time.Sleep(100 * time.Millisecond)
}

func Debug(v ...interface{}) {
	if IS_DEBUG == false {
		return
	}
	fmt.Print(GREEN)
	fmt.Println(fmt.Sprint(v...), CLEAR)
	time.Sleep(100 * time.Millisecond)
}