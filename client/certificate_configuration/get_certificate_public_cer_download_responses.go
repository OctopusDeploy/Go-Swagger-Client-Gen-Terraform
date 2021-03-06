// Code generated by go-swagger; DO NOT EDIT.

package certificate_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetCertificatePublicCerDownloadReader is a Reader for the GetCertificatePublicCerDownload structure.
type GetCertificatePublicCerDownloadReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetCertificatePublicCerDownloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCertificatePublicCerDownloadOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCertificatePublicCerDownloadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCertificatePublicCerDownloadOK creates a GetCertificatePublicCerDownloadOK with default headers values
func NewGetCertificatePublicCerDownloadOK(writer io.Writer) *GetCertificatePublicCerDownloadOK {
	return &GetCertificatePublicCerDownloadOK{
		Payload: writer,
	}
}

/*GetCertificatePublicCerDownloadOK handles this case with default header values.

OK
*/
type GetCertificatePublicCerDownloadOK struct {
	Payload io.Writer
}

func (o *GetCertificatePublicCerDownloadOK) Error() string {
	return fmt.Sprintf("[GET /api/configuration/certificates/{id}/public-cer][%d] getCertificatePublicCerDownloadOK  %+v", 200, o.Payload)
}

func (o *GetCertificatePublicCerDownloadOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetCertificatePublicCerDownloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCertificatePublicCerDownloadBadRequest creates a GetCertificatePublicCerDownloadBadRequest with default headers values
func NewGetCertificatePublicCerDownloadBadRequest() *GetCertificatePublicCerDownloadBadRequest {
	return &GetCertificatePublicCerDownloadBadRequest{}
}

/*GetCertificatePublicCerDownloadBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetCertificatePublicCerDownloadBadRequest struct {
}

func (o *GetCertificatePublicCerDownloadBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/configuration/certificates/{id}/public-cer][%d] getCertificatePublicCerDownloadBadRequest ", 400)
}

func (o *GetCertificatePublicCerDownloadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
