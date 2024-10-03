package main

import (
	"fmt"
	"sync"
)

// struct data bank
type BankAccount struct {
	Saldo int
	mutex sync.Mutex
}

// struct userbank
type UserBank struct {
	ID     int
	Nama   string
	Saldo  int
	Status string
}

// function implementasi mutex
func ViewSaldo(account *BankAccount, wg *sync.WaitGroup, saldoChannel chan<- string) {
	defer wg.Done()
	account.mutex.Lock()
	defer account.mutex.Unlock()
	//validasi saldo
	if account.Saldo <= 0 {
		saldoChannel <- "Gagal mengambil saldo: Saldo tidak mencukupi."
		return
	}
	saldoChannel <- fmt.Sprintf("Saldo saat ini: %d", account.Saldo)
}

func main() {
	//membuat tipe data struct
	Bank := BankAccount{
		Saldo: 200,
	}
	//membuat tipe data strut
	transaksi := []UserBank{
		{ID: 1, Nama: "andi", Saldo: 100, Status: "tarik"},
		{ID: 2, Nama: "nabila", Saldo: 300, Status: "setor"},
		{ID: 3, Nama: "anton", Saldo: 700, Status: "tarik"},
		{ID: 4, Nama: "sena", Saldo: 100, Status: "tarik"},
	}

	var wg sync.WaitGroup
	saldoChannel := make(chan string) //membuat chanel
	//membuat antrian dengan goroutine
	for _, t := range transaksi {
		wg.Add(1)
		go func(t UserBank) {
			defer wg.Done()
			Bank.mutex.Lock()
			defer Bank.mutex.Unlock()
			//validasi saldo
			if t.Status == "tarik" {
				if Bank.Saldo < t.Saldo {
					saldoChannel <- fmt.Sprintf("Gagal mengambil saldo: Saldo tidak mencukupi untuk %s", t.Nama)
				} else {
					Bank.Saldo -= t.Saldo
					saldoChannel <- fmt.Sprintf("%s berhasil menarik %d. Saldo sekarang: %d", t.Nama, t.Saldo, Bank.Saldo)
				}
			} else if t.Status == "setor" {
				Bank.Saldo += t.Saldo
				saldoChannel <- fmt.Sprintf("%s berhasil setor %d. Saldo sekarang: %d", t.Nama, t.Saldo, Bank.Saldo)
			}
		}(t)
	}
	//menunggu goroutine selesai
	go func() {
		wg.Wait()
		close(saldoChannel)
	}()
	//menampilkan pesan
	for msg := range saldoChannel {
		fmt.Println(msg)
	}
}
