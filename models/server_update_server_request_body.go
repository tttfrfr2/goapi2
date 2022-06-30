// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServerUpdateServerRequestBody ServerUpdateServerRequestBody
// Example: {"additionalInfo":"This server is expensive","defaultUserID":1,"roleID":1,"serverName":"new-server-name"}
//
// swagger:model ServerUpdateServerRequestBody
type ServerUpdateServerRequestBody struct {

	// Additional information of the server
	// Example: This server is expensive
	AdditionalInfo string `json:"additionalInfo,omitempty"`

	// DefaultUserID of server
	// Example: 1
	DefaultUserID int64 `json:"defaultUserID,omitempty"`

	// ServerRoleID of server
	// Example: 1
	RoleID int64 `json:"roleID,omitempty"`

	// ServerName of server
	// Example: new-server-name
	ServerName string `json:"serverName,omitempty"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *ServerUpdateServerRequestBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// Additional information of the server
		// Example: This server is expensive
		AdditionalInfo string `json:"additionalInfo,omitempty"`

		// DefaultUserID of server
		// Example: 1
		DefaultUserID int64 `json:"defaultUserID,omitempty"`

		// ServerRoleID of server
		// Example: 1
		RoleID int64 `json:"roleID,omitempty"`

		// ServerName of server
		// Example: new-server-name
		ServerName string `json:"serverName,omitempty"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.AdditionalInfo = props.AdditionalInfo
	m.DefaultUserID = props.DefaultUserID
	m.RoleID = props.RoleID
	m.ServerName = props.ServerName
	return nil
}

// Validate validates this server update server request body
func (m *ServerUpdateServerRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this server update server request body based on context it is used
func (m *ServerUpdateServerRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerUpdateServerRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerUpdateServerRequestBody) UnmarshalBinary(b []byte) error {
	var res ServerUpdateServerRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
