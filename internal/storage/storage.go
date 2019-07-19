package storage

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/atomic"
	"sync"
)

type NoSuchMoneyErr int

func (nsm NoSuchMoneyErr) Error() string {
	return fmt.Sprintf("not enough money in account with id %d", nsm)
}

type AccountNotFoundErr int64

func (anf AccountNotFoundErr) Error() string {
	return fmt.Sprintf("account with id %d not found", anf)
}

func NewStorage() *Storage {
	return &Storage{
		sequence: atomic.Int64{},
		accounts: make(map[int64]*account),
		m:        sync.RWMutex{},
	}
}

// We can use here sync.Map, but it will be much slower in processor with less than 16 cores
type Storage struct {
	sequence atomic.Int64
	accounts map[int64]*account
	m        sync.RWMutex
}

func (s *Storage) CreateAccount(balance float64) int64 {
	id := s.sequence.Inc()

	a := account{
		id:      id,
		balance: balance,
		m:       sync.RWMutex{},
	}

	s.m.Lock()
	s.accounts[id] = &a
	s.m.Unlock()
	return id
}

func (s *Storage) GetBalance(id int64) (float64, error) {
	s.m.RLock()
	a, ok := s.accounts[id]
	s.m.RUnlock()

	if !ok {
		return 0, AccountNotFoundErr(id)
	}
	return a.getBalance(), nil
}

func (s *Storage) Transfer(fromID, toID int64, sum float64) error {
	if sum <= 0 {
		return errors.New("sum of transfer must be positive value")
	}

	s.m.RLock()
	fromAccount, fromOk := s.accounts[fromID]
	toAccount, toOk := s.accounts[toID]
	s.m.RUnlock()

	if !fromOk {
		return AccountNotFoundErr(fromID)
	}
	if !toOk {
		return AccountNotFoundErr(toID)
	}

	err := fromAccount.minusBalance(sum)
	if err != nil {
		return err
	}

	toAccount.addBalance(sum)
	return nil
}

type account struct {
	id      int64
	balance float64
	m       sync.RWMutex
}

func (a *account) addBalance(sum float64) {
	a.m.Lock()
	a.balance += sum
	a.m.Unlock()
}

func (a *account) minusBalance(sum float64) error {
	a.m.Lock()
	defer a.m.Unlock()

	if a.balance < sum {
		return NoSuchMoneyErr(a.id)
	}

	a.balance -= sum
	return nil
}

func (a *account) getBalance() float64 {
	a.m.RLock()
	balance := a.balance
	a.m.RUnlock()
	return balance
}
