// Code generated by go-swagger; DO NOT EDIT.

package groupings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/crayon-api-go/models"
)

// CreateGroupingsReader is a Reader for the CreateGroupings structure.
type CreateGroupingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGroupingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateGroupingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateGroupingsOK creates a CreateGroupingsOK with default headers values
func NewCreateGroupingsOK() *CreateGroupingsOK {
	return &CreateGroupingsOK{}
}

/*CreateGroupingsOK handles this case with default header values.

Success
*/
type CreateGroupingsOK struct {
	Payload *models.Grouping
}

func (o *CreateGroupingsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/Groupings][%d] createGroupingsOK  %+v", 200, o.Payload)
}

func (o *CreateGroupingsOK) GetPayload() *models.Grouping {
	return o.Payload
}

func (o *CreateGroupingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Grouping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
