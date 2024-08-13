package helper

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
)

const SigningKey = "testingkey"

type UserId string

type HttpHandlerFunc func(w http.ResponseWriter, r *http.Response)

type ControllerFunc func(context.Context, http.ResponseWriter, *http.Request) error

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

type CtxRequestId string

const RequestId CtxRequestId = "requestId"

// save panic to file
func SaveLog(logs ...string) {
	dir := "logs"
	fileName := time.Now().Format("2006-01-02") + ".log"
	filePath := dir + "/" + fileName

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal("error creating directory", err)
		return
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("error occured while opening the file", err)
		return
	}

	printLog := time.Now().String()
	for _, txt := range logs {
		printLog = printLog + " " + txt + " "
	}
	if _, err := file.Write([]byte("\n" + printLog)); err != nil {
		log.Fatal("error occured while writing the log file", err)
	}
}
