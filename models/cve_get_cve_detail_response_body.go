// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CveGetCveDetailResponseBody CveGetCveDetailResponseBody
// Example: {"advisories":[{"advisoryID":"advisoryID","osFamily":"redhat"},{"advisoryID":"advisoryID","osFamily":"redhat"},{"advisoryID":"advisoryID","osFamily":"redhat"},{"advisoryID":"advisoryID","osFamily":"redhat"}],"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","cvss":"Et aut quae vero fugit.","cwes":[{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}],"envMetricV2s":[{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"},{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"},{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"}],"envMetricV3s":[{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"}],"references":{"nvd":"https://xxxxxx"},"secMetrics":[{"ar":"","cr":"","createdAt":"2018-07-14T08:13:28Z","ir":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"ar":"","cr":"","createdAt":"2018-07-14T08:13:28Z","ir":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"}],"serverOsFamilies":["Est est.","Dolor natus.","Dolores reiciendis exercitationem labore dolor.","Magnam odit sapiente."],"tmpMetricV2":{"createdAt":"2018-07-14T08:13:28Z","e":"","rc":"","rl":"","updatedAt":"2018-07-14T08:13:28Z"},"tmpMetricV3":{"createdAt":"2018-07-14T08:13:28Z","e":"","rc":"","rl":"","updatedAt":"2018-07-14T08:13:28Z"},"updatedAt":"2018-07-14T08:13:28Z"}
//
// swagger:model CveGetCveDetailResponseBody
type CveGetCveDetailResponseBody struct {

	// advisory of cve
	// Example: [{"advisoryID":"advisoryID","osFamily":"redhat"},{"advisoryID":"advisoryID","osFamily":"redhat"},{"advisoryID":"advisoryID","osFamily":"redhat"}]
	Advisories []*AdvisoryResponseBody `json:"advisories"`

	// created time
	// Example: 2018-07-14T08:13:28Z
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// Cve ID string of cve
	// Example: CVE-2018-1234
	// Required: true
	CveID *string `json:"cveID"`

	// cvss of cve
	// Example: Unde dolorum distinctio voluptatem nemo ut.
	// Required: true
	// Format: binary
	Cvss io.ReadCloser `json:"cvss"`

	// cwes of cve
	// Example: [{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}]
	// Required: true
	Cwes []*CweResponseBody `json:"cwes"`

	// envMetricV2 of cve
	// Example: [{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"},{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"}]
	// Required: true
	EnvMetricV2s []*EnvMetricV2ResponseBody `json:"envMetricV2s"`

	// envMetricV3 of cve
	// Example: [{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"}]
	// Required: true
	EnvMetricV3s []*EnvMetricV3ResponseBody `json:"envMetricV3s"`

	// references of cve
	// Example: {"nvd":"https://xxxxxx"}
	// Required: true
	References map[string]string `json:"references"`

	// secMetric of cve
	// Example: [{"ar":"","cr":"","createdAt":"2018-07-14T08:13:28Z","ir":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"ar":"","cr":"","createdAt":"2018-07-14T08:13:28Z","ir":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"ar":"","cr":"","createdAt":"2018-07-14T08:13:28Z","ir":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"}]
	// Required: true
	SecMetrics []*SecMetricResponseBody `json:"secMetrics"`

	// serverOsFamilies of cve
	// Example: ["Odit suscipit suscipit.","Dolor reiciendis.","Est accusamus et repudiandae.","Illum atque dicta sapiente optio commodi."]
	// Required: true
	ServerOsFamilies []string `json:"serverOsFamilies"`

	// tmp metric v2
	// Required: true
	TmpMetricV2 *TmpMetricResponseBody `json:"tmpMetricV2"`

	// tmp metric v3
	// Required: true
	TmpMetricV3 *TmpMetricResponseBody `json:"tmpMetricV3"`

	// updated time
	// Example: 2018-07-14T08:13:28Z
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *CveGetCveDetailResponseBody) UnmarshalJSON(data []byte) error {
	var props struct {

		// advisory of cve
		// Example: [{"advisoryID":"advisoryID","osFamily":"redhat"},{"advisoryID":"advisoryID","osFamily":"redhat"},{"advisoryID":"advisoryID","osFamily":"redhat"}]
		Advisories []*AdvisoryResponseBody `json:"advisories"`

		// created time
		// Example: 2018-07-14T08:13:28Z
		// Required: true
		// Format: date-time
		CreatedAt *strfmt.DateTime `json:"createdAt"`

		// Cve ID string of cve
		// Example: CVE-2018-1234
		// Required: true
		CveID *string `json:"cveID"`

		// cvss of cve
		// Example: Unde dolorum distinctio voluptatem nemo ut.
		// Required: true
		// Format: binary
		Cvss io.ReadCloser `json:"cvss"`

		// cwes of cve
		// Example: [{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"},{"cweID":"CWE-416","english":"english summary","japanese":"japanese summary","owaspTopTen2017":"10","sourceType":"nvd"}]
		// Required: true
		Cwes []*CweResponseBody `json:"cwes"`

		// envMetricV2 of cve
		// Example: [{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"},{"cdp":"","createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","roleID":1,"roleName":"roleName","td":"","updatedAt":"2018-07-14T08:13:28Z"}]
		// Required: true
		EnvMetricV2s []*EnvMetricV2ResponseBody `json:"envMetricV2s"`

		// envMetricV3 of cve
		// Example: [{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"createdAt":"2018-07-14T08:13:28Z","cveID":"CVE-2018-1234","ma":"","mac":"","mav":"","mc":"","mpr":"","ms":"","mui":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"}]
		// Required: true
		EnvMetricV3s []*EnvMetricV3ResponseBody `json:"envMetricV3s"`

		// references of cve
		// Example: {"nvd":"https://xxxxxx"}
		// Required: true
		References map[string]string `json:"references"`

		// secMetric of cve
		// Example: [{"ar":"","cr":"","createdAt":"2018-07-14T08:13:28Z","ir":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"ar":"","cr":"","createdAt":"2018-07-14T08:13:28Z","ir":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"},{"ar":"","cr":"","createdAt":"2018-07-14T08:13:28Z","ir":"","roleID":1,"roleName":"roleName","updatedAt":"2018-07-14T08:13:28Z"}]
		// Required: true
		SecMetrics []*SecMetricResponseBody `json:"secMetrics"`

		// serverOsFamilies of cve
		// Example: ["Odit suscipit suscipit.","Dolor reiciendis.","Est accusamus et repudiandae.","Illum atque dicta sapiente optio commodi."]
		// Required: true
		ServerOsFamilies []string `json:"serverOsFamilies"`

		// tmp metric v2
		// Required: true
		TmpMetricV2 *TmpMetricResponseBody `json:"tmpMetricV2"`

		// tmp metric v3
		// Required: true
		TmpMetricV3 *TmpMetricResponseBody `json:"tmpMetricV3"`

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

	m.Advisories = props.Advisories
	m.CreatedAt = props.CreatedAt
	m.CveID = props.CveID
	m.Cvss = props.Cvss
	m.Cwes = props.Cwes
	m.EnvMetricV2s = props.EnvMetricV2s
	m.EnvMetricV3s = props.EnvMetricV3s
	m.References = props.References
	m.SecMetrics = props.SecMetrics
	m.ServerOsFamilies = props.ServerOsFamilies
	m.TmpMetricV2 = props.TmpMetricV2
	m.TmpMetricV3 = props.TmpMetricV3
	m.UpdatedAt = props.UpdatedAt
	return nil
}

// Validate validates this cve get cve detail response body
func (m *CveGetCveDetailResponseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvisories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCveID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCvss(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCwes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvMetricV2s(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvMetricV3s(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerOsFamilies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTmpMetricV2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTmpMetricV3(formats); err != nil {
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

func (m *CveGetCveDetailResponseBody) validateAdvisories(formats strfmt.Registry) error {
	if swag.IsZero(m.Advisories) { // not required
		return nil
	}

	for i := 0; i < len(m.Advisories); i++ {
		if swag.IsZero(m.Advisories[i]) { // not required
			continue
		}

		if m.Advisories[i] != nil {
			if err := m.Advisories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("advisories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("advisories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CveGetCveDetailResponseBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CveGetCveDetailResponseBody) validateCveID(formats strfmt.Registry) error {

	if err := validate.Required("cveID", "body", m.CveID); err != nil {
		return err
	}

	return nil
}

func (m *CveGetCveDetailResponseBody) validateCvss(formats strfmt.Registry) error {

	if err := validate.Required("cvss", "body", io.ReadCloser(m.Cvss)); err != nil {
		return err
	}

	return nil
}

func (m *CveGetCveDetailResponseBody) validateCwes(formats strfmt.Registry) error {

	if err := validate.Required("cwes", "body", m.Cwes); err != nil {
		return err
	}

	for i := 0; i < len(m.Cwes); i++ {
		if swag.IsZero(m.Cwes[i]) { // not required
			continue
		}

		if m.Cwes[i] != nil {
			if err := m.Cwes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cwes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cwes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CveGetCveDetailResponseBody) validateEnvMetricV2s(formats strfmt.Registry) error {

	if err := validate.Required("envMetricV2s", "body", m.EnvMetricV2s); err != nil {
		return err
	}

	for i := 0; i < len(m.EnvMetricV2s); i++ {
		if swag.IsZero(m.EnvMetricV2s[i]) { // not required
			continue
		}

		if m.EnvMetricV2s[i] != nil {
			if err := m.EnvMetricV2s[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("envMetricV2s" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("envMetricV2s" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CveGetCveDetailResponseBody) validateEnvMetricV3s(formats strfmt.Registry) error {

	if err := validate.Required("envMetricV3s", "body", m.EnvMetricV3s); err != nil {
		return err
	}

	for i := 0; i < len(m.EnvMetricV3s); i++ {
		if swag.IsZero(m.EnvMetricV3s[i]) { // not required
			continue
		}

		if m.EnvMetricV3s[i] != nil {
			if err := m.EnvMetricV3s[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("envMetricV3s" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("envMetricV3s" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CveGetCveDetailResponseBody) validateReferences(formats strfmt.Registry) error {

	if err := validate.Required("references", "body", m.References); err != nil {
		return err
	}

	return nil
}

func (m *CveGetCveDetailResponseBody) validateSecMetrics(formats strfmt.Registry) error {

	if err := validate.Required("secMetrics", "body", m.SecMetrics); err != nil {
		return err
	}

	for i := 0; i < len(m.SecMetrics); i++ {
		if swag.IsZero(m.SecMetrics[i]) { // not required
			continue
		}

		if m.SecMetrics[i] != nil {
			if err := m.SecMetrics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secMetrics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("secMetrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CveGetCveDetailResponseBody) validateServerOsFamilies(formats strfmt.Registry) error {

	if err := validate.Required("serverOsFamilies", "body", m.ServerOsFamilies); err != nil {
		return err
	}

	return nil
}

func (m *CveGetCveDetailResponseBody) validateTmpMetricV2(formats strfmt.Registry) error {

	if err := validate.Required("tmpMetricV2", "body", m.TmpMetricV2); err != nil {
		return err
	}

	if m.TmpMetricV2 != nil {
		if err := m.TmpMetricV2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tmpMetricV2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tmpMetricV2")
			}
			return err
		}
	}

	return nil
}

func (m *CveGetCveDetailResponseBody) validateTmpMetricV3(formats strfmt.Registry) error {

	if err := validate.Required("tmpMetricV3", "body", m.TmpMetricV3); err != nil {
		return err
	}

	if m.TmpMetricV3 != nil {
		if err := m.TmpMetricV3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tmpMetricV3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tmpMetricV3")
			}
			return err
		}
	}

	return nil
}

func (m *CveGetCveDetailResponseBody) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cve get cve detail response body based on the context it is used
func (m *CveGetCveDetailResponseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdvisories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCwes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvMetricV2s(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvMetricV3s(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTmpMetricV2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTmpMetricV3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CveGetCveDetailResponseBody) contextValidateAdvisories(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Advisories); i++ {

		if m.Advisories[i] != nil {
			if err := m.Advisories[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("advisories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("advisories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CveGetCveDetailResponseBody) contextValidateCwes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Cwes); i++ {

		if m.Cwes[i] != nil {
			if err := m.Cwes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cwes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cwes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CveGetCveDetailResponseBody) contextValidateEnvMetricV2s(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EnvMetricV2s); i++ {

		if m.EnvMetricV2s[i] != nil {
			if err := m.EnvMetricV2s[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("envMetricV2s" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("envMetricV2s" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CveGetCveDetailResponseBody) contextValidateEnvMetricV3s(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EnvMetricV3s); i++ {

		if m.EnvMetricV3s[i] != nil {
			if err := m.EnvMetricV3s[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("envMetricV3s" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("envMetricV3s" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CveGetCveDetailResponseBody) contextValidateSecMetrics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SecMetrics); i++ {

		if m.SecMetrics[i] != nil {
			if err := m.SecMetrics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secMetrics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("secMetrics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CveGetCveDetailResponseBody) contextValidateTmpMetricV2(ctx context.Context, formats strfmt.Registry) error {

	if m.TmpMetricV2 != nil {
		if err := m.TmpMetricV2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tmpMetricV2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tmpMetricV2")
			}
			return err
		}
	}

	return nil
}

func (m *CveGetCveDetailResponseBody) contextValidateTmpMetricV3(ctx context.Context, formats strfmt.Registry) error {

	if m.TmpMetricV3 != nil {
		if err := m.TmpMetricV3.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tmpMetricV3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tmpMetricV3")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CveGetCveDetailResponseBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CveGetCveDetailResponseBody) UnmarshalBinary(b []byte) error {
	var res CveGetCveDetailResponseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}