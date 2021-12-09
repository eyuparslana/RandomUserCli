package user

import (
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	bolt "go.etcd.io/bbolt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"user-cli/repo"
)

func CreateUsers(reqParams Params, start bool) {
	params := url.Values{}
	params.Add("results", reqParams.Count)
	params.Add("noinfo", "")
	params.Add("inc", "gender,name,email,dob,picture,nat")

	if reqParams.Gender != "" {
		params.Add("gender", reqParams.Gender)
	}

	if reqParams.Nationality != "" {
		params.Add("nat", reqParams.Nationality)
	}

	reqUrl, err := url.Parse(BASEURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	reqUrl.RawQuery = params.Encode()
	resp, err := http.Get(reqUrl.String())
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	db := repo.GetRepo()
	tx, err := db.Begin(true)
	if err != nil {
		fmt.Println(err)
	}

	if start {
		err := tx.DeleteBucket([]byte("users"))
		if err != nil && err != bolt.ErrBucketNotFound {
			tx.Rollback()
			fmt.Println(err)
		}
	}
	tx.Commit()

	for _, user := range response.Results {
		AddUserToDataBase(user)
	}

}

func FilterUser(filter Filter, count bool) string {
	var userId *uint64
	if filter.UserId != 0 {
		userId = &filter.UserId
	}
	result := GetUsers(userId)

	if filter.Nationality != "" {
		result = FilterByNationality(filter.Nationality, result)
	}

	if filter.Gender != "" {
		if !(filter.Gender == "male" || filter.Gender == "female") {
			return "gender param must be male or female."
		}
		result = FilterByGender(filter.Gender, result)
	}

	if filter.Age > 0 {
		if filter.AgeOperator == "" {
			return "If the age field is entered, the operator field must also be entered."
		}

		if !IsInMapKeys(filter.AgeOperator) {
			return "operator value must be one of [e, gt, gte, lt, lte]"
		}
		result = FilterByAge(filter.AgeOperator, filter.Age, result)
	}

	var log string
	if filter.UserId != 0 {
		result = FilterByUserId(&filter.UserId, result)
		if len(result) == 1 {
			log = fmt.Sprintf("User found: %d", filter.UserId)
		} else {
			log = fmt.Sprintf("User not found: %d", filter.UserId)
		}
	}

	if count {
		return fmt.Sprintf("User count: %d", len(result))
	} else {
		ShowUsers(result)
		if log != "" {
			return log
		}
		return fmt.Sprintf("%d users are shown in the table", len(result))
	}
}

func DeleteUser(userId uint64, all bool) (string, error) {
	db := repo.GetRepo()
	tx, err := db.Begin(true)
	if err != nil {
		return "", err
	}

	if all {
		tx.DeleteBucket([]byte("users"))
		tx.Commit()
		return "all user data deleted.", nil
	}

	b := tx.Bucket([]byte("users"))
	if b == nil {
		return "bucket not found.", nil
	}
	err = b.Delete([]byte(strconv.FormatUint(userId, 10)))
	if err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()

	return fmt.Sprintf("user deleted: %d", userId), nil
}

func ShowUsers(userList []*User) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "First Name", "Last Name", "Email", "Age", "Nationality", "Gender", "Picture"})
	for _, user := range userList {
		t.AppendRow(table.Row{*user.UserId, user.Name.First, user.Name.Last, user.Email, user.Dob.Age, user.Nat, user.Gender, user.Picture.Large})
	}
	t.Render()
}
