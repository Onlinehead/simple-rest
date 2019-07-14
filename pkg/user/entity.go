package user

type User struct {
	tableName struct{} `sql:"users"`
	Username string `sql:"username,pk"`
	Birthday int64 `sql:"birthday"`
}
