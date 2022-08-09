package chapter3

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Run Multiple Func:
// go test -v -timeout 30s -run "^(TestSomething|TestJSONPayload)"$ learning/unit-tests/examples/chapter3

// To Do: Contoh penggunaan assert dan require dengan berbagai fungsi yang dapat digunakan.
func TestSomething(t *testing.T) {
	// Penggunaan assert, ketika diketemukan ketidak samaan package akan mengirimkan signal error pada
	// unit test dan melanjutkan pada baris code berikutnya ( Fail() ).

	// Nilai ekspektasi dengan aktual harus sama, tidak hanya nilai untuk tipe data yang digunakan harus sama.
	assert.Equal(t, 162, 162, "Nilai harus sama")

	// Nilai ekspektasi dengan aktual harus tidak sama. Tidak sama bisa jadi tidak sama nilai atau tipe datanya.
	// assert.NotEqual(t, 162, float64(162), "Nilai harus tidak sama") -> TRUE
	assert.NotEqual(t, 162, 1024, "Nilai harus tidak sama")

	// Nilai harus nil
	var err1 = func() error {
		return nil
	}()
	assert.Nil(t, err1, "Nilai harus nil")
	// Dapat juga menggunakan pengecekan error
	assert.NoError(t, err1)

	// Nilai harus tidak nil. Untuk nilai kembali dari pemanggilan fungsi di assert
	// akan mengembalikan nilai boolean dan dapat digunakan untuk melakukan pengkondisian.
	var err2 = func() error {
		return errors.New("something error")
	}()
	if assert.NotNil(t, err2) {
		assert.Equal(t, "something error", err2.Error())
	}
	// Dapat juga menggunakan pengecekan error
	assert.Error(t, err2)

	// Nilai harus berisikan (mengandung) kata sesuai dengan ekspektasi
	text := "nilai harus lebih kecil sama dengan 2"
	assert.Contains(t, text, "harus lebih kecil", "harus mengandung kata: harus lebih kecil")

	t.Log("Pengujian dengan assert testify selesai!")

	// Penggunaan require, memiliki fungsi yang sama dengan assert. Perbedaannya
	// package ini tidak mengembalikan nilai boolean (true/false) jadi ketika diketemukan
	// ketidak samaan akan langsung mengirimkan signal error ke proses unit test dan mengakhir
	// proses pada fungsi tsb ( FailNow() )
	var session = func() any {
		return "example session"
	}()
	require.Equal(t, "example session", session.(string), "require, nilai harus sama")

	t.Log("Pengujian dengan require testify selesai!")
}

// To Do: Melakukan perbandingan nilai dengan format data JSON
func TestJSONPayload(t *testing.T) {
	payload := `
		{
			"id": 1,
			"title": "iPhone 9",
			"description": "An apple mobile which is nothing like apple",
			"price": 549,
			"stock": 94,
			"brand": "Apple",
			"category": "smartphones"
		}
	`

	var result = func() any {
		return struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Price       int    `json:"price"`
			Stock       int    `json:"stock"`
			Brand       string `json:"brand"`
			Category    string `json:"category"`
		}{
			ID:          1,
			Title:       "iPhone 9",
			Description: "An apple mobile which is nothing like apple",
			Price:       549,
			Stock:       94,
			Brand:       "Apple",
			Category:    "smartphones",
		}
	}()

	resMarshal, err := json.Marshal(result)
	if assert.NoError(t, err) {
		// Nilai harus dalam format JSON, dan akan dibandingkan dengan menggunakan sistem conversi ke map[string]interface{}
		assert.JSONEq(t, payload, string(resMarshal))
	}
}
