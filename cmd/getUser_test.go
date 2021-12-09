package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_GetUsers(t *testing.T) {
	getUserCmd := NewGetUserCmd()
	b := bytes.NewBufferString("")
	getUserCmd.SetOut(b)
	getUserCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	successMessageSuffix := "users are shown in the table"
	if !strings.HasSuffix(string(out), successMessageSuffix) {
		t.Errorf("Test Failed: %s", string(out))
	}
}

func Test_GetUsersWithUserId(t *testing.T) {
	getUserCmd := NewGetUserCmd()
	getUserCmd.Flags().Uint64VarP(&UserFilter.UserId, "userid", "u", 0, "the userID to be retrieved.")
	userId := "15"
	getUserCmd.SetArgs([]string{"-u", userId})
	b := bytes.NewBufferString("")
	getUserCmd.SetOut(b)
	getUserCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	if string(out) != fmt.Sprintf("User found: %s", userId) {
		t.Errorf("Test Failed: %s", string(out))
	}
}

func Test_GetUsersWithWrongUserId(t *testing.T) {
	getUserCmd := NewGetUserCmd()
	getUserCmd.Flags().Uint64VarP(&UserFilter.UserId, "userid", "u", 0, "the userID to be retrieved.")
	userId := "999999"
	getUserCmd.SetArgs([]string{"-u", userId})
	b := bytes.NewBufferString("")
	getUserCmd.SetOut(b)
	getUserCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	if string(out) != fmt.Sprintf("User not found: %s", userId) {
		t.Errorf("Test Failed: %s", string(out))
	}
}
