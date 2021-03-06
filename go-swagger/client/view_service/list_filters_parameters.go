// Code generated by go-swagger; DO NOT EDIT.

package view_service

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

// NewListFiltersParams creates a new ListFiltersParams object
// with the default values initialized.
func NewListFiltersParams() *ListFiltersParams {
	var ()
	return &ListFiltersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListFiltersParamsWithTimeout creates a new ListFiltersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListFiltersParamsWithTimeout(timeout time.Duration) *ListFiltersParams {
	var ()
	return &ListFiltersParams{

		timeout: timeout,
	}
}

// NewListFiltersParamsWithContext creates a new ListFiltersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListFiltersParamsWithContext(ctx context.Context) *ListFiltersParams {
	var ()
	return &ListFiltersParams{

		Context: ctx,
	}
}

// NewListFiltersParamsWithHTTPClient creates a new ListFiltersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListFiltersParamsWithHTTPClient(client *http.Client) *ListFiltersParams {
	var ()
	return &ListFiltersParams{
		HTTPClient: client,
	}
}

/*ListFiltersParams contains all the parameters to send to the API endpoint
for the list filters operation typically these are written to a http.Request
*/
type ListFiltersParams struct {

	/*Body*/
	Body *models.ConquestAPIListFiltersQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list filters params
func (o *ListFiltersParams) WithTimeout(timeout time.Duration) *ListFiltersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list filters params
func (o *ListFiltersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list filters params
func (o *ListFiltersParams) WithContext(ctx context.Context) *ListFiltersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list filters params
func (o *ListFiltersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list filters params
func (o *ListFiltersParams) WithHTTPClient(client *http.Client) *ListFiltersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list filters params
func (o *ListFiltersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list filters params
func (o *ListFiltersParams) WithBody(body *models.ConquestAPIListFiltersQuery) *ListFiltersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list filters params
func (o *ListFiltersParams) SetBody(body *models.ConquestAPIListFiltersQuery) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListFiltersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
