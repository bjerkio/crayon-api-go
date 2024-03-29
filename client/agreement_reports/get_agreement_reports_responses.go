// Code generated by go-swagger; DO NOT EDIT.

package agreement_reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bjerkio/crayon-api-go/models"
)

// GetAgreementReportsReader is a Reader for the GetAgreementReports structure.
type GetAgreementReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAgreementReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAgreementReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAgreementReportsOK creates a GetAgreementReportsOK with default headers values
func NewGetAgreementReportsOK() *GetAgreementReportsOK {
	return &GetAgreementReportsOK{}
}

/*GetAgreementReportsOK handles this case with default header values.

Success
*/
type GetAgreementReportsOK struct {
	Payload *models.APICollectionOfAgreementReport
}

func (o *GetAgreementReportsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/AgreementReports/{productContainerId}][%d] getAgreementReportsOK  %+v", 200, o.Payload)
}

func (o *GetAgreementReportsOK) GetPayload() *models.APICollectionOfAgreementReport {
	return o.Payload
}

func (o *GetAgreementReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APICollectionOfAgreementReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
