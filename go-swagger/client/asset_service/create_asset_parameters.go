// Code generated by go-swagger; DO NOT EDIT.

package asset_service

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

// NewCreateAssetParams creates a new CreateAssetParams object
// with the default values initialized.
func NewCreateAssetParams() *CreateAssetParams {
	var ()
	return &CreateAssetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAssetParamsWithTimeout creates a new CreateAssetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAssetParamsWithTimeout(timeout time.Duration) *CreateAssetParams {
	var ()
	return &CreateAssetParams{

		timeout: timeout,
	}
}

// NewCreateAssetParamsWithContext creates a new CreateAssetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAssetParamsWithContext(ctx context.Context) *CreateAssetParams {
	var ()
	return &CreateAssetParams{

		Context: ctx,
	}
}

// NewCreateAssetParamsWithHTTPClient creates a new CreateAssetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAssetParamsWithHTTPClient(client *http.Client) *CreateAssetParams {
	var ()
	return &CreateAssetParams{
		HTTPClient: client,
	}
}

/*CreateAssetParams contains all the parameters to send to the API endpoint
for the create asset operation typically these are written to a http.Request
*/
type CreateAssetParams struct {

	/*Body*/
	Body *models.ConquestAPICreateAssetCommand

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create asset params
func (o *CreateAssetParams) WithTimeout(timeout time.Duration) *CreateAssetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create asset params
func (o *CreateAssetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create asset params
func (o *CreateAssetParams) WithContext(ctx context.Context) *CreateAssetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create asset params
func (o *CreateAssetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create asset params
func (o *CreateAssetParams) WithHTTPClient(client *http.Client) *CreateAssetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create asset params
func (o *CreateAssetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create asset params
func (o *CreateAssetParams) WithBody(body *models.ConquestAPICreateAssetCommand) *CreateAssetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create asset params
func (o *CreateAssetParams) SetBody(body *models.ConquestAPICreateAssetCommand) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAssetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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