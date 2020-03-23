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

// NewListViewFieldsParams creates a new ListViewFieldsParams object
// with the default values initialized.
func NewListViewFieldsParams() *ListViewFieldsParams {
	var ()
	return &ListViewFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListViewFieldsParamsWithTimeout creates a new ListViewFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListViewFieldsParamsWithTimeout(timeout time.Duration) *ListViewFieldsParams {
	var ()
	return &ListViewFieldsParams{

		timeout: timeout,
	}
}

// NewListViewFieldsParamsWithContext creates a new ListViewFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListViewFieldsParamsWithContext(ctx context.Context) *ListViewFieldsParams {
	var ()
	return &ListViewFieldsParams{

		Context: ctx,
	}
}

// NewListViewFieldsParamsWithHTTPClient creates a new ListViewFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListViewFieldsParamsWithHTTPClient(client *http.Client) *ListViewFieldsParams {
	var ()
	return &ListViewFieldsParams{
		HTTPClient: client,
	}
}

/*ListViewFieldsParams contains all the parameters to send to the API endpoint
for the list view fields operation typically these are written to a http.Request
*/
type ListViewFieldsParams struct {

	/*Body*/
	Body *models.ConquestAPIListViewFieldsQuery

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list view fields params
func (o *ListViewFieldsParams) WithTimeout(timeout time.Duration) *ListViewFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list view fields params
func (o *ListViewFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list view fields params
func (o *ListViewFieldsParams) WithContext(ctx context.Context) *ListViewFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list view fields params
func (o *ListViewFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list view fields params
func (o *ListViewFieldsParams) WithHTTPClient(client *http.Client) *ListViewFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list view fields params
func (o *ListViewFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list view fields params
func (o *ListViewFieldsParams) WithBody(body *models.ConquestAPIListViewFieldsQuery) *ListViewFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list view fields params
func (o *ListViewFieldsParams) SetBody(body *models.ConquestAPIListViewFieldsQuery) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListViewFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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