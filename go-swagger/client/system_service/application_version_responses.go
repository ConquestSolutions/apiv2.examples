// Code generated by go-swagger; DO NOT EDIT.

package system_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
)

// ApplicationVersionReader is a Reader for the ApplicationVersion structure.
type ApplicationVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewApplicationVersionOK creates a ApplicationVersionOK with default headers values
func NewApplicationVersionOK() *ApplicationVersionOK {
	return &ApplicationVersionOK{}
}

/*ApplicationVersionOK handles this case with default header values.

A successful response.
*/
type ApplicationVersionOK struct {
	Payload *models.ConquestAPISystemVersionResponse
}

func (o *ApplicationVersionOK) Error() string {
	return fmt.Sprintf("[GET /api/system/version][%d] applicationVersionOK  %+v", 200, o.Payload)
}

func (o *ApplicationVersionOK) GetPayload() *models.ConquestAPISystemVersionResponse {
	return o.Payload
}

func (o *ApplicationVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPISystemVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
