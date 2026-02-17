package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func registrarLog(logs []Log, tipo string, mensagem string) []Log {
	log := Log{
		Tipo:     tipo,
		Mensagem: mensagem,
	}

	logs = append(logs, log)

	err := salvarLogs(logs)
	if err != nil {
		fmt.Println("Erro ao salvar logs:", err)
	}

	return logs
}

func salvarLogs(logs []Log) error {
	arquivo, err := json.MarshalIndent(logs, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("logs.json", arquivo, 0644)
}

func carregarLogs() ([]Log, error) {
	var logs []Log

	dados, err := os.ReadFile("logs.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Log{}, nil
		}
		return nil, err
	}

	err = json.Unmarshal(dados, &logs)
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func listarLogs(logs []Log) {
	if len(logs) == 0 {
		fmt.Println("Nenhum log registrado.")
		return
	}

	for i, log := range logs {
		fmt.Printf("%d - [%s] %s\n", i+1, log.Tipo, log.Mensagem)
	}
}
