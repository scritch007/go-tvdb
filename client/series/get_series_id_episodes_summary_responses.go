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

// GetSeriesIDEpisodesSummaryReader is a Reader for the GetSeriesIDEpisodesSummary structure.
type GetSeriesIDEpisodesSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSeriesIDEpisodesSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSeriesIDEpisodesSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetSeriesIDEpisodesSummaryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSeriesIDEpisodesSummaryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSeriesIDEpisodesSummaryOK creates a GetSeriesIDEpisodesSummaryOK with default headers values
func NewGetSeriesIDEpisodesSummaryOK() *GetSeriesIDEpisodesSummaryOK {
	return &GetSeriesIDEpisodesSummaryOK{}
}

/*GetSeriesIDEpisodesSummaryOK handles this case with default header values.

A summary of the episodes and seasons avaialable for the given series.
*/
type GetSeriesIDEpisodesSummaryOK struct {
	Payload *models.SeriesEpisodesSummary
}

func (o *GetSeriesIDEpisodesSummaryOK) Error() string {
	return fmt.Sprintf("[GET /series/{id}/episodes/summary][%d] getSeriesIdEpisodesSummaryOK  %+v", 200, o.Payload)
}

func (o *GetSeriesIDEpisodesSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SeriesEpisodesSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSeriesIDEpisodesSummaryUnauthorized creates a GetSeriesIDEpisodesSummaryUnauthorized with default headers values
func NewGetSeriesIDEpisodesSummaryUnauthorized() *GetSeriesIDEpisodesSummaryUnauthorized {
	return &GetSeriesIDEpisodesSummaryUnauthorized{}
}

/*GetSeriesIDEpisodesSummaryUnauthorized handles this case with default header values.

Returned if your JWT token is missing or expired
*/
type GetSeriesIDEpisodesSummaryUnauthorized struct {
	Payload *models.NotAuthorized
}

func (o *GetSeriesIDEpisodesSummaryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /series/{id}/episodes/summary][%d] getSeriesIdEpisodesSummaryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSeriesIDEpisodesSummaryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotAuthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSeriesIDEpisodesSummaryNotFound creates a GetSeriesIDEpisodesSummaryNotFound with default headers values
func NewGetSeriesIDEpisodesSummaryNotFound() *GetSeriesIDEpisodesSummaryNotFound {
	return &GetSeriesIDEpisodesSummaryNotFound{}
}

/*GetSeriesIDEpisodesSummaryNotFound handles this case with default header values.

Returned if the given series ID does not exist
*/
type GetSeriesIDEpisodesSummaryNotFound struct {
	Payload *models.NotFound
}

func (o *GetSeriesIDEpisodesSummaryNotFound) Error() string {
	return fmt.Sprintf("[GET /series/{id}/episodes/summary][%d] getSeriesIdEpisodesSummaryNotFound  %+v", 404, o.Payload)
}

func (o *GetSeriesIDEpisodesSummaryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
