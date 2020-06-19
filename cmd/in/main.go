package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/arekmano/restic-store/exec"
	"github.com/arekmano/restic-store/store"
	"github.com/chimano/restic-resource/common"
	"github.com/sirupsen/logrus"
)

func main() {
	c := &InCommand{CommandReader: &common.Reader{}}

	command, err := c.generateResticCommand()
	if err != nil {
		logrus.Fatal(err)
	}

	output, err := command.Execute()
	if err != nil {
		logrus.Fatal(err)
	}

	versionID, err := c.parseCommandOutput(output)
	if err != nil {
		logrus.Fatal(err)
	}
	response := InResponse{Version: common.Version{
		VersionID: versionID,
	},
	}
	json.NewEncoder(os.Stdout).Encode(response)
}

type InCommand struct {
	CommandReader common.InputReader
}

func (c *InCommand) generateResticCommand() (*exec.ResticCommand, error) {
	destDir, err := c.CommandReader.GetInputDirectory()
	if err != nil {
		return nil, err
	}
	request, err := c.CommandReader.GetRequest()
	if err != nil {
		return nil, err
	}

	resticConfig := &store.ResticConfiguration{Host: request.Source.Host, Repository: request.Source.Repository}
	restic := store.NewRestic(resticConfig)
	resticInput := &store.ResticOptions{Options: request.Source.Options, Tags: request.Source.Tags}
	versionID := request.Version.VersionID
	return restic.Get(destDir, resticInput, versionID)
}

func (c *InCommand) parseCommandOutput(output []byte) (string, error) {
	// SAMPLE OUTPUT
	// restoring <Snapshot 37e24a73 of [/blah/blah/go.mod] at 2020-06-17 17:10:12.843022846 -0400 EDT by user@host> to destDir/

	pattern := regexp.MustCompile(`Snapshot\s[a-zA-Z0-9]+\s`)
	match := string(pattern.Find(output))
	fmt.Print(string(output))
	splittedMatch := strings.Split(match, " ")
	if len(splittedMatch) < 2 {
		return "", errors.New("Could not identify snapshot version from output")
	}
	id := strings.Split(match, " ")[1]
	return id, nil
}

type InResponse struct {
	Version  common.Version    `json:"version"`
	Metadata map[string]string `json:"metadata"`
}
