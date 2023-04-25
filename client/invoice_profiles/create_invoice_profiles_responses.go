// Code generated by go-swagger; DO NOT EDIT.

package invoice_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/crayon-api-go/models"
)

// CreateInvoiceProfilesReader is a Reader for the CreateInvoiceProfiles structure.
type CreateInvoiceProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInvoiceProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateInvoiceProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateInvoiceProfilesOK creates a CreateInvoiceProfilesOK with default headers values
func NewCreateInvoiceProfilesOK() *CreateInvoiceProfilesOK {
	return &CreateInvoiceProfilesOK{}
}

/*CreateInvoiceProfilesOK handles this case with default header values.

Success
*/
type CreateInvoiceProfilesOK struct {
	Payload *models.InvoiceProfile
}

func (o *CreateInvoiceProfilesOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/InvoiceProfiles][%d] createInvoiceProfilesOK  %+v", 200, o.Payload)
}

func (o *CreateInvoiceProfilesOK) GetPayload() *models.InvoiceProfile {
	return o.Payload
}

func (o *CreateInvoiceProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvoiceProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}