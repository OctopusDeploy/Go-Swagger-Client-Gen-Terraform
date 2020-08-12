// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateCertificateArchiveReader is a Reader for the CreateCertificateArchive structure.
type CreateCertificateArchiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCertificateArchiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCertificateArchiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCertificateArchiveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCertificateArchiveOK creates a CreateCertificateArchiveOK with default headers values
func NewCreateCertificateArchiveOK() *CreateCertificateArchiveOK {
	return &CreateCertificateArchiveOK{}
}

/*CreateCertificateArchiveOK handles this case with default header values.

OK
*/
type CreateCertificateArchiveOK struct {
}

func (o *CreateCertificateArchiveOK) Error() string {
	return fmt.Sprintf("[POST /api/certificates/{id}/archive][%d] createCertificateArchiveOK ", 200)
}

func (o *CreateCertificateArchiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCertificateArchiveBadRequest creates a CreateCertificateArchiveBadRequest with default headers values
func NewCreateCertificateArchiveBadRequest() *CreateCertificateArchiveBadRequest {
	return &CreateCertificateArchiveBadRequest{}
}

/*CreateCertificateArchiveBadRequest handles this case with default header values.

No id parameter was provided.
*/
type CreateCertificateArchiveBadRequest struct {
}

func (o *CreateCertificateArchiveBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/certificates/{id}/archive][%d] createCertificateArchiveBadRequest ", 400)
}

func (o *CreateCertificateArchiveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}