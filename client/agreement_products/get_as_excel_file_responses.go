// Code generated by go-swagger; DO NOT EDIT.

package agreement_products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAsExcelFileReader is a Reader for the GetAsExcelFile structure.
type GetAsExcelFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAsExcelFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAsExcelFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAsExcelFileOK creates a GetAsExcelFileOK with default headers values
func NewGetAsExcelFileOK() *GetAsExcelFileOK {
	return &GetAsExcelFileOK{}
}

/*GetAsExcelFileOK handles this case with default header values.

Success
*/
type GetAsExcelFileOK struct {
}

func (o *GetAsExcelFileOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/AgreementProducts/file/xlsx][%d] getAsExcelFileOK ", 200)
}

func (o *GetAsExcelFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
