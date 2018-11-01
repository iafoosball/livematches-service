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
