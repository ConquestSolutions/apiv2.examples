// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIAttributeSetField conquest api attribute set field
//
// swagger:model conquest_apiAttributeSetField
type ConquestAPIAttributeSetField struct {

	// caption
	Caption string `json:"Caption,omitempty"`

	// dimension value
	DimensionValue *ConquestAPIAttributeSetDimensionValue `json:"DimensionValue,omitempty"`

	// FieldName is the Field that appears on the ObjectType, for example, Asset.UserText1
	FieldName string `json:"FieldName,omitempty"`

	// hierarchy value
	HierarchyValue *ConquestAPIAttributeSetHierarchyValue `json:"HierarchyValue,omitempty"`

	// list value
	ListValue *ConquestAPIAttributeSetListValue `json:"ListValue,omitempty"`

	// notes
	Notes string `json:"Notes,omitempty"`

	// order
	Order int32 `json:"Order,omitempty"`

	// required
	Required bool `json:"Required,omitempty"`

	// simple value
	SimpleValue interface{} `json:"SimpleValue,omitempty"`

	// weighted value
	WeightedValue *ConquestAPIAttributeSetWeightedValue `json:"WeightedValue,omitempty"`
}

// Validate validates this conquest api attribute set field
func (m *ConquestAPIAttributeSetField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDimensionValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHierarchyValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateListValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeightedValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIAttributeSetField) validateDimensionValue(formats strfmt.Registry) error {

	if swag.IsZero(m.DimensionValue) { // not required
		return nil
	}

	if m.DimensionValue != nil {
		if err := m.DimensionValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DimensionValue")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAttributeSetField) validateHierarchyValue(formats strfmt.Registry) error {

	if swag.IsZero(m.HierarchyValue) { // not required
		return nil
	}

	if m.HierarchyValue != nil {
		if err := m.HierarchyValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HierarchyValue")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAttributeSetField) validateListValue(formats strfmt.Registry) error {

	if swag.IsZero(m.ListValue) { // not required
		return nil
	}

	if m.ListValue != nil {
		if err := m.ListValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ListValue")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAttributeSetField) validateWeightedValue(formats strfmt.Registry) error {

	if swag.IsZero(m.WeightedValue) { // not required
		return nil
	}

	if m.WeightedValue != nil {
		if err := m.WeightedValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("WeightedValue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIAttributeSetField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIAttributeSetField) UnmarshalBinary(b []byte) error {
	var res ConquestAPIAttributeSetField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}