// Code generated by go-swagger; DO NOT EDIT.

package billing_statements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetReconciliationFileReader is a Reader for the GetReconciliationFile structure.
type GetReconciliationFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReconciliationFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReconciliationFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReconciliationFileOK creates a GetReconciliationFileOK with default headers values
func NewGetReconciliationFileOK() *GetReconciliationFileOK {
	return &GetReconciliationFileOK{}
}

/*GetReconciliationFileOK handles this case with default header values.

Success
*/
type GetReconciliationFileOK struct {
}

func (o *GetReconciliationFileOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/BillingStatements/{id}/reconciliationfile][%d] getReconciliationFileOK ", 200)
}

func (o *GetReconciliationFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
