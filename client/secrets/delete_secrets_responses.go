// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteSecretsReader is a Reader for the DeleteSecrets structure.
type DeleteSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSecretsOK creates a DeleteSecretsOK with default headers values
func NewDeleteSecretsOK() *DeleteSecretsOK {
	return &DeleteSecretsOK{}
}

/*DeleteSecretsOK handles this case with default header values.

Success
*/
type DeleteSecretsOK struct {
	Payload bool
}

func (o *DeleteSecretsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/Secrets][%d] deleteSecretsOK  %+v", 200, o.Payload)
}

func (o *DeleteSecretsOK) GetPayload() bool {
	return o.Payload
}

func (o *DeleteSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}