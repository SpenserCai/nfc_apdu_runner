// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FormatTemplate format template
//
// swagger:model FormatTemplate
type FormatTemplate struct {

	// Description of the format template
	Description string `json:"description,omitempty"`

	// ID of the format template
	ID string `json:"id,omitempty"`

	// Name of the format template
	Name string `json:"name,omitempty"`

	// Path to the format template file
	Path string `json:"path,omitempty"`
}

// Validate validates this format template
func (m *FormatTemplate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this format template based on context it is used
func (m *FormatTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FormatTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FormatTemplate) UnmarshalBinary(b []byte) error {
	var res FormatTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
