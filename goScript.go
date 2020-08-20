package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Println("This is a message")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("------------")
	
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)
		
		if strings.Compare("hi", text) == 0{
			fmt.Println("hello Yourself")
		}
	}
}