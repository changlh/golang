package main

import (
	"fmt"
	"io/ioutil"
	"log"
	//"net/http"
	"os"
	"strings"
)

func useage() {
	fmt.Println("useage:\n" +
		"   COMMAND FILENAME1 [ FILENAME2 ... ]")
	fmt.Println(" OPTIONS:\n" +
		"\tFILENAME\t 可以是文件，也可以是http url")
}

//打印文件内容
func Printfile(name string) {
	filename, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(filename))
}
//打印http_url内容
func Printhttp(url string) {
	//urlcont, err := http.Get(url)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//web, err := ioutil.ReadAll(urlcont.Body)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println(string(web))
	return
}

func main() {
	if len(os.Args) != 2 {
		useage()
		return
	}
	//后续会增加多个文件处理
	filename := os.Args[1]
	if strings.HasPrefix(filename, "http://") {
		Printhttp(filename)
	} else {
		Printfile(filename)
	}
}
