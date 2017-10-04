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

// DeleteUserFavoritesIDReader is a Reader for the DeleteUserFavoritesID structure.
type DeleteUserFavoritesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserFavoritesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteUserFavoritesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteUserFavoritesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserFavoritesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteUserFavoritesIDConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserFavoritesIDOK creates a DeleteUserFavoritesIDOK with default headers values
func NewDeleteUserFavoritesIDOK() *DeleteUserFavoritesIDOK {
	return &DeleteUserFavoritesIDOK{}
}

/*DeleteUserFavoritesIDOK handles this case with default header values.

List of user favorites.
*/
type DeleteUserFavoritesIDOK struct {
	Payload *models.UserFavoritesData
}

func (o *DeleteUserFavoritesIDOK) Error() string {
	return fmt.Sprintf("[DELETE /user/favorites/{id}][%d] deleteUserFavoritesIdOK  %+v", 200, o.Payload)
}

func (o *DeleteUserFavoritesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserFavoritesData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserFavoritesIDUnauthorized creates a DeleteUserFavoritesIDUnauthorized with default headers values
func NewDeleteUserFavoritesIDUnauthorized() *DeleteUserFavoritesIDUnauthorized {
	return &DeleteUserFavoritesIDUnauthorized{}
}

/*DeleteUserFavoritesIDUnauthorized handles this case with default header values.

Returned if your JWT token is missing or expired
*/
type DeleteUserFavoritesIDUnauthorized struct {
	Payload *models.NotAuthorized
}

func (o *DeleteUserFavoritesIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /user/favorites/{id}][%d] deleteUserFavoritesIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUserFavoritesIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotAuthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserFavoritesIDNotFound creates a DeleteUserFavoritesIDNotFound with default headers values
func NewDeleteUserFavoritesIDNotFound() *DeleteUserFavoritesIDNotFound {
	return &DeleteUserFavoritesIDNotFound{}
}

/*DeleteUserFavoritesIDNotFound handles this case with default header values.

Returned if no information exists for the current user
*/
type DeleteUserFavoritesIDNotFound struct {
	Payload *models.NotFound
}

func (o *DeleteUserFavoritesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /user/favorites/{id}][%d] deleteUserFavoritesIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserFavoritesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserFavoritesIDConflict creates a DeleteUserFavoritesIDConflict with default headers values
func NewDeleteUserFavoritesIDConflict() *DeleteUserFavoritesIDConflict {
	return &DeleteUserFavoritesIDConflict{}
}

/*DeleteUserFavoritesIDConflict handles this case with default header values.

Returned if requested record could not be deleted
*/
type DeleteUserFavoritesIDConflict struct {
	Payload *models.Conflict
}

func (o *DeleteUserFavoritesIDConflict) Error() string {
	return fmt.Sprintf("[DELETE /user/favorites/{id}][%d] deleteUserFavoritesIdConflict  %+v", 409, o.Payload)
}

func (o *DeleteUserFavoritesIDConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Conflict)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
