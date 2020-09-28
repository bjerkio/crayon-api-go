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

// GetSubscriptionTagsReader is a Reader for the GetSubscriptionTags structure.
type GetSubscriptionTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubscriptionTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSubscriptionTagsOK creates a GetSubscriptionTagsOK with default headers values
func NewGetSubscriptionTagsOK() *GetSubscriptionTagsOK {
	return &GetSubscriptionTagsOK{}
}

/*GetSubscriptionTagsOK handles this case with default header values.

Success
*/
type GetSubscriptionTagsOK struct {
	Payload *models.SubscriptionTags
}

func (o *GetSubscriptionTagsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/Subscriptions/{subscriptionId}/tags][%d] getSubscriptionTagsOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionTagsOK) GetPayload() *models.SubscriptionTags {
	return o.Payload
}

func (o *GetSubscriptionTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionTags)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}