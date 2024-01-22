package util

import "math/rand"

var alphanum []string

func init() {
	generateAlphaNum()
}

func generateAlphaNum() {
	alphanum = []string{}
	for i := 'A'; i <= 'Z'; i++ {
		alphanum = append(alphanum, string(i))
	}
	for i := 'a'; i <= 'z'; i++ {
		alphanum = append(alphanum, string(i))
	}
	for i := '0'; i <= '9'; i++ {
		alphanum = append(alphanum, string(i))
	}
}

func UniqueName(length int) string {
	name := ""
	for i := 0; i < length; i++ {
		name += alphanum[rand.Intn(len(alphanum))]
	}
	return name
}
