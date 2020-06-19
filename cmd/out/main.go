package main

import (
	"encoding/json"
	"os"
	"regexp"

	"github.com/arekmano/restic-store/exec"
	"github.com/arekmano/restic-store/store"
	"github.com/chimano/restic-resource/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	c := &OutCommand{CommandReader: &common.Reader{}}

	command, err := c.generateResticCommand()
	if err != nil {
		logrus.Fatal(err)
	}

	output, err := command.Execute()
	if err != nil {
		logrus.Fatal(err)
	}

	parsedOutput, err := c.parseCommandOutput(output)
	if err != nil {
		logrus.Fatal(err)
	}
	response := OutResponse{Version: common.Version{
		VersionID: parsedOutput.SnapshotId,
	},
	}
	json.NewEncoder(os.Stdout).Encode(response)
}

type OutCommand struct {
	CommandReader common.InputReader
}

func (c *OutCommand) generateResticCommand() (*exec.ResticCommand, error) {
	sourceDir, err := c.CommandReader.GetInputDirectory()
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
	return restic.Put(sourceDir, resticInput)
}

func (c *OutCommand) parseCommandOutput(output []byte) (*BackupResponse, error) {
	pattern := regexp.MustCompile(`\{.*\"message_type\":\"summary\".*\}\n*`)
	matches := pattern.FindAll(output, -1)
	completed := matches[len(matches)-1]
	var parsed BackupResponse
	err := json.Unmarshal(completed, &parsed)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing byte output")
	}
	return &parsed, nil
}

type BackupResponse struct {
	SnapshotId string `json:"snapshot_id"`
}

type OutResponse struct {
	Version  common.Version    `json:"version"`
	Metadata map[string]string `json:"metadata"`
}
