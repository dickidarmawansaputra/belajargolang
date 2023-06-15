package test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)

		// pastikan ada pengirim dan penerima data channel
		// kirim data ke channel
		channel <- "Dicki Darmawan Saputra"

		fmt.Println("selesai kirim data ke channel")
	}()

	// menerima data dari channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

	// mengirim data ke channel
	// channel <- "Dicki"

	// menerima data dari channel
	// data := <- channel
}

// tidak butuh pointer jika di channel
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Dicki Darmawan Saputra"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	// menerima data dari channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// hanya bisa mengirim data ke channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Dicki Darmawan Saputra"
}

// hanya bisa menerima data dari channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	// anonymous func
	go func() {
		channel <- "Dicki"
		channel <- "Darmawan"
		channel <- "Saputra"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		// jika tidak diclose maka akan error deadlock (setelah data ke 9 akan ngecek terus)
		close(channel)
	}()

	// gunakan range channel jika datanya banyak atau lebih dari satu
	for data := range channel {
		fmt.Println("menerima data", data)
	}
}

func TestSelectChannel(t *testing.T) {
	channelSatu := make(chan string)
	channelDua := make(chan string)
	defer close(channelSatu)
	defer close(channelDua)

	go GiveMeResponse(channelSatu)
	go GiveMeResponse(channelDua)

	counter := 0
	// jika tanpa perulangan maka yang keambil hanya 1 channel
	for {
		select {
		case data := <-channelSatu:
			fmt.Println("data channel 1:", data)
			counter++
		case data := <-channelDua:
			fmt.Println("data channel 2:", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channelSatu := make(chan string)
	channelDua := make(chan string)
	defer close(channelSatu)
	defer close(channelDua)

	go GiveMeResponse(channelSatu)
	go GiveMeResponse(channelDua)

	counter := 0
	// jika tanpa perulangan maka yang keambil hanya 1 channel
	for {
		select {
		case data := <-channelSatu:
			fmt.Println("data channel 1:", data)
			counter++
		case data := <-channelDua:
			fmt.Println("data channel 2:", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
