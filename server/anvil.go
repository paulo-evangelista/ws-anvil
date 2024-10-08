package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"net/http"

)

func handleAnvilRestart(w http.ResponseWriter, r *http.Request) {
	err := startAnvil()
	if err != nil {
		http.Error(w, fmt.Sprintf("Falha ao iniciar/reiniciar Anvil: %v", err), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Anvil iniciado/reiniciado com sucesso"))
}

func startAnvil() error {
	anvilMux.Lock()
	defer anvilMux.Unlock()

	if anvilCmd != nil {
		anvilCmd.Process.Kill()
	}

	anvilCmd = exec.Command("anvil", "--no-cors", "--block-time", "3", "--code-size-limit", "8888888888888888")

	stdout, err := anvilCmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("erro ao obter stdout: %w", err)
	}

	stderr, err := anvilCmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("erro ao obter stdout: %w", err)
	}
	
	if err := anvilCmd.Start(); err != nil {
		return fmt.Errorf("erro ao iniciar anvil: %w", err)
	}

	go func() {
		scannerOut := bufio.NewScanner(stdout)
		scannerErr := bufio.NewScanner(stderr)
		for {
			if scannerOut.Scan() {
				incomingLogs <- scannerOut.Text()
				continue
			}
			if scannerErr.Scan() {
				fmt.Println("anvil err:", scannerErr.Text())
				incomingLogs <- "----- ANVIL ERROR -----"
				incomingLogs <- scannerErr.Text()
				continue
			}
		}
	}()

	return nil
}