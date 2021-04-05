package os

import "testing"

func TestRunBashCommand(t *testing.T) {
	err := RunBashCommand("ls")
	if err != nil {
		t.Error("RunBashCommand error must be nil")
	}

	err = RunBashCommand("asasa")
	if err == nil {
		t.Error("RunBashCommand error must not be nil")
	}
}
