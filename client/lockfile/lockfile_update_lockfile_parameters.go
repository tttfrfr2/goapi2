// Code generated by go-swagger; DO NOT EDIT.

package lockfile

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

// NewLockfileUpdateLockfileParams creates a new LockfileUpdateLockfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLockfileUpdateLockfileParams() *LockfileUpdateLockfileParams {
	return &LockfileUpdateLockfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLockfileUpdateLockfileParamsWithTimeout creates a new LockfileUpdateLockfileParams object
// with the ability to set a timeout on a request.
func NewLockfileUpdateLockfileParamsWithTimeout(timeout time.Duration) *LockfileUpdateLockfileParams {
	return &LockfileUpdateLockfileParams{
		timeout: timeout,
	}
}

// NewLockfileUpdateLockfileParamsWithContext creates a new LockfileUpdateLockfileParams object
// with the ability to set a context for a request.
func NewLockfileUpdateLockfileParamsWithContext(ctx context.Context) *LockfileUpdateLockfileParams {
	return &LockfileUpdateLockfileParams{
		Context: ctx,
	}
}

// NewLockfileUpdateLockfileParamsWithHTTPClient creates a new LockfileUpdateLockfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewLockfileUpdateLockfileParamsWithHTTPClient(client *http.Client) *LockfileUpdateLockfileParams {
	return &LockfileUpdateLockfileParams{
		HTTPClient: client,
	}
}

/* LockfileUpdateLockfileParams contains all the parameters to send to the API endpoint
   for the lockfile update lockfile operation.

   Typically these are written to a http.Request.
*/
type LockfileUpdateLockfileParams struct {

	/* Authorization.

	   api key auth
	*/
	Authorization *string

	// UpdateLockfileRequestBody.
	UpdateLockfileRequestBody *models.LockfileUpdateLockfileRequestBody

	/* LockfileID.

	   Lockfile ID

	   Format: int64
	*/
	LockfileID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lockfile update lockfile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LockfileUpdateLockfileParams) WithDefaults() *LockfileUpdateLockfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lockfile update lockfile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LockfileUpdateLockfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lockfile update lockfile params
func (o *LockfileUpdateLockfileParams) WithTimeout(timeout time.Duration) *LockfileUpdateLockfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lockfile update lockfile params
func (o *LockfileUpdateLockfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lockfile update lockfile params
func (o *LockfileUpdateLockfileParams) WithContext(ctx context.Context) *LockfileUpdateLockfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lockfile update lockfile params
func (o *LockfileUpdateLockfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lockfile update lockfile params
func (o *LockfileUpdateLockfileParams) WithHTTPClient(client *http.Client) *LockfileUpdateLockfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lockfile update lockfile params
func (o *LockfileUpdateLockfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the lockfile update lockfile params
func (o *LockfileUpdateLockfileParams) WithAuthorization(authorization *string) *LockfileUpdateLockfileParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the lockfile update lockfile params
func (o *LockfileUpdateLockfileParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithUpdateLockfileRequestBody adds the updateLockfileRequestBody to the lockfile update lockfile params
func (o *LockfileUpdateLockfileParams) WithUpdateLockfileRequestBody(updateLockfileRequestBody *models.LockfileUpdateLockfileRequestBody) *LockfileUpdateLockfileParams {
	o.SetUpdateLockfileRequestBody(updateLockfileRequestBody)
	return o
}

// SetUpdateLockfileRequestBody adds the updateLockfileRequestBody to the lockfile update lockfile params
func (o *LockfileUpdateLockfileParams) SetUpdateLockfileRequestBody(updateLockfileRequestBody *models.LockfileUpdateLockfileRequestBody) {
	o.UpdateLockfileRequestBody = updateLockfileRequestBody
}

// WithLockfileID adds the lockfileID to the lockfile update lockfile params
func (o *LockfileUpdateLockfileParams) WithLockfileID(lockfileID int64) *LockfileUpdateLockfileParams {
	o.SetLockfileID(lockfileID)
	return o
}

// SetLockfileID adds the lockfileId to the lockfile update lockfile params
func (o *LockfileUpdateLockfileParams) SetLockfileID(lockfileID int64) {
	o.LockfileID = lockfileID
}

// WriteToRequest writes these params to a swagger request
func (o *LockfileUpdateLockfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.UpdateLockfileRequestBody != nil {
		if err := r.SetBodyParam(o.UpdateLockfileRequestBody); err != nil {
			return err
		}
	}

	// path param lockfileID
	if err := r.SetPathParam("lockfileID", swag.FormatInt64(o.LockfileID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
