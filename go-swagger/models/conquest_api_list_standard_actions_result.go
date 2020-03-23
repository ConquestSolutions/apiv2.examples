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

// ConquestAPIListStandardActionsResult conquest api list standard actions result
//
// swagger:model conquest_apiListStandardActionsResult
type ConquestAPIListStandardActionsResult struct {

	// standard actions
	StandardActions []*ConquestAPIStandardActionEntity `json:"StandardActions"`
}

// Validate validates this conquest api list standard actions result
func (m *ConquestAPIListStandardActionsResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStandardActions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIListStandardActionsResult) validateStandardActions(formats strfmt.Registry) error {

	if swag.IsZero(m.StandardActions) { // not required
		return nil
	}

	for i := 0; i < len(m.StandardActions); i++ {
		if swag.IsZero(m.StandardActions[i]) { // not required
			continue
		}

		if m.StandardActions[i] != nil {
			if err := m.StandardActions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StandardActions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIListStandardActionsResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIListStandardActionsResult) UnmarshalBinary(b []byte) error {
	var res ConquestAPIListStandardActionsResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}