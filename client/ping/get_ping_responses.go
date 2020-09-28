// Code generated by go-swagger; DO NOT EDIT.

package ping

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetPingReader is a Reader for the GetPing structure.
type GetPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPingOK creates a GetPingOK with default headers values
func NewGetPingOK() *GetPingOK {
	return &GetPingOK{}
}

/*GetPingOK handles this case with default header values.

Success
*/
type GetPingOK struct {
}

func (o *GetPingOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/Ping][%d] getPingOK ", 200)
}

func (o *GetPingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
