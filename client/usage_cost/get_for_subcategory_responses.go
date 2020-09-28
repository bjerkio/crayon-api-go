// Code generated by go-swagger; DO NOT EDIT.

package usage_cost

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/crayon-api-go/models"
)

// GetForSubcategoryReader is a Reader for the GetForSubcategory structure.
type GetForSubcategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetForSubcategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetForSubcategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetForSubcategoryOK creates a GetForSubcategoryOK with default headers values
func NewGetForSubcategoryOK() *GetForSubcategoryOK {
	return &GetForSubcategoryOK{}
}

/*GetForSubcategoryOK handles this case with default header values.

Success
*/
type GetForSubcategoryOK struct {
	Payload *models.APICollectionOfSubcategoryUsageCost
}

func (o *GetForSubcategoryOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/UsageCost/resellerCustomer/{resellerCustomerId}/subscription/{subscriptionId}/category/{category}/subcategory/{subcategory}/currency/{currencyCode}][%d] getForSubcategoryOK  %+v", 200, o.Payload)
}

func (o *GetForSubcategoryOK) GetPayload() *models.APICollectionOfSubcategoryUsageCost {
	return o.Payload
}

func (o *GetForSubcategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APICollectionOfSubcategoryUsageCost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
