package main 

import(
	"fmt"
	"math/rand"
)

func main(){
	numeroRandomico := rand.Intn(99)

	var numeroDigitado int

	for numeroDigitado != numeroRandomico{

		fmt.Print("Digite um valor entre 0 e 100: ")
		fmt.Scan(&numeroDigitado)

		if numeroDigitado > numeroRandomico{
			fmt.Println("Você errou! O número digitado é maior.")
		}else if  numeroDigitado < numeroRandomico {
			fmt.Println("Você errou! O número digitado é menor.")
		}else{
			fmt.Println("Parabéns você acertou, fim de jogo.")
		}
	}

	
}