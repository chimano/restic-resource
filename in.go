// +build in

package main

import (
	"os"

	"github.com/chimano/restic-resource/common"
)

func init() {
	inCommand := &InCommand{}
	subCommands = append(subCommands, inCommand)
}

type InCommand struct {
}

func (c *InCommand) Execute(request *common.Request) {
	destinationDir := os.Args[1]
}
