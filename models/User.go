package models

type AuthorityLevel int

const (
	AUTHORITY_USER = iota
	AUTHORITY_ADMIN
)

type User struct {
	Id          int
	Name        string
	AuthMethods []*AuthMethod
	Authority   AuthorityLevel
}

// Checks if the user is authorized to access the requested protection level
// The higher the value of the user's Authority the more access to the system he has
// The function checks if the user has higher or equal the provided AuthorityLevel.
func (u User) CheckAuthority(lvl AuthorityLevel) bool {
	return u.Authority >= lvl
}

func (u User) IsAdmin() bool {
	return u.Authority >= AUTHORITY_ADMIN
}
