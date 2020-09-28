// Code generated by go-swagger; DO NOT EDIT.

package regions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/crayon-api-go/models"
)

// GetRegionsReader is a Reader for the GetRegions structure.
type GetRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRegionsOK creates a GetRegionsOK with default headers values
func NewGetRegionsOK() *GetRegionsOK {
	return &GetRegionsOK{}
}

/*GetRegionsOK handles this case with default header values.

Success
*/
type GetRegionsOK struct {
	Payload *models.APICollectionOfRegion
}

func (o *GetRegionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/Regions][%d] getRegionsOK  %+v", 200, o.Payload)
}

func (o *GetRegionsOK) GetPayload() *models.APICollectionOfRegion {
	return o.Payload
}

func (o *GetRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APICollectionOfRegion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}