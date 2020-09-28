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

// PostAPIV1UsageCostGetForSubcategoryReader is a Reader for the PostAPIV1UsageCostGetForSubcategory structure.
type PostAPIV1UsageCostGetForSubcategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV1UsageCostGetForSubcategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAPIV1UsageCostGetForSubcategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAPIV1UsageCostGetForSubcategoryOK creates a PostAPIV1UsageCostGetForSubcategoryOK with default headers values
func NewPostAPIV1UsageCostGetForSubcategoryOK() *PostAPIV1UsageCostGetForSubcategoryOK {
	return &PostAPIV1UsageCostGetForSubcategoryOK{}
}

/*PostAPIV1UsageCostGetForSubcategoryOK handles this case with default header values.

Success
*/
type PostAPIV1UsageCostGetForSubcategoryOK struct {
	Payload *models.APICollectionOfSubcategoryUsageCost
}

func (o *PostAPIV1UsageCostGetForSubcategoryOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/UsageCost/getForSubcategory][%d] postApiV1UsageCostGetForSubcategoryOK  %+v", 200, o.Payload)
}

func (o *PostAPIV1UsageCostGetForSubcategoryOK) GetPayload() *models.APICollectionOfSubcategoryUsageCost {
	return o.Payload
}

func (o *PostAPIV1UsageCostGetForSubcategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APICollectionOfSubcategoryUsageCost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}