package chapter2

import (
	"testing"
)

var numberString = "   -46"

// Pada bagian ini dicontohkan nilai dari result sesuai dengan ekspektasi
func TestMyAtoiPassed(t *testing.T) {
	result := MyAtoi(numberString)
	if result != -46 {
		// Fungsi ini akan menjalankan fungsi Fail() dan Logf() dimana ketika fungsi ini dipanggil akan memberikan
		// signal error (fail) pada proses unit testing dan menampilkan pesan sesuai dengan yang ditulis pada argument fungsi.
		// Dan akan dilanjutkan pada bari kode berikutnya.
		t.Errorf("Nilai yang dimasukan: %s, ekspektasi -46 tetapi hasil %d", numberString, result)
	}

	t.Log("Proses testing MyAtoi example 1 selesai!")
}

// Pada bagian ini dicontohkan nilai dari result tidak sesuai dengan ekspektasi dan akan memanggil fungsi Errorf().
// Fungsi Errorf() dan Error() memiliki kesamaan dimana ketika fungsi ini dipanggail akan memanggil fungsi Fail() dan dilanjutkan
// dengan memberi pesan, Error() -> Log() sedangkan Errorf() -> Logf().
func TestMyAtoiFailedError(t *testing.T) {
	result := MyAtoi(numberString)
	if result != -461 {
		// Fungsi ini akan menjalankan fungsi Fail() dan Logf() dimana ketika fungsi ini dipanggil akan memberikan
		// signal error (fail) pada proses unit testing dan menampilkan pesan sesuai dengan yang ditulis pada argument fungsi.
		// Dan akan dilanjutkan pada bari kode berikutnya.
		t.Errorf("Nilai yang dimasukan: -461, ekspektasi -461 tetapi hasil %d", result)
	}

	t.Log("Proses testing MyAtoi example 2 selesai!")
}

// Pada bagian ini dicontohkan nilai dari result tidak sesuai dengan ekspektasi dan akan memanggil fungsi Fatalf().
// Fungsi Fatalf() dan Fatal() memiliki kesamaan dimana ketika fungsi ini dipanggail akan memanggil fungsi FailNow() dan dilanjutkan
// dengan memberi pesan, Fatal() -> Log() sedangkan Fatalf() -> Logf().
func TestMyAtoiFailedFatal(t *testing.T) {
	result := MyAtoi(numberString)
	if result != -461 {
		// Fungsi ini akan menjalankan fungsi FailNow() dan Logf() dimana ketika fungsi ini dipanggil akan memberikan
		// signal error (fail) pada proses unit testing dan menampilkan pesan sesuai dengan yang ditulis pada argument fungsi.
		// Dan mengabaikan baris kode berikutnya (exit) pada fungsi ini (dilanjutkan pada fungsi test berikutnya).
		t.Fatalf("Nilai yang dimasukan: -461, ekspektasi -461 tetapi hasil %d", result)
	}

	t.Log("Proses testing MyAtoi example 3 selesai!")
}

// Pada bagian ini menguji semua kondisi yang mungkin terjadi ke fungsi MyAtoi untuk memastikan fungsi berjalan sesuai dengan
// apa yang diharapkan. Dimana menerapkan koleksi nilai pengujian yang ditampung pada slice yang bertipe struct dimana struct
// yang digunakan memiliki properti Input (masukan ke fungsi) dan Want (nilai yang diharapkan).
func TestMyAtoiCollectionValues(t *testing.T) {
	var tests = []struct {
		Input string
		Want  int
	}{
		{
			Input: "",
			Want:  0,
		},
		{
			Input: "     -46",
			Want:  -46,
		},
		{
			Input: "+46",
			Want:  46,
		},
		{
			Input: "     4567/1234",
			Want:  4567,
		},
		{
			Input: "89876Hello Word",
			Want:  89876,
		},
		{
			Input: "1234567891032",
			Want:  2147483647,
		},
		{
			Input: "-1234567891032",
			Want:  -2147483648,
		},
		{
			Input: "1024",
			Want:  1024,
		},
		{
			Input: "-1024",
			Want:  -1024,
		},
	}

	for _, test := range tests {
		if result := MyAtoi(test.Input); result != test.Want {
			t.Fatalf("Testing with input value %s, expected %d but result %d", test.Input, test.Want, result)
		}
	}

	t.Log("Testing function MyAtoi finished")
}
