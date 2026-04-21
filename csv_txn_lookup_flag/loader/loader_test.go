package loader

import (
	"fmt"
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
	os.WriteFile(filepath.Join(dir, "c.txt"), []byte(""), 0644)

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

// --- LoadWithWorkerPool ---

func TestLoadWithWorkerPool_FindsRecord(t *testing.T) {
	dir := t.TempDir()
	f := writeTempCSV(t, dir, "test.csv", "TXN001,REF001,UPI\nTXN002,REF002,Wallet\n")

	index := LoadWithWorkerPool([]string{f}, 2)

	val, ok := index["TXN001"]
	if !ok {
		t.Fatal("expected TXN001 in index")
	}
	if val.RefNo != "REF001" {
		t.Errorf("expected REF001, got %s", val.RefNo)
	}
	if val.PaymentInfo != "UPI" {
		t.Errorf("expected UPI, got %s", val.PaymentInfo)
	}
}

func TestLoadWithWorkerPool_NotFound(t *testing.T) {
	dir := t.TempDir()
	f := writeTempCSV(t, dir, "test.csv", "TXN001,REF001,UPI\n")

	index := LoadWithWorkerPool([]string{f}, 2)
	if _, ok := index["TXN999"]; ok {
		t.Error("TXN999 should not be in index")
	}
}

func TestLoadWithWorkerPool_MultipleFiles(t *testing.T) {
	dir := t.TempDir()
	f1 := writeTempCSV(t, dir, "file1.csv", "TXN001,REF001,UPI\n")
	f2 := writeTempCSV(t, dir, "file2.csv", "TXN002,REF002,Wallet\n")

	index := LoadWithWorkerPool([]string{f1, f2}, 2)

	if _, ok := index["TXN001"]; !ok {
		t.Error("expected TXN001 in index")
	}
	if _, ok := index["TXN002"]; !ok {
		t.Error("expected TXN002 in index")
	}
}

func TestLoadWithWorkerPool_EmptyFile(t *testing.T) {
	dir := t.TempDir()
	f := writeTempCSV(t, dir, "empty.csv", "")

	index := LoadWithWorkerPool([]string{f}, 2)
	if len(index) != 0 {
		t.Errorf("expected empty index, got %d entries", len(index))
	}
}

func TestLoadWithWorkerPool_StoresFileName(t *testing.T) {
	dir := t.TempDir()
	f := writeTempCSV(t, dir, "myfile.csv", "TXN001,REF001,UPI\n")

	index := LoadWithWorkerPool([]string{f}, 2)

	val, ok := index["TXN001"]
	if !ok {
		t.Fatal("expected TXN001 in index")
	}
	if val.FileName != f {
		t.Errorf("expected FileName=%s, got %s", f, val.FileName)
	}
}

func TestLoadWithWorkerPool_ManyRecords(t *testing.T) {
	dir := t.TempDir()
	content := ""
	for i := 0; i < 100; i++ {
		content += fmt.Sprintf("TXN%03d,REF%03d,UPI\n", i, i)
	}
	f := writeTempCSV(t, dir, "big.csv", content)

	index := LoadWithWorkerPool([]string{f}, 4)
	if len(index) != 100 {
		t.Errorf("expected 100 records, got %d", len(index))
	}
}
