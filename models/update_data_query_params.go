// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpdateDataQueryParams update data query params
// swagger:model UpdateDataQueryParams

type UpdateDataQueryParams struct {

	// data
	Data []string `json:"data"`
}

/* polymorph UpdateDataQueryParams data false */

// Validate validates this update data query params
func (m *UpdateDataQueryParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateDataQueryParams) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateDataQueryParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateDataQueryParams) UnmarshalBinary(b []byte) error {
	var res UpdateDataQueryParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
