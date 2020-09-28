// Code generated by go-swagger; DO NOT EDIT.

package clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteClientsReader is a Reader for the DeleteClients structure.
type DeleteClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteClientsOK creates a DeleteClientsOK with default headers values
func NewDeleteClientsOK() *DeleteClientsOK {
	return &DeleteClientsOK{}
}

/*DeleteClientsOK handles this case with default header values.

Success
*/
type DeleteClientsOK struct {
	Payload bool
}

func (o *DeleteClientsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/Clients/{clientId}][%d] deleteClientsOK  %+v", 200, o.Payload)
}

func (o *DeleteClientsOK) GetPayload() bool {
	return o.Payload
}

func (o *DeleteClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
