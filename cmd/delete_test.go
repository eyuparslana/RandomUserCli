package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_Delete(t *testing.T) {
	deleteCmd := NewDeleteCmd()
	deleteCmd.Flags().Uint64VarP(&userId, "userid", "u", 0, "id of the user to be deleted.")
	deleteCmd.SetArgs([]string{"-u", "15"})
	b := bytes.NewBufferString("")
	deleteCmd.SetOut(b)
	deleteCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	successMessageLayout := "user deleted: %d"
	if string(out) != fmt.Sprintf(successMessageLayout, userId) {
		t.Errorf("expected \"%s\" got \"%s\"", fmt.Sprintf(successMessageLayout, userId), string(out))
	}
}

func Test_DeleteAll(t *testing.T) {
	deleteCmd := NewDeleteCmd()
	deleteCmd.Flags().BoolVarP(&all, "all", "a", false, "deletes all users.")
	deleteCmd.SetArgs([]string{"-a"})
	b := bytes.NewBufferString("")
	deleteCmd.SetOut(b)
	deleteCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	successMessage := "all user data deleted."
	bucketNotFountErr := "bucket not found."
	if !(string(out) != successMessage || bucketNotFountErr != string(out)) {
		t.Errorf("Test failed.")
	}
}
