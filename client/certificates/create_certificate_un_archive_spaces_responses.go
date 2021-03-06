// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateCertificateUnArchiveSpacesReader is a Reader for the CreateCertificateUnArchiveSpaces structure.
type CreateCertificateUnArchiveSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCertificateUnArchiveSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCertificateUnArchiveSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCertificateUnArchiveSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCertificateUnArchiveSpacesOK creates a CreateCertificateUnArchiveSpacesOK with default headers values
func NewCreateCertificateUnArchiveSpacesOK() *CreateCertificateUnArchiveSpacesOK {
	return &CreateCertificateUnArchiveSpacesOK{}
}

/*CreateCertificateUnArchiveSpacesOK handles this case with default header values.

OK
*/
type CreateCertificateUnArchiveSpacesOK struct {
}

func (o *CreateCertificateUnArchiveSpacesOK) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/certificates/{id}/unarchive][%d] createCertificateUnArchiveSpacesOK ", 200)
}

func (o *CreateCertificateUnArchiveSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCertificateUnArchiveSpacesBadRequest creates a CreateCertificateUnArchiveSpacesBadRequest with default headers values
func NewCreateCertificateUnArchiveSpacesBadRequest() *CreateCertificateUnArchiveSpacesBadRequest {
	return &CreateCertificateUnArchiveSpacesBadRequest{}
}

/*CreateCertificateUnArchiveSpacesBadRequest handles this case with default header values.

No id parameter was provided.
You cannot unarchive a certificate that has been replaced.
*/
type CreateCertificateUnArchiveSpacesBadRequest struct {
}

func (o *CreateCertificateUnArchiveSpacesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/certificates/{id}/unarchive][%d] createCertificateUnArchiveSpacesBadRequest ", 400)
}

func (o *CreateCertificateUnArchiveSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
