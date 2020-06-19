package main

import (
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/chimano/restic-resource/common"
	"github.com/stretchr/testify/assert"
)

type MockReader struct {
}

func (r *MockReader) GetInputDirectory() (string, error) {
	return os.Getenv("TEST_INDIR"), nil
}

var mockRequest = &common.Request{
	Source: common.Source{
		Repository: os.Getenv("TEST_REPO"),
		Host:       os.Getenv("TEST_HOST"),
	},
	Version: common.Version{VersionID: strconv.FormatInt(rand.Int63(), 36)},
}

func (r *MockReader) GetRequest() (*common.Request, error) {
	return mockRequest, nil
}

func TestIn(t *testing.T) {
	c := InCommand{CommandReader: &MockReader{}}
	command, _ := c.generateResticCommand()
	output, _ := command.Execute()
	parsed, _ := c.parseCommandOutput(output)
	assert.NotEmpty(t, parsed)
}

func TestInCommand(t *testing.T) {
	c := InCommand{CommandReader: &MockReader{}}
	command, _ := c.generateResticCommand()
	expectedArgs := []string{
		"--repo",
		os.Getenv("TEST_REPO"),
		"--host",
		os.Getenv("TEST_HOST"),
		"--verbose",
		"--json",
		"restore",
		mockRequest.Version.VersionID,
		"--target",
		os.Getenv("TEST_INDIR"),
	}
	assert.Equal(t, expectedArgs, command.Arguments)
}
