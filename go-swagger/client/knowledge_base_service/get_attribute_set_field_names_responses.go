// Code generated by go-swagger; DO NOT EDIT.

package knowledge_base_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
)

// GetAttributeSetFieldNamesReader is a Reader for the GetAttributeSetFieldNames structure.
type GetAttributeSetFieldNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAttributeSetFieldNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAttributeSetFieldNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAttributeSetFieldNamesOK creates a GetAttributeSetFieldNamesOK with default headers values
func NewGetAttributeSetFieldNamesOK() *GetAttributeSetFieldNamesOK {
	return &GetAttributeSetFieldNamesOK{}
}

/*GetAttributeSetFieldNamesOK handles this case with default header values.

A successful response.
*/
type GetAttributeSetFieldNamesOK struct {
	Payload *models.ConquestAPIGetAttributeSetFieldNamesResponse
}

func (o *GetAttributeSetFieldNamesOK) Error() string {
	return fmt.Sprintf("[POST /api/knowledge_base/list_field_names_for_attribute_set][%d] getAttributeSetFieldNamesOK  %+v", 200, o.Payload)
}

func (o *GetAttributeSetFieldNamesOK) GetPayload() *models.ConquestAPIGetAttributeSetFieldNamesResponse {
	return o.Payload
}

func (o *GetAttributeSetFieldNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConquestAPIGetAttributeSetFieldNamesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
