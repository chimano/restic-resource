package exec

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/sirupsen/logrus"
)

// ResticCommand represents the restic command that can be executed
type ResticCommand struct {
	BinaryPath  string
	Arguments   []string
	Environment []string
}

// InitCommand will create a new ResticCommand.
// This function will panic if restic is not installed and found in the PATH
func InitCommand(Arguments []string) *ResticCommand {
	command := ResticCommand{
		Environment: os.Environ(),
		Arguments:   Arguments,
	}
	binary, lookErr := exec.LookPath("restic")
	if lookErr != nil {
		panic(lookErr)
	}
	command.BinaryPath = binary
	return &command
}

// Execute will execute the ResticCommand. After this command executes,
// restic-secret-store's process will be replaced with restic's.
// (i.e. no more restic-secret-store code will execute)
func (r *ResticCommand) Execute() {
	logrus.
		WithField("Arguments", r.Arguments).
		Info("Executing Command")
	execErr := syscall.Exec(r.BinaryPath, r.Arguments, r.Environment)
	if execErr != nil {
		panic(execErr)
	}
}

// Print logs the command using the logger. Useful for debugging.
func (r *ResticCommand) Print() {
	logrus.
		WithField("BinaryPath", r.BinaryPath).
		WithField("Arguments", r.Arguments).
		WithField("Environment", r.Environment).
		Info("Restic Command Printed")
}
