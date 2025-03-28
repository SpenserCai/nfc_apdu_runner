// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TlvParseResult tlv parse result
//
// swagger:model TlvParseResult
type TlvParseResult struct {

	// Parsed TLV structure
	Structure []*TlvItem `json:"structure"`
}

// Validate validates this tlv parse result
func (m *TlvParseResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStructure(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TlvParseResult) validateStructure(formats strfmt.Registry) error {
	if swag.IsZero(m.Structure) { // not required
		return nil
	}

	for i := 0; i < len(m.Structure); i++ {
		if swag.IsZero(m.Structure[i]) { // not required
			continue
		}

		if m.Structure[i] != nil {
			if err := m.Structure[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("structure" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("structure" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tlv parse result based on the context it is used
func (m *TlvParseResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStructure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TlvParseResult) contextValidateStructure(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Structure); i++ {

		if m.Structure[i] != nil {

			if swag.IsZero(m.Structure[i]) { // not required
				return nil
			}

			if err := m.Structure[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("structure" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("structure" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TlvParseResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TlvParseResult) UnmarshalBinary(b []byte) error {
	var res TlvParseResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
