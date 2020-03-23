// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConquestAPIAddDocumentCommand AddDocumentCommand works in two steps, first you add a record using this command. The command may return AddDocumentCommand.UploadURI
// use that URI to post file content. See AddDocumentResult
//
// swagger:model conquest_apiAddDocumentCommand
type ConquestAPIAddDocumentCommand struct {

	// Address is a URI with a supported scheme (blob://, file://, https://, conquest://, trim://).
	//
	// - URIs use the canonical '/' path separator (not '\').
	//
	// - If the path component of a URI contains hex encoded characters, eg. '%20' for a space ' ', the path will be decoded and saved as ' '.
	//
	// - The [blob://] scheme represents a filesystem that is managed by Conquest. This is the scheme used in "conquest.live"
	//   All urls take the form 'blob:/${configured-container}/path/to/file.txt'
	//
	// - The [conquest://] scheme provides a way to encode links to Conquest Objects and Views. See the documentation to learn
	//   how to construct a conquest link. (TODO documentation link, conquest://? is an alias for "https://link.conquest.live/?")
	//
	// - The [https://] scheme can be used to link websites, "http" is not supported.
	//
	// - The [file://] scheme is not available in "conquest.live", it is the default for in-house installations.
	//
	// - When a [file://] scheme is used to *submit* a document, it refers to a location available on the API Server filesystem. Only configured locations will be accepted.
	//   A configured location takes the form: 'file://${configured-location}/path/to/file.txt'
	//
	// - When *submitting* and *serving*, "conquest_documents", is a configured location for the [file://] scheme:
	//     - "conquest_documents" points to the Conquest Document directory
	//       For example:
	//         "file://conquest_documents/relative/path/file.txt" will point to
	//         "${ConquestDocumentDirectory}/relative/path/file.txt
	//
	// [Special handling]: For local filesystem paths are as follows. Consumers of the API / End-User apps cannot use local filesystem paths.
	//
	// - The [file://] scheme is used to *serve* a file, it refers to a location available on the API Server filesystem.
	//
	//      - If a [file://] scheme uses the prefix "file://MACHINE/", this path path aligns with the underlying filesystem.
	//        For example: "file://MACHINE/shared/folder/file.txt" translates to the UNC path "\\MACHINE\shared\folder\test.txt"
	//
	//      - "file://MACHINE/" locations cannot be submitted via. the API.
	//
	// Address is a protected property. It is not exposed to the end-user in subsequent requests. To reference this document
	// use the returned Document.DocumentID
	Address string `json:"Address,omitempty"`

	// content length
	ContentLength string `json:"ContentLength,omitempty"`

	// content type
	ContentType string `json:"ContentType,omitempty"`

	// CreateTime is unique. When adding a document, there is no DocumentID yet, the CreateTime should be used as a key until the DocumentID is retrieved.
	// Format: date-time
	CreateTime strfmt.DateTime `json:"CreateTime,omitempty"`

	// document description
	DocumentDescription string `json:"DocumentDescription,omitempty"`

	// A list of calculated hashes / checksum of the file to be uploaded.
	Hashes []string `json:"Hashes"`

	// link existing document
	LinkExistingDocument bool `json:"LinkExistingDocument,omitempty"`

	// object key
	ObjectKey *ConquestAPIObjectKey `json:"ObjectKey,omitempty"`
}

// Validate validates this conquest api add document command
func (m *ConquestAPIAddDocumentCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConquestAPIAddDocumentCommand) validateCreateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("CreateTime", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConquestAPIAddDocumentCommand) validateObjectKey(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectKey) { // not required
		return nil
	}

	if m.ObjectKey != nil {
		if err := m.ObjectKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ObjectKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConquestAPIAddDocumentCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConquestAPIAddDocumentCommand) UnmarshalBinary(b []byte) error {
	var res ConquestAPIAddDocumentCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}