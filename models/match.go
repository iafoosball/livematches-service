// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Match match
// swagger:model Match
type Match struct {

	// The match id which is the collection + "/" + the key
	ID string `json:"_id,omitempty"`

	// The match key
	Key string `json:"_key,omitempty"`

	// Was the game completed.
	Completed bool `json:"completed,omitempty"`

	// the datetime when the match ends
	EndTime string `json:"endTime,omitempty"`

	// positions
	Positions *MatchPositions `json:"positions,omitempty"`

	// score blue
	ScoreBlue int64 `json:"scoreBlue,omitempty"`

	// score red
	ScoreRed int64 `json:"scoreRed,omitempty"`

	// settings
	Settings *MatchSettings `json:"settings,omitempty"`

	// the datetime when the game ends
	StartTime string `json:"startTime,omitempty"`

	// the id of table
	TableID string `json:"tableID,omitempty"`

	// users
	Users []*MatchUsersItems0 `json:"users"`

	// Can be either "red" or "blue"
	Winner string `json:"winner,omitempty"`
}

// Validate validates this match
func (m *Match) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePositions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Match) validatePositions(formats strfmt.Registry) error {

	if swag.IsZero(m.Positions) { // not required
		return nil
	}

	if m.Positions != nil {
		if err := m.Positions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("positions")
			}
			return err
		}
	}

	return nil
}

func (m *Match) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

func (m *Match) validateUsers(formats strfmt.Registry) error {

	if swag.IsZero(m.Users) { // not required
		return nil
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Match) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Match) UnmarshalBinary(b []byte) error {
	var res Match
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MatchPositions match positions
// swagger:model MatchPositions
type MatchPositions struct {

	// THe UID.
	BlueAttack string `json:"blueAttack,omitempty"`

	// THe UID.
	BlueDefense string `json:"blueDefense,omitempty"`

	// THe UID.
	RedAttack string `json:"redAttack,omitempty"`

	// THe UID.
	RedDefense string `json:"redDefense,omitempty"`
}

// Validate validates this match positions
func (m *MatchPositions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MatchPositions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MatchPositions) UnmarshalBinary(b []byte) error {
	var res MatchPositions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MatchSettings match settings
// swagger:model MatchSettings
type MatchSettings struct {

	// Is this game with bets
	Bet bool `json:"bet,omitempty"`

	// drunk
	Drunk bool `json:"drunk,omitempty"`

	// free game
	FreeGame bool `json:"freeGame,omitempty"`

	// The maximum number of goals for this game. If a time is specified the first criteria which is true will stop the match.
	MaxGoals int64 `json:"maxGoals,omitempty"`

	// The maximum tim in sec for this game. If a max goals is specified the first criteria which is true will stop the match.
	MaxTime int64 `json:"maxTime,omitempty"`

	// one on one
	OneOnOne bool `json:"oneOnOne,omitempty"`

	// payed
	Payed bool `json:"payed,omitempty"`

	// A match can be rated, ie a ranked game with points, or without.
	Rated bool `json:"rated,omitempty"`

	// start match
	StartMatch bool `json:"startMatch,omitempty"`

	// Switch the position after 50% of the goal (time or goals) is reached.
	SwitchPositions bool `json:"switchPositions,omitempty"`

	// tournament
	Tournament bool `json:"tournament,omitempty"`

	// two on one
	TwoOnOne bool `json:"twoOnOne,omitempty"`

	// two on two
	TwoOnTwo bool `json:"twoOnTwo,omitempty"`
}

// Validate validates this match settings
func (m *MatchSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MatchSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MatchSettings) UnmarshalBinary(b []byte) error {
	var res MatchSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MatchUsersItems0 match users items0
// swagger:model MatchUsersItems0
type MatchUsersItems0 struct {

	// admin
	Admin bool `json:"admin,omitempty"`

	// is the amount a user wants to bet on this game
	Bet int64 `json:"bet,omitempty"`

	// color
	Color string `json:"color,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// Can either be attack or defense
	Position string `json:"position,omitempty"`

	// ready
	Ready bool `json:"ready,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this match users items0
func (m *MatchUsersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MatchUsersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MatchUsersItems0) UnmarshalBinary(b []byte) error {
	var res MatchUsersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
