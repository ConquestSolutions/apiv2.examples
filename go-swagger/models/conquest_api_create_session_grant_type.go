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

// ConquestAPICreateSessionGrantType  - authorization_code: code is used for exchanging an authorization code for an access token
// when an application is configured for the 'authorization code flow'
//  - refresh_token: a refresh_token is obtained when a session has been created.
// It is used for obtaining a new access token when a session has expired.
//
// A refresh token may be provided when a session is created with the 'offline' scope.
//  - password: DO NOT USE THIS. This flow has changed. Instead of calling the API, obtain an authorization code from services.
//
// password is a **custom** flow where a USER provides their own credentials directly
// to the client application the application itself generates an access token.
//
// For when an application is configured for the 'resource owner flow'.
//  - pass_code: pass_code is a **custom** flow for conquest III
// so that users can open the workplanner
//
// this flow must be used with the "/redirect/" endpoint
//
// swagger:model conquest_apiCreateSessionGrantType
type ConquestAPICreateSessionGrantType string

const (

	// ConquestAPICreateSessionGrantTypeUnknown captures enum value "unknown"
	ConquestAPICreateSessionGrantTypeUnknown ConquestAPICreateSessionGrantType = "unknown"

	// ConquestAPICreateSessionGrantTypeAuthorizationCode captures enum value "authorization_code"
	ConquestAPICreateSessionGrantTypeAuthorizationCode ConquestAPICreateSessionGrantType = "authorization_code"

	// ConquestAPICreateSessionGrantTypeRefreshToken captures enum value "refresh_token"
	ConquestAPICreateSessionGrantTypeRefreshToken ConquestAPICreateSessionGrantType = "refresh_token"

	// ConquestAPICreateSessionGrantTypePassword captures enum value "password"
	ConquestAPICreateSessionGrantTypePassword ConquestAPICreateSessionGrantType = "password"

	// ConquestAPICreateSessionGrantTypePassCode captures enum value "pass_code"
	ConquestAPICreateSessionGrantTypePassCode ConquestAPICreateSessionGrantType = "pass_code"
)

// for schema
var conquestApiCreateSessionGrantTypeEnum []interface{}

func init() {
	var res []ConquestAPICreateSessionGrantType
	if err := json.Unmarshal([]byte(`["unknown","authorization_code","refresh_token","password","pass_code"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conquestApiCreateSessionGrantTypeEnum = append(conquestApiCreateSessionGrantTypeEnum, v)
	}
}

func (m ConquestAPICreateSessionGrantType) validateConquestAPICreateSessionGrantTypeEnum(path, location string, value ConquestAPICreateSessionGrantType) error {
	if err := validate.Enum(path, location, value, conquestApiCreateSessionGrantTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this conquest api create session grant type
func (m ConquestAPICreateSessionGrantType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConquestAPICreateSessionGrantTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
