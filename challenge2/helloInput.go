package challenge2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func HelloInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Digite aqui seu nome: ")
	textLine, err := reader.ReadString('\n')

	if err == nil {
		fullName := strings.TrimSpace(textLine)
		fmt.Printf("Olá pessoa com nome %s!", fullName)
	} else {
		log.Fatal(err)
	}
}
