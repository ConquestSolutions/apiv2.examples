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

// NewReverseActionCompletionParams creates a new ReverseActionCompletionParams object
// with the default values initialized.
func NewReverseActionCompletionParams() *ReverseActionCompletionParams {
	var ()
	return &ReverseActionCompletionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReverseActionCompletionParamsWithTimeout creates a new ReverseActionCompletionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReverseActionCompletionParamsWithTimeout(timeout time.Duration) *ReverseActionCompletionParams {
	var ()
	return &ReverseActionCompletionParams{

		timeout: timeout,
	}
}

// NewReverseActionCompletionParamsWithContext creates a new ReverseActionCompletionParams object
// with the default values initialized, and the ability to set a context for a request
func NewReverseActionCompletionParamsWithContext(ctx context.Context) *ReverseActionCompletionParams {
	var ()
	return &ReverseActionCompletionParams{

		Context: ctx,
	}
}

// NewReverseActionCompletionParamsWithHTTPClient creates a new ReverseActionCompletionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReverseActionCompletionParamsWithHTTPClient(client *http.Client) *ReverseActionCompletionParams {
	var ()
	return &ReverseActionCompletionParams{
		HTTPClient: client,
	}
}

/*ReverseActionCompletionParams contains all the parameters to send to the API endpoint
for the reverse action completion operation typically these are written to a http.Request
*/
type ReverseActionCompletionParams struct {

	/*Body*/
	Body *models.ConquestAPIReverseActionCompletionCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reverse action completion params
func (o *ReverseActionCompletionParams) WithTimeout(timeout time.Duration) *ReverseActionCompletionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reverse action completion params
func (o *ReverseActionCompletionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reverse action completion params
func (o *ReverseActionCompletionParams) WithContext(ctx context.Context) *ReverseActionCompletionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reverse action completion params
func (o *ReverseActionCompletionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reverse action completion params
func (o *ReverseActionCompletionParams) WithHTTPClient(client *http.Client) *ReverseActionCompletionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reverse action completion params
func (o *ReverseActionCompletionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the reverse action completion params
func (o *ReverseActionCompletionParams) WithBody(body *models.ConquestAPIReverseActionCompletionCommand) *ReverseActionCompletionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the reverse action completion params
func (o *ReverseActionCompletionParams) SetBody(body *models.ConquestAPIReverseActionCompletionCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ReverseActionCompletionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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