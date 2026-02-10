package main

import "fmt"

func adicionarProduto(
	estoque map[int]Produto,
	nome string,
	quantidade int,
	preco float64,
	logs []Log,
) ([]Log, error) {

	if nome == "" {
		logs = adicionarLog(logs, "ERROR", "Nome do produto nao pode ser vazio")
		return logs, fmt.Errorf("nome do produto nao pode ser vazio")
	}

	if quantidade < 0 {
		logs = adicionarLog(logs, "ERROR", "Quantidade nao pode ser negativa")
		return logs, fmt.Errorf("quantidade nao pode ser negativa")
	}

	if preco <= 0 {
		logs = adicionarLog(logs, "ERROR", "Preco deve ser maior que zero")
		return logs, fmt.Errorf("preco deve ser maior que zero")
	}

	produto := Produto{
		ID:         proxID,
		Nome:       nome,
		Quantidade: quantidade,
		Preco:      preco,
	}

	estoque[produto.ID] = produto
	proxID++

	logs = adicionarLog(
		logs,
		"INFO",
		fmt.Sprintf("Produto adicionado: ID=%d | Nome=%s", produto.ID, produto.Nome),
	)

	return logs, nil
}

func listarProdutosSimples(estoque map[int]Produto) {
	if len(estoque) == 0 {
		fmt.Println("Nenhum produto cadastrado.")
		return
	}

	fmt.Println("Produtos disponíveis:")
	for id, produto := range estoque {
		fmt.Printf("ID %d - %s\n", id, produto.Nome)
	}
}

// READ
func listarProdutos(estoque map[int]Produto) {
	if len(estoque) == 0 {
		fmt.Println("Estoque vazio")
		return
	}

	for _, produto := range estoque {
		fmt.Printf(
			"ID: %d | Nome: %s | Qtd: %d | Preço: %.2f\n",
			produto.ID,
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
