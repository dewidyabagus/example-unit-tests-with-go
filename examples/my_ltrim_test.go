package examples

import (
	"testing"
)

var input = "    Hello World!"

func TestMyLtrimExample1(t *testing.T) {

	result := MyLtrim(input)
	if result != "Hello World!" {
		// Ketika fungsi ini dipanggil maka memberikan signal ke proses testing bahwa terdapat proses test
		// yang gagal atau tidak sesuai dengan ekspetasi dan akan dilanjutkan ke baris kode berikutnya.
		t.Fail()
	}

	// Akan dicetak ketika terdapat fail atau tambahan argument -v ketika menjalankan go test
	t.Log("Proses testing example 1 selesai!")
}

// Pada bagian ini dicontohkan nilai dari result tidak sesuai dengan ekspektasi dan akan mengeksekusi
// fungsi Fail() untuk memberikan signal error test pada proses testing
func TestMyLtrimExample2(t *testing.T) {

	result := MyLtrim(input)
	if result != "Hello World" {
		// Ketika fungsi ini dipanggil maka memberikan signal ke proses testing bahwa terdapat proses test
		// yang gagal atau tidak sesuai dengan ekspetasi dan akan dilanjutkan ke baris kode berikutnya.
		t.Fail()
	}

	// Akan dicetak ketika terdapat fail atau tambahan argument -v ketika menjalankan go test
	t.Log("Proses testing example 2 selesai!")
}

// Pada bagian ini dicontohkan nilai dari result tidak sesuai dengan ekspektasi dan akan mengeksekusi
// fungsi FailNow() untuk memberikan signal error test pada proses testing
func TestMyLtrimExample3(t *testing.T) {

	result := MyLtrim(input)
	if result != "Hello World" {
		// Ketika fungsi ini dipanggil maka memberikan signal ke proses testing bahwa terdapat proses test
		// yang gagal atau tidak sesuai dengan ekspetasi dan menghentikan proses testing pada fungsi tersebut.
		t.FailNow()
	}

	// Akan dicetak ketika terdapat fail atau tambahan argument -v ketika menjalankan go test
	t.Log("Proses testing example 3 selesai!")
}

// Pada bagian ini digunakan untuk memperlihatkan bahwa ketika fungsi sebelumnya gagal menjalankan proses test dan
// memanggil FailNow() akan tetap dilanjutkan untuk fungsi test berikutnya.
func TestMyLtrimExample4(t *testing.T) {
	if result := MyLtrim(input); result != "Hello World!" {
		t.FailNow()
	}
	t.Log("Proses testing example 4 selesai!")
}
