package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

var logFile *os.File
var logger *slog.Logger

type blob struct {
	name string
	age  int
}

func init() {
	log.Println("init function ...")
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal("Log file cant be open or create ", err)
	}
	logFile = file
	log.SetOutput(file)
	logger = slog.New(slog.NewJSONHandler(file, nil))
}

func main() {
	user := blob{name: "daddydemir", age: -1}

	slog.Debug("started")
	logger.Error(" user created ", "userId", 1, "userName", "daddydemir")
	logger.Warn("warning ", "userId", 1)
	logger.Info("user ", "person", fmt.Sprint(user))

	err := logFile.Close()
	if err != nil {
		log.Println("error: ", err)
	}
}
