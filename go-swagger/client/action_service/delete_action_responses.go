// Code generated by go-swagger; DO NOT EDIT.

package action_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteActionReader is a Reader for the DeleteAction structure.
type DeleteActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteActionOK creates a DeleteActionOK with default headers values
func NewDeleteActionOK() *DeleteActionOK {
	return &DeleteActionOK{}
}

/*DeleteActionOK handles this case with default header values.

A successful response.
*/
type DeleteActionOK struct {
	Payload interface{}
}

func (o *DeleteActionOK) Error() string {
	return fmt.Sprintf("[POST /api/actions/delete_action][%d] deleteActionOK  %+v", 200, o.Payload)
}

func (o *DeleteActionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}