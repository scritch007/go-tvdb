// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/scritch007/go-tvdb/models"
)

// GetUserRatingsReader is a Reader for the GetUserRatings structure.
type GetUserRatingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserRatingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserRatingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetUserRatingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUserRatingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserRatingsOK creates a GetUserRatingsOK with default headers values
func NewGetUserRatingsOK() *GetUserRatingsOK {
	return &GetUserRatingsOK{}
}

/*GetUserRatingsOK handles this case with default header values.

Array of user ratings.
*/
type GetUserRatingsOK struct {
	Payload *models.UserRatingsData
}

func (o *GetUserRatingsOK) Error() string {
	return fmt.Sprintf("[GET /user/ratings][%d] getUserRatingsOK  %+v", 200, o.Payload)
}

func (o *GetUserRatingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserRatingsData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRatingsUnauthorized creates a GetUserRatingsUnauthorized with default headers values
func NewGetUserRatingsUnauthorized() *GetUserRatingsUnauthorized {
	return &GetUserRatingsUnauthorized{}
}

/*GetUserRatingsUnauthorized handles this case with default header values.

Returned if your JWT token is missing or expired
*/
type GetUserRatingsUnauthorized struct {
	Payload *models.NotAuthorized
}

func (o *GetUserRatingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /user/ratings][%d] getUserRatingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserRatingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotAuthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRatingsNotFound creates a GetUserRatingsNotFound with default headers values
func NewGetUserRatingsNotFound() *GetUserRatingsNotFound {
	return &GetUserRatingsNotFound{}
}

/*GetUserRatingsNotFound handles this case with default header values.

Returned if no information exists for the current user
*/
type GetUserRatingsNotFound struct {
	Payload *models.NotFound
}

func (o *GetUserRatingsNotFound) Error() string {
	return fmt.Sprintf("[GET /user/ratings][%d] getUserRatingsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserRatingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
