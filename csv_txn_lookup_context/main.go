package main

import (
	"flag"
	"fmt"
	"runtime"

	"csv_txn_lookup_context/loader"
)

func main() {
	txn := flag.String("txn", "", "Transaction number")
	dir := flag.String("dir", "data", "CSV directory")

	flag.Parse()

	if *txn == "" {
		fmt.Println("Usage: go run main.go -txn=TXN001")
		return
	}

	// scan files
	files, err := loader.GetCSVFiles(*dir)

	if err != nil {
		fmt.Println("Error scanning directory:", err)
		return
	}

	fmt.Println("Files found:", len(files))
	fmt.Println("Searching...")

	fmt.Println("Before loading NumGoroutine:", runtime.NumGoroutine())

	res, ok := loader.SearchTxnParallel(files, *txn)

	fmt.Println("After loading NumGoroutine:", runtime.NumGoroutine())

	if ok {
		fmt.Println("Found")
		fmt.Println("Txn:", res.Txn)
		fmt.Println("Ref:", res.RefNo)
		fmt.Println("Info:", res.PaymentInfo)
		fmt.Println("File:", res.FileName)
	} else {
		fmt.Println("Not found")
	}
}
