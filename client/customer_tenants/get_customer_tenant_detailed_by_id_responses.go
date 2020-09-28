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

// GetCustomerTenantDetailedByIDReader is a Reader for the GetCustomerTenantDetailedByID structure.
type GetCustomerTenantDetailedByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomerTenantDetailedByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomerTenantDetailedByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCustomerTenantDetailedByIDOK creates a GetCustomerTenantDetailedByIDOK with default headers values
func NewGetCustomerTenantDetailedByIDOK() *GetCustomerTenantDetailedByIDOK {
	return &GetCustomerTenantDetailedByIDOK{}
}

/*GetCustomerTenantDetailedByIDOK handles this case with default header values.

Success
*/
type GetCustomerTenantDetailedByIDOK struct {
	Payload *models.CustomerTenantDetailed
}

func (o *GetCustomerTenantDetailedByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/CustomerTenants/{id}/detailed][%d] getCustomerTenantDetailedByIdOK  %+v", 200, o.Payload)
}

func (o *GetCustomerTenantDetailedByIDOK) GetPayload() *models.CustomerTenantDetailed {
	return o.Payload
}

func (o *GetCustomerTenantDetailedByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerTenantDetailed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
