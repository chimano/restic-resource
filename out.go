// +build out
package main

import (
	"os"

	"github.com/arekmano/restic-store/store"
	"github.com/chimano/restic-resource/common"
)

func init() {
	outCommand := &OutCommand{}
	subCommands = append(subCommands, outCommand)
}

type OutCommand struct {
}

func (c *OutCommand) Execute(request *common.Request) {
	sourceDir := os.Args[1]
	resticConfig := &store.ResticConfiguration{Host: request.Source.Host, Repository: request.Source.Repository}
	restic := store.NewRestic(resticConfig)
	resticInput := &store.ResticOptions{Options: request.Source.Options, Tags: request.Source.Tags}
	restic.Put(sourceDir, resticInput)
}
