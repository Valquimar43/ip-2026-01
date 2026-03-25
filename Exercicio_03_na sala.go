package main
import "fmt"

func main() {
    var quantidade int
    var nota, soma, media float64
    
    fmt.Printf("Digite a quantidade de alunos da turma: ")
    fmt.Scan(&quantidade)
    
    //Valida se a qtd é maior que zero para evitar divisão por zero
    if quantidade <=0{
        fmt.Printf("Quantidade inválida.")
		return
	}

	// Laço para capturar as notas
	for i := 1; i <= quantidade; i++ {
		fmt.Printf("Digite a nota do aluno %d: ", i)
		fmt.Scan(&nota)
		soma += nota // Acumula o valor da nota na variável soma
	

	// Cálculo da média
	media = soma / float64(quantidade)

	fmt.Printf("\n--- Resultado ---\n")
	fmt.Printf("Soma total: %.2f\n", soma)
	fmt.Printf("Média da turma: %.2f\n", media)
}
}