package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	wordCounter := make(
		map[string]uint8,
	)

	fmt.Println("Insira aqui sua frase: ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	wordArray := strings.Split(
		strings.Replace(
			strings.TrimSpace(line), `,`, "", -1,
		), ` `)

	if err == nil && len(wordArray) > 0 {
		for _, word := range wordArray {
			wordLowerCase := strings.ToLower(word)

			if len(wordLowerCase) > 1 {
				if wordCounter[wordLowerCase] == 0 {
					wordCounter[wordLowerCase] = 1
				} else {
					wordCounter[wordLowerCase] += 1
				}
			}
		}

		for word, counter := range wordCounter {
			if counter == 1 {
				fmt.Printf(
					"\tA palavra '%s' apareceu apenas uma vez\n", word)
			} else {
				fmt.Printf(
					"\tA palavra '%s' apareceu %d vezes\n", word, counter)
			}
		}
	} else {
		println("Erro na leitura!")
	}
}
