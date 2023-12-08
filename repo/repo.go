package repo

import (
	"go-bank/models"
	"sync"

	"github.com/gin-gonic/gin"
)

var AuthTable gin.H

type RepositoryInterface interface {
	UploadAccounts(a []models.Account) error
	GetAccounts() []models.Account
	GetBalance(id string) string
	UpdateAccount(id, am string)
	GetAccount(id string) models.Account
	CheckUser(uname string) string
}

type repository struct {
	Database models.DB
	Mutex    sync.Mutex
}

func NewRepository() RepositoryInterface {
	AuthTable = gin.H{}
	db := models.DB{
		Store: make(map[string]*models.Account),
	}

	return &repository{
		Database: db,
		Mutex:    sync.Mutex{},
	}
}

func (r *repository) UploadAccounts(a []models.Account) error {
	for _, ac := range a {
		ptr := &models.Account{
			Id:      ac.Id,
			Name:    ac.Name,
			Balance: ac.Balance,
		}
		r.Database.Store[ac.Id] = ptr
		AuthTable[ac.Name] = ac.Id

	}
	return nil
}

func (r *repository) GetAccount(id string) models.Account {
	if r.Database.Store[id] == nil {
		return models.Account{}
	}
	return *r.Database.Store[id]
}

func (r *repository) GetAccounts() []models.Account {
	var accounts []models.Account
	for _, v := range r.Database.Store {
		accounts = append(accounts, *v)
	}
	return accounts
}

func (r *repository) GetBalance(id string) string {
	return r.Database.Store[id].Balance
}

func (r *repository) UpdateAccount(id, am string) {
	r.Mutex.Lock()
	r.Database.Store[id].Balance = am
	r.Mutex.Unlock()
}

func (r *repository) CheckUser(uname string) string {

	v, ok := AuthTable[uname]
	if ok {
		return v.(string)
	}
	return ""
}
