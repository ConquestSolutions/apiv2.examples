// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIActionCompletionDetails conquest api action completion details
//
// swagger:model conquest_apiActionCompletionDetails
type ConquestAPIActionCompletionDetails struct {

	// action
	Action *ConquestAPIObjectHeader `json:"Action,omitempty"`

	// action type
	ActionType ConquestAPIActionType `json:"ActionType,omitempty"`

	// actual cost
	ActualCost ConquestAPIDecimal `json:"ActualCost,omitempty"`

	// asset
	Asset *ConquestAPIObjectHeader `json:"Asset,omitempty"`

	// estimated cost
	EstimatedCost ConquestAPIDecimal `json:"EstimatedCost,omitempty"`
}

// Validate validates this conquest api action completion details
func (m *ConquestAPIActionCompletionDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActualCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAsset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEstimatedCost(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIActionCompletionDetails) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Action")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) validateActionType(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionType) { // not required
		return nil
	}

	if err := m.ActionType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ActionType")
		}
		return err
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) validateActualCost(formats strfmt.Registry) error {

	if swag.IsZero(m.ActualCost) { // not required
		return nil
	}

	if err := m.ActualCost.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ActualCost")
		}
		return err
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) validateAsset(formats strfmt.Registry) error {

	if swag.IsZero(m.Asset) { // not required
		return nil
	}

	if m.Asset != nil {
		if err := m.Asset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Asset")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPIActionCompletionDetails) validateEstimatedCost(formats strfmt.Registry) error {

	if swag.IsZero(m.EstimatedCost) { // not required
		return nil
	}

	if err := m.EstimatedCost.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("EstimatedCost")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIActionCompletionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIActionCompletionDetails) UnmarshalBinary(b []byte) error {
	var res ConquestAPIActionCompletionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
