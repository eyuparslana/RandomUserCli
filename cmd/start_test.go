package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
	"user-cli/user"
)

func Test_Start(t *testing.T) {
	cmd := NewStartCmd()
	count := "50"
	cmd.Flags().StringVarP(&Params.Count, "count", "c", "500", "the count of the users to be created.")
	cmd.Flags().StringVarP(&Params.Gender, "gender", "g", "", "the gender of the user to be created")
	cmd.Flags().StringVarP(&Params.Nationality, "nat", "n", "", "the nationality of the user to be created")
	cmd.SetArgs([]string{"-c", count})
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	if string(out) != fmt.Sprintf("%s user(s) added to db.", count) {
		t.Errorf("expected \"%s\" got \"%s\"", fmt.Sprintf("%s user(s) added to db.", count), string(out))
	}
	users := user.GetUsers(nil)
	if fmt.Sprint(len(users)) != count {
		t.Errorf("expected user count %s got %d, ", count, len(users))
	}
}
