package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/chimano/restic-resource/common"
)

var subCommands []Command

func main() {
	var request *common.Request
	parseRequest(request)

	for _, command := range subCommands {
		command.Execute(request)
	}
}

func parseRequest(request *common.Request) {
	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		log.Fatal(err.Error())
	}
}

type Command interface {
	Execute(request *common.Request)
}
