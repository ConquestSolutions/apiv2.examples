// Code generated by go-swagger; DO NOT EDIT.

package request_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
)

// UpdateRequestReader is a Reader for the UpdateRequest structure.
type UpdateRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateRequestOK creates a UpdateRequestOK with default headers values
func NewUpdateRequestOK() *UpdateRequestOK {
	return &UpdateRequestOK{}
}

/*UpdateRequestOK handles this case with default header values.

A successful response.
*/
type UpdateRequestOK struct {
	Payload *models.ConquestAPIChangeSetResult
}

func (o *UpdateRequestOK) Error() string {
	return fmt.Sprintf("[POST /api/requests/update_request][%d] updateRequestOK  %+v", 200, o.Payload)
}

func (o *UpdateRequestOK) GetPayload() *models.ConquestAPIChangeSetResult {
	return o.Payload
}

func (o *UpdateRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPIChangeSetResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
