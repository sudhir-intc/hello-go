package main
import (
	"fmt"
	"github.com/sudhir-intc/hello/string"
	"time"
)

func  main(){
 fmt.Println("Hello, new world!")	
 fmt.Println(string.Reverse("Hello, new world!"))
 time.Sleep(5000)
 fmt.Println("After sleep of 5 seconds")
}
