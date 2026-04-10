package service

import (
	"csv-txn-lookup-gin-api/internal/loader"
	"csv-txn-lookup-gin-api/internal/model"
	"fmt"
	"os"
	"path/filepath"
)

type TxnService struct{}

func NewTxnService() *TxnService {
	return &TxnService{}
}

func (s *TxnService) Lookup(txnID string) (*model.Payment, error) {

	wd, _ := os.Getwd()

	dirPath := filepath.Join(wd, "../../internal/data")

	files, err := loader.GetCSVFiles(dirPath)

	if err != nil {
		return nil, err
	}

	result, _ := loader.SearchTxnParallel(files, txnID)

	if result == nil {
		return nil, fmt.Errorf("txn not found")
	}

	return result, nil
}
