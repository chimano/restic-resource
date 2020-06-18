package main

import (
	"os"
	"testing"

	"github.com/chimano/restic-resource/common"
	"github.com/stretchr/testify/assert"
)

type MockReader struct {
}

func (r *MockReader) GetInputDirectory() (string, error) {
	return os.Getenv("TEST_INDIR"), nil
}

func (r *MockReader) GetRequest() (*common.Request, error) {
	return &common.Request{
		Source: common.Source{
			Repository: os.Getenv("TEST_REPO"),
			Host:       os.Getenv("TEST_HOST"),
		},
	}, nil
}

func TestIn(t *testing.T) {
	c := InCommand{CommandReader: &MockReader{}}
	command, _ := c.generateResticCommand()
	output, _ := command.Execute()
	parsed, _ := c.parseCommandOutput(output)
	assert.NotEmpty(t, parsed)
}
