package common

import "github.com/chimano/restic-resource/exec"

// ResticStore represents a restic secret store.
type ResticStore struct {
	Repository string
	Tags       []string
	Region     string
	Host       string
}

// ResticConfiguration represents all configuration required to create a
// ResticStore
type ResticConfiguration struct {
	Repository string
	Tags       []string
	Region     string
	Host       string
}

// NewRestic Creates a ResticStore, based on the given ResticConfiguration
func NewRestic(config *ResticConfiguration) *ResticStore {
	return &ResticStore{
		Host:       config.Host,
		Region:     config.Region,
		Repository: config.Repository,
		Tags:       config.Tags,
	}
}

// Put will insert a secret into the restic repository, by crafting the
// relevant restic command.
func (r *ResticStore) Put(secretName, inputDir string) *exec.ResticCommand {
	args := []string{
		"restic",
		"--repo",
		r.Repository,
		"--host",
		r.Host,
		"--option",
		"s3.region=" + r.Region,
		"--verbose",
		"--tag",
		secretName,
	}
	for _, tag := range r.Tags {
		args = append(args, "--tag", tag)
	}
	args = append(
		args,
		"backup",
		inputDir,
	)
	return exec.InitCommand(args)
}
