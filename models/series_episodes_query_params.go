// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SeriesEpisodesQueryParams series episodes query params
// swagger:model SeriesEpisodesQueryParams

type SeriesEpisodesQueryParams struct {

	// data
	Data []string `json:"data"`
}

/* polymorph SeriesEpisodesQueryParams data false */

// Validate validates this series episodes query params
func (m *SeriesEpisodesQueryParams) Validate(formats strfmt.Registry) error {
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

func (m *SeriesEpisodesQueryParams) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SeriesEpisodesQueryParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SeriesEpisodesQueryParams) UnmarshalBinary(b []byte) error {
	var res SeriesEpisodesQueryParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}