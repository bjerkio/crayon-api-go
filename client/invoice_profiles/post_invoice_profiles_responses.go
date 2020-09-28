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

// PostInvoiceProfilesReader is a Reader for the PostInvoiceProfiles structure.
type PostInvoiceProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInvoiceProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostInvoiceProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostInvoiceProfilesOK creates a PostInvoiceProfilesOK with default headers values
func NewPostInvoiceProfilesOK() *PostInvoiceProfilesOK {
	return &PostInvoiceProfilesOK{}
}

/*PostInvoiceProfilesOK handles this case with default header values.

Success
*/
type PostInvoiceProfilesOK struct {
	Payload *models.InvoiceProfile
}

func (o *PostInvoiceProfilesOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/InvoiceProfiles][%d] postInvoiceProfilesOK  %+v", 200, o.Payload)
}

func (o *PostInvoiceProfilesOK) GetPayload() *models.InvoiceProfile {
	return o.Payload
}

func (o *PostInvoiceProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InvoiceProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
