package cmd

import (
	"fmt"
	"main/internal/analyzer"
	"main/internal/config"
	"main/internal/reporter"
	"sync"

	"github.com/spf13/cobra"
)

var (
	configPath string
	outputPath string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse les logs du fichier de configuration",
	Run: func(cmd *cobra.Command, args []string) {
		configs, err := config.LoadConfig(configPath)
		if err != nil {
			fmt.Errorf("Erreur de lecture de configuration: %v", err)
		}

		var wg sync.WaitGroup
		resultChan := make(chan analyzer.Result, len(configs))

		for _, logConf := range configs {
			wg.Add(1)
			go func(conf config.LogConfig) {
				defer wg.Done()
				analyzer.AnalyzeLog(conf.ID, conf.Path, resultChan)
			}(logConf)
		}

		wg.Wait()
		close(resultChan)

		var results []analyzer.Result
		for res := range resultChan {
			results = append(results, res)
			fmt.Printf("ID: %s | Chemin: %s | Statut: %s | Message: %s\n", res.LogID, res.FilePath, res.Status, res.Message)
			if res.Status == "FAILED" {
				fmt.Printf("  Détails: %s\n", res.ErrorDetails)
			}
		}

		if outputPath != "" {
			if err := reporter.Export(results, outputPath); err != nil {
				fmt.Errorf("Erreur d’export JSON: %v", err)
			}
			fmt.Printf("Rapport exporté vers %s\n", outputPath)
		}
	},
}

func init() {
	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Chemin du fichier de configuration JSON (obligatoire)")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Chemin pour exporter le rapport JSON")
	err := analyzeCmd.MarkFlagRequired("config")
	if err != nil {
		return
	}
	rootCmd.AddCommand(analyzeCmd)
}
