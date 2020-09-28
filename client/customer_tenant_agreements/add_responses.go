// Code generated by go-swagger; DO NOT EDIT.

package customer_tenant_agreements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/crayon-api-go/models"
)

// AddReader is a Reader for the Add structure.
type AddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddOK creates a AddOK with default headers values
func NewAddOK() *AddOK {
	return &AddOK{}
}

/*AddOK handles this case with default header values.

Success
*/
type AddOK struct {
	Payload *models.ServiceAccountAgreement
}

func (o *AddOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/customertenants/{customerTenantId}/agreements][%d] addOK  %+v", 200, o.Payload)
}

func (o *AddOK) GetPayload() *models.ServiceAccountAgreement {
	return o.Payload
}

func (o *AddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceAccountAgreement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
