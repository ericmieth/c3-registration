package user

import ()

type User struct {
	FirstName   string
	LastName    string
	MailAddress string
	PassHash64  string
	PassHash    []byte
	Salt64      string
	Salt        []byte
}
