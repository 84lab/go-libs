package email

import "regexp"

var EmailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func IsValid(email string) bool {
	return EmailRegex.MatchString(email)
}
