// Code generated by go-swagger; DO NOT EDIT.

package pkg_cpe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tttfrfr2/tempfiles/openapi/gen/models"
)

// PkgCpeAddCpeReader is a Reader for the PkgCpeAddCpe structure.
type PkgCpeAddCpeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PkgCpeAddCpeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPkgCpeAddCpeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPkgCpeAddCpeOK creates a PkgCpeAddCpeOK with default headers values
func NewPkgCpeAddCpeOK() *PkgCpeAddCpeOK {
	return &PkgCpeAddCpeOK{}
}

/* PkgCpeAddCpeOK describes a response with status code 200, with default header values.

OK response.
*/
type PkgCpeAddCpeOK struct {
	Payload *models.PkgCpeAddCpeResponseBody
}

func (o *PkgCpeAddCpeOK) Error() string {
	return fmt.Sprintf("[POST /v1/pkgCpe/cpe][%d] pkgCpeAddCpeOK  %+v", 200, o.Payload)
}
func (o *PkgCpeAddCpeOK) GetPayload() *models.PkgCpeAddCpeResponseBody {
	return o.Payload
}

func (o *PkgCpeAddCpeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PkgCpeAddCpeResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}