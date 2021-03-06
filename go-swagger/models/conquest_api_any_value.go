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

// ConquestAPIAnyValue Each field corresponds to exactly ValueType
// For example:
//   stringValue and nStringValue (nullable) is for ValueType_String
//
// swagger:model conquest_apiAnyValue
type ConquestAPIAnyValue struct {

	// boolean value
	BooleanValue bool `json:"booleanValue,omitempty"`

	// code value
	CodeValue int32 `json:"codeValue,omitempty"`

	// date time value
	// Format: date-time
	DateTimeValue strfmt.DateTime `json:"dateTimeValue,omitempty"`

	// date value
	// Format: date-time
	DateValue strfmt.DateTime `json:"dateValue,omitempty"`

	// decimal value
	DecimalValue ConquestAPIDecimal `json:"decimalValue,omitempty"`

	// duration value
	DurationValue *ConquestAPIDuration `json:"durationValue,omitempty"`

	// enumeration value
	EnumerationValue int32 `json:"enumerationValue,omitempty"`

	// geography coordinates value
	GeographyCoordinatesValue *ConquestAPIGeographyCoordinates `json:"geographyCoordinatesValue,omitempty"`

	// hierarchy value
	HierarchyValue *ConquestAPIObjectKey `json:"hierarchyValue,omitempty"`

	// int32 value
	Int32Value int32 `json:"int32Value,omitempty"`

	// int64 value
	Int64Value string `json:"int64Value,omitempty"`

	// n boolean value
	NBooleanValue string `json:"nBooleanValue,omitempty"`

	// n code value
	NCodeValue int32 `json:"nCodeValue,omitempty"`

	// n date time value
	NDateTimeValue *ConquestAPITimestampValue `json:"nDateTimeValue,omitempty"`

	// n date value
	NDateValue *ConquestAPITimestampValue `json:"nDateValue,omitempty"`

	// n decimal value
	NDecimalValue *ConquestAPIDecimalValue `json:"nDecimalValue,omitempty"`

	// n duration value
	NDurationValue *ConquestAPIDurationValue `json:"nDurationValue,omitempty"`

	// n enumeration value
	NEnumerationValue int32 `json:"nEnumerationValue,omitempty"`

	// n geography coordinates value
	NGeographyCoordinatesValue *ConquestAPIGeographyCoordinatesValue `json:"nGeographyCoordinatesValue,omitempty"`

	// n hierarchy value
	NHierarchyValue *ConquestAPIObjectKeyValue `json:"nHierarchyValue,omitempty"`

	// n int32 value
	NInt32Value int32 `json:"nInt32Value,omitempty"`

	// n int64 value
	NInt64Value string `json:"nInt64Value,omitempty"`

	// n number value
	NNumberValue float64 `json:"nNumberValue,omitempty"`

	// n object key value
	NObjectKeyValue *ConquestAPIObjectKeyValue `json:"nObjectKeyValue,omitempty"`

	// n string value
	NStringValue string `json:"nStringValue,omitempty"`

	// number value
	NumberValue float64 `json:"numberValue,omitempty"`

	// object key value
	ObjectKeyValue *ConquestAPIObjectKey `json:"objectKeyValue,omitempty"`

	// string value
	StringValue string `json:"stringValue,omitempty"`

	// unknown value
	UnknownValue bool `json:"unknownValue,omitempty"`
}

// Validate validates this conquest api any value
func (m *ConquestAPIAnyValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateTimeValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDecimalValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDurationValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeographyCoordinatesValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHierarchyValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNDateTimeValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNDateValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNDecimalValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNDurationValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNGeographyCoordinatesValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNHierarchyValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNObjectKeyValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectKeyValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIAnyValue) validateDateTimeValue(formats strfmt.Registry) error {

	if swag.IsZero(m.DateTimeValue) { // not required
		return nil
	}

	if err := validate.FormatOf("dateTimeValue", "body", "date-time", m.DateTimeValue.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConquestAPIAnyValue) validateDateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.DateValue) { // not required
		return nil
	}

	if err := validate.FormatOf("dateValue", "body", "date-time", m.DateValue.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConquestAPIAnyValue) validateDecimalValue(formats strfmt.Registry) error {

	if swag.IsZero(m.DecimalValue) { // not required
		return nil
	}

	if err := m.DecimalValue.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("decimalValue")
		}
		return err
	}

	return nil
}

func (m *ConquestAPIAnyValue) validateDurationValue(formats strfmt.Registry) error {

	if swag.IsZero(m.DurationValue) { // not required
		return nil
	}

	if m.DurationValue != nil {
		if err := m.DurationValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("durationValue")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAnyValue) validateGeographyCoordinatesValue(formats strfmt.Registry) error {

	if swag.IsZero(m.GeographyCoordinatesValue) { // not required
		return nil
	}

	if m.GeographyCoordinatesValue != nil {
		if err := m.GeographyCoordinatesValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geographyCoordinatesValue")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAnyValue) validateHierarchyValue(formats strfmt.Registry) error {

	if swag.IsZero(m.HierarchyValue) { // not required
		return nil
	}

	if m.HierarchyValue != nil {
		if err := m.HierarchyValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hierarchyValue")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAnyValue) validateNDateTimeValue(formats strfmt.Registry) error {

	if swag.IsZero(m.NDateTimeValue) { // not required
		return nil
	}

	if m.NDateTimeValue != nil {
		if err := m.NDateTimeValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nDateTimeValue")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAnyValue) validateNDateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.NDateValue) { // not required
		return nil
	}

	if m.NDateValue != nil {
		if err := m.NDateValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nDateValue")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAnyValue) validateNDecimalValue(formats strfmt.Registry) error {

	if swag.IsZero(m.NDecimalValue) { // not required
		return nil
	}

	if m.NDecimalValue != nil {
		if err := m.NDecimalValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nDecimalValue")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAnyValue) validateNDurationValue(formats strfmt.Registry) error {

	if swag.IsZero(m.NDurationValue) { // not required
		return nil
	}

	if m.NDurationValue != nil {
		if err := m.NDurationValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nDurationValue")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAnyValue) validateNGeographyCoordinatesValue(formats strfmt.Registry) error {

	if swag.IsZero(m.NGeographyCoordinatesValue) { // not required
		return nil
	}

	if m.NGeographyCoordinatesValue != nil {
		if err := m.NGeographyCoordinatesValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nGeographyCoordinatesValue")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAnyValue) validateNHierarchyValue(formats strfmt.Registry) error {

	if swag.IsZero(m.NHierarchyValue) { // not required
		return nil
	}

	if m.NHierarchyValue != nil {
		if err := m.NHierarchyValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nHierarchyValue")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAnyValue) validateNObjectKeyValue(formats strfmt.Registry) error {

	if swag.IsZero(m.NObjectKeyValue) { // not required
		return nil
	}

	if m.NObjectKeyValue != nil {
		if err := m.NObjectKeyValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nObjectKeyValue")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAnyValue) validateObjectKeyValue(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectKeyValue) { // not required
		return nil
	}

	if m.ObjectKeyValue != nil {
		if err := m.ObjectKeyValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objectKeyValue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIAnyValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIAnyValue) UnmarshalBinary(b []byte) error {
	var res ConquestAPIAnyValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
