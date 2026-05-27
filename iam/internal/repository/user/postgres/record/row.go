package record

type UserRow struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Image    string `db:"image"`
}

const (
	UsersTable               string = "users"
	UsersTableColumnID       string = "id"
	UsersTableColumnUsername string = "username"
	UsersTableColumnPassword string = "password"
	UsersTableColumnImage    string = "image"
)

var UsersTableColumns = []string{
	UsersTableColumnID,
	UsersTableColumnUsername,
	UsersTableColumnPassword,
	UsersTableColumnImage,
}
