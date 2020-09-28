// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/crayon-api-go/models"
)

// GetOrganizationsReader is a Reader for the GetOrganizations structure.
type GetOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationsOK creates a GetOrganizationsOK with default headers values
func NewGetOrganizationsOK() *GetOrganizationsOK {
	return &GetOrganizationsOK{}
}

/*GetOrganizationsOK handles this case with default header values.

Success
*/
type GetOrganizationsOK struct {
	Payload *models.APICollectionOfOrganization
}

func (o *GetOrganizationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/Organizations][%d] getOrganizationsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsOK) GetPayload() *models.APICollectionOfOrganization {
	return o.Payload
}

func (o *GetOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APICollectionOfOrganization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}