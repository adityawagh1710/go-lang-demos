package main

import (
	"flag"
	"fmt"
	"runtime"

	"csv_txn_lookup_flag/loader"
)

func main() {
	txn := flag.String("txn", "", "Transaction number")
	workers := flag.Int("workers", 0, "Worker count (default auto)")
	dir := flag.String("dir", "data", "Data directory")

	flag.Parse()

	if *txn == "" {
		fmt.Println("Usage: go run main.go -txn=TXN001 [-workers=10]")
		return
	}

	files, err := loader.GetCSVFiles(*dir)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	maxWorkers := runtime.NumCPU() * 2

	if *workers == 0 {
		if len(files) < maxWorkers {
			*workers = len(files)
		} else {
			*workers = maxWorkers
		}
	}

	fmt.Println("Files:", len(files))

	fmt.Println("Workers:", *workers)

	index := loader.LoadWithWorkerPool(files, *workers)

	if val, ok := index[*txn]; ok {
		fmt.Println("Found:", val.RefNo, val.PaymentInfo)
	} else {
		fmt.Println("Not found")
	}
}
