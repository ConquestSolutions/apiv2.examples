// Code generated by go-swagger; DO NOT EDIT.

package knowledge_base_service

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

// NewListAttributeSetsParams creates a new ListAttributeSetsParams object
// with the default values initialized.
func NewListAttributeSetsParams() *ListAttributeSetsParams {
	var ()
	return &ListAttributeSetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAttributeSetsParamsWithTimeout creates a new ListAttributeSetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAttributeSetsParamsWithTimeout(timeout time.Duration) *ListAttributeSetsParams {
	var ()
	return &ListAttributeSetsParams{

		timeout: timeout,
	}
}

// NewListAttributeSetsParamsWithContext creates a new ListAttributeSetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAttributeSetsParamsWithContext(ctx context.Context) *ListAttributeSetsParams {
	var ()
	return &ListAttributeSetsParams{

		Context: ctx,
	}
}

// NewListAttributeSetsParamsWithHTTPClient creates a new ListAttributeSetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAttributeSetsParamsWithHTTPClient(client *http.Client) *ListAttributeSetsParams {
	var ()
	return &ListAttributeSetsParams{
		HTTPClient: client,
	}
}

/*ListAttributeSetsParams contains all the parameters to send to the API endpoint
for the list attribute sets operation typically these are written to a http.Request
*/
type ListAttributeSetsParams struct {

	/*Body*/
	Body *models.ConquestAPIListAttributeSetsQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list attribute sets params
func (o *ListAttributeSetsParams) WithTimeout(timeout time.Duration) *ListAttributeSetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list attribute sets params
func (o *ListAttributeSetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list attribute sets params
func (o *ListAttributeSetsParams) WithContext(ctx context.Context) *ListAttributeSetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list attribute sets params
func (o *ListAttributeSetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list attribute sets params
func (o *ListAttributeSetsParams) WithHTTPClient(client *http.Client) *ListAttributeSetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list attribute sets params
func (o *ListAttributeSetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list attribute sets params
func (o *ListAttributeSetsParams) WithBody(body *models.ConquestAPIListAttributeSetsQuery) *ListAttributeSetsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list attribute sets params
func (o *ListAttributeSetsParams) SetBody(body *models.ConquestAPIListAttributeSetsQuery) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListAttributeSetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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