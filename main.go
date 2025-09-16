package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Qual seu ano de nascimento?")

	scanner := bufio.NewScanner((os.Stdin))
	scanner.Scan()
	anoNascimentoTexto := scanner.Text()

	anoNascimentoInt, err := strconv.ParseInt(anoNascimentoTexto, 10, 64)

	if err != nil {
		fmt.Println("Digite um ano válido.")
		return
	}

	anoAtual := time.Now().Year()

	idadeUser := anoAtual - int(anoNascimentoInt)

	switch {
	case idadeUser < 18:
		fmt.Printf("Você é menor de idade, sua idade é: %d\n", idadeUser)
	case idadeUser >= 18:
		fmt.Printf("Você é maior de idade, sua idade é: %d\n", idadeUser)
	}
}
