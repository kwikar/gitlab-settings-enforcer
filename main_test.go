package main

import (
	"github.com/libri-gmbh/gitlab-settings-enforcer/cmd"
	"github.com/stretchr/testify/require"
	"os"
	"os/exec"
	"testing"
)

// TestCrasher tests the command from the outside with a special config.json
// modelled after https://talks.golang.org/2014/testing.slide#23
func TestCrasher(t *testing.T) {

	// This variable is set by this test function so it calls itself
	// to test the actual execution
	if os.Getenv("BE_CRASHER") == "1" {
		cmd.Execute()
		return
	}

	const configJSON = `{
  "group_name": "some/group",
  "project_whitelist": ["some/group/some-project"],
  "include_subgroups": false,
  "create_default_branch": true,
  "protected_branches": [
    {
      "name": "master",
      "push_access_level": "noone",
      "merge_access_level": "developer"
    }
  ],
  "project_settings": {
    "default_branch": "master"
  }
}`

	cfgFile := "config.json"
	require.NoError(t, os.WriteFile(cfgFile, ([]byte)(configJSON), 0666))
	defer func() {
		if err := os.Remove(cfgFile); err != nil {
			t.Logf("failed to delete config: %v", err)
		}
	}()

	command := exec.Command(os.Args[0], "-test.run=TestCrasher", "sync")
	command.Env = append(os.Environ(), "BE_CRASHER=1")
	out, err := command.CombinedOutput()
	if err != nil {
		t.Log(string(out))
		t.Errorf("run failed: %v", err)
	}
}
