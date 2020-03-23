// Code generated by go-swagger; DO NOT EDIT.

package action_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new action service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for action service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ApplyStandardAction(params *ApplyStandardActionParams) (*ApplyStandardActionOK, error)

	CompleteAction(params *CompleteActionParams) (*CompleteActionOK, error)

	CreateSucceedingAction(params *CreateSucceedingActionParams) (*CreateSucceedingActionOK, error)

	DeleteAction(params *DeleteActionParams) (*DeleteActionOK, error)

	GetAction(params *GetActionParams) (*GetActionOK, error)

	GetActionCompletionDetails(params *GetActionCompletionDetailsParams) (*GetActionCompletionDetailsOK, error)

	MoveAction(params *MoveActionParams) (*MoveActionOK, error)

	ReverseActionCompletion(params *ReverseActionCompletionParams) (*ReverseActionCompletionOK, error)

	UpdateAction(params *UpdateActionParams) (*UpdateActionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ApplyStandardAction apply standard action API
*/
func (a *Client) ApplyStandardAction(params *ApplyStandardActionParams) (*ApplyStandardActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplyStandardActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApplyStandardAction",
		Method:             "POST",
		PathPattern:        "/api/actions/apply_standard_action",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ApplyStandardActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplyStandardActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ApplyStandardAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CompleteAction complete action API
*/
func (a *Client) CompleteAction(params *CompleteActionParams) (*CompleteActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CompleteAction",
		Method:             "POST",
		PathPattern:        "/api/actions/complete_action",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CompleteActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CompleteActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CompleteAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateSucceedingAction create succeeding action API
*/
func (a *Client) CreateSucceedingAction(params *CreateSucceedingActionParams) (*CreateSucceedingActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSucceedingActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateSucceedingAction",
		Method:             "POST",
		PathPattern:        "/api/actions/create_succeeding_action",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSucceedingActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSucceedingActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateSucceedingAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAction delete action API
*/
func (a *Client) DeleteAction(params *DeleteActionParams) (*DeleteActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAction",
		Method:             "POST",
		PathPattern:        "/api/actions/delete_action",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAction get action API
*/
func (a *Client) GetAction(params *GetActionParams) (*GetActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAction",
		Method:             "GET",
		PathPattern:        "/api/actions/{value}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetActionCompletionDetails get action completion details API
*/
func (a *Client) GetActionCompletionDetails(params *GetActionCompletionDetailsParams) (*GetActionCompletionDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActionCompletionDetailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetActionCompletionDetails",
		Method:             "POST",
		PathPattern:        "/api/actions/completion_details",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActionCompletionDetailsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetActionCompletionDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetActionCompletionDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MoveAction move action API
*/
func (a *Client) MoveAction(params *MoveActionParams) (*MoveActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMoveActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "MoveAction",
		Method:             "POST",
		PathPattern:        "/api/actions/move_action",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MoveActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MoveActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for MoveAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReverseActionCompletion reverse action completion API
*/
func (a *Client) ReverseActionCompletion(params *ReverseActionCompletionParams) (*ReverseActionCompletionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReverseActionCompletionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ReverseActionCompletion",
		Method:             "POST",
		PathPattern:        "/api/actions/reverse_action_completion",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReverseActionCompletionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReverseActionCompletionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ReverseActionCompletion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateAction update action API
*/
func (a *Client) UpdateAction(params *UpdateActionParams) (*UpdateActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateAction",
		Method:             "POST",
		PathPattern:        "/api/actions/update_action",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateAction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}