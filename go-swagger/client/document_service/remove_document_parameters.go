// Code generated by go-swagger; DO NOT EDIT.

package document_service

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

// NewRemoveDocumentParams creates a new RemoveDocumentParams object
// with the default values initialized.
func NewRemoveDocumentParams() *RemoveDocumentParams {
	var ()
	return &RemoveDocumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveDocumentParamsWithTimeout creates a new RemoveDocumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveDocumentParamsWithTimeout(timeout time.Duration) *RemoveDocumentParams {
	var ()
	return &RemoveDocumentParams{

		timeout: timeout,
	}
}

// NewRemoveDocumentParamsWithContext creates a new RemoveDocumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveDocumentParamsWithContext(ctx context.Context) *RemoveDocumentParams {
	var ()
	return &RemoveDocumentParams{

		Context: ctx,
	}
}

// NewRemoveDocumentParamsWithHTTPClient creates a new RemoveDocumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveDocumentParamsWithHTTPClient(client *http.Client) *RemoveDocumentParams {
	var ()
	return &RemoveDocumentParams{
		HTTPClient: client,
	}
}

/*RemoveDocumentParams contains all the parameters to send to the API endpoint
for the remove document operation typically these are written to a http.Request
*/
type RemoveDocumentParams struct {

	/*Body*/
	Body *models.ConquestAPIRemoveDocumentCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove document params
func (o *RemoveDocumentParams) WithTimeout(timeout time.Duration) *RemoveDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove document params
func (o *RemoveDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove document params
func (o *RemoveDocumentParams) WithContext(ctx context.Context) *RemoveDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove document params
func (o *RemoveDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove document params
func (o *RemoveDocumentParams) WithHTTPClient(client *http.Client) *RemoveDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove document params
func (o *RemoveDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the remove document params
func (o *RemoveDocumentParams) WithBody(body *models.ConquestAPIRemoveDocumentCommand) *RemoveDocumentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remove document params
func (o *RemoveDocumentParams) SetBody(body *models.ConquestAPIRemoveDocumentCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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