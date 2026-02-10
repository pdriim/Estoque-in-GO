package main

// Produto representa um item do estoque
type Produto struct {
	ID         int
	Nome       string
	Quantidade int
	Preco      float64
}

// proxID controla a geração automática de IDs
var proxID int = 1
