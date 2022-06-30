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
)

// NewLockfileGetLockfileDetailParams creates a new LockfileGetLockfileDetailParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLockfileGetLockfileDetailParams() *LockfileGetLockfileDetailParams {
	return &LockfileGetLockfileDetailParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLockfileGetLockfileDetailParamsWithTimeout creates a new LockfileGetLockfileDetailParams object
// with the ability to set a timeout on a request.
func NewLockfileGetLockfileDetailParamsWithTimeout(timeout time.Duration) *LockfileGetLockfileDetailParams {
	return &LockfileGetLockfileDetailParams{
		timeout: timeout,
	}
}

// NewLockfileGetLockfileDetailParamsWithContext creates a new LockfileGetLockfileDetailParams object
// with the ability to set a context for a request.
func NewLockfileGetLockfileDetailParamsWithContext(ctx context.Context) *LockfileGetLockfileDetailParams {
	return &LockfileGetLockfileDetailParams{
		Context: ctx,
	}
}

// NewLockfileGetLockfileDetailParamsWithHTTPClient creates a new LockfileGetLockfileDetailParams object
// with the ability to set a custom HTTPClient for a request.
func NewLockfileGetLockfileDetailParamsWithHTTPClient(client *http.Client) *LockfileGetLockfileDetailParams {
	return &LockfileGetLockfileDetailParams{
		HTTPClient: client,
	}
}

/* LockfileGetLockfileDetailParams contains all the parameters to send to the API endpoint
   for the lockfile get lockfile detail operation.

   Typically these are written to a http.Request.
*/
type LockfileGetLockfileDetailParams struct {

	/* Authorization.

	   api key auth
	*/
	Authorization *string

	/* LockfileID.

	   Lockfile ID

	   Format: int64
	*/
	LockfileID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lockfile get lockfile detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LockfileGetLockfileDetailParams) WithDefaults() *LockfileGetLockfileDetailParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lockfile get lockfile detail params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LockfileGetLockfileDetailParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lockfile get lockfile detail params
func (o *LockfileGetLockfileDetailParams) WithTimeout(timeout time.Duration) *LockfileGetLockfileDetailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lockfile get lockfile detail params
func (o *LockfileGetLockfileDetailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lockfile get lockfile detail params
func (o *LockfileGetLockfileDetailParams) WithContext(ctx context.Context) *LockfileGetLockfileDetailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lockfile get lockfile detail params
func (o *LockfileGetLockfileDetailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lockfile get lockfile detail params
func (o *LockfileGetLockfileDetailParams) WithHTTPClient(client *http.Client) *LockfileGetLockfileDetailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lockfile get lockfile detail params
func (o *LockfileGetLockfileDetailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the lockfile get lockfile detail params
func (o *LockfileGetLockfileDetailParams) WithAuthorization(authorization *string) *LockfileGetLockfileDetailParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the lockfile get lockfile detail params
func (o *LockfileGetLockfileDetailParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithLockfileID adds the lockfileID to the lockfile get lockfile detail params
func (o *LockfileGetLockfileDetailParams) WithLockfileID(lockfileID int64) *LockfileGetLockfileDetailParams {
	o.SetLockfileID(lockfileID)
	return o
}

// SetLockfileID adds the lockfileId to the lockfile get lockfile detail params
func (o *LockfileGetLockfileDetailParams) SetLockfileID(lockfileID int64) {
	o.LockfileID = lockfileID
}

// WriteToRequest writes these params to a swagger request
func (o *LockfileGetLockfileDetailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param lockfileID
	if err := r.SetPathParam("lockfileID", swag.FormatInt64(o.LockfileID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}