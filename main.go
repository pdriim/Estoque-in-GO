package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func lerString(mensagem string) string {
	fmt.Print(mensagem)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

func lerInt(mensagem string) (int, error) {
	input := lerString(mensagem)
	return strconv.Atoi(input)
}

func lerFloat(mensagem string) (float64, error) {
	input := lerString(mensagem)
	return strconv.ParseFloat(input, 64)
}

func main() {
	fmt.Println("Sistema de Controle de Estoque")

	estoque, err := carregarEstoque()
	if err != nil {
		fmt.Println("Erro ao carregar estoque:", err)
		return
	}

	logs, err := carregarLogs()
	if err != nil {
		fmt.Println("Erro ao carregar logs:", err)
		return
	}

	for {
		fmt.Println("\n==============================")
		fmt.Println(" SISTEMA DE CONTROLE ESTOQUE ")
		fmt.Println("==============================")
		fmt.Printf("Total de produtos: %d\n", len(estoque))
		fmt.Println("1 - Adicionar Produto")
		fmt.Println("2 - Listar Produtos")
		fmt.Println("3 - Atualizar Produto")
		fmt.Println("4 - Remover Produto")
		fmt.Println("5 - Listar Logs")
		fmt.Println("0 - Sair")
		fmt.Println("==============================")

		var opcao int
		fmt.Print("Escolha uma opcao: ")
		_, err := fmt.Scanln(&opcao)
		if err != nil {
			fmt.Println("Entrada invalida.")
			continue
		}

		switch opcao {

		case 1:
			var nome string
			var quantidade int
			var preco float64

			nome = lerString("Nome: ")

			if nome == "" {
				fmt.Println("Nome nao pode ser vazio.")
				continue
			}

			quantidade, err := lerInt("Quantidade: ")
			if err != nil || quantidade < 0 {
				fmt.Println("Quantidade invalida.")
				continue
			}

			preco, err = lerFloat("Preco: ")
			if err != nil || preco <= 0 {
				fmt.Println("Preco invalido.")
				continue
			}

			estoque, logs, err = adicionarProduto(estoque, logs, nome, quantidade, preco)
			if err != nil {
				fmt.Println("Erro:", err)
				logs = registrarLog(logs, "ERROR", err.Error())
				continue
			}

			logs = registrarLog(logs, "INFO", "Produto adicionado: "+nome)

			err = salvarEstoque(estoque)
			if err != nil {
				fmt.Println("Erro ao salvar estoque:", err)
			}

			fmt.Println("Produto adicionado com sucesso!")

		case 2:
			listarProdutos(estoque)

		case 3:
			listarProdutosSimples(estoque)

			var id int
			var nome string
			var quantidade int
			var preco float64

			fmt.Print("ID: ")
			fmt.Scanln(&id)

			fmt.Print("Novo Nome: ")
			fmt.Scanln(&nome)

			fmt.Print("Nova Quantidade: ")
			fmt.Scanln(&quantidade)

			fmt.Print("Novo Preco: ")
			fmt.Scanln(&preco)

			err = atualizarProduto(estoque, id, nome, quantidade, preco)
			if err != nil {
				fmt.Println("Erro:", err)
				logs = registrarLog(logs, "ERROR", err.Error())
				continue
			}

			logs = registrarLog(logs, "UPDATE", fmt.Sprintf("Produto %d atualizado", id))

			salvarEstoque(estoque)

			fmt.Println("Produto atualizado!")

		case 4:
			listarProdutosSimples(estoque)

			var id int
			fmt.Print("ID para remover: ")
			fmt.Scanln(&id)

			err = removerProduto(estoque, id)
			if err != nil {
				fmt.Println("Erro:", err)
				logs = registrarLog(logs, "ERROR", err.Error())
				continue
			}

			logs = registrarLog(logs, "INFO", fmt.Sprintf("Produto %d removido", id))

			salvarEstoque(estoque)

			fmt.Println("Produto removido!")

		case 5:
			listarLogs(logs)

		case 0:
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opcao invalida.")
			logs = registrarLog(logs, "ERROR", "Opcao invalida selecionada")
		}
	}
}

func gerarNovoID(estoque map[int]Produto) int {
	maiorID := 0

	for id := range estoque {
		if id > maiorID {
			maiorID = id
		}
	}

	return maiorID + 1
}
