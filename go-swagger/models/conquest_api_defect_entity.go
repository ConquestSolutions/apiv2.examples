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

// ConquestAPIDefectEntity DefectEntity is comprised of the Defect's Record and auxiliary fields (calculated, related)
//
// swagger:model conquest_apiDefectEntity
type ConquestAPIDefectEntity struct {

	// action ID
	ActionID int32 `json:"ActionID,omitempty"`

	// asset ID
	AssetID int32 `json:"AssetID,omitempty"`

	// AttributeID for the AttributeSet configured on the AssetType
	AttributeID int32 `json:"AttributeID,omitempty"`

	// completion date
	CompletionDate *ConquestAPITimestampValue `json:"CompletionDate,omitempty"`

	// defect ID
	DefectID int32 `json:"DefectID,omitempty"`

	// edit date
	// Format: date-time
	EditDate strfmt.DateTime `json:"EditDate,omitempty"`

	// editor
	Editor string `json:"Editor,omitempty"`

	// inspection ID
	InspectionID int32 `json:"InspectionID,omitempty"`

	// inspection notes
	InspectionNotes string `json:"InspectionNotes,omitempty"`

	// permission
	Permission ConquestAPIPermission `json:"Permission,omitempty"`

	// record
	Record *ConquestAPIDefectRecord `json:"Record,omitempty"`

	// ResponseDate for the Selected Severity
	ResponseDate *ConquestAPITimestampValue `json:"ResponseDate,omitempty"`

	// The Severity Score of the Standard Defect Response of the selected Severity
	SeverityScore int32 `json:"SeverityScore,omitempty"`

	// The Asset Priority (aka. Asset.SysUserHierarchy1)
	SysUserHierarchy1 int32 `json:"SysUserHierarchy1,omitempty"`

	// lock
	Lock *ConquestAPILock `json:"lock,omitempty"`
}

// Validate validates this conquest api defect entity
func (m *ConquestAPIDefectEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEditDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecord(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLock(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIDefectEntity) validateCompletionDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletionDate) { // not required
		return nil
	}

	if m.CompletionDate != nil {
		if err := m.CompletionDate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CompletionDate")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIDefectEntity) validateEditDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EditDate) { // not required
		return nil
	}

	if err := validate.FormatOf("EditDate", "body", "date-time", m.EditDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConquestAPIDefectEntity) validatePermission(formats strfmt.Registry) error {

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

func (m *ConquestAPIDefectEntity) validateRecord(formats strfmt.Registry) error {

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

func (m *ConquestAPIDefectEntity) validateResponseDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ResponseDate) { // not required
		return nil
	}

	if m.ResponseDate != nil {
		if err := m.ResponseDate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResponseDate")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIDefectEntity) validateLock(formats strfmt.Registry) error {

	if swag.IsZero(m.Lock) { // not required
		return nil
	}

	if m.Lock != nil {
		if err := m.Lock.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lock")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIDefectEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIDefectEntity) UnmarshalBinary(b []byte) error {
	var res ConquestAPIDefectEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
