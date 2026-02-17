package main

import (
	"encoding/json"
	"os"
)

func salvarEstoque(estoque map[int]Produto) error {
	data, err := json.MarshalIndent(estoque, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("estoque.json", data, 0644)
}

func carregarEstoque() (map[int]Produto, error) {
	estoque := make(map[int]Produto)

	data, err := os.ReadFile("estoque.json")
	if err != nil {
		if os.IsNotExist(err) {
			return estoque, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, &estoque)
	if err != nil {
		return nil, err
	}

	// Ajustar ultimoID ao maior ID existente
	for id := range estoque {
		if id > ultimoID {
			ultimoID = id
		}
	}

	return estoque, nil
}
