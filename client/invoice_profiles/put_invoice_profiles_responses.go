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

// PutInvoiceProfilesReader is a Reader for the PutInvoiceProfiles structure.
type PutInvoiceProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutInvoiceProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutInvoiceProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutInvoiceProfilesOK creates a PutInvoiceProfilesOK with default headers values
func NewPutInvoiceProfilesOK() *PutInvoiceProfilesOK {
	return &PutInvoiceProfilesOK{}
}

/*PutInvoiceProfilesOK handles this case with default header values.

Success
*/
type PutInvoiceProfilesOK struct {
	Payload *models.InvoiceProfile
}

func (o *PutInvoiceProfilesOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/InvoiceProfiles/{id}][%d] putInvoiceProfilesOK  %+v", 200, o.Payload)
}

func (o *PutInvoiceProfilesOK) GetPayload() *models.InvoiceProfile {
	return o.Payload
}

func (o *PutInvoiceProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvoiceProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
