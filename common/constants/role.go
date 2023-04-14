package constants

type Role int8

const (
	AdminAccount Role = 1
	UserAccount  Role = 2
	TestAccount  Role = 3
)

var roleMap = map[Role]string{
	AdminAccount: "admin",
	UserAccount:  "user",
	TestAccount:  "test",
}

func (r Role) ToString() string {
	return roleMap[r]
}
