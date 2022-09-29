package main 

import(
	"fmt"
	"math/rand"
)

func main(){
	var nivelEscolhido int

	fmt.Println("1. Nível Fácil, números entre 0 e 100")
	fmt.Println("2. Nivel médio, números entre 0 e 1000")
	fmt.Println("3. Nível difícil, números entre 0 e 10000")
	fmt.Print("Digite o nivel: ")
	fmt.Scan(&nivelEscolhido)

	
	var nivel int

	switch nivel{
		case 1:
			nivelEscolhido = 300
	
		case 2:
			nivelEscolhido = 1000
		
		case 3:
			nivelEscolhido=10000

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