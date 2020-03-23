// Code generated by go-swagger; DO NOT EDIT.

package action_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
)

// CompleteActionReader is a Reader for the CompleteAction structure.
type CompleteActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCompleteActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCompleteActionOK creates a CompleteActionOK with default headers values
func NewCompleteActionOK() *CompleteActionOK {
	return &CompleteActionOK{}
}

/*CompleteActionOK handles this case with default header values.

A successful response.
*/
type CompleteActionOK struct {
	Payload *models.ConquestAPICompleteActionResult
}

func (o *CompleteActionOK) Error() string {
	return fmt.Sprintf("[POST /api/actions/complete_action][%d] completeActionOK  %+v", 200, o.Payload)
}

func (o *CompleteActionOK) GetPayload() *models.ConquestAPICompleteActionResult {
	return o.Payload
}

func (o *CompleteActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPICompleteActionResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}