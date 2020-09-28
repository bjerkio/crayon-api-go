// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/crayon-api-go/models"
)

// GetSubscriptionPriceTypesReader is a Reader for the GetSubscriptionPriceTypes structure.
type GetSubscriptionPriceTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionPriceTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubscriptionPriceTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSubscriptionPriceTypesOK creates a GetSubscriptionPriceTypesOK with default headers values
func NewGetSubscriptionPriceTypesOK() *GetSubscriptionPriceTypesOK {
	return &GetSubscriptionPriceTypesOK{}
}

/*GetSubscriptionPriceTypesOK handles this case with default header values.

Success
*/
type GetSubscriptionPriceTypesOK struct {
	Payload *models.APICollectionOfObjectReference
}

func (o *GetSubscriptionPriceTypesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/Subscriptions/subscriptionpricetypes][%d] getSubscriptionPriceTypesOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionPriceTypesOK) GetPayload() *models.APICollectionOfObjectReference {
	return o.Payload
}

func (o *GetSubscriptionPriceTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APICollectionOfObjectReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
