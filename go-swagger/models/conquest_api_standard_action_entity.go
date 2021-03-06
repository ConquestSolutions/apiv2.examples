// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIStandardActionEntity conquest api standard action entity
//
// swagger:model conquest_apiStandardActionEntity
type ConquestAPIStandardActionEntity struct {

	// deleted
	Deleted bool `json:"Deleted,omitempty"`

	// record
	Record *ConquestAPIStandardActionRecord `json:"Record,omitempty"`

	// std action ID
	StdActionID int32 `json:"StdActionID,omitempty"`

	// type ID
	TypeID int32 `json:"TypeID,omitempty"`
}

// Validate validates this conquest api standard action entity
func (m *ConquestAPIStandardActionEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecord(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIStandardActionEntity) validateRecord(formats strfmt.Registry) error {

	if swag.IsZero(m.Record) { // not required
		return nil
	}

	if m.Record != nil {
		if err := m.Record.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Record")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIStandardActionEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIStandardActionEntity) UnmarshalBinary(b []byte) error {
	var res ConquestAPIStandardActionEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
