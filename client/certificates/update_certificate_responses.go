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

// UpdateCertificateReader is a Reader for the UpdateCertificate structure.
type UpdateCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCertificateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCertificateOK creates a UpdateCertificateOK with default headers values
func NewUpdateCertificateOK() *UpdateCertificateOK {
	return &UpdateCertificateOK{}
}

/*UpdateCertificateOK handles this case with default header values.

CertificateResource Modified
*/
type UpdateCertificateOK struct {
	Payload *models.CertificateResource
}

func (o *UpdateCertificateOK) Error() string {
	return fmt.Sprintf("[PUT /api/certificates/{id}][%d] updateCertificateOK  %+v", 200, o.Payload)
}

func (o *UpdateCertificateOK) GetPayload() *models.CertificateResource {
	return o.Payload
}

func (o *UpdateCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCertificateBadRequest creates a UpdateCertificateBadRequest with default headers values
func NewUpdateCertificateBadRequest() *UpdateCertificateBadRequest {
	return &UpdateCertificateBadRequest{}
}

/*UpdateCertificateBadRequest handles this case with default header values.

Model validation failed.
No id parameter was provided.
No request body was supplied.
This resource is read-only and cannot be modified.
*/
type UpdateCertificateBadRequest struct {
}

func (o *UpdateCertificateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/certificates/{id}][%d] updateCertificateBadRequest ", 400)
}

func (o *UpdateCertificateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
