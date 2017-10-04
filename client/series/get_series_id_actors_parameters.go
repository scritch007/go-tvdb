// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSeriesIDActorsParams creates a new GetSeriesIDActorsParams object
// with the default values initialized.
func NewGetSeriesIDActorsParams() *GetSeriesIDActorsParams {
	var ()
	return &GetSeriesIDActorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSeriesIDActorsParamsWithTimeout creates a new GetSeriesIDActorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSeriesIDActorsParamsWithTimeout(timeout time.Duration) *GetSeriesIDActorsParams {
	var ()
	return &GetSeriesIDActorsParams{

		timeout: timeout,
	}
}

// NewGetSeriesIDActorsParamsWithContext creates a new GetSeriesIDActorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSeriesIDActorsParamsWithContext(ctx context.Context) *GetSeriesIDActorsParams {
	var ()
	return &GetSeriesIDActorsParams{

		Context: ctx,
	}
}

// NewGetSeriesIDActorsParamsWithHTTPClient creates a new GetSeriesIDActorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSeriesIDActorsParamsWithHTTPClient(client *http.Client) *GetSeriesIDActorsParams {
	var ()
	return &GetSeriesIDActorsParams{
		HTTPClient: client,
	}
}

/*GetSeriesIDActorsParams contains all the parameters to send to the API endpoint
for the get series ID actors operation typically these are written to a http.Request
*/
type GetSeriesIDActorsParams struct {

	/*ID
	  ID of the series

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get series ID actors params
func (o *GetSeriesIDActorsParams) WithTimeout(timeout time.Duration) *GetSeriesIDActorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get series ID actors params
func (o *GetSeriesIDActorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get series ID actors params
func (o *GetSeriesIDActorsParams) WithContext(ctx context.Context) *GetSeriesIDActorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get series ID actors params
func (o *GetSeriesIDActorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get series ID actors params
func (o *GetSeriesIDActorsParams) WithHTTPClient(client *http.Client) *GetSeriesIDActorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get series ID actors params
func (o *GetSeriesIDActorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get series ID actors params
func (o *GetSeriesIDActorsParams) WithID(id int64) *GetSeriesIDActorsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get series ID actors params
func (o *GetSeriesIDActorsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSeriesIDActorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
