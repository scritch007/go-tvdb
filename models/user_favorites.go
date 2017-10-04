// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UserFavorites user favorites
// swagger:model UserFavorites

type UserFavorites struct {

	// favorites
	Favorites []string `json:"favorites"`
}

/* polymorph UserFavorites favorites false */

// Validate validates this user favorites
func (m *UserFavorites) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFavorites(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserFavorites) validateFavorites(formats strfmt.Registry) error {

	if swag.IsZero(m.Favorites) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserFavorites) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserFavorites) UnmarshalBinary(b []byte) error {
	var res UserFavorites
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}