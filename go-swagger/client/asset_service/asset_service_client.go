// Code generated by go-swagger; DO NOT EDIT.

package asset_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new asset service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for asset service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ChangeAssetType(params *ChangeAssetTypeParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeAssetTypeOK, error)

	CreateActionForAsset(params *CreateActionForAssetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateActionForAssetOK, error)

	CreateAsset(params *CreateAssetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateAssetOK, error)

	CreateDefectForAsset(params *CreateDefectForAssetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDefectForAssetOK, error)

	DeleteAsset(params *DeleteAssetParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAssetOK, error)

	GetAsset(params *GetAssetParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssetOK, error)

	ListInspectionsForAsset(params *ListInspectionsForAssetParams, authInfo runtime.ClientAuthInfoWriter) (*ListInspectionsForAssetOK, error)

	MoveAsset(params *MoveAssetParams, authInfo runtime.ClientAuthInfoWriter) (*MoveAssetOK, error)

	NewConditionInspection(params *NewConditionInspectionParams, authInfo runtime.ClientAuthInfoWriter) (*NewConditionInspectionOK, error)

	UpdateAsset(params *UpdateAssetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAssetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ChangeAssetType change asset type API
*/
func (a *Client) ChangeAssetType(params *ChangeAssetTypeParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeAssetTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeAssetTypeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeAssetType",
		Method:             "POST",
		PathPattern:        "/api/assets/change_asset_type",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangeAssetTypeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeAssetTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ChangeAssetType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateActionForAsset create action for asset API
*/
func (a *Client) CreateActionForAsset(params *CreateActionForAssetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateActionForAssetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateActionForAssetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateActionForAsset",
		Method:             "POST",
		PathPattern:        "/api/assets/create_action",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateActionForAssetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateActionForAssetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateActionForAsset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateAsset create asset API
*/
func (a *Client) CreateAsset(params *CreateAssetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateAssetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAssetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateAsset",
		Method:             "POST",
		PathPattern:        "/api/assets/create_asset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAssetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateAssetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateAsset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateDefectForAsset create defect for asset API
*/
func (a *Client) CreateDefectForAsset(params *CreateDefectForAssetParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDefectForAssetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDefectForAssetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDefectForAsset",
		Method:             "POST",
		PathPattern:        "/api/assets/create_defect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDefectForAssetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDefectForAssetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateDefectForAsset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAsset delete asset API
*/
func (a *Client) DeleteAsset(params *DeleteAssetParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAssetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAssetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAsset",
		Method:             "POST",
		PathPattern:        "/api/assets/delete_asset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAssetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAssetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAsset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAsset get asset API
*/
func (a *Client) GetAsset(params *GetAssetParams, authInfo runtime.ClientAuthInfoWriter) (*GetAssetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAssetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAsset",
		Method:             "GET",
		PathPattern:        "/api/assets/{value}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAssetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAssetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAsset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListInspectionsForAsset list inspections for asset API
*/
func (a *Client) ListInspectionsForAsset(params *ListInspectionsForAssetParams, authInfo runtime.ClientAuthInfoWriter) (*ListInspectionsForAssetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListInspectionsForAssetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListInspectionsForAsset",
		Method:             "POST",
		PathPattern:        "/api/assets/list_inspections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListInspectionsForAssetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListInspectionsForAssetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListInspectionsForAsset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MoveAsset move asset API
*/
func (a *Client) MoveAsset(params *MoveAssetParams, authInfo runtime.ClientAuthInfoWriter) (*MoveAssetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMoveAssetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "MoveAsset",
		Method:             "POST",
		PathPattern:        "/api/assets/move_asset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MoveAssetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MoveAssetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for MoveAsset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  NewConditionInspection new condition inspection API
*/
func (a *Client) NewConditionInspection(params *NewConditionInspectionParams, authInfo runtime.ClientAuthInfoWriter) (*NewConditionInspectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNewConditionInspectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NewConditionInspection",
		Method:             "POST",
		PathPattern:        "/api/assets/new_condition_inspection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &NewConditionInspectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NewConditionInspectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for NewConditionInspection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateAsset update asset API
*/
func (a *Client) UpdateAsset(params *UpdateAssetParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAssetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAssetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateAsset",
		Method:             "POST",
		PathPattern:        "/api/assets/update_asset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAssetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateAssetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateAsset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
