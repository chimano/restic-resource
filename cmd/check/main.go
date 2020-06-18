package main

// SORTED IN ASCENDING ORDER BY DATE
import (
	"encoding/json"
	"os"

	"github.com/arekmano/restic-store/exec"
	"github.com/arekmano/restic-store/store"
	"github.com/chimano/restic-resource/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	c := &CheckCommand{commandReader: &common.Reader{}}

	command, err := c.generateResticCommand()
	if err != nil {
		logrus.Fatal(err)
	}

	output, err := command.Execute()
	if err != nil {
		logrus.Fatal(err)
	}
	request, err := c.commandReader.GetRequest()
	if err != nil {
		logrus.Fatal(err)
	}
	versionID := request.Version

	allVersions, err := c.parseCommandOutput(output)
	if err != nil {
		logrus.Fatal(err)
	}
	filteredVersions, err := c.keepNewerVersions(allVersions, versionID)
	if err != nil {
		logrus.Fatal(err)
	}
	json.NewEncoder(os.Stdout).Encode(filteredVersions)
}

type CheckCommand struct {
	commandReader common.InputReader
}

func (c *CheckCommand) generateResticCommand() (*exec.ResticCommand, error) {

	request, err := c.commandReader.GetRequest()
	if err != nil {
		return nil, err
	}

	resticConfig := &store.ResticConfiguration{Host: request.Source.Host, Repository: request.Source.Repository}
	restic := store.NewRestic(resticConfig)
	resticInput := &store.ResticOptions{Options: request.Source.Options, Tags: request.Source.Tags}
	return restic.ListSnapshots(resticInput), nil
}

func (c *CheckCommand) parseCommandOutput(output []byte) ([]common.Version, error) {

	var parsed []Snapshot
	err := json.Unmarshal(output, &parsed)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing byte output")
	}
	versions := make([]common.Version, len(parsed))
	for i, snapshot := range parsed {
		versions[i].VersionID = snapshot.SnapshotId
	}
	return versions, nil
}

func (c *CheckCommand) keepNewerVersions(versions []common.Version, base common.Version) ([]common.Version, error) {
	foundIndex := -1
	for i, version := range versions {
		if version.VersionID == base.VersionID {
			foundIndex = i
		}
	}
	if foundIndex == -1 {
		return nil, errors.New("Could not find version amongst list of snapshots")
	}
	return versions[foundIndex:], nil
}

type Snapshot struct {
	SnapshotId string `json:"short_id"`
}
