package common

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

type InputReader interface {
	GetInputDirectory() (string, error)
	GetRequest() (*Request, error)
}

type Reader struct{}

func (r *Reader) GetInputDirectory() (string, error) {
	if len(os.Args) <= 1 {
		return "", errors.New("Directory is not being passed as a command line argument")
	}
	return os.Args[1], nil
}

func (r *Reader) GetRequest() (*Request, error) {
	var request *Request

	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "Error parsing json request from stdin")
	}

	return request, nil
}
