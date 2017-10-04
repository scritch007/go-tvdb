// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/scritch007/go-tvdb/models"
)

// GetSeriesIDFilterParamsReader is a Reader for the GetSeriesIDFilterParams structure.
type GetSeriesIDFilterParamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSeriesIDFilterParamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSeriesIDFilterParamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetSeriesIDFilterParamsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSeriesIDFilterParamsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSeriesIDFilterParamsOK creates a GetSeriesIDFilterParamsOK with default headers values
func NewGetSeriesIDFilterParamsOK() *GetSeriesIDFilterParamsOK {
	return &GetSeriesIDFilterParamsOK{}
}

/*GetSeriesIDFilterParamsOK handles this case with default header values.

A list of keys to filter by
*/
type GetSeriesIDFilterParamsOK struct {
	Payload *models.FilterKeys
}

func (o *GetSeriesIDFilterParamsOK) Error() string {
	return fmt.Sprintf("[GET /series/{id}/filter/params][%d] getSeriesIdFilterParamsOK  %+v", 200, o.Payload)
}

func (o *GetSeriesIDFilterParamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FilterKeys)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSeriesIDFilterParamsUnauthorized creates a GetSeriesIDFilterParamsUnauthorized with default headers values
func NewGetSeriesIDFilterParamsUnauthorized() *GetSeriesIDFilterParamsUnauthorized {
	return &GetSeriesIDFilterParamsUnauthorized{}
}

/*GetSeriesIDFilterParamsUnauthorized handles this case with default header values.

Returned if your JWT token is missing or expired
*/
type GetSeriesIDFilterParamsUnauthorized struct {
	Payload *models.NotAuthorized
}

func (o *GetSeriesIDFilterParamsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /series/{id}/filter/params][%d] getSeriesIdFilterParamsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSeriesIDFilterParamsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotAuthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSeriesIDFilterParamsNotFound creates a GetSeriesIDFilterParamsNotFound with default headers values
func NewGetSeriesIDFilterParamsNotFound() *GetSeriesIDFilterParamsNotFound {
	return &GetSeriesIDFilterParamsNotFound{}
}

/*GetSeriesIDFilterParamsNotFound handles this case with default header values.

Returned if the given series ID does not exist
*/
type GetSeriesIDFilterParamsNotFound struct {
	Payload *models.NotFound
}

func (o *GetSeriesIDFilterParamsNotFound) Error() string {
	return fmt.Sprintf("[GET /series/{id}/filter/params][%d] getSeriesIdFilterParamsNotFound  %+v", 404, o.Payload)
}

func (o *GetSeriesIDFilterParamsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
