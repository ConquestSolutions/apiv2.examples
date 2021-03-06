// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConquestAPIObjectKey conquest api object key
//
// swagger:model conquest_apiObjectKey
type ConquestAPIObjectKey struct {

	// int32 value
	Int32Value int32 `json:"int32Value,omitempty"`

	// object type
	ObjectType ConquestAPIObjectType `json:"objectType,omitempty"`

	// string value
	StringValue string `json:"stringValue,omitempty"`

	// timestamp value
	// Format: date-time
	TimestampValue *strfmt.DateTime `json:"timestampValue,omitempty"`
}

// Validate validates this conquest api object key
func (m *ConquestAPIObjectKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestampValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIObjectKey) validateObjectType(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectType) { // not required
		return nil
	}

	if err := m.ObjectType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("objectType")
		}
		return err
	}

	return nil
}

func (m *ConquestAPIObjectKey) validateTimestampValue(formats strfmt.Registry) error {

	if m.TimestampValue == nil || swag.IsZero(m.TimestampValue) { // not required
		return nil
	}

	if err := validate.FormatOf("timestampValue", "body", "date-time", m.TimestampValue.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIObjectKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIObjectKey) UnmarshalBinary(b []byte) error {
	var res ConquestAPIObjectKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
