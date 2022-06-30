// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tttfrfr2/goapi2/models"
)

// TaskAddTaskCommentReader is a Reader for the TaskAddTaskComment structure.
type TaskAddTaskCommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaskAddTaskCommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaskAddTaskCommentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTaskAddTaskCommentOK creates a TaskAddTaskCommentOK with default headers values
func NewTaskAddTaskCommentOK() *TaskAddTaskCommentOK {
	return &TaskAddTaskCommentOK{}
}

/* TaskAddTaskCommentOK describes a response with status code 200, with default header values.

OK response.
*/
type TaskAddTaskCommentOK struct {
	Payload *models.TaskAddTaskCommentResponseBody
}

func (o *TaskAddTaskCommentOK) Error() string {
	return fmt.Sprintf("[POST /v1/task/{taskID}/comment][%d] taskAddTaskCommentOK  %+v", 200, o.Payload)
}
func (o *TaskAddTaskCommentOK) GetPayload() *models.TaskAddTaskCommentResponseBody {
	return o.Payload
}

func (o *TaskAddTaskCommentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskAddTaskCommentResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
