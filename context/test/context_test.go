package test

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	// nanti ketika bikin web dg golang tidak perlu buat context manual, nanti secara otomatis setiap request yg masuk akan ada context
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	// membuat context baru / menambah child (immutable)
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	// context get value
	fmt.Println(contextF.Value("f")) // dapat
	fmt.Println(contextF.Value("c")) // dapat milik parent
	fmt.Println(contextF.Value("b")) // tidak dapat, beda parent
	fmt.Println(contextA.Value("b")) // tidak bisa mengambil data child
}

// solusi goroutine leak dengan context with cancel (dengan buat parameter context)
func CreateConter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		// melakukan perulangan yg tidak henti, mengirim data ke channel
		for {
			// lakukan pengecekan context with cancel
			select {
			// jika sudah done, maka return
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++

				// simulasi slow response untuk context with timeout
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

// simulasi goroutine leak
func TestContextWithCancel(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine()) // total = 2

	// kirim contextnya
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateConter(ctx)

	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}

	}
	// jika sudah break akan panggil cancel, ngirim sinyal cancel ke context
	cancel()

	// menunggu 2 detik untuk lihat hasil total goroutine, biar ndak kecepatan
	time.Sleep(2 * time.Second)

	fmt.Println("total goroutine", runtime.NumGoroutine()) // total = 3, setelah if = 10 goroutinenya masih nyala harusnya sudah mati kembali 2 (goroutinenya jalan terus). ini bahaya tiap request ada leak goroutine 1 saja, jika ada 1000 request jadi ada 1000 goroutine yg berjalan terus. hal ini dapat menyebabkan aplikasi berjalan lambat

	// solusinya dengan implementasi context with cancel
	// sekarang total goroutinenya sudah 2
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine()) // total = 2

	// kirim contextnya
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	// kenapa defer cancel() untuk jaga-jaga misal tidak sampe 5 detik juga harus tetap panggil cancel() agar goroutine berhenti
	defer cancel()

	destination := CreateConter(ctx)

	for n := range destination {
		fmt.Println("counter", n)
	}

	// menunggu 2 detik untuk lihat hasil total goroutine, biar ndak kecepatan
	time.Sleep(2 * time.Second)

	fmt.Println("total goroutine", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine()) // total = 2

	// kirim contextnya
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	// kenapa defer cancel() untuk jaga-jaga misal tidak sampe 5 detik juga harus tetap panggil cancel() agar goroutine berhenti
	defer cancel()

	destination := CreateConter(ctx)

	for n := range destination {
		fmt.Println("counter", n)
	}

	// menunggu 2 detik untuk lihat hasil total goroutine, biar ndak kecepatan
	time.Sleep(2 * time.Second)

	fmt.Println("total goroutine", runtime.NumGoroutine())
}
