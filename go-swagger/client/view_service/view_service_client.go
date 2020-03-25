// Code generated by go-swagger; DO NOT EDIT.

package view_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new view service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for view service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	Find(params *FindParams, authInfo runtime.ClientAuthInfoWriter) (*FindOK, error)

	ListFilters(params *ListFiltersParams, authInfo runtime.ClientAuthInfoWriter) (*ListFiltersOK, error)

	ListHierarchyNodes(params *ListHierarchyNodesParams, authInfo runtime.ClientAuthInfoWriter) (*ListHierarchyNodesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Find find API
*/
func (a *Client) Find(params *FindParams, authInfo runtime.ClientAuthInfoWriter) (*FindOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Find",
		Method:             "POST",
		PathPattern:        "/api/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FindReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Find: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListFilters list filters API
*/
func (a *Client) ListFilters(params *ListFiltersParams, authInfo runtime.ClientAuthInfoWriter) (*ListFiltersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListFiltersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListFilters",
		Method:             "POST",
		PathPattern:        "/api/list_filters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListFiltersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListFiltersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListFilters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListHierarchyNodes list hierarchy nodes API
*/
func (a *Client) ListHierarchyNodes(params *ListHierarchyNodesParams, authInfo runtime.ClientAuthInfoWriter) (*ListHierarchyNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListHierarchyNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListHierarchyNodes",
		Method:             "POST",
		PathPattern:        "/api/list_hierarchy_nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListHierarchyNodesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListHierarchyNodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListHierarchyNodes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
