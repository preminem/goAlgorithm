package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetAwardUserName(users map[string]int64) (name string) {
	sizeOfUsers := len(users)
	awardIndex := rand.Intn(sizeOfUsers)

	var index int
	for uname, _ := range users {
		if index == awardIndex {
			name = uname
			return
		}
		index += 1
	}
	return
}

func GetAwardUserNameByWeights(users map[string]int64) (name string) {
	type User struct {
		Name   string
		Offset int64
		Num    int64
	}

	userArr := make([]*User, 0)
	var sum int64
	for name, num := range users {
		user := &User{
			Name:   name,
			Offset: sum,
			Num:    num,
		}
		userArr = append(userArr, user)
		sum += num
	}

	awardNum := rand.Int63n(sum)

	for index, _ := range userArr {
		user := userArr[index]
		if user.Offset+user.Num > awardNum {
			name = user.Name
			return
		}
	}
	return
}

func main() {
	var users map[string]int64 = map[string]int64{
		"Alice": 10,
		"Bob":   5,
		"Cindy": 15,
		"Dale":  20,
		"Eric":  10,
		"Frank": 30,
	}

	rand.Seed(time.Now().Unix())
	award_stat := make(map[string]int64)
	for i := 0; i < 10000; i += 1 {
		//name := GetAwardUserName(users)
		name := GetAwardUserNameByWeights(users)
		if count, ok := award_stat[name]; ok {
			award_stat[name] = count + 1
		} else {
			award_stat[name] = 1
		}
	}

	for name, count := range award_stat {
		fmt.Printf("user: %s, award count: %d\n", name, count)
	}

	return
}
