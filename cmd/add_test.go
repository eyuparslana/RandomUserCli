package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_Add(t *testing.T) {
	addCmd := NewAddCmd()
	count := "50"
	nat := "TR"
	gender := "male"
	addCmd.Flags().StringVarP(&AddParams.Count, "count", "c", "500", "the count of the users to be added.")
	addCmd.Flags().StringVarP(&AddParams.Gender, "gender", "g", "", "the gender of the user to be added")
	addCmd.Flags().StringVarP(&AddParams.Nationality, "nat", "n", "", "the nationality of the user to be added")
	addCmd.SetArgs([]string{"-c", count, "-n", nat, "-g", gender})
	b := bytes.NewBufferString("")
	addCmd.SetOut(b)
	addCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	successMessageLayout := "%s %s %s user(s) added to db."
	if string(out) != fmt.Sprintf(successMessageLayout, count, nat, gender) {
		t.Errorf("expected \"%s\" got \"%s\"", fmt.Sprintf(successMessageLayout, count, nat, gender), string(out))
	}
}
