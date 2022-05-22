package chat

const (
	ServerId = 0
	Non      = -1
)

type User struct {
	// 持久性
	uid      string
	password string
	name     string

	// 易失性

}
