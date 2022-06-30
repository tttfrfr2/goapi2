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

// TaskUpdateTaskRequestBody TaskUpdateTaskRequestBody
// Example: {"applyingPatchOn":"2018-07-14","mainUserID":1,"priority":"none","status":"new","subUserID":2}
//
// swagger:model TaskUpdateTaskRequestBody
type TaskUpdateTaskRequestBody struct {

	// applyingPatchOn (YYYY-MM-DD) UTC
	// Example: 2018-07-14
	// Format: date
	ApplyingPatchOn strfmt.Date `json:"applyingPatchOn,omitempty"`

	// mainUserID of task
	// Example: 1
	MainUserID int64 `json:"mainUserID,omitempty"`

	// Priority of task
	// Example: medium
	// Enum: [none high medium low]
	Priority string `json:"priority,omitempty"`

	// Status of task
	// Example: ongoing
	// Enum: [new investigating ongoing not_affected risk_accepted workaround]
	Status string `json:"status,omitempty"`

	// subUserID of task
	// Example: 2
	SubUserID int64 `json:"subUserID,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *TaskUpdateTaskRequestBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// applyingPatchOn (YYYY-MM-DD) UTC
		// Example: 2018-07-14
		// Format: date
		ApplyingPatchOn strfmt.Date `json:"applyingPatchOn,omitempty"`

		// mainUserID of task
		// Example: 1
		MainUserID int64 `json:"mainUserID,omitempty"`

		// Priority of task
		// Example: medium
		// Enum: [none high medium low]
		Priority string `json:"priority,omitempty"`

		// Status of task
		// Example: ongoing
		// Enum: [new investigating ongoing not_affected risk_accepted workaround]
		Status string `json:"status,omitempty"`

		// subUserID of task
		// Example: 2
		SubUserID int64 `json:"subUserID,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.ApplyingPatchOn = props.ApplyingPatchOn
	m.MainUserID = props.MainUserID
	m.Priority = props.Priority
	m.Status = props.Status
	m.SubUserID = props.SubUserID
	return nil
}

// Validate validates this task update task request body
func (m *TaskUpdateTaskRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplyingPatchOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskUpdateTaskRequestBody) validateApplyingPatchOn(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplyingPatchOn) { // not required
		return nil
	}

	if err := validate.FormatOf("applyingPatchOn", "body", "date", m.ApplyingPatchOn.String(), formats); err != nil {
		return err
	}

	return nil
}

var taskUpdateTaskRequestBodyTypePriorityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","high","medium","low"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taskUpdateTaskRequestBodyTypePriorityPropEnum = append(taskUpdateTaskRequestBodyTypePriorityPropEnum, v)
	}
}

const (

	// TaskUpdateTaskRequestBodyPriorityNone captures enum value "none"
	TaskUpdateTaskRequestBodyPriorityNone string = "none"

	// TaskUpdateTaskRequestBodyPriorityHigh captures enum value "high"
	TaskUpdateTaskRequestBodyPriorityHigh string = "high"

	// TaskUpdateTaskRequestBodyPriorityMedium captures enum value "medium"
	TaskUpdateTaskRequestBodyPriorityMedium string = "medium"

	// TaskUpdateTaskRequestBodyPriorityLow captures enum value "low"
	TaskUpdateTaskRequestBodyPriorityLow string = "low"
)

// prop value enum
func (m *TaskUpdateTaskRequestBody) validatePriorityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, taskUpdateTaskRequestBodyTypePriorityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TaskUpdateTaskRequestBody) validatePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	// value enum
	if err := m.validatePriorityEnum("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

var taskUpdateTaskRequestBodyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["new","investigating","ongoing","not_affected","risk_accepted","workaround"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taskUpdateTaskRequestBodyTypeStatusPropEnum = append(taskUpdateTaskRequestBodyTypeStatusPropEnum, v)
	}
}

const (

	// TaskUpdateTaskRequestBodyStatusNew captures enum value "new"
	TaskUpdateTaskRequestBodyStatusNew string = "new"

	// TaskUpdateTaskRequestBodyStatusInvestigating captures enum value "investigating"
	TaskUpdateTaskRequestBodyStatusInvestigating string = "investigating"

	// TaskUpdateTaskRequestBodyStatusOngoing captures enum value "ongoing"
	TaskUpdateTaskRequestBodyStatusOngoing string = "ongoing"

	// TaskUpdateTaskRequestBodyStatusNotAffected captures enum value "not_affected"
	TaskUpdateTaskRequestBodyStatusNotAffected string = "not_affected"

	// TaskUpdateTaskRequestBodyStatusRiskAccepted captures enum value "risk_accepted"
	TaskUpdateTaskRequestBodyStatusRiskAccepted string = "risk_accepted"

	// TaskUpdateTaskRequestBodyStatusWorkaround captures enum value "workaround"
	TaskUpdateTaskRequestBodyStatusWorkaround string = "workaround"
)

// prop value enum
func (m *TaskUpdateTaskRequestBody) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, taskUpdateTaskRequestBodyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TaskUpdateTaskRequestBody) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this task update task request body based on context it is used
func (m *TaskUpdateTaskRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskUpdateTaskRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskUpdateTaskRequestBody) UnmarshalBinary(b []byte) error {
	var res TaskUpdateTaskRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
