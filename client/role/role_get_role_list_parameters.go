// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewRoleGetRoleListParams creates a new RoleGetRoleListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRoleGetRoleListParams() *RoleGetRoleListParams {
	return &RoleGetRoleListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRoleGetRoleListParamsWithTimeout creates a new RoleGetRoleListParams object
// with the ability to set a timeout on a request.
func NewRoleGetRoleListParamsWithTimeout(timeout time.Duration) *RoleGetRoleListParams {
	return &RoleGetRoleListParams{
		timeout: timeout,
	}
}

// NewRoleGetRoleListParamsWithContext creates a new RoleGetRoleListParams object
// with the ability to set a context for a request.
func NewRoleGetRoleListParamsWithContext(ctx context.Context) *RoleGetRoleListParams {
	return &RoleGetRoleListParams{
		Context: ctx,
	}
}

// NewRoleGetRoleListParamsWithHTTPClient creates a new RoleGetRoleListParams object
// with the ability to set a custom HTTPClient for a request.
func NewRoleGetRoleListParamsWithHTTPClient(client *http.Client) *RoleGetRoleListParams {
	return &RoleGetRoleListParams{
		HTTPClient: client,
	}
}

/* RoleGetRoleListParams contains all the parameters to send to the API endpoint
   for the role get role list operation.

   Typically these are written to a http.Request.
*/
type RoleGetRoleListParams struct {

	/* Authorization.

	   api key auth
	*/
	Authorization *string

	/* FilterCveID.

	   CveID filter
	*/
	FilterCveID *string

	/* Limit.

	   Limit (default: 20, max: 100)

	   Default: 20
	*/
	Limit *int64

	/* Offset.

	   Offset
	*/
	Offset *int64

	/* Page.

	   Page Number (default: 1)

	   Default: 1
	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the role get role list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RoleGetRoleListParams) WithDefaults() *RoleGetRoleListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the role get role list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RoleGetRoleListParams) SetDefaults() {
	var (
		limitDefault = int64(20)

		offsetDefault = int64(0)

		pageDefault = int64(1)
	)

	val := RoleGetRoleListParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
		Page:   &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the role get role list params
func (o *RoleGetRoleListParams) WithTimeout(timeout time.Duration) *RoleGetRoleListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the role get role list params
func (o *RoleGetRoleListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the role get role list params
func (o *RoleGetRoleListParams) WithContext(ctx context.Context) *RoleGetRoleListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the role get role list params
func (o *RoleGetRoleListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the role get role list params
func (o *RoleGetRoleListParams) WithHTTPClient(client *http.Client) *RoleGetRoleListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the role get role list params
func (o *RoleGetRoleListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the role get role list params
func (o *RoleGetRoleListParams) WithAuthorization(authorization *string) *RoleGetRoleListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the role get role list params
func (o *RoleGetRoleListParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithFilterCveID adds the filterCveID to the role get role list params
func (o *RoleGetRoleListParams) WithFilterCveID(filterCveID *string) *RoleGetRoleListParams {
	o.SetFilterCveID(filterCveID)
	return o
}

// SetFilterCveID adds the filterCveId to the role get role list params
func (o *RoleGetRoleListParams) SetFilterCveID(filterCveID *string) {
	o.FilterCveID = filterCveID
}

// WithLimit adds the limit to the role get role list params
func (o *RoleGetRoleListParams) WithLimit(limit *int64) *RoleGetRoleListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the role get role list params
func (o *RoleGetRoleListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the role get role list params
func (o *RoleGetRoleListParams) WithOffset(offset *int64) *RoleGetRoleListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the role get role list params
func (o *RoleGetRoleListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPage adds the page to the role get role list params
func (o *RoleGetRoleListParams) WithPage(page *int64) *RoleGetRoleListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the role get role list params
func (o *RoleGetRoleListParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *RoleGetRoleListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// header param Authorization
		if err := r.SetHeaderParam("Authorization", *o.Authorization); err != nil {
			return err
		}
	}

	if o.FilterCveID != nil {

		// query param filterCveID
		var qrFilterCveID string

		if o.FilterCveID != nil {
			qrFilterCveID = *o.FilterCveID
		}
		qFilterCveID := qrFilterCveID
		if qFilterCveID != "" {

			if err := r.SetQueryParam("filterCveID", qFilterCveID); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
