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

// PkgCpeAddCpeRequestBody PkgCpeAddCpeRequestBody
// Example: {"cpeName":"cpe:/a:berlios:discussion_forum_2k:3.3","isURI":true,"serverID":1}
//
// swagger:model PkgCpeAddCpeRequestBody
type PkgCpeAddCpeRequestBody struct {

	// Cpe Name(cpe uri or cpe format string)
	// Example: cpe:/a:berlios:discussion_forum_2k:3.3
	// Required: true
	CpeName *string `json:"cpeName"`

	// isURI specifies whether cpeName is in URI format or FormatString format.
	// Example: true
	IsURI *bool `json:"isURI,omitempty"`

	// ServerID
	// Example: 1
	// Required: true
	ServerID *int64 `json:"serverID"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *PkgCpeAddCpeRequestBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// Cpe Name(cpe uri or cpe format string)
		// Example: cpe:/a:berlios:discussion_forum_2k:3.3
		// Required: true
		CpeName *string `json:"cpeName"`

		// isURI specifies whether cpeName is in URI format or FormatString format.
		// Example: true
		IsURI *bool `json:"isURI,omitempty"`

		// ServerID
		// Example: 1
		// Required: true
		ServerID *int64 `json:"serverID"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.CpeName = props.CpeName
	m.IsURI = props.IsURI
	m.ServerID = props.ServerID
	return nil
}

// Validate validates this pkg cpe add cpe request body
func (m *PkgCpeAddCpeRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCpeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgCpeAddCpeRequestBody) validateCpeName(formats strfmt.Registry) error {

	if err := validate.Required("cpeName", "body", m.CpeName); err != nil {
		return err
	}

	return nil
}

func (m *PkgCpeAddCpeRequestBody) validateServerID(formats strfmt.Registry) error {

	if err := validate.Required("serverID", "body", m.ServerID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pkg cpe add cpe request body based on context it is used
func (m *PkgCpeAddCpeRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PkgCpeAddCpeRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgCpeAddCpeRequestBody) UnmarshalBinary(b []byte) error {
	var res PkgCpeAddCpeRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}