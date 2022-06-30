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

// EnvMetricV2ResponseBody EnvMetricV2ResponseBody
//
// EnvMetricV2 describes a envMetricV2
// Example: {"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"}
//
// swagger:model EnvMetricV2ResponseBody
type EnvMetricV2ResponseBody struct {

	// CDP of envMetricV2
	// Required: true
	Cdp *string `json:"cdp"`

	// created time
	// Example: 2018-07-14T08:13:28Z
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// CveID of envMetricV2
	// Example: CVE-2018-1234
	// Required: true
	CveID *string `json:"cveID"`

	// ServerRoleID of envMetricV2
	// Example: 1
	// Required: true
	RoleID *int64 `json:"roleID"`

	// ServerRoleName of envMetricV2
	// Example: roleName
	// Required: true
	RoleName *string `json:"roleName"`

	// TD of envMetricV2
	// Required: true
	Td *string `json:"td"`

	// updated time
	// Example: 2018-07-14T08:13:28Z
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *EnvMetricV2ResponseBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// CDP of envMetricV2
		// Required: true
		Cdp *string `json:"cdp"`

		// created time
		// Example: 2018-07-14T08:13:28Z
		// Required: true
		// Format: date-time
		CreatedAt *strfmt.DateTime `json:"createdAt"`

		// CveID of envMetricV2
		// Example: CVE-2018-1234
		// Required: true
		CveID *string `json:"cveID"`

		// ServerRoleID of envMetricV2
		// Example: 1
		// Required: true
		RoleID *int64 `json:"roleID"`

		// ServerRoleName of envMetricV2
		// Example: roleName
		// Required: true
		RoleName *string `json:"roleName"`

		// TD of envMetricV2
		// Required: true
		Td *string `json:"td"`

		// updated time
		// Example: 2018-07-14T08:13:28Z
		// Required: true
		// Format: date-time
		UpdatedAt *strfmt.DateTime `json:"updatedAt"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Cdp = props.Cdp
	m.CreatedAt = props.CreatedAt
	m.CveID = props.CveID
	m.RoleID = props.RoleID
	m.RoleName = props.RoleName
	m.Td = props.Td
	m.UpdatedAt = props.UpdatedAt
	return nil
}

// Validate validates this env metric v2 response body
func (m *EnvMetricV2ResponseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCdp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCveID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvMetricV2ResponseBody) validateCdp(formats strfmt.Registry) error {

	if err := validate.Required("cdp", "body", m.Cdp); err != nil {
		return err
	}

	return nil
}

func (m *EnvMetricV2ResponseBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EnvMetricV2ResponseBody) validateCveID(formats strfmt.Registry) error {

	if err := validate.Required("cveID", "body", m.CveID); err != nil {
		return err
	}

	return nil
}

func (m *EnvMetricV2ResponseBody) validateRoleID(formats strfmt.Registry) error {

	if err := validate.Required("roleID", "body", m.RoleID); err != nil {
		return err
	}

	return nil
}

func (m *EnvMetricV2ResponseBody) validateRoleName(formats strfmt.Registry) error {

	if err := validate.Required("roleName", "body", m.RoleName); err != nil {
		return err
	}

	return nil
}

func (m *EnvMetricV2ResponseBody) validateTd(formats strfmt.Registry) error {

	if err := validate.Required("td", "body", m.Td); err != nil {
		return err
	}

	return nil
}

func (m *EnvMetricV2ResponseBody) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this env metric v2 response body based on context it is used
func (m *EnvMetricV2ResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EnvMetricV2ResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvMetricV2ResponseBody) UnmarshalBinary(b []byte) error {
	var res EnvMetricV2ResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
