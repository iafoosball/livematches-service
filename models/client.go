package models

type Client interface {
	SetID(s string)
	GetID() string
}

type Table struct {

	// The user ID
	ID string `json:"id"`

	//In match
	Match bool `json:"match"`

	//In lobby
	Lobby bool `json:"lobby"`

	//In free
	Free bool `json:"free"`

	// last change. Could be used for reset time.
	Datetime string
}

func (t *Table) SetID(s string) {
	t.ID = s
}

func (t *Table) GetID() string {
	return t.ID
}

type User struct {

	// The user ID
	ID string `json:"id"`

	// The username
	Username string `json:"username"`

	// The first to join the lobby is the admin.
	Admin bool `json:"admin"`

	// for people streaming the results
	Visitor bool `json:"visitor"`
}

func (u *User) SetID(s string) {
	u.ID = s
}
func (u *User) GetID() string {
	return u.ID
}
