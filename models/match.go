// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Match match
// swagger:model Match
type Match struct {

	// The match id which is the collection + "/" + the key
	ID string `json:"_id,omitempty"`

	// The match key
	Key string `json:"_key,omitempty"`

	// This is the user id of the first player playing on the blue side.
	BlueUserIDOne string `json:"blue_user_id_one,omitempty"`

	// This is the user id of the second player playing on the blue side. (Not used in single)
	BlueUserIDTwo string `json:"blue_user_id_two,omitempty"`

	// If the match is a double, the position can be attack. (Not used at the moment)
	PositionAttack bool `json:"position_attack,omitempty"`

	// A match can be rated, ie a ranked game with points, or without.
	RatedMatch bool `json:"rated_match,omitempty"`

	// This is the user id of the first player playing on the red side.
	RedUserIDOne string `json:"red_user_id_one,omitempty"`

	// This is the user id of the second player playing on the red side. (Not used in single)
	RedUserIDTwo string `json:"red_user_id_two,omitempty"`
}

// Validate validates this match
func (m *Match) Validate(formats strfmt.Registry) error {
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
