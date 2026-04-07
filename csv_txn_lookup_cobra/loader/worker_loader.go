package loader

import (
	"csv_txn_lookup/model"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
)

type Record struct {
	Txn  string
	Ref  string
	Info string
}

func LoadWithWorkerPool(files []string, workerCount int) map[string]model.Payment {

	result := make(map[string]model.Payment)

	var mu sync.RWMutex

	recordCh := make(chan Record, 1000)

	var wg sync.WaitGroup

	fmt.Println("No of workers ", workerCount)

	// Workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for r := range recordCh {
				mu.Lock()
				result[r.Txn] = model.Payment{
					RefNo:       r.Ref,
					PaymentInfo: r.Info,
				}
				mu.Unlock()
			}
		}()
	}

	// File readers
	var fileWg sync.WaitGroup

	for _, f := range files {
		fileWg.Add(1)

		go func(filePath string) {
			defer fileWg.Done()

			file, err := os.Open(filePath)

			if err != nil {
				log.Println(err)
				return
			}

			defer file.Close()

			reader := csv.NewReader(file)

			for {
				rec, err := reader.Read()

				if err != nil {
					break
				}

				if len(rec) < 3 {
					continue
				}

				recordCh <- Record{
					Txn:  rec[0],
					Ref:  rec[1],
					Info: rec[2],
				}
			}
		}(f)
	}

	fileWg.Wait()
	close(recordCh)
	wg.Wait()

	return result
}
