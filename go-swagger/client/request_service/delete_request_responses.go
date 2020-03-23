// Code generated by go-swagger; DO NOT EDIT.

package request_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteRequestReader is a Reader for the DeleteRequest structure.
type DeleteRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRequestOK creates a DeleteRequestOK with default headers values
func NewDeleteRequestOK() *DeleteRequestOK {
	return &DeleteRequestOK{}
}

/*DeleteRequestOK handles this case with default header values.

A successful response.
*/
type DeleteRequestOK struct {
	Payload interface{}
}

func (o *DeleteRequestOK) Error() string {
	return fmt.Sprintf("[POST /api/requests/delete_request][%d] deleteRequestOK  %+v", 200, o.Payload)
}

func (o *DeleteRequestOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}