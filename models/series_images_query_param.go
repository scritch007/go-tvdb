// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SeriesImagesQueryParam series images query param
// swagger:model SeriesImagesQueryParam

type SeriesImagesQueryParam struct {

	// key type
	KeyType string `json:"keyType,omitempty"`

	// language Id
	LanguageID string `json:"languageId,omitempty"`

	// resolution
	Resolution []string `json:"resolution"`

	// sub key
	SubKey []string `json:"subKey"`
}

/* polymorph SeriesImagesQueryParam keyType false */

/* polymorph SeriesImagesQueryParam languageId false */

/* polymorph SeriesImagesQueryParam resolution false */

/* polymorph SeriesImagesQueryParam subKey false */

// Validate validates this series images query param
func (m *SeriesImagesQueryParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResolution(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubKey(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SeriesImagesQueryParam) validateResolution(formats strfmt.Registry) error {

	if swag.IsZero(m.Resolution) { // not required
		return nil
	}

	return nil
}

func (m *SeriesImagesQueryParam) validateSubKey(formats strfmt.Registry) error {

	if swag.IsZero(m.SubKey) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SeriesImagesQueryParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SeriesImagesQueryParam) UnmarshalBinary(b []byte) error {
	var res SeriesImagesQueryParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
