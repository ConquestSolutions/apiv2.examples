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

// ConquestAPIListAttributeSetsQuery AttributeSets can be queried using a number of different objects types.
// The AttributeSet that will be returned is the one that relates to the ObjectKey
//
// swagger:model conquest_apiListAttributeSetsQuery
type ConquestAPIListAttributeSetsQuery struct {

	// attribute set keys
	AttributeSetKeys []*ConquestAPIObjectKey `json:"AttributeSetKeys"`

	// include unset fields
	IncludeUnsetFields bool `json:"IncludeUnsetFields,omitempty"`

	// object type
	ObjectType ConquestAPIObjectType `json:"ObjectType,omitempty"`
}

// Validate validates this conquest api list attribute sets query
func (m *ConquestAPIListAttributeSetsQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeSetKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIListAttributeSetsQuery) validateAttributeSetKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.AttributeSetKeys) { // not required
		return nil
	}

	for i := 0; i < len(m.AttributeSetKeys); i++ {
		if swag.IsZero(m.AttributeSetKeys[i]) { // not required
			continue
		}

		if m.AttributeSetKeys[i] != nil {
			if err := m.AttributeSetKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AttributeSetKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConquestAPIListAttributeSetsQuery) validateObjectType(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectType) { // not required
		return nil
	}

	if err := m.ObjectType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ObjectType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIListAttributeSetsQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIListAttributeSetsQuery) UnmarshalBinary(b []byte) error {
	var res ConquestAPIListAttributeSetsQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}