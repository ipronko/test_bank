package storage

import (
	"sync"
	"testing"
)

const startMoney = 1000000
const transferSum = 1

func TestStorageConcurrent(t *testing.T) {
	storage := NewStorage()
	firstA := storage.CreateAccount(startMoney)
	secondA := storage.CreateAccount(0)

	wg := sync.WaitGroup{}
	for i := 0; i < startMoney/transferSum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := storage.Transfer(firstA, secondA, transferSum)
			if err != nil {
				t.Fatal(err)
			}
		}()
	}
	wg.Wait()

	if sum, err := storage.GetBalance(firstA); err != nil {

	} else if sum != 0 {
		t.Error("first account balance is not 0")
	}

	if sum, err := storage.GetBalance(secondA); err != nil {
		t.Error(err)
	} else if sum != startMoney {
		t.Error("second account balance is not ", startMoney)
	}
}

func BenchmarkStorageTransfer(b *testing.B) {
	storage := NewStorage()
	firstA := storage.CreateAccount(float64(b.N))
	secondA := storage.CreateAccount(0)

	for i := 0; i < b.N; i++ {
		err := storage.Transfer(firstA, secondA, 1)
		if err != nil {
			b.Fatal(err)
		}
	}
}
