package models

type User struct {

	// The user ID
	ID *string `json:"id"`

	// The username
	Username *string `json:"username"`

	// Only admin can create new matches.
	// Usually only the tablet on the table is admin
	Raspberry bool `json:"raspberry"`

	// for people streaming the results
	Visitor bool `json:"visitor"`
}
