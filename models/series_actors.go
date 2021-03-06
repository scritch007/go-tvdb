// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SeriesActors series actors
// swagger:model SeriesActors

type SeriesActors struct {

	// data
	Data SeriesActorsData `json:"data"`

	// errors
	Errors *JSONErrors `json:"errors,omitempty"`
}

/* polymorph SeriesActors data false */

/* polymorph SeriesActors errors false */

// Validate validates this series actors
func (m *SeriesActors) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SeriesActors) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	if m.Errors != nil {

		if err := m.Errors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errors")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SeriesActors) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SeriesActors) UnmarshalBinary(b []byte) error {
	var res SeriesActors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
