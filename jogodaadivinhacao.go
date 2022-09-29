package main 

import(
	"fmt"
	"math/rand"
)

func main(){
	var nivelEscolhido int

	fmt.Println("1. Nível Fácil, números entre 0 e 100")
	fmt.Println("2. Nivel médio, números entre 0 e 1000")
	fmt.Println("Ou outro número: Nível difícil, números entre 0 e 10000")
	fmt.Print("Digite o nivel: ")
	
	var nivel int
	fmt.Scan(&nivel)
	

	switch nivel{
		case 1:
			nivelEscolhido = 100
		
		case 2:
			nivelEscolhido = 1000
		default:
			nivelEscolhido = 10000
	}	
	
	numeroRandomico := rand.Intn(nivelEscolhido)
	
	var numeroDigitado int
	
	for numeroDigitado != numeroRandomico{
		
		fmt.Print("Digite um número: ")
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