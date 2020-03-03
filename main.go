package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
 
func main() {
    fmt.Println("Вы находитесь в темной пещере.")
	
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Куда идти? :")
	command, _ := input.ReadString('\n')

    var exit = strings.Contains(command, "наружу")
	
    fmt.Println("Вы покидаете пещеру:", exit) // Выводит: Вы покидаете пещеру: true     
}
