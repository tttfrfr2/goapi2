// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaskAddTaskCommentRequestBody TaskAddTaskCommentRequestBody
// Example: {"commentContent":"comment"}
//
// swagger:model TaskAddTaskCommentRequestBody
type TaskAddTaskCommentRequestBody struct {

	// comment content
	// Example: comment
	// Required: true
	CommentContent *string `json:"commentContent"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *TaskAddTaskCommentRequestBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// comment content
		// Example: comment
		// Required: true
		CommentContent *string `json:"commentContent"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.CommentContent = props.CommentContent
	return nil
}

// Validate validates this task add task comment request body
func (m *TaskAddTaskCommentRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommentContent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAddTaskCommentRequestBody) validateCommentContent(formats strfmt.Registry) error {

	if err := validate.Required("commentContent", "body", m.CommentContent); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this task add task comment request body based on context it is used
func (m *TaskAddTaskCommentRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskAddTaskCommentRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskAddTaskCommentRequestBody) UnmarshalBinary(b []byte) error {
	var res TaskAddTaskCommentRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
