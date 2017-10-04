// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Auth auth
// swagger:model Auth

type Auth struct {

	// apikey
	Apikey string `json:"apikey,omitempty"`

	// userkey
	Userkey string `json:"userkey,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

/* polymorph Auth apikey false */

/* polymorph Auth userkey false */

/* polymorph Auth username false */

// Validate validates this auth
func (m *Auth) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Auth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Auth) UnmarshalBinary(b []byte) error {
	var res Auth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
