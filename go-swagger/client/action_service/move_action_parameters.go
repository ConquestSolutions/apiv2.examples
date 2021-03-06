// Code generated by go-swagger; DO NOT EDIT.

package action_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
)

// NewMoveActionParams creates a new MoveActionParams object
// with the default values initialized.
func NewMoveActionParams() *MoveActionParams {
	var ()
	return &MoveActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMoveActionParamsWithTimeout creates a new MoveActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMoveActionParamsWithTimeout(timeout time.Duration) *MoveActionParams {
	var ()
	return &MoveActionParams{

		timeout: timeout,
	}
}

// NewMoveActionParamsWithContext creates a new MoveActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewMoveActionParamsWithContext(ctx context.Context) *MoveActionParams {
	var ()
	return &MoveActionParams{

		Context: ctx,
	}
}

// NewMoveActionParamsWithHTTPClient creates a new MoveActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMoveActionParamsWithHTTPClient(client *http.Client) *MoveActionParams {
	var ()
	return &MoveActionParams{
		HTTPClient: client,
	}
}

/*MoveActionParams contains all the parameters to send to the API endpoint
for the move action operation typically these are written to a http.Request
*/
type MoveActionParams struct {

	/*Body*/
	Body *models.ConquestAPIMoveActionCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the move action params
func (o *MoveActionParams) WithTimeout(timeout time.Duration) *MoveActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the move action params
func (o *MoveActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the move action params
func (o *MoveActionParams) WithContext(ctx context.Context) *MoveActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the move action params
func (o *MoveActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the move action params
func (o *MoveActionParams) WithHTTPClient(client *http.Client) *MoveActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the move action params
func (o *MoveActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the move action params
func (o *MoveActionParams) WithBody(body *models.ConquestAPIMoveActionCommand) *MoveActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the move action params
func (o *MoveActionParams) SetBody(body *models.ConquestAPIMoveActionCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *MoveActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
