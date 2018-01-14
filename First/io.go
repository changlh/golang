package main

import (
	"io/ioutil"
	"os"
	"fmt"
)

func main() {
	body, _ := ioutil.ReadAll(os.Stdin)
	//fmt.Println(len(body))
	fmt.Fprintf(os.Stdout,"%d", len(body))
}
