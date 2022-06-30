// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tttfrfr2/tempfiles/openapi/gen/models"
)

// RoleGetRoleListReader is a Reader for the RoleGetRoleList structure.
type RoleGetRoleListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoleGetRoleListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRoleGetRoleListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRoleGetRoleListOK creates a RoleGetRoleListOK with default headers values
func NewRoleGetRoleListOK() *RoleGetRoleListOK {
	return &RoleGetRoleListOK{}
}

/* RoleGetRoleListOK describes a response with status code 200, with default header values.

OK response.
*/
type RoleGetRoleListOK struct {
	Payload *models.RoleGetRoleListResponseBody
}

func (o *RoleGetRoleListOK) Error() string {
	return fmt.Sprintf("[GET /v1/roles][%d] roleGetRoleListOK  %+v", 200, o.Payload)
}
func (o *RoleGetRoleListOK) GetPayload() *models.RoleGetRoleListResponseBody {
	return o.Payload
}

func (o *RoleGetRoleListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleGetRoleListResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
