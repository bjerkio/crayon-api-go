// Code generated by go-swagger; DO NOT EDIT.

package aws_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/crayon-api-go/models"
)

// GetAwsAccountByIDReader is a Reader for the GetAwsAccountByID structure.
type GetAwsAccountByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAwsAccountByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAwsAccountByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAwsAccountByIDOK creates a GetAwsAccountByIDOK with default headers values
func NewGetAwsAccountByIDOK() *GetAwsAccountByIDOK {
	return &GetAwsAccountByIDOK{}
}

/*GetAwsAccountByIDOK handles this case with default header values.

Success
*/
type GetAwsAccountByIDOK struct {
	Payload *models.AwsAccount
}

func (o *GetAwsAccountByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/AwsAccounts/{id}][%d] getAwsAccountByIdOK  %+v", 200, o.Payload)
}

func (o *GetAwsAccountByIDOK) GetPayload() *models.AwsAccount {
	return o.Payload
}

func (o *GetAwsAccountByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}