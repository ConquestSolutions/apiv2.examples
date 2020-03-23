// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIAttributeSetsResult conquest api attribute sets result
//
// swagger:model conquest_apiAttributeSetsResult
type ConquestAPIAttributeSetsResult struct {

	// attribute sets
	AttributeSets []*ConquestAPIAttributeSet `json:"AttributeSets"`
}

// Validate validates this conquest api attribute sets result
func (m *ConquestAPIAttributeSetsResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeSets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIAttributeSetsResult) validateAttributeSets(formats strfmt.Registry) error {

	if swag.IsZero(m.AttributeSets) { // not required
		return nil
	}

	for i := 0; i < len(m.AttributeSets); i++ {
		if swag.IsZero(m.AttributeSets[i]) { // not required
			continue
		}

		if m.AttributeSets[i] != nil {
			if err := m.AttributeSets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AttributeSets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIAttributeSetsResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIAttributeSetsResult) UnmarshalBinary(b []byte) error {
	var res ConquestAPIAttributeSetsResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}