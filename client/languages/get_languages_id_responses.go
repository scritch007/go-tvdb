// Code generated by go-swagger; DO NOT EDIT.

package languages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/scritch007/go-tvdb/models"
)

// GetLanguagesIDReader is a Reader for the GetLanguagesID structure.
type GetLanguagesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguagesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLanguagesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetLanguagesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetLanguagesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLanguagesIDOK creates a GetLanguagesIDOK with default headers values
func NewGetLanguagesIDOK() *GetLanguagesIDOK {
	return &GetLanguagesIDOK{}
}

/*GetLanguagesIDOK handles this case with default header values.

An array of language objects.
*/
type GetLanguagesIDOK struct {
	Payload *models.Language
}

func (o *GetLanguagesIDOK) Error() string {
	return fmt.Sprintf("[GET /languages/{id}][%d] getLanguagesIdOK  %+v", 200, o.Payload)
}

func (o *GetLanguagesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Language)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesIDUnauthorized creates a GetLanguagesIDUnauthorized with default headers values
func NewGetLanguagesIDUnauthorized() *GetLanguagesIDUnauthorized {
	return &GetLanguagesIDUnauthorized{}
}

/*GetLanguagesIDUnauthorized handles this case with default header values.

Returned if your JWT token is missing or expired
*/
type GetLanguagesIDUnauthorized struct {
	Payload *models.NotAuthorized
}

func (o *GetLanguagesIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /languages/{id}][%d] getLanguagesIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguagesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotAuthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesIDNotFound creates a GetLanguagesIDNotFound with default headers values
func NewGetLanguagesIDNotFound() *GetLanguagesIDNotFound {
	return &GetLanguagesIDNotFound{}
}

/*GetLanguagesIDNotFound handles this case with default header values.

Returned if the given language ID does not exist.
*/
type GetLanguagesIDNotFound struct {
	Payload *models.NotFound
}

func (o *GetLanguagesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /languages/{id}][%d] getLanguagesIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguagesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}