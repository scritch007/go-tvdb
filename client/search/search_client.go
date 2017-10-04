// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new search API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for search API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetSearchSeries Allows the user to search for a series based on the following parameters.
*/
func (a *Client) GetSearchSeries(params *GetSearchSeriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetSearchSeriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchSeriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchSeries",
		Method:             "GET",
		PathPattern:        "/search/series",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSearchSeriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchSeriesOK), nil

}

/*
GetSearchSeriesParams Returns an array of parameters to query by in the `/search/series` route.
*/
func (a *Client) GetSearchSeriesParams(params *GetSearchSeriesParamsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSearchSeriesParamsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchSeriesParamsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchSeriesParams",
		Method:             "GET",
		PathPattern:        "/search/series/params",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSearchSeriesParamsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchSeriesParamsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
