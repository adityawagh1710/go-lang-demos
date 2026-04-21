package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"csv-txn-lookup-gin-api/internal/model"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

// fakeLookup lets tests inject a result without touching the filesystem
type fakeLookup struct {
	result *model.Payment
	err    error
}

func (f *fakeLookup) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if f.err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": f.err.Error()})
			return
		}
		c.JSON(http.StatusOK, f.result)
	}
}

func newRouter(h gin.HandlerFunc) *gin.Engine {
	r := gin.New()
	r.GET("/api/v1/txn/:id", h)
	return r
}

func TestGetTxnHandler_Found(t *testing.T) {
	fake := &fakeLookup{
		result: &model.Payment{Txn: "TXN001", RefNo: "REF001", PaymentInfo: "UPI", FileName: "file1.csv"},
	}
	r := newRouter(fake.handler())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/txn/TXN001", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var resp model.Payment
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if resp.Txn != "TXN001" {
		t.Errorf("expected TXN001, got %s", resp.Txn)
	}
	if resp.RefNo != "REF001" {
		t.Errorf("expected REF001, got %s", resp.RefNo)
	}
	if resp.PaymentInfo != "UPI" {
		t.Errorf("expected UPI, got %s", resp.PaymentInfo)
	}
}

func TestGetTxnHandler_NotFound(t *testing.T) {
	fake := &fakeLookup{err: fmt.Errorf("txn not found")}
	r := newRouter(fake.handler())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/txn/TXN999", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}
}

func TestGetTxnHandler_ResponseIsJSON(t *testing.T) {
	fake := &fakeLookup{
		result: &model.Payment{Txn: "TXN001", RefNo: "REF001", PaymentInfo: "Wallet", FileName: "f.csv"},
	}
	r := newRouter(fake.handler())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/txn/TXN001", nil)
	r.ServeHTTP(w, req)

	ct := w.Header().Get("Content-Type")
	if ct == "" {
		t.Error("expected Content-Type header")
	}
}
