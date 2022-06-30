// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewServerGetServerListParams creates a new ServerGetServerListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServerGetServerListParams() *ServerGetServerListParams {
	return &ServerGetServerListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServerGetServerListParamsWithTimeout creates a new ServerGetServerListParams object
// with the ability to set a timeout on a request.
func NewServerGetServerListParamsWithTimeout(timeout time.Duration) *ServerGetServerListParams {
	return &ServerGetServerListParams{
		timeout: timeout,
	}
}

// NewServerGetServerListParamsWithContext creates a new ServerGetServerListParams object
// with the ability to set a context for a request.
func NewServerGetServerListParamsWithContext(ctx context.Context) *ServerGetServerListParams {
	return &ServerGetServerListParams{
		Context: ctx,
	}
}

// NewServerGetServerListParamsWithHTTPClient creates a new ServerGetServerListParams object
// with the ability to set a custom HTTPClient for a request.
func NewServerGetServerListParamsWithHTTPClient(client *http.Client) *ServerGetServerListParams {
	return &ServerGetServerListParams{
		HTTPClient: client,
	}
}

/* ServerGetServerListParams contains all the parameters to send to the API endpoint
   for the server get server list operation.

   Typically these are written to a http.Request.
*/
type ServerGetServerListParams struct {

	/* Authorization.

	   api key auth
	*/
	Authorization *string

	/* FilterCveID.

	   CveID filter
	*/
	FilterCveID *string

	/* FilterRoleID.

	   ServerRoleID filter
	*/
	FilterRoleID *int64

	/* FilterTagName.

	   ServerTagName filter
	*/
	FilterTagName *string

	/* Limit.

	   Limit

	   Default: 20
	*/
	Limit *int64

	/* Offset.

	   Offset
	*/
	Offset *int64

	/* Page.

	   Page Number

	   Default: 1
	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the server get server list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServerGetServerListParams) WithDefaults() *ServerGetServerListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the server get server list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServerGetServerListParams) SetDefaults() {
	var (
		limitDefault = int64(20)

		offsetDefault = int64(0)

		pageDefault = int64(1)
	)

	val := ServerGetServerListParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,
		Page:   &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the server get server list params
func (o *ServerGetServerListParams) WithTimeout(timeout time.Duration) *ServerGetServerListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the server get server list params
func (o *ServerGetServerListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the server get server list params
func (o *ServerGetServerListParams) WithContext(ctx context.Context) *ServerGetServerListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the server get server list params
func (o *ServerGetServerListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the server get server list params
func (o *ServerGetServerListParams) WithHTTPClient(client *http.Client) *ServerGetServerListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the server get server list params
func (o *ServerGetServerListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the server get server list params
func (o *ServerGetServerListParams) WithAuthorization(authorization *string) *ServerGetServerListParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the server get server list params
func (o *ServerGetServerListParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithFilterCveID adds the filterCveID to the server get server list params
func (o *ServerGetServerListParams) WithFilterCveID(filterCveID *string) *ServerGetServerListParams {
	o.SetFilterCveID(filterCveID)
	return o
}

// SetFilterCveID adds the filterCveId to the server get server list params
func (o *ServerGetServerListParams) SetFilterCveID(filterCveID *string) {
	o.FilterCveID = filterCveID
}

// WithFilterRoleID adds the filterRoleID to the server get server list params
func (o *ServerGetServerListParams) WithFilterRoleID(filterRoleID *int64) *ServerGetServerListParams {
	o.SetFilterRoleID(filterRoleID)
	return o
}

// SetFilterRoleID adds the filterRoleId to the server get server list params
func (o *ServerGetServerListParams) SetFilterRoleID(filterRoleID *int64) {
	o.FilterRoleID = filterRoleID
}

// WithFilterTagName adds the filterTagName to the server get server list params
func (o *ServerGetServerListParams) WithFilterTagName(filterTagName *string) *ServerGetServerListParams {
	o.SetFilterTagName(filterTagName)
	return o
}

// SetFilterTagName adds the filterTagName to the server get server list params
func (o *ServerGetServerListParams) SetFilterTagName(filterTagName *string) {
	o.FilterTagName = filterTagName
}

// WithLimit adds the limit to the server get server list params
func (o *ServerGetServerListParams) WithLimit(limit *int64) *ServerGetServerListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the server get server list params
func (o *ServerGetServerListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the server get server list params
func (o *ServerGetServerListParams) WithOffset(offset *int64) *ServerGetServerListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the server get server list params
func (o *ServerGetServerListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPage adds the page to the server get server list params
func (o *ServerGetServerListParams) WithPage(page *int64) *ServerGetServerListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the server get server list params
func (o *ServerGetServerListParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *ServerGetServerListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilterRoleID != nil {

		// query param filterRoleID
		var qrFilterRoleID int64

		if o.FilterRoleID != nil {
			qrFilterRoleID = *o.FilterRoleID
		}
		qFilterRoleID := swag.FormatInt64(qrFilterRoleID)
		if qFilterRoleID != "" {

			if err := r.SetQueryParam("filterRoleID", qFilterRoleID); err != nil {
				return err
			}
		}
	}

	if o.FilterTagName != nil {

		// query param filterTagName
		var qrFilterTagName string

		if o.FilterTagName != nil {
			qrFilterTagName = *o.FilterTagName
		}
		qFilterTagName := qrFilterTagName
		if qFilterTagName != "" {

			if err := r.SetQueryParam("filterTagName", qFilterTagName); err != nil {
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
