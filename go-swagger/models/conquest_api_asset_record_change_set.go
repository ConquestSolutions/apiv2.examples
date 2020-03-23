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

// ConquestAPIAssetRecordChangeSet conquest api asset record change set
//
// swagger:model conquest_apiAssetRecordChangeSet
type ConquestAPIAssetRecordChangeSet struct {

	// asset ID
	AssetID int32 `json:"AssetID,omitempty"`

	// changes
	Changes []string `json:"Changes"`

	// last edit
	// Format: date-time
	LastEdit strfmt.DateTime `json:"LastEdit,omitempty"`

	// original
	Original *ConquestAPIAssetRecord `json:"Original,omitempty"`

	// updated
	Updated *ConquestAPIAssetRecord `json:"Updated,omitempty"`
}

// Validate validates this conquest api asset record change set
func (m *ConquestAPIAssetRecordChangeSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastEdit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIAssetRecordChangeSet) validateLastEdit(formats strfmt.Registry) error {

	if swag.IsZero(m.LastEdit) { // not required
		return nil
	}

	if err := validate.FormatOf("LastEdit", "body", "date-time", m.LastEdit.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConquestAPIAssetRecordChangeSet) validateOriginal(formats strfmt.Registry) error {

	if swag.IsZero(m.Original) { // not required
		return nil
	}

	if m.Original != nil {
		if err := m.Original.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Original")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIAssetRecordChangeSet) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if m.Updated != nil {
		if err := m.Updated.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Updated")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIAssetRecordChangeSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIAssetRecordChangeSet) UnmarshalBinary(b []byte) error {
	var res ConquestAPIAssetRecordChangeSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}