// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPITypePermission conquest api type permission
//
// swagger:model conquest_apiTypePermission
type ConquestAPITypePermission struct {

	// account ID
	AccountID *ConquestAPIAccountID `json:"AccountID,omitempty"`

	// permission
	Permission ConquestAPIPermission `json:"Permission,omitempty"`

	// type description
	TypeDescription string `json:"TypeDescription,omitempty"`

	// type ID
	TypeID int32 `json:"TypeID,omitempty"`
}

// Validate validates this conquest api type permission
func (m *ConquestAPITypePermission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPITypePermission) validateAccountID(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if m.AccountID != nil {
		if err := m.AccountID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AccountID")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPITypePermission) validatePermission(formats strfmt.Registry) error {

	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	if err := m.Permission.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Permission")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPITypePermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPITypePermission) UnmarshalBinary(b []byte) error {
	var res ConquestAPITypePermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
