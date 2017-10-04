// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SeriesImagesCount series images count
// swagger:model SeriesImagesCount

type SeriesImagesCount struct {

	// fanart
	Fanart int64 `json:"fanart,omitempty"`

	// poster
	Poster int64 `json:"poster,omitempty"`

	// season
	Season int64 `json:"season,omitempty"`

	// seasonwide
	Seasonwide int64 `json:"seasonwide,omitempty"`

	// series
	Series int64 `json:"series,omitempty"`
}

/* polymorph SeriesImagesCount fanart false */

/* polymorph SeriesImagesCount poster false */

/* polymorph SeriesImagesCount season false */

/* polymorph SeriesImagesCount seasonwide false */

/* polymorph SeriesImagesCount series false */

// Validate validates this series images count
func (m *SeriesImagesCount) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SeriesImagesCount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SeriesImagesCount) UnmarshalBinary(b []byte) error {
	var res SeriesImagesCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
