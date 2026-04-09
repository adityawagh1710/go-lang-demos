package loader

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"sync"

	"csv_txn_lookup_context/model"
)

func SearchTxnParallel(files []string, txn string) (*model.Payment, bool) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	var mu sync.Mutex

	var result *model.Payment
	found := false

	sem := make(chan struct{}, runtime.NumCPU()*2)

	for i, f := range files {
		wg.Add(1)

		go func(id int, filePath string) {
			defer wg.Done()

			// acquire slot
			sem <- struct{}{}
			defer func() { <-sem }()

			file, err := os.Open(filePath)

			if err != nil {
				return
			}

			defer file.Close()

			reader := csv.NewReader(file)

			for {
				select {
				case <-ctx.Done():
					return
				default:
				}

				record, err := reader.Read()

				if err != nil {
					return
				}
				if record[0] == txn {
					mu.Lock()
					if !found {
						found = true
						result = &model.Payment{
							Txn:         record[0],
							RefNo:       record[1],
							PaymentInfo: record[2],
							FileName:    filePath,
						}

						fmt.Println("Found by goroutine:", id, "File:", filePath)

						cancel()
					}
					mu.Unlock()
					return
				}
			}
		}(i, f)
	}

	wg.Wait()
	return result, found
}
