package main

import (
	"fmt"
	"sort"
)

var ultimoID int

func adicionarProduto(
	estoque map[int]Produto,
	logs []Log,
	nome string,
	quantidade int,
	preco float64,
) (map[int]Produto, []Log, error) {

	if nome == "" {
		return estoque, logs, fmt.Errorf("nome do produto nao pode ser vazio")
	}

	if quantidade < 0 {
		return estoque, logs, fmt.Errorf("quantidade nao pode ser negativa")
	}

	if preco <= 0 {
		return estoque, logs, fmt.Errorf("preco deve ser maior que zero")
	}

	ultimoID++
	id := ultimoID

	estoque[id] = Produto{
		Nome:       nome,
		Quantidade: quantidade,
		Preco:      preco,
	}

	return estoque, logs, nil
}

// READ
func listarProdutos(estoque map[int]Produto) {
	if len(estoque) == 0 {
		fmt.Println("Estoque vazio")
		return
	}

	// Criar slice de IDs
	var ids []int
	for id := range estoque {
		ids = append(ids, id)
	}

	// Ordenar IDs
	sort.Ints(ids)

	// Listar ordenado
	for _, id := range ids {
		produto := estoque[id]
		fmt.Printf(
			"ID: %d | Nome: %s | Qtd: %d | Preço: %.2f\n",
			id,
			produto.Nome,
			produto.Quantidade,
			produto.Preco,
		)
	}
}

// DELETE
func removerProduto(estoque map[int]Produto, id int) error {
	if _, existe := estoque[id]; !existe {
		return fmt.Errorf("produto com ID %d nao encontrado", id)
	}

	delete(estoque, id)
	return nil
}

// UPDATE
func atualizarProduto(
	estoque map[int]Produto,
	id int,
	novoNome string,
	novaQuantidade int,
	novoPreco float64,
) error {

	produto, existe := estoque[id]
	if !existe {
		return fmt.Errorf("produto com ID %d nao encontrado", id)
	}

	if novoNome == "" {
		return fmt.Errorf("nome do produto nao pode ser vazio")
	}
	if novaQuantidade < 0 {
		return fmt.Errorf("quantidade nao pode ser negativa")
	}
	if novoPreco <= 0 {
		return fmt.Errorf("preco deve ser maior que zero")
	}

	estoque[id] = Produto{
		ID:         produto.ID,
		Nome:       novoNome,
		Quantidade: novaQuantidade,
		Preco:      novoPreco,
	}
	return nil
}

func listarProdutosSimples(estoque map[int]Produto) {
	if len(estoque) == 0 {
		fmt.Println("Nenhum produto cadastrado.")
		return
	}

	var ids []int
	for id := range estoque {
		ids = append(ids, id)
	}

	sort.Ints(ids)

	fmt.Println("Produtos disponíveis:")
	for _, id := range ids {
		fmt.Printf("ID %d - %s\n", id, estoque[id].Nome)
	}
}
