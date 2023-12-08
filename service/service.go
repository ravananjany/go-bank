package service

import (
	"fmt"
	inerror "go-bank/error"
	"go-bank/models"
	"go-bank/repo"
	"log"
	"strconv"
)

type ServiceInterface interface {
	GetAccounts() []models.Account
	Transfer(trequest models.TransferRequest) error
	GetAccount(id string) models.Account
	CheckUser(uname string) string
}

type service struct {
	Repo repo.RepositoryInterface
}

func NewService(r repo.RepositoryInterface) ServiceInterface {
	return &service{Repo: r}
}

func (s *service) CheckUser(uname string) string {
	return s.Repo.CheckUser(uname)
}

func (s *service) GetAccounts() []models.Account {
	return s.Repo.GetAccounts()
}

func (s *service) GetAccount(id string) models.Account {
	return s.Repo.GetAccount(id)
}

func (s *service) Transfer(trequest models.TransferRequest) error {
	functionDesc := "Transfer service"
	fa := s.Repo.GetAccount(trequest.FromAccountID)
	if (fa == models.Account{}) {
		log.Println("invalid account id", inerror.ErrINvalidAcc)
		return inerror.ErrINvalidAcc
	}
	accountBalanceFrom := s.Repo.GetBalance(trequest.FromAccountID)
	amountTobeDebited := trequest.Amount
	updatedBalanceFrom, err := updateFromAcc(accountBalanceFrom, amountTobeDebited)
	if err != nil {
		log.Println("Error in updating"+functionDesc, err.Error())
		return err
	}
	s.Repo.UpdateAccount(trequest.FromAccountID, updatedBalanceFrom)

	accountBalanceTo := s.Repo.GetBalance(trequest.ToAccountID)
	amountTobeCredit := trequest.Amount
	//TODO need to to check err
	updatedBalanceTo, _ := updateToAcc(trequest.ToAccountID, accountBalanceTo, amountTobeCredit)
	s.Repo.UpdateAccount(trequest.ToAccountID, updatedBalanceTo)
	return nil
}

func updateFromAcc(amount string, tamount string) (string, error) {
	var updatedAmount string
	amountIn := Float(amount)
	amountOut := Float(tamount)
	if amountIn-amountOut < 0 {
		return updatedAmount, inerror.ErrInsuff
	}
	updatedFloatAmount := amountIn - amountOut
	updatedAmount = String(updatedFloatAmount)
	return updatedAmount, nil
}

func updateToAcc(aid string, amount string, tamount string) (string, error) {
	var updatedAmount string
	amountIn := Float(amount)
	amountOut := Float(tamount)
	updatedFloatAmount := amountIn + amountOut
	updatedAmount = String(updatedFloatAmount)
	return updatedAmount, nil
}

func Float(am string) float64 {
	fl, err := strconv.ParseFloat(am, 64)
	if err != nil {
		return fl
	}
	return fl
}

func String(v float64) string {
	fmt64 := fmt.Sprintf("%.2f", v)
	return fmt64
}
