// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ConquestAPIObjectAttributeParameterType An ObjectAttribute may be parameterised, ObjectAttributeParameterType indicates how a parameter is interpreted
//
// swagger:model conquest_apiObjectAttributeParameterType
type ConquestAPIObjectAttributeParameterType string

const (

	// ConquestAPIObjectAttributeParameterTypeObjectAttributeParameterTypeUnknown captures enum value "ObjectAttributeParameterType_Unknown"
	ConquestAPIObjectAttributeParameterTypeObjectAttributeParameterTypeUnknown ConquestAPIObjectAttributeParameterType = "ObjectAttributeParameterType_Unknown"

	// ConquestAPIObjectAttributeParameterTypeObjectAttributeParameterTypeCodeListID captures enum value "ObjectAttributeParameterType_CodeListID"
	ConquestAPIObjectAttributeParameterTypeObjectAttributeParameterTypeCodeListID ConquestAPIObjectAttributeParameterType = "ObjectAttributeParameterType_CodeListID"

	// ConquestAPIObjectAttributeParameterTypeObjectAttributeParameterTypeCodeListName captures enum value "ObjectAttributeParameterType_CodeListName"
	ConquestAPIObjectAttributeParameterTypeObjectAttributeParameterTypeCodeListName ConquestAPIObjectAttributeParameterType = "ObjectAttributeParameterType_CodeListName"

	// ConquestAPIObjectAttributeParameterTypeObjectAttributeParameterTypeEnumerationValue captures enum value "ObjectAttributeParameterType_EnumerationValue"
	ConquestAPIObjectAttributeParameterTypeObjectAttributeParameterTypeEnumerationValue ConquestAPIObjectAttributeParameterType = "ObjectAttributeParameterType_EnumerationValue"

	// ConquestAPIObjectAttributeParameterTypeObjectAttributeParameterTypeEnumerationDisplayName captures enum value "ObjectAttributeParameterType_EnumerationDisplayName"
	ConquestAPIObjectAttributeParameterTypeObjectAttributeParameterTypeEnumerationDisplayName ConquestAPIObjectAttributeParameterType = "ObjectAttributeParameterType_EnumerationDisplayName"

	// ConquestAPIObjectAttributeParameterTypeObjectAttributeParameterTypeCodeID captures enum value "ObjectAttributeParameterType_CodeID"
	ConquestAPIObjectAttributeParameterTypeObjectAttributeParameterTypeCodeID ConquestAPIObjectAttributeParameterType = "ObjectAttributeParameterType_CodeID"
)

// for schema
var conquestApiObjectAttributeParameterTypeEnum []interface{}

func init() {
	var res []ConquestAPIObjectAttributeParameterType
	if err := json.Unmarshal([]byte(`["ObjectAttributeParameterType_Unknown","ObjectAttributeParameterType_CodeListID","ObjectAttributeParameterType_CodeListName","ObjectAttributeParameterType_EnumerationValue","ObjectAttributeParameterType_EnumerationDisplayName","ObjectAttributeParameterType_CodeID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conquestApiObjectAttributeParameterTypeEnum = append(conquestApiObjectAttributeParameterTypeEnum, v)
	}
}

func (m ConquestAPIObjectAttributeParameterType) validateConquestAPIObjectAttributeParameterTypeEnum(path, location string, value ConquestAPIObjectAttributeParameterType) error {
	if err := validate.Enum(path, location, value, conquestApiObjectAttributeParameterTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this conquest api object attribute parameter type
func (m ConquestAPIObjectAttributeParameterType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConquestAPIObjectAttributeParameterTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
