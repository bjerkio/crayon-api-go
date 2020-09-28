// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/crayon-api-go/models"
)

// PutClientsReader is a Reader for the PutClients structure.
type PutClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutClientsOK creates a PutClientsOK with default headers values
func NewPutClientsOK() *PutClientsOK {
	return &PutClientsOK{}
}

/*PutClientsOK handles this case with default header values.

Success
*/
type PutClientsOK struct {
	Payload *models.Client
}

func (o *PutClientsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/Clients/{clientId}][%d] putClientsOK  %+v", 200, o.Payload)
}

func (o *PutClientsOK) GetPayload() *models.Client {
	return o.Payload
}

func (o *PutClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Client)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}