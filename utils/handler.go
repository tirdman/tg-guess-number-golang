package utils

import (
	"errors"
	"fmt"
	"github.com/tirdman/tg-guess-number-golang/models"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateNum(length int) string {

	var unknownNumber string
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {

		unknownNumber += strconv.Itoa(rand.Intn(9))
	}

	fmt.Println(unknownNumber)
	return unknownNumber
}

func IsNumber(text string) bool {

	if _, err := strconv.Atoi(text); err == nil {
		return true
	}

	return false
}

func CheckInputNumber(text string, unknownNumber string) string {

	var answer string
	for i := 0; i < len(text); i++ {
		if text[i] == unknownNumber[i] {
			answer += "B"
		} else if strings.Contains(unknownNumber, string(text[i])) {
			answer += "K"
		} else {
			answer += "-"
		}

	}

	return answer
}

func isContainInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetAllUserInCurrentQuest(usersAttempts []int) []int {
	var uniqUser []int
	for _, v := range usersAttempts {

		if !isContainInt(uniqUser, v) {
			uniqUser = append(uniqUser, v)
		}

	}

	return uniqUser
}

func GetUser(id int, users []*models.User) (user *models.User, err error) {
	for _, u := range users {
		if u.Id == id {
			return u, nil
		}
	}
	return nil, errors.New("User no founded")
}
