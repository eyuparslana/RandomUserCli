package user

func FilterByUserId(userId *uint64, users []*User) []*User {
	result := make([]*User, 0)
	for _, user := range users {
		if *user.UserId == *userId {
			result = append(result, user)
			return result
		}
	}
	return nil
}

func FilterByNationality(nat string, users []*User) []*User {
	result := make([]*User, 0)
	for _, user := range users {
		if user.Nat == nat {
			result = append(result, user)
		}
	}
	return result
}

func FilterByGender(gender string, users []*User) []*User {
	result := make([]*User, 0)
	for _, user := range users {
		if user.Gender == gender {
			result = append(result, user)
		}
	}
	return result
}

func FilterByAge(funcKey string, limit int, users []*User) []*User {
	result := make([]*User, 0)
	f := OperatorMap[funcKey]
	for _, user := range users {
		if f(limit, user.Dob.Age) {
			result = append(result, user)
		}
	}
	return result
}

func IsInMapKeys(key string) bool {
	for k, _ := range OperatorMap {
		if key == k {
			return true
		}
	}
	return false
}

func DeleteByNationality(nat string, users []*User) []*User {
	result := make([]*User, 0)
	for _, user := range users {
		if user.Nat != nat {
			result = append(result, user)
		}
	}
	return result
}

func DeleteByGender(gender string, users []*User) []*User {
	result := make([]*User, 0)
	for _, user := range users {
		if user.Gender != gender {
			result = append(result, user)
		}
	}
	return result
}

func DeleteByAge(funcKey string, limit int, users []*User) []*User {
	result := make([]*User, 0)
	f := OperatorMap[funcKey]
	for _, user := range users {
		if !f(limit, user.Dob.Age) {
			result = append(result, user)
		}
	}
	return result
}
