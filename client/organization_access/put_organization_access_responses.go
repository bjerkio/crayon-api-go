// Code generated by go-swagger; DO NOT EDIT.

package organization_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/crayon-api-go/models"
)

// PutOrganizationAccessReader is a Reader for the PutOrganizationAccess structure.
type PutOrganizationAccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrganizationAccessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOrganizationAccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOrganizationAccessOK creates a PutOrganizationAccessOK with default headers values
func NewPutOrganizationAccessOK() *PutOrganizationAccessOK {
	return &PutOrganizationAccessOK{}
}

/*PutOrganizationAccessOK handles this case with default header values.

Success
*/
type PutOrganizationAccessOK struct {
	Payload []*models.OrganizationAccess
}

func (o *PutOrganizationAccessOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/OrganizationAccess][%d] putOrganizationAccessOK  %+v", 200, o.Payload)
}

func (o *PutOrganizationAccessOK) GetPayload() []*models.OrganizationAccess {
	return o.Payload
}

func (o *PutOrganizationAccessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
