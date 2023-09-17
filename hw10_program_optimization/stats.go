package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"

	"github.com/mailru/easyjson"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	scanner := bufio.NewScanner(r)
	result := make(DomainStat)
	var user User
	for scanner.Scan() {
		if err := easyjson.Unmarshal(scanner.Bytes(), &user); err != nil {
			continue
		}
		if strings.HasSuffix(user.Email, domain) {
			lIdx := strings.LastIndex(user.Email, "@")
			result[strings.ToLower(user.Email[lIdx+1:])]++
		}
	}
	return result, nil
}
