// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ServerUpdatePkgPasteServerReader is a Reader for the ServerUpdatePkgPasteServer structure.
type ServerUpdatePkgPasteServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServerUpdatePkgPasteServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServerUpdatePkgPasteServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServerUpdatePkgPasteServerOK creates a ServerUpdatePkgPasteServerOK with default headers values
func NewServerUpdatePkgPasteServerOK() *ServerUpdatePkgPasteServerOK {
	return &ServerUpdatePkgPasteServerOK{}
}

/* ServerUpdatePkgPasteServerOK describes a response with status code 200, with default header values.

OK response.
*/
type ServerUpdatePkgPasteServerOK struct {
}

func (o *ServerUpdatePkgPasteServerOK) Error() string {
	return fmt.Sprintf("[PUT /v1/server/paste/{serverID}][%d] serverUpdatePkgPasteServerOK ", 200)
}

func (o *ServerUpdatePkgPasteServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
