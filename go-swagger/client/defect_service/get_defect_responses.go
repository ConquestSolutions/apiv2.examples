// Code generated by go-swagger; DO NOT EDIT.

package defect_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
)

// GetDefectReader is a Reader for the GetDefect structure.
type GetDefectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDefectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDefectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDefectOK creates a GetDefectOK with default headers values
func NewGetDefectOK() *GetDefectOK {
	return &GetDefectOK{}
}

/*GetDefectOK handles this case with default header values.

A successful response.
*/
type GetDefectOK struct {
	Payload *models.ConquestAPIDefectEntity
}

func (o *GetDefectOK) Error() string {
	return fmt.Sprintf("[GET /api/defects/{value}][%d] getDefectOK  %+v", 200, o.Payload)
}

func (o *GetDefectOK) GetPayload() *models.ConquestAPIDefectEntity {
	return o.Payload
}

func (o *GetDefectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPIDefectEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}