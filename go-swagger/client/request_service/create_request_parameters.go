// Code generated by go-swagger; DO NOT EDIT.

package request_service

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

// NewCreateRequestParams creates a new CreateRequestParams object
// with the default values initialized.
func NewCreateRequestParams() *CreateRequestParams {
	var ()
	return &CreateRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRequestParamsWithTimeout creates a new CreateRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRequestParamsWithTimeout(timeout time.Duration) *CreateRequestParams {
	var ()
	return &CreateRequestParams{

		timeout: timeout,
	}
}

// NewCreateRequestParamsWithContext creates a new CreateRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRequestParamsWithContext(ctx context.Context) *CreateRequestParams {
	var ()
	return &CreateRequestParams{

		Context: ctx,
	}
}

// NewCreateRequestParamsWithHTTPClient creates a new CreateRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRequestParamsWithHTTPClient(client *http.Client) *CreateRequestParams {
	var ()
	return &CreateRequestParams{
		HTTPClient: client,
	}
}

/*CreateRequestParams contains all the parameters to send to the API endpoint
for the create request operation typically these are written to a http.Request
*/
type CreateRequestParams struct {

	/*Body*/
	Body *models.ConquestAPICreateRequestCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create request params
func (o *CreateRequestParams) WithTimeout(timeout time.Duration) *CreateRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create request params
func (o *CreateRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create request params
func (o *CreateRequestParams) WithContext(ctx context.Context) *CreateRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create request params
func (o *CreateRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create request params
func (o *CreateRequestParams) WithHTTPClient(client *http.Client) *CreateRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create request params
func (o *CreateRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create request params
func (o *CreateRequestParams) WithBody(body *models.ConquestAPICreateRequestCommand) *CreateRequestParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create request params
func (o *CreateRequestParams) SetBody(body *models.ConquestAPICreateRequestCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
