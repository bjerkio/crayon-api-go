// Code generated by go-swagger; DO NOT EDIT.

package product_containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/crayon-api-go/models"
)

// GetOrCreateShoppingCartReader is a Reader for the GetOrCreateShoppingCart structure.
type GetOrCreateShoppingCartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrCreateShoppingCartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrCreateShoppingCartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrCreateShoppingCartOK creates a GetOrCreateShoppingCartOK with default headers values
func NewGetOrCreateShoppingCartOK() *GetOrCreateShoppingCartOK {
	return &GetOrCreateShoppingCartOK{}
}

/*GetOrCreateShoppingCartOK handles this case with default header values.

Success
*/
type GetOrCreateShoppingCartOK struct {
	Payload *models.ProductContainer
}

func (o *GetOrCreateShoppingCartOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/ProductContainers/getorcreateshoppingcart][%d] getOrCreateShoppingCartOK  %+v", 200, o.Payload)
}

func (o *GetOrCreateShoppingCartOK) GetPayload() *models.ProductContainer {
	return o.Payload
}

func (o *GetOrCreateShoppingCartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProductContainer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
