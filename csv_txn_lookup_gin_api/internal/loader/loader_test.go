package loader

import (
	"os"
	"path/filepath"
	"testing"
)

// --- helpers ---

func writeTempCSV(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp CSV: %v", err)
	}
	return path
}

// --- GetCSVFiles ---

func TestGetCSVFiles_ReturnsOnlyCSV(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "a.csv"), []byte(""), 0644)
	os.WriteFile(filepath.Join(dir, "b.csv"), []byte(""), 0644)
	os.WriteFile(filepath.Join(dir, "readme.txt"), []byte(""), 0644)

	files, err := GetCSVFiles(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(files) != 2 {
		t.Errorf("expected 2 CSV files, got %d", len(files))
	}
}

func TestGetCSVFiles_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	files, err := GetCSVFiles(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(files) != 0 {
		t.Errorf("expected 0 files, got %d", len(files))
	}
}

func TestGetCSVFiles_InvalidDir(t *testing.T) {
	_, err := GetCSVFiles("/nonexistent/path/xyz")
	if err == nil {
		t.Error("expected error for invalid directory")
	}
}

func TestGetCSVFiles_NoCSVFiles(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "data.json"), []byte(""), 0644)

	files, err := GetCSVFiles(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(files) != 0 {
		t.Errorf("expected 0 CSV files, got %d", len(files))
	}
}

// --- SearchTxnParallel ---

func TestSearchTxnParallel_Found(t *testing.T) {
	dir := t.TempDir()
	f := writeTempCSV(t, dir, "test.csv", "TXN001,REF001,UPI\nTXN002,REF002,Wallet\n")

	result, ok := SearchTxnParallel([]string{f}, "TXN001")
	if !ok {
		t.Fatal("expected to find TXN001")
	}
	if result.Txn != "TXN001" {
		t.Errorf("expected Txn=TXN001, got %s", result.Txn)
	}
	if result.RefNo != "REF001" {
		t.Errorf("expected RefNo=REF001, got %s", result.RefNo)
	}
	if result.PaymentInfo != "UPI" {
		t.Errorf("expected PaymentInfo=UPI, got %s", result.PaymentInfo)
	}
	// gin_api version stores only the base filename
	if result.FileName != "test.csv" {
		t.Errorf("expected FileName=test.csv, got %s", result.FileName)
	}
}

func TestSearchTxnParallel_NotFound(t *testing.T) {
	dir := t.TempDir()
	f := writeTempCSV(t, dir, "test.csv", "TXN001,REF001,UPI\n")

	result, ok := SearchTxnParallel([]string{f}, "TXN999")
	if ok || result != nil {
		t.Error("expected not found")
	}
}

func TestSearchTxnParallel_EmptyFileList(t *testing.T) {
	result, ok := SearchTxnParallel([]string{}, "TXN001")
	if ok || result != nil {
		t.Error("expected no result for empty file list")
	}
}

func TestSearchTxnParallel_MultipleFiles(t *testing.T) {
	dir := t.TempDir()
	f1 := writeTempCSV(t, dir, "file1.csv", "TXN001,REF001,UPI\n")
	f2 := writeTempCSV(t, dir, "file2.csv", "TXN002,REF002,Wallet\n")

	result, ok := SearchTxnParallel([]string{f1, f2}, "TXN002")
	if !ok {
		t.Fatal("expected to find TXN002")
	}
	if result.Txn != "TXN002" {
		t.Errorf("expected TXN002, got %s", result.Txn)
	}
}

func TestSearchTxnParallel_EmptyCSVFile(t *testing.T) {
	dir := t.TempDir()
	f := writeTempCSV(t, dir, "empty.csv", "")

	result, ok := SearchTxnParallel([]string{f}, "TXN001")
	if ok || result != nil {
		t.Error("expected no result from empty file")
	}
}

func TestSearchTxnParallel_LastRecordInFile(t *testing.T) {
	dir := t.TempDir()
	f := writeTempCSV(t, dir, "test.csv", "TXN001,REF001,UPI\nTXN002,REF002,Wallet\nTXN003,REF003,NetBanking\n")

	result, ok := SearchTxnParallel([]string{f}, "TXN003")
	if !ok {
		t.Fatal("expected to find TXN003 at end of file")
	}
	if result.Txn != "TXN003" {
		t.Errorf("expected TXN003, got %s", result.Txn)
	}
}

func TestSearchTxnParallel_DuplicateAcrossFiles_ReturnsOne(t *testing.T) {
	dir := t.TempDir()
	f1 := writeTempCSV(t, dir, "file1.csv", "TXN001,REF001,UPI\n")
	f2 := writeTempCSV(t, dir, "file2.csv", "TXN001,REF001,UPI\n")

	result, ok := SearchTxnParallel([]string{f1, f2}, "TXN001")
	if !ok {
		t.Fatal("expected to find TXN001")
	}
	if result.Txn != "TXN001" {
		t.Errorf("expected TXN001, got %s", result.Txn)
	}
}
