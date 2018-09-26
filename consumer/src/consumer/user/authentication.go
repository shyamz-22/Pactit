package user

import "os"

func Authenticate(username string, password string) bool {
	expectedUsername := os.Getenv("QUOKI_USERNAME")
	expectedPassword := os.Getenv("QUOKI_PASSWORD")
	return username == expectedUsername && password == expectedPassword
}
