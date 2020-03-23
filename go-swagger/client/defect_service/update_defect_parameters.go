// Code generated by go-swagger; DO NOT EDIT.

package defect_service

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

// NewUpdateDefectParams creates a new UpdateDefectParams object
// with the default values initialized.
func NewUpdateDefectParams() *UpdateDefectParams {
	var ()
	return &UpdateDefectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDefectParamsWithTimeout creates a new UpdateDefectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDefectParamsWithTimeout(timeout time.Duration) *UpdateDefectParams {
	var ()
	return &UpdateDefectParams{

		timeout: timeout,
	}
}

// NewUpdateDefectParamsWithContext creates a new UpdateDefectParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDefectParamsWithContext(ctx context.Context) *UpdateDefectParams {
	var ()
	return &UpdateDefectParams{

		Context: ctx,
	}
}

// NewUpdateDefectParamsWithHTTPClient creates a new UpdateDefectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDefectParamsWithHTTPClient(client *http.Client) *UpdateDefectParams {
	var ()
	return &UpdateDefectParams{
		HTTPClient: client,
	}
}

/*UpdateDefectParams contains all the parameters to send to the API endpoint
for the update defect operation typically these are written to a http.Request
*/
type UpdateDefectParams struct {

	/*Body*/
	Body *models.ConquestAPIDefectRecordChangeSet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update defect params
func (o *UpdateDefectParams) WithTimeout(timeout time.Duration) *UpdateDefectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update defect params
func (o *UpdateDefectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update defect params
func (o *UpdateDefectParams) WithContext(ctx context.Context) *UpdateDefectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update defect params
func (o *UpdateDefectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update defect params
func (o *UpdateDefectParams) WithHTTPClient(client *http.Client) *UpdateDefectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update defect params
func (o *UpdateDefectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update defect params
func (o *UpdateDefectParams) WithBody(body *models.ConquestAPIDefectRecordChangeSet) *UpdateDefectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update defect params
func (o *UpdateDefectParams) SetBody(body *models.ConquestAPIDefectRecordChangeSet) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDefectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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