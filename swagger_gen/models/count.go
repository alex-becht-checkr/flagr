// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Count count
//
// swagger:model count
type Count struct {

	// total flags
	// Required: true
	TotalFlags *int64 `json:"totalFlags"`
}

// Validate validates this count
func (m *Count) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTotalFlags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Count) validateTotalFlags(formats strfmt.Registry) error {

	if err := validate.Required("totalFlags", "body", m.TotalFlags); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this count based on context it is used
func (m *Count) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Count) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Count) UnmarshalBinary(b []byte) error {
	var res Count
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
