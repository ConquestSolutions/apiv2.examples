// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPICreateSessionResponse conquest api create session response
//
// swagger:model conquest_apiCreateSessionResponse
type ConquestAPICreateSessionResponse struct {

	// error
	Error *ConquestAPIErrorResponse `json:"Error,omitempty"`

	// token response
	TokenResponse *ConquestAPIAuthTokenResponse `json:"TokenResponse,omitempty"`
}

// Validate validates this conquest api create session response
func (m *ConquestAPICreateSessionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPICreateSessionResponse) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Error")
			}
			return err
		}
	}

	return nil
}

func (m *ConquestAPICreateSessionResponse) validateTokenResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.TokenResponse) { // not required
		return nil
	}

	if m.TokenResponse != nil {
		if err := m.TokenResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TokenResponse")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPICreateSessionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPICreateSessionResponse) UnmarshalBinary(b []byte) error {
	var res ConquestAPICreateSessionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
