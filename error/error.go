package error

import "errors"

type ErrTypes string

const (
	INSUFF            ErrTypes = "Insufficent Balance"
	NOT_VALID_USER    ErrTypes = "Not a valid User"
	NOT_VALID_AMOUNT           = "Not a valid Amount , Please Enter a valid Amount"
	NOT_VALID_ACCOUNT          = "Not a valid acount , Invalid account number"
)

var (
	ErrInsuff       error = errors.New("insufficient Balance")
	ErrNotValidUsr  error = errors.New("invalid user")
	ErrNotValidAmnt error = errors.New("invalid amount")
	ErrINvalidAcc   error = errors.New("invalid account")
)
