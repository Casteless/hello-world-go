package hello

import (
	"bufio"
	"fmt"
	"os"
)

func Main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
