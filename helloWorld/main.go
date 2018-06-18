package main

import (
	"fmt"

	"github.com/vpsidar/GoTraining/helloWorld/vis"
)

func main() {
	fmt.Println("Hello,  Go Lang !") // inline msg
	fmt.Println(vis.MYName)          // printing variable from imported package the variable name starting character should be capital
	vis.PrintVar()                   //You can access package local variable through its Function

}
