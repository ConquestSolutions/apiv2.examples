// Code generated by go-swagger; DO NOT EDIT.

package system_service

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
)

// NewApplicationVersionParams creates a new ApplicationVersionParams object
// with the default values initialized.
func NewApplicationVersionParams() *ApplicationVersionParams {

	return &ApplicationVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationVersionParamsWithTimeout creates a new ApplicationVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewApplicationVersionParamsWithTimeout(timeout time.Duration) *ApplicationVersionParams {

	return &ApplicationVersionParams{

		timeout: timeout,
	}
}

// NewApplicationVersionParamsWithContext creates a new ApplicationVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewApplicationVersionParamsWithContext(ctx context.Context) *ApplicationVersionParams {

	return &ApplicationVersionParams{

		Context: ctx,
	}
}

// NewApplicationVersionParamsWithHTTPClient creates a new ApplicationVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewApplicationVersionParamsWithHTTPClient(client *http.Client) *ApplicationVersionParams {

	return &ApplicationVersionParams{
		HTTPClient: client,
	}
}

/*ApplicationVersionParams contains all the parameters to send to the API endpoint
for the application version operation typically these are written to a http.Request
*/
type ApplicationVersionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the application version params
func (o *ApplicationVersionParams) WithTimeout(timeout time.Duration) *ApplicationVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application version params
func (o *ApplicationVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application version params
func (o *ApplicationVersionParams) WithContext(ctx context.Context) *ApplicationVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application version params
func (o *ApplicationVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application version params
func (o *ApplicationVersionParams) WithHTTPClient(client *http.Client) *ApplicationVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application version params
func (o *ApplicationVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
