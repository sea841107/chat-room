package database

const (
	UserTable = "user"
)

type User struct {
	Name     string
	Password string
	Friend   []string
	Hint     []string
}

type Message struct {
	Name string
	Text string
	Time int64
}
