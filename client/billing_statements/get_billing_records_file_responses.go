// Code generated by go-swagger; DO NOT EDIT.

package billing_statements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetBillingRecordsFileReader is a Reader for the GetBillingRecordsFile structure.
type GetBillingRecordsFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBillingRecordsFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBillingRecordsFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBillingRecordsFileOK creates a GetBillingRecordsFileOK with default headers values
func NewGetBillingRecordsFileOK() *GetBillingRecordsFileOK {
	return &GetBillingRecordsFileOK{}
}

/*GetBillingRecordsFileOK handles this case with default header values.

Success
*/
type GetBillingRecordsFileOK struct {
}

func (o *GetBillingRecordsFileOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/BillingStatements/{id}/billingrecordsfile][%d] getBillingRecordsFileOK ", 200)
}

func (o *GetBillingRecordsFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
