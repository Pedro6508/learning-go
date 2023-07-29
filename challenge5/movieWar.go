package movieWar

import (
	"fmt"
	"github.com/bxcodec/faker/v4"
	"time"
)

func verOppenheimer(verCanal chan string) {
	for {
		select {
		case msg := <-verCanal:
			fmt.Printf("💣 %s assistiu Oppenheimer\n", msg)
			time.Sleep(time.Second * 3)
		default:
			fmt.Printf("Ninguem está vendo Oppenheimer\n")
		}
	}
}

func verBarbie(Canal chan string) {
	for {
		select {
		case nomePessoa := <-Canal:
			fmt.Printf("🌸 %s assistiu Barbie\n", nomePessoa)
			time.Sleep(time.Second * 4)
		default:
			fmt.Printf("Ninguem está vendo Barbie\n")
		}
	}
}

func EnviarTarefas() {
	nome := make(chan string)
	go verOppenheimer(nome)
	go verBarbie(nome)

	for {
		nomeAleatorio := faker.Name()
		nome <- nomeAleatorio
	}
}
