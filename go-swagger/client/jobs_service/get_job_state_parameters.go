// Code generated by go-swagger; DO NOT EDIT.

package jobs_service

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

// NewGetJobStateParams creates a new GetJobStateParams object
// with the default values initialized.
func NewGetJobStateParams() *GetJobStateParams {
	var ()
	return &GetJobStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetJobStateParamsWithTimeout creates a new GetJobStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetJobStateParamsWithTimeout(timeout time.Duration) *GetJobStateParams {
	var ()
	return &GetJobStateParams{

		timeout: timeout,
	}
}

// NewGetJobStateParamsWithContext creates a new GetJobStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetJobStateParamsWithContext(ctx context.Context) *GetJobStateParams {
	var ()
	return &GetJobStateParams{

		Context: ctx,
	}
}

// NewGetJobStateParamsWithHTTPClient creates a new GetJobStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetJobStateParamsWithHTTPClient(client *http.Client) *GetJobStateParams {
	var ()
	return &GetJobStateParams{
		HTTPClient: client,
	}
}

/*GetJobStateParams contains all the parameters to send to the API endpoint
for the get job state operation typically these are written to a http.Request
*/
type GetJobStateParams struct {

	/*Body*/
	Body *models.ConquestAPIGetJobStateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get job state params
func (o *GetJobStateParams) WithTimeout(timeout time.Duration) *GetJobStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get job state params
func (o *GetJobStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get job state params
func (o *GetJobStateParams) WithContext(ctx context.Context) *GetJobStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get job state params
func (o *GetJobStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get job state params
func (o *GetJobStateParams) WithHTTPClient(client *http.Client) *GetJobStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get job state params
func (o *GetJobStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get job state params
func (o *GetJobStateParams) WithBody(body *models.ConquestAPIGetJobStateRequest) *GetJobStateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get job state params
func (o *GetJobStateParams) SetBody(body *models.ConquestAPIGetJobStateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetJobStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
