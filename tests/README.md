# Go Basic Unit Tests
Pada bagian ini akan dijelaskan penggunaan package `testing` khususnya untuk `testing.T` dan `assertion` menggunakan library `testify`.

Penulisan fungsi testing di Go sendiri memiliki aturan dimana nama fungsi harus diawali dengan `Test`, memiliki satu parameter dengan tipe `*testing.T`, dan tanpa mengembalikan nilai (return value).
```go
func TestMyLtrim(t *testing.T) {
    // code unit testing
}
```

Pembagian pembahasan :
- Basic Unit Tests pada [chapter1](https://github.com/dewidyabagus/example-unit-tests-with-go/tree/master/examples/chapter1) dan [chapter2](https://github.com/dewidyabagus/example-unit-tests-with-go/tree/master/examples/chapter2) :
  - Fungsi `Fail()` dan `FailNow()`
  - Fungsi `Log()` dan `Logf()`
  - Fungsi `Error()` dan `Errorf()`
  - Fungsi `Fatal()` dan `Fatalf()`
- Assertion pada [chapter3](https://github.com/dewidyabagus/example-unit-tests-with-go/tree/master/examples/chapter3) :
  - Penggunaan `assert` dan `require`
  - Penggunaan fungsi `Equal()`, `Nil()`, `NotNil()`, `NoError()` dan lain - lain.
  
## Requirements
- Go Language 1.18.3

## Dependency
- Package [Testify](https://github.com/stretchr/testify)

## Menjalankan unit testing
Pada bahasa Go secara bawaan sudah di support untuk melakukan unit testing dengan menjalankan perintah:
```console
go test
```
Berikut attribut yang dapat digunakan:
- `-v` verbose, menampilkan pesan/log yang dijalankan.
- `-cover` mengaktifkan analisa tingkat cakupan kode (coverage).
- `-run` menjalankan fungsi test tertentu, contoh `-run=TestMyLtrimExample1`.
- `-race` memastikan code yang dibuat tidak berpotensi `race condition`.

Selain attribut yang dijelaskan diatas, masih banyak attribut yang dapat digunakan. Berikut [dokumentasi](https://pkg.go.dev/cmd/go/internal/test) menjalankan command `go test`.

## Referensi
- https://pkg.go.dev/cmd/go/internal/test
- https://github.com/stretchr/testify#readme
