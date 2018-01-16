package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal()
	}
	defer f.Close() //函数执行结束之后执行这个操作，一般用于http、socket、句柄、unlock，所有和资源释放的情况都需要
	body, _ := ioutil.ReadAll(f)
	fmt.Println(string(body))
}
