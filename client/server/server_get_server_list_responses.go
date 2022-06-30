// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tttfrfr2/goapi2/models"
)

// ServerGetServerListReader is a Reader for the ServerGetServerList structure.
type ServerGetServerListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServerGetServerListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServerGetServerListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServerGetServerListOK creates a ServerGetServerListOK with default headers values
func NewServerGetServerListOK() *ServerGetServerListOK {
	return &ServerGetServerListOK{}
}

/* ServerGetServerListOK describes a response with status code 200, with default header values.

OK response.
*/
type ServerGetServerListOK struct {
	Payload *models.ServerGetServerListResponseBody
}

func (o *ServerGetServerListOK) Error() string {
	return fmt.Sprintf("[GET /v1/servers][%d] serverGetServerListOK  %+v", 200, o.Payload)
}
func (o *ServerGetServerListOK) GetPayload() *models.ServerGetServerListResponseBody {
	return o.Payload
}

func (o *ServerGetServerListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerGetServerListResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
