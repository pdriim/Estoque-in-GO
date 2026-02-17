package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sistema de Controle de Estoque")

	// Criando o estoque (map vazio)
	estoque := make(map[int]Produto)
	logs := []Log{}
	fmt.Println("Estoque iniciado:")

	for {
		fmt.Println("======MENU======")
		fmt.Println("1 - Adicionar Produto")
		fmt.Println("2 - Listar Produtos")
		fmt.Println("3 - Atualizar Produto")
		fmt.Println("4 - Remover Produto")
		fmt.Println("5 - Listar Logs")
		fmt.Println("0 - Sair")
		fmt.Println("================")

		var opcao int
		fmt.Println("Escolha uma opcao:")

		_, err := fmt.Scanln(&opcao)
		if err != nil {
			fmt.Println("Entrada Invalida. Digite um numero.")
			continue
		}

		fmt.Println("Opcao escolhida:", opcao)

		switch opcao {
		case 1:
			fmt.Println("Adicionar Produto selecionado")
			var nome string
			var quantidade int
			var preco float64

			fmt.Println("Digite o nome do produto:")
			_, err := fmt.Scanln(&nome)
			if err != nil {
				fmt.Println("Erro na leitura do nome:", err)
				logs = append(logs, Log{
					Tipo:     "ERROR",
					Mensagem: fmt.Sprintf("Erro na leitura do nome do produto: %v", err),
				})
				continue
			}

			fmt.Println("digite a quantidade do produto:")
			_, err = fmt.Scanln(&quantidade)
			if err != nil {
				fmt.Println("Erro na leitura da quantidade:", err)
				logs = append(logs, Log{
					Tipo:     "ERROR",
					Mensagem: fmt.Sprintf("Erro na leitura da quantidade do produto: %v", err),
				})
				continue
			}

			fmt.Println("digite o preco do produto:")
			_, err = fmt.Scanln(&preco)
			if err != nil {
				fmt.Println("Erro na leitura do preco:", err)
				logs = append(logs, Log{
					Tipo:     "ERROR",
					Mensagem: fmt.Sprintf("Erro na leitura do preco do produto: %v", err),
				})
				continue
			}

			fmt.Printf("Produto: %s | Quantidade: %d | Preco: %.2f\n", nome, quantidade, preco)
			estoque, logs, err = adicionarProduto(estoque, logs, nome, quantidade, preco)

			if err != nil {
				fmt.Println("Erro ao adicionar produto:", err)
				logs = append(logs, Log{
					Tipo:     "ERROR",
					Mensagem: fmt.Sprintf("Erro ao adicionar produto: %v", err),
				})
				continue
			} else {
				fmt.Println("Produto adicionado com sucesso!")
				logs = append(logs, Log{
					Tipo:     "INFO",
					Mensagem: fmt.Sprintf("Produto adicionado: Nome=%s", nome),
				})
			}

		case 2:
			fmt.Println("Listar Produtos")
			listarProdutos(estoque)

		case 3:
			listarProdutosSimples(estoque)

			var id int
			fmt.Println("Digite o ID do produto a ser atualizado:")
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println("Erro na leitura do ID:", err)
				logs = append(logs, Log{
					Tipo:     "ERROR",
					Mensagem: fmt.Sprintf("Erro na leitura do ID do produto: %v", err),
				})
				continue
			}
			var novoNome string
			var novaQuantidade int
			var novoPreco float64
			fmt.Println("Digite o novo nome do produto:")
			_, err = fmt.Scanln(&novoNome)
			if err != nil {
				fmt.Println("Erro na leitura do novo nome:", err)
				logs = append(logs, Log{
					Tipo:     "ERROR",
					Mensagem: fmt.Sprintf("Erro na leitura do novo nome do produto: %v", err),
				})
				continue
			}

			fmt.Println("Digite a nova quantidade do produto:")
			_, err = fmt.Scanln(&novaQuantidade)
			if err != nil {
				fmt.Println("Erro na leitura da nova quantidade:", err)
				logs = append(logs, Log{
					Tipo:     "ERROR",
					Mensagem: fmt.Sprintf("Erro na leitura da nova quantidade do produto: %v", err),
				})
				continue
			}

			fmt.Println("Digite o novo preco do produto:")
			_, err = fmt.Scanln(&novoPreco)
			if err != nil {
				fmt.Println("Erro na leitura do novo preco:", err)
				logs = append(logs, Log{
					Tipo:     "ERROR",
					Mensagem: fmt.Sprintf("Erro na leitura do novo preco do produto: %v", err),
				})
				continue
			}

			err = atualizarProduto(estoque, id, novoNome, novaQuantidade, novoPreco)
			if err != nil {
				fmt.Println("Erro ao atualizar produto:", err)
				logs = append(logs, Log{
					Tipo:     "ERROR",
					Mensagem: fmt.Sprintf("Erro ao atualizar produto ID %d: %v", id, err),
				})
			} else {
				logs = append(logs, Log{
					Tipo:     "UPDATE",
					Mensagem: fmt.Sprintf("Produto atualizado: ID=%d | Nome=%s", id, novoNome),
				})
				fmt.Println("Produto atualizado com sucesso!")
			}

		case 4:
			listarProdutosSimples(estoque)

			var id int
			fmt.Println("Digite o ID do produto a ser removido:")
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println("Erro na leitura do ID:", err)
				continue
			}

			err = removerProduto(estoque, id)
			if err != nil {
				fmt.Println("Erro ao remover produto:", err)
				logs = append(logs, Log{
					Tipo:     "ERROR",
					Mensagem: fmt.Sprintf("Erro ao remover produto ID %d: %v", id, err),
				})
			} else {
				fmt.Println("Produto removido com sucesso!")
			}

		case 5:
			fmt.Println("Listar Logs")
			listarLogs(logs)

		case 0:
			fmt.Println("Saindo do sistema")
			return // Sai do main e encerra o programa
		default:
			fmt.Println("Opcao invalida")
			logs = append(logs, Log{
				Tipo:     "ERROR",
				Mensagem: fmt.Sprintf("Opcao invalida selecionada: %d", opcao),
			})
		}
	}
}
