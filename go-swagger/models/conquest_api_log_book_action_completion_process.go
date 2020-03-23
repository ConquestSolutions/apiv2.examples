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

// ConquestAPILogBookActionCompletionProcess conquest api log book action completion process
//
// swagger:model conquest_apiLogBookActionCompletionProcess
type ConquestAPILogBookActionCompletionProcess struct {

	// next log action due date
	// Format: date-time
	NextLogActionDueDate strfmt.DateTime `json:"NextLogActionDueDate,omitempty"`

	// next log meter reading
	NextLogMeterReading float64 `json:"NextLogMeterReading,omitempty"`

	// next log std action ID
	NextLogStdActionID int32 `json:"NextLogStdActionID,omitempty"`
}

// Validate validates this conquest api log book action completion process
func (m *ConquestAPILogBookActionCompletionProcess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNextLogActionDueDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPILogBookActionCompletionProcess) validateNextLogActionDueDate(formats strfmt.Registry) error {

	if swag.IsZero(m.NextLogActionDueDate) { // not required
		return nil
	}

	if err := validate.FormatOf("NextLogActionDueDate", "body", "date-time", m.NextLogActionDueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPILogBookActionCompletionProcess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPILogBookActionCompletionProcess) UnmarshalBinary(b []byte) error {
	var res ConquestAPILogBookActionCompletionProcess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}