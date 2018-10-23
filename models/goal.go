// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Goal goal
// swagger:model Goal
type Goal struct {

	// DO NOT USE! Only use match id
	From string `json:"_from,omitempty"`

	// The goal id, which is the collection + "/" + the key
	ID string `json:"_id,omitempty"`

	// The goal key
	Key string `json:"_key,omitempty"`

	// DO NOT USE! Only use match id
	To string `json:"_to,omitempty"`

	// The datetime in nanoseconds of the goal.
	Datetime string `json:"datetime,omitempty"`

	// the corresponding match id
	MatchID string `json:"match_id,omitempty"`

	// This could be used in a double game, if the person attacking scores.
	PositionAttack bool `json:"position_attack,omitempty"`

	// The side who scored the goal. Usually either red or blue.
	Side string `json:"side,omitempty"`

	// The speed of the goal
	Speed string `json:"speed,omitempty"`
}

// Validate validates this goal
func (m *Goal) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Goal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Goal) UnmarshalBinary(b []byte) error {
	var res Goal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
