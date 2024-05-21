package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pessoa struct {
	Nome        string
	Sobrenome   string
	Idade       int
	Endereco    string
	Email       string
	Telefone    string
	TimeFutebol string
}

var pessoas []Pessoa

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("===== MENU DE OPÇÕES =====")
		fmt.Println("[1] - Cadastrar pessoas")
		fmt.Println("[2] - Pesquisar pessoa") //primeira pesquisa pede a variavel que sera usada para encontrar o cadastrado//
		fmt.Println("[3] - Listar todas pessoas")
		fmt.Println("[e] - Sair do programa")        //apaga a base de dados ao sair, sair do programa precisa que o e minusculo//
		fmt.Println("[R] - Resetar a base de dados") //necessita ser R maiusculo para funcionar, s/n funciona de qqr caixa de texto (maiscula ou minuscula) //
		fmt.Print("Escolha uma opção: ")

		opcao, _ := reader.ReadString('\n')
		opcao = strings.TrimSpace(opcao)

		switch opcao {
		case "1":
			cadastrarPessoa(reader)
		case "2":
			pesquisarPessoa(reader)
		case "3":
			listarPessoas()
		case "e":
			if confirmarAcao(reader, "Tem certeza que deseja sair do programa? (s/n): ") {
				fmt.Println("Obrigado por usar o programa!")
				return
			}
		case "R":
			if confirmarAcao(reader, "Tem certeza que deseja resetar a base de dados? (s/n): ") {
				resetarBaseDeDados()
			}
		default:
			fmt.Println("Opção inválida! Tente novamente.")
		}
	}
}

func cadastrarPessoa(reader *bufio.Reader) {
	var p Pessoa

	fmt.Print("Nome*: ")
	p.Nome = lerEntradaObrigatoria(reader)
	if p.Nome == "" {
		fmt.Println("Nome é obrigatório.")
		return
	}

	fmt.Print("Sobrenome: ")
	p.Sobrenome, _ = reader.ReadString('\n')
	p.Sobrenome = strings.TrimSpace(p.Sobrenome)

	fmt.Print("Idade*: ")
	fmt.Scanf("%d\n", &p.Idade)
	if p.Idade <= 0 {
		fmt.Println("Idade é obrigatória.")
		return
	}

	fmt.Print("Endereço*: ")
	p.Endereco = lerEntradaObrigatoria(reader)
	if p.Endereco == "" {
		fmt.Println("Endereço é obrigatório.")
		return
	}

	fmt.Print("Email: ")
	p.Email, _ = reader.ReadString('\n')
	p.Email = strings.TrimSpace(p.Email)

	fmt.Print("Telefone: ")
	p.Telefone, _ = reader.ReadString('\n')
	p.Telefone = strings.TrimSpace(p.Telefone)

	fmt.Print("Time de futebol*: ")
	p.TimeFutebol = lerEntradaObrigatoria(reader)
	if p.TimeFutebol == "" {
		fmt.Println("Time de futebol é obrigatório.")
		return
	}

	pessoas = append(pessoas, p)
	fmt.Println("Pessoa cadastrada com sucesso!")
}

func pesquisarPessoa(reader *bufio.Reader) {
	if len(pessoas) == 0 {
		fmt.Println("Não há pessoas cadastradas.")
		return
	}

	fmt.Print("Digite o campo a ser pesquisado (Nome, Endereço, ...): ")
	campo, _ := reader.ReadString('\n')
	campo = strings.TrimSpace(strings.ToLower(campo))

	fmt.Print("Digite o valor a ser pesquisado: ")
	valor, _ := reader.ReadString('\n')
	valor = strings.TrimSpace(valor)

	encontrado := false
	for _, p := range pessoas {
		if (campo == "nome" && strings.EqualFold(p.Nome, valor)) ||
			(campo == "sobrenome" && strings.EqualFold(p.Sobrenome, valor)) ||
			(campo == "idade" && fmt.Sprintf("%d", p.Idade) == valor) ||
			(campo == "endereco" && strings.EqualFold(p.Endereco, valor)) ||
			(campo == "email" && strings.EqualFold(p.Email, valor)) ||
			(campo == "telefone" && strings.EqualFold(p.Telefone, valor)) ||
			(campo == "time de futebol" && strings.EqualFold(p.TimeFutebol, valor)) {
			fmt.Println("Pessoa encontrada:")
			fmt.Printf("Nome: %s\n", p.Nome)
			fmt.Printf("Sobrenome: %s\n", p.Sobrenome)
			fmt.Printf("Idade: %d\n", p.Idade)
			fmt.Printf("Endereço: %s\n", p.Endereco)
			fmt.Printf("Email: %s\n", p.Email)
			fmt.Printf("Telefone: %s\n", p.Telefone)
			fmt.Printf("Time de futebol: %s\n", p.TimeFutebol)
			encontrado = true
		}
	}
	if !encontrado {
		fmt.Println("Pessoa não encontrada.")
	}
}

func listarPessoas() {
	if len(pessoas) == 0 {
		fmt.Println("Não há pessoas cadastradas.")
		return
	}

	for i, p := range pessoas {
		fmt.Printf("\nPessoa %d:\n", i+1)
		fmt.Printf("Nome: %s\n", p.Nome)
		fmt.Printf("Sobrenome: %s\n", p.Sobrenome)
		fmt.Printf("Idade: %d\n", p.Idade)
		fmt.Printf("Endereço: %s\n", p.Endereco)
		fmt.Printf("Email: %s\n", p.Email)
		fmt.Printf("Telefone: %s\n", p.Telefone)
		fmt.Printf("Time de futebol: %s\n", p.TimeFutebol)
	}
}

func resetarBaseDeDados() {
	pessoas = []Pessoa{}
	fmt.Println("Base de dados resetada com sucesso.")
}

func lerEntradaObrigatoria(reader *bufio.Reader) string {
	entrada, _ := reader.ReadString('\n')
	return strings.TrimSpace(entrada)
}

func confirmarAcao(reader *bufio.Reader, mensagem string) bool {
	fmt.Print(mensagem)
	resposta, _ := reader.ReadString('\n')
	resposta = strings.TrimSpace(strings.ToLower(resposta))
	return resposta == "s"
}
