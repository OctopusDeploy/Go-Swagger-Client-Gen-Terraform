// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetCertificateUsageReader is a Reader for the GetCertificateUsage structure.
type GetCertificateUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCertificateUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCertificateUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCertificateUsageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCertificateUsageOK creates a GetCertificateUsageOK with default headers values
func NewGetCertificateUsageOK() *GetCertificateUsageOK {
	return &GetCertificateUsageOK{}
}

/*GetCertificateUsageOK handles this case with default header values.

CertificateUsageResource resource returned
*/
type GetCertificateUsageOK struct {
	Payload *models.CertificateUsageResource
}

func (o *GetCertificateUsageOK) Error() string {
	return fmt.Sprintf("[GET /api/certificates/{id}/usages][%d] getCertificateUsageOK  %+v", 200, o.Payload)
}

func (o *GetCertificateUsageOK) GetPayload() *models.CertificateUsageResource {
	return o.Payload
}

func (o *GetCertificateUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateUsageResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertificateUsageBadRequest creates a GetCertificateUsageBadRequest with default headers values
func NewGetCertificateUsageBadRequest() *GetCertificateUsageBadRequest {
	return &GetCertificateUsageBadRequest{}
}

/*GetCertificateUsageBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetCertificateUsageBadRequest struct {
}

func (o *GetCertificateUsageBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/certificates/{id}/usages][%d] getCertificateUsageBadRequest ", 400)
}

func (o *GetCertificateUsageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
