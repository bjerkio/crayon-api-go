// Code generated by go-swagger; DO NOT EDIT.

package customer_tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/crayon-api-go/models"
)

// AddExistingReader is a Reader for the AddExisting structure.
type AddExistingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddExistingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddExistingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddExistingOK creates a AddExistingOK with default headers values
func NewAddExistingOK() *AddExistingOK {
	return &AddExistingOK{}
}

/*AddExistingOK handles this case with default header values.

Success
*/
type AddExistingOK struct {
	Payload *models.CustomerTenantDetailed
}

func (o *AddExistingOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/CustomerTenants/existing][%d] addExistingOK  %+v", 200, o.Payload)
}

func (o *AddExistingOK) GetPayload() *models.CustomerTenantDetailed {
	return o.Payload
}

func (o *AddExistingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerTenantDetailed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
