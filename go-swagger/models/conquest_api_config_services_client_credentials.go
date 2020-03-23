// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConquestAPIConfigServicesClientCredentials ServicesCredentials are the used to connect the Conquest API to services.
// These are required to host the API. It is used for tasks like
// - Authentication
// - Matching Mobile users with self-hosted APIs
// - Retrieving user and company preferences
// - Submitting service statistics and error logs
//
// These can be obtained by either of the following methods:
// - Registering your instance of the Conquest API in the Services Portal and adding
//   the credentials manually here
// - Using the [Link] task in the Conquest Management Console
//
// swagger:model conquest_api_configServicesClientCredentials
type ConquestAPIConfigServicesClientCredentials struct {

	// This is a "Username" for your instance of the Conquest API
	ClientID string `json:"client_id,omitempty"`

	// This is a password for your instance of the Conquest API
	ClientSecret string `json:"client_secret,omitempty"`
}

// Validate validates this conquest api config services client credentials
func (m *ConquestAPIConfigServicesClientCredentials) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIConfigServicesClientCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIConfigServicesClientCredentials) UnmarshalBinary(b []byte) error {
	var res ConquestAPIConfigServicesClientCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}