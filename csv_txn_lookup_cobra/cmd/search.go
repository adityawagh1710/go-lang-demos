package cmd

import (
	"csv_txn_lookup/loader"
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var txn string
var workers int

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search transaction",
	Run: func(cmd *cobra.Command, args []string) {

		files, err := loader.GetAllCSVFiles("data")

		if err != nil {
			fmt.Println("Error scanning directory:", err)
			return
		}

		cmd.Println("Initial workers", workers)

		maxWorkers := runtime.NumCPU() * 2

		if workers == 0 {
			if len(files) < maxWorkers {
				workers = len(files)
			} else {
				workers = maxWorkers
			}
		}

		cmd.Println("Workers after len", workers)

		index := loader.LoadWithWorkerPool(files, workers)

		if val, ok := index[txn]; ok {
			fmt.Println("Found:", val.RefNo, val.PaymentInfo)
		} else {
			fmt.Println("Transaction not found")
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&txn, "txn", "t", "", "Transaction number")

	// By command line
	// searchCmd.Flags().IntVarP(&workers, "workers", "w", 0, "Worker count")

	if err := searchCmd.MarkFlagRequired("txn"); err != nil {
		panic(err)
	}
}
