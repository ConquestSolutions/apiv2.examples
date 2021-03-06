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

// ConquestAPICodeList tblCodeField
//
// swagger:model conquest_apiCodeList
type ConquestAPICodeList struct {

	// CodeListID aka. CodeFieldID
	CodeListID int32 `json:"CodeListID,omitempty"`

	// CodeListName is set for System lists (aka. CodeField)
	CodeListName string `json:"CodeListName,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// items
	Items []*ConquestAPICodeListItem `json:"Items"`

	// type
	Type ConquestAPICodeListType `json:"Type,omitempty"`
}

// Validate validates this conquest api code list
func (m *ConquestAPICodeList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPICodeList) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConquestAPICodeList) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPICodeList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPICodeList) UnmarshalBinary(b []byte) error {
	var res ConquestAPICodeList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
