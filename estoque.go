package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func adicionarProduto(
	estoque map[int]Produto,
	logs []Log,
	nome string,
	quantidade int,
	preco float64,
) (map[int]Produto, []Log, error) {

	if nome == "" {
		return estoque, logs, errors.New("nome nao pode ser vazio")
	}

	if quantidade < 0 {
		return estoque, logs, errors.New("quantidade nao pode ser negativa")
	}

	if preco <= 0 {
		return estoque, logs, errors.New("preco deve ser maior que zero")
	}

	novoID := len(estoque) + 1

	produto := Produto{
		ID:         novoID,
		Nome:       nome,
		Quantidade: quantidade,
		Preco:      preco,
	}

	estoque[novoID] = produto

	return estoque, logs, nil
}

func atualizarProduto(estoque map[int]Produto, id int, nome string, quantidade int, preco float64) error {
	produto, existe := estoque[id]
	if !existe {
		return errors.New("produto nao encontrado")
	}

	produto.Nome = nome
	produto.Quantidade = quantidade
	produto.Preco = preco

	estoque[id] = produto
	return nil
}

func removerProduto(estoque map[int]Produto, id int) error {
	_, existe := estoque[id]
	if !existe {
		return errors.New("produto nao encontrado")
	}

	delete(estoque, id)
	return nil
}

func listarProdutos(estoque map[int]Produto) {
	if len(estoque) == 0 {
		fmt.Println("Estoque vazio.")
		return
	}

	for _, produto := range estoque {
		fmt.Printf("ID: %d | Nome: %s | Quantidade: %d | Preco: %.2f\n",
			produto.ID, produto.Nome, produto.Quantidade, produto.Preco)
	}
}

func listarProdutosSimples(estoque map[int]Produto) {
	for _, produto := range estoque {
		fmt.Printf("ID: %d | Nome: %s\n", produto.ID, produto.Nome)
	}
}

func salvarEstoque(estoque map[int]Produto) error {
	arquivo, err := json.MarshalIndent(estoque, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("estoque.json", arquivo, 0644)
}

func carregarEstoque() (map[int]Produto, error) {
	var estoque map[int]Produto

	dados, err := os.ReadFile("estoque.json")
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[int]Produto), nil
		}
		return nil, err
	}

	err = json.Unmarshal(dados, &estoque)
	if err != nil {
		return nil, err
	}

	return estoque, nil
}
