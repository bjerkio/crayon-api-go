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

// GetCustomerTenantsReader is a Reader for the GetCustomerTenants structure.
type GetCustomerTenantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomerTenantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomerTenantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCustomerTenantsOK creates a GetCustomerTenantsOK with default headers values
func NewGetCustomerTenantsOK() *GetCustomerTenantsOK {
	return &GetCustomerTenantsOK{}
}

/*GetCustomerTenantsOK handles this case with default header values.

Success
*/
type GetCustomerTenantsOK struct {
	Payload *models.APICollectionOfCustomerTenant
}

func (o *GetCustomerTenantsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/CustomerTenants][%d] getCustomerTenantsOK  %+v", 200, o.Payload)
}

func (o *GetCustomerTenantsOK) GetPayload() *models.APICollectionOfCustomerTenant {
	return o.Payload
}

func (o *GetCustomerTenantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APICollectionOfCustomerTenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
