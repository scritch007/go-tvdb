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

// NewGetSeriesIDFilterParamsParams creates a new GetSeriesIDFilterParamsParams object
// with the default values initialized.
func NewGetSeriesIDFilterParamsParams() *GetSeriesIDFilterParamsParams {
	var ()
	return &GetSeriesIDFilterParamsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSeriesIDFilterParamsParamsWithTimeout creates a new GetSeriesIDFilterParamsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSeriesIDFilterParamsParamsWithTimeout(timeout time.Duration) *GetSeriesIDFilterParamsParams {
	var ()
	return &GetSeriesIDFilterParamsParams{

		timeout: timeout,
	}
}

// NewGetSeriesIDFilterParamsParamsWithContext creates a new GetSeriesIDFilterParamsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSeriesIDFilterParamsParamsWithContext(ctx context.Context) *GetSeriesIDFilterParamsParams {
	var ()
	return &GetSeriesIDFilterParamsParams{

		Context: ctx,
	}
}

// NewGetSeriesIDFilterParamsParamsWithHTTPClient creates a new GetSeriesIDFilterParamsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSeriesIDFilterParamsParamsWithHTTPClient(client *http.Client) *GetSeriesIDFilterParamsParams {
	var ()
	return &GetSeriesIDFilterParamsParams{
		HTTPClient: client,
	}
}

/*GetSeriesIDFilterParamsParams contains all the parameters to send to the API endpoint
for the get series ID filter params operation typically these are written to a http.Request
*/
type GetSeriesIDFilterParamsParams struct {

	/*AcceptLanguage
	  Records are returned with the Episode name and Overview in the desired language, if it exists. If there is no translation for the given language, then the record is still returned but with empty values for the translated fields.

	*/
	AcceptLanguage *string
	/*ID
	  ID of the series

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get series ID filter params params
func (o *GetSeriesIDFilterParamsParams) WithTimeout(timeout time.Duration) *GetSeriesIDFilterParamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get series ID filter params params
func (o *GetSeriesIDFilterParamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get series ID filter params params
func (o *GetSeriesIDFilterParamsParams) WithContext(ctx context.Context) *GetSeriesIDFilterParamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get series ID filter params params
func (o *GetSeriesIDFilterParamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get series ID filter params params
func (o *GetSeriesIDFilterParamsParams) WithHTTPClient(client *http.Client) *GetSeriesIDFilterParamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get series ID filter params params
func (o *GetSeriesIDFilterParamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get series ID filter params params
func (o *GetSeriesIDFilterParamsParams) WithAcceptLanguage(acceptLanguage *string) *GetSeriesIDFilterParamsParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get series ID filter params params
func (o *GetSeriesIDFilterParamsParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithID adds the id to the get series ID filter params params
func (o *GetSeriesIDFilterParamsParams) WithID(id int64) *GetSeriesIDFilterParamsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get series ID filter params params
func (o *GetSeriesIDFilterParamsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSeriesIDFilterParamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptLanguage != nil {

		// header param Accept-Language
		if err := r.SetHeaderParam("Accept-Language", *o.AcceptLanguage); err != nil {
			return err
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
