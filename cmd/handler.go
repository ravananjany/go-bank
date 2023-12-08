package server

import (
	inerror "go-bank/error"
	"go-bank/middleware"
	"go-bank/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @GetAccounts API
// path /api/accounts
// response obj
func GetAccounts(ctx *gin.Context) {
	functionDesc := "GetAccounts"
	log.Println(functionDesc)
	uname := ctx.Request.Header.Get("Auth")
	id := Service.CheckUser(uname)
	if len(id) <= 0 {
		log.Println("No matching account id", id, uname)
		ctx.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	auth, err := middleware.AuthCheck(uname, id)
	if err != nil {
		log.Println("Error in generating token"+functionDesc, err.Error())
		ctx.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	accounts := Service.GetAccounts()
	ctx.Header("token", auth)
	ctx.JSON(http.StatusOK, accounts)
}

// @Transfer API
// Path /api/v1/transfer
// Body {"from_id" :"" , "to_id" : "" , "amount" : "" }
// response success msg
func Transfer(ctx *gin.Context) {
	functionDesc := "Transfer"
	log.Println(functionDesc)
	var req models.TransferRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Println("Error in marshal "+functionDesc, err.Error())
		ctx.JSON(http.StatusBadRequest, "Bad Request")
		return
	}
	if !ValidateAmount(req.Amount) {
		log.Println("Error in converting "+functionDesc, req.Amount)
		ctx.JSON(http.StatusBadRequest, inerror.NOT_VALID_AMOUNT)
		return
	}
	err = Service.Transfer(req)
	if err != nil {
		switch err {
		case inerror.ErrINvalidAcc:
			ctx.JSON(http.StatusBadRequest, inerror.NOT_VALID_ACCOUNT)
			return
		case inerror.ErrInsuff:
			ctx.JSON(http.StatusForbidden, inerror.INSUFF)
			return
		}
	}
	ctx.JSON(http.StatusAccepted, gin.H{"Status": "Success"})
}

func ValidateAmount(a string) bool {
	_, err := strconv.ParseFloat(a, 64)
	return err == nil
}
