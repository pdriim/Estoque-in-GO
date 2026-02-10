package main

import (
	"fmt"
	"time"
)

type Log struct {
	Tipo     string
	Mensagem string
	Horario  time.Time
}

const (
	logInfo   = "INFO"
	logError  = "ERROR"
	logUpdate = "UPDATE"
)

func adicionarLog(logs []Log, tipo string, mensagem string) []Log {
	log := Log{
		Tipo:     tipo,
		Mensagem: mensagem,
		Horario:  time.Now(),
	}
	return append(logs, log)
}

func listarLogs(logs []Log) {
	if len(logs) == 0 {
		fmt.Println("Nenhum log registrado.")
		return
	}

	fmt.Println("===== LOGS DO SISTEMA =====")
	for _, log := range logs {
		fmt.Printf(
			"[%s] %s - %s\n",
			log.Tipo,
			log.Horario.Format("02/01/2006 15:04:05"),
			log.Mensagem,
		)
	}
}
