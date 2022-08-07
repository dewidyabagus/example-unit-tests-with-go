# Unit Tests With testing Package
Pada bagian ini akan dijelaskan penggunaan package `testing` khususnya untuk `testing.T` yang diteruskan (parameter) melalui fungsi `Test`.
- Fungsi `Fail()` dan `FailNow()`
- Fungsi `Log()` dan `Logf()`
- Fungsi `Error()` dan `Errorf()`
- Fungsi `Fatal()` dan `Fatalf()`

## Requirements
- Go Language 1.18.3

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

Selain attribut yang dijelaskan diatas masih banyak lagi attribut yang dapat digunakan, berikut [dokumentasi go](https://pkg.go.dev/cmd/go/internal/test).

## Referensi
- https://pkg.go.dev/cmd/go/internal/test
