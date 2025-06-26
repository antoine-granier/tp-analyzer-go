package analyzer

import (
	"math/rand"
	"time"
)

type Result struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

func AnalyzeLog(id, path string, resultChan chan<- Result) {
	var result Result
	result.LogID = id
	result.FilePath = path

	time.Sleep(time.Duration(rand.Intn(150)+50) * time.Millisecond)

	if rand.Float32() < 0.1 {
		parseErr := &ParseError{ID: id}
		result.Status = "FAILED"
		result.Message = "Erreur de parsing."
		result.ErrorDetails = parseErr.Error()
		resultChan <- result
		return
	}

	result.Status = "OK"
	result.Message = "Analyse terminée avec succès."
	resultChan <- result
}
