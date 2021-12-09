package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
	"user-cli/user"
)

func Test_CountValidGender(t *testing.T) {
	CountFilter = user.Filter{}
	countCmd := NewCountCmd()
	countCmd.Flags().StringVarP(&CountFilter.Gender, "gender", "g", "", "filter user by gender.")
	countCmd.SetArgs([]string{"-g", "male"})
	b := bytes.NewBufferString("")
	countCmd.SetOut(b)
	countCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	successMessage := "user count: "
	if strings.HasPrefix(string(out), successMessage) {
		t.Errorf("Test Failed: %s", string(out))
	}
}

func Test_CountInValidGender(t *testing.T) {
	CountFilter = user.Filter{}
	countCmd := NewCountCmd()
	countCmd.Flags().StringVarP(&CountFilter.Gender, "gender", "g", "", "filter user by gender.")
	countCmd.SetArgs([]string{"-g", "other"})
	b := bytes.NewBufferString("")
	countCmd.SetOut(b)
	countCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	successMessage := "gender param must be male or female."
	if string(out) != successMessage {
		t.Errorf("Test Failed: %s", string(out))
	}
}

func Test_CountAge(t *testing.T) {
	CountFilter = user.Filter{}
	countCmd := NewCountCmd()
	countCmd.Flags().IntVarP(&CountFilter.Age, "age", "a", 0, "filter user by age.")
	countCmd.Flags().StringVarP(&CountFilter.AgeOperator, "operator", "o", "", "rules of the age filter. (gt, lt, gte, lte, e)")
	countCmd.SetArgs([]string{"-a", "50", "-o", "gt"})
	b := bytes.NewBufferString("")
	countCmd.SetOut(b)
	countCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	successMessage := "user count: "
	if strings.HasPrefix(string(out), successMessage) {
		t.Errorf("Test Failed: %s", string(out))
	}
}

func Test_CountAgeWithoutOperator(t *testing.T) {
	CountFilter = user.Filter{}
	countCmd := NewCountCmd()
	countCmd.Flags().IntVarP(&CountFilter.Age, "age", "a", 0, "filter user by age.")
	countCmd.Flags().StringVarP(&CountFilter.AgeOperator, "operator", "o", "", "rules of the age filter. (gt, lt, gte, lte, e)")
	countCmd.SetArgs([]string{"-a", "50"})
	b := bytes.NewBufferString("")
	countCmd.SetOut(b)
	countCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	successMessage := "If the age field is entered, the operator field must also be entered."
	if string(out) != successMessage {
		t.Errorf("Test Failed: %s", string(out))
	}
}

func Test_CountAgeWithOperator(t *testing.T) {
	CountFilter = user.Filter{}
	countCmd := NewCountCmd()
	countCmd.Flags().IntVarP(&CountFilter.Age, "age", "a", 0, "filter user by age.")
	countCmd.Flags().StringVarP(&CountFilter.AgeOperator, "operator", "o", "", "rules of the age filter. (gt, lt, gte, lte, e)")
	countCmd.SetArgs([]string{"-a", "50"})
	b := bytes.NewBufferString("")
	countCmd.SetOut(b)
	countCmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	successMessage := "If the age field is entered, the operator field must also be entered."
	if string(out) != successMessage {
		t.Errorf("Test Failed: %s", string(out))
	}
}
