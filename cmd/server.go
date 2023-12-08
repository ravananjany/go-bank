package server

import (
	"encoding/json"
	"go-bank/middleware"
	"go-bank/models"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func init() {
	os.Setenv("API_SECRET", "secret_api_protected")
	os.Setenv("TOKEN_HOUR_LIFESPAN", "1")
}

func Start() {
	r := gin.Default()
	r.Use(middleware.Logger())
	buildRoutes()
	wc := make(chan int)
	go func(wc chan int) {
		err := InitDB()
		if err != nil {
			log.Fatalln("error in injecting accounts , please reload")
			return
		}
		wc <- 1
	}(wc)
	<-wc
	AddRoutes(r)
	ServerRun(r)
}

// Reads the json file and uplodas the data
func InitDB() error {
	log.Println("In initDB")
	var accounts []models.Account
	file, err := os.Open("accounts-mock.json")
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &accounts)
	if err != nil {
		return err
	}
	err = Repository.UploadAccounts(accounts)
	if err != nil {
		return err
	}
	return nil
}

func ServerRun(r *gin.Engine) {
	go func(srv *gin.Engine) {
		log.Println("Ready to transafer")
		if err := srv.Run(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}

	}(r)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
}
