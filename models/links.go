// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Links links
// swagger:model Links

type Links struct {

	// first
	First int64 `json:"first,omitempty"`

	// last
	Last int64 `json:"last,omitempty"`

	// next
	Next int64 `json:"next,omitempty"`

	// previous
	Previous int64 `json:"previous,omitempty"`
}

/* polymorph Links first false */

/* polymorph Links last false */

/* polymorph Links next false */

/* polymorph Links previous false */

// Validate validates this links
func (m *Links) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Links) UnmarshalBinary(b []byte) error {
	var res Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
