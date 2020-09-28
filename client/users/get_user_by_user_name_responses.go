// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/crayon-api-go/models"
)

// GetUserByUserNameReader is a Reader for the GetUserByUserName structure.
type GetUserByUserNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserByUserNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserByUserNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserByUserNameOK creates a GetUserByUserNameOK with default headers values
func NewGetUserByUserNameOK() *GetUserByUserNameOK {
	return &GetUserByUserNameOK{}
}

/*GetUserByUserNameOK handles this case with default header values.

Success
*/
type GetUserByUserNameOK struct {
	Payload *models.User
}

func (o *GetUserByUserNameOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/Users/user][%d] getUserByUserNameOK  %+v", 200, o.Payload)
}

func (o *GetUserByUserNameOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetUserByUserNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}