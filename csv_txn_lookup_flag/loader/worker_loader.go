package loader

import (
	"encoding/csv"
	"log"
	"os"
	"sync"

	"csv_txn_lookup_flag/model"
)

type Record struct {
	Txn      string
	Ref      string
	Info     string
	FileName string
}

func LoadWithWorkerPool(files []string, workerCount int) map[string]model.Payment {

	result := make(map[string]model.Payment)
	var mu sync.Mutex

	recordCh := make(chan Record, 1000)
	var wg sync.WaitGroup

	for range workerCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for r := range recordCh {
				mu.Lock()
				result[r.Txn] = model.Payment{
					RefNo:       r.Ref,
					PaymentInfo: r.Info,
					FileName:    r.FileName,
				}
				mu.Unlock()
			}
		}()
	}

	var fileWeightGroup sync.WaitGroup

	for _, f := range files {
		fileWeightGroup.Add(1)
		go func(filePath string) {

			defer fileWeightGroup.Done()

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
					Txn:      rec[0],
					Ref:      rec[1],
					Info:     rec[2],
					FileName: filePath,
				}
			}
		}(f)
	}

	fileWeightGroup.Wait()
	close(recordCh)
	wg.Wait()

	return result
}
