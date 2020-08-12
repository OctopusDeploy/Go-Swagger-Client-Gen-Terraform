// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateCertificateArchiveSpacesReader is a Reader for the CreateCertificateArchiveSpaces structure.
type CreateCertificateArchiveSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCertificateArchiveSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCertificateArchiveSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCertificateArchiveSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCertificateArchiveSpacesOK creates a CreateCertificateArchiveSpacesOK with default headers values
func NewCreateCertificateArchiveSpacesOK() *CreateCertificateArchiveSpacesOK {
	return &CreateCertificateArchiveSpacesOK{}
}

/*CreateCertificateArchiveSpacesOK handles this case with default header values.

OK
*/
type CreateCertificateArchiveSpacesOK struct {
}

func (o *CreateCertificateArchiveSpacesOK) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/certificates/{id}/archive][%d] createCertificateArchiveSpacesOK ", 200)
}

func (o *CreateCertificateArchiveSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCertificateArchiveSpacesBadRequest creates a CreateCertificateArchiveSpacesBadRequest with default headers values
func NewCreateCertificateArchiveSpacesBadRequest() *CreateCertificateArchiveSpacesBadRequest {
	return &CreateCertificateArchiveSpacesBadRequest{}
}

/*CreateCertificateArchiveSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type CreateCertificateArchiveSpacesBadRequest struct {
}

func (o *CreateCertificateArchiveSpacesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/certificates/{id}/archive][%d] createCertificateArchiveSpacesBadRequest ", 400)
}

func (o *CreateCertificateArchiveSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}