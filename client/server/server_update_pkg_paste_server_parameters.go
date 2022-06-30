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

	"github.com/tttfrfr2/tempfiles/openapi/gen/models"
)

// NewServerUpdatePkgPasteServerParams creates a new ServerUpdatePkgPasteServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServerUpdatePkgPasteServerParams() *ServerUpdatePkgPasteServerParams {
	return &ServerUpdatePkgPasteServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServerUpdatePkgPasteServerParamsWithTimeout creates a new ServerUpdatePkgPasteServerParams object
// with the ability to set a timeout on a request.
func NewServerUpdatePkgPasteServerParamsWithTimeout(timeout time.Duration) *ServerUpdatePkgPasteServerParams {
	return &ServerUpdatePkgPasteServerParams{
		timeout: timeout,
	}
}

// NewServerUpdatePkgPasteServerParamsWithContext creates a new ServerUpdatePkgPasteServerParams object
// with the ability to set a context for a request.
func NewServerUpdatePkgPasteServerParamsWithContext(ctx context.Context) *ServerUpdatePkgPasteServerParams {
	return &ServerUpdatePkgPasteServerParams{
		Context: ctx,
	}
}

// NewServerUpdatePkgPasteServerParamsWithHTTPClient creates a new ServerUpdatePkgPasteServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewServerUpdatePkgPasteServerParamsWithHTTPClient(client *http.Client) *ServerUpdatePkgPasteServerParams {
	return &ServerUpdatePkgPasteServerParams{
		HTTPClient: client,
	}
}

/* ServerUpdatePkgPasteServerParams contains all the parameters to send to the API endpoint
   for the server update pkg paste server operation.

   Typically these are written to a http.Request.
*/
type ServerUpdatePkgPasteServerParams struct {

	/* Authorization.

	   api key auth
	*/
	Authorization *string

	// UpdatePkgPasteServerRequestBody.
	UpdatePkgPasteServerRequestBody *models.ServerUpdatePkgPasteServerRequestBody

	/* ServerID.

	   Server ID

	   Format: int64
	*/
	ServerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the server update pkg paste server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServerUpdatePkgPasteServerParams) WithDefaults() *ServerUpdatePkgPasteServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the server update pkg paste server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServerUpdatePkgPasteServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the server update pkg paste server params
func (o *ServerUpdatePkgPasteServerParams) WithTimeout(timeout time.Duration) *ServerUpdatePkgPasteServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the server update pkg paste server params
func (o *ServerUpdatePkgPasteServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the server update pkg paste server params
func (o *ServerUpdatePkgPasteServerParams) WithContext(ctx context.Context) *ServerUpdatePkgPasteServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the server update pkg paste server params
func (o *ServerUpdatePkgPasteServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the server update pkg paste server params
func (o *ServerUpdatePkgPasteServerParams) WithHTTPClient(client *http.Client) *ServerUpdatePkgPasteServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the server update pkg paste server params
func (o *ServerUpdatePkgPasteServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the server update pkg paste server params
func (o *ServerUpdatePkgPasteServerParams) WithAuthorization(authorization *string) *ServerUpdatePkgPasteServerParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the server update pkg paste server params
func (o *ServerUpdatePkgPasteServerParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithUpdatePkgPasteServerRequestBody adds the updatePkgPasteServerRequestBody to the server update pkg paste server params
func (o *ServerUpdatePkgPasteServerParams) WithUpdatePkgPasteServerRequestBody(updatePkgPasteServerRequestBody *models.ServerUpdatePkgPasteServerRequestBody) *ServerUpdatePkgPasteServerParams {
	o.SetUpdatePkgPasteServerRequestBody(updatePkgPasteServerRequestBody)
	return o
}

// SetUpdatePkgPasteServerRequestBody adds the updatePkgPasteServerRequestBody to the server update pkg paste server params
func (o *ServerUpdatePkgPasteServerParams) SetUpdatePkgPasteServerRequestBody(updatePkgPasteServerRequestBody *models.ServerUpdatePkgPasteServerRequestBody) {
	o.UpdatePkgPasteServerRequestBody = updatePkgPasteServerRequestBody
}

// WithServerID adds the serverID to the server update pkg paste server params
func (o *ServerUpdatePkgPasteServerParams) WithServerID(serverID int64) *ServerUpdatePkgPasteServerParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the server update pkg paste server params
func (o *ServerUpdatePkgPasteServerParams) SetServerID(serverID int64) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *ServerUpdatePkgPasteServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.UpdatePkgPasteServerRequestBody != nil {
		if err := r.SetBodyParam(o.UpdatePkgPasteServerRequestBody); err != nil {
			return err
		}
	}

	// path param serverID
	if err := r.SetPathParam("serverID", swag.FormatInt64(o.ServerID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
