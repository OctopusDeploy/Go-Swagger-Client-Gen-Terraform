// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteCertificateSpacesReader is a Reader for the DeleteCertificateSpaces structure.
type DeleteCertificateSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCertificateSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCertificateSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCertificateSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCertificateSpacesOK creates a DeleteCertificateSpacesOK with default headers values
func NewDeleteCertificateSpacesOK() *DeleteCertificateSpacesOK {
	return &DeleteCertificateSpacesOK{}
}

/*DeleteCertificateSpacesOK handles this case with default header values.

OK
*/
type DeleteCertificateSpacesOK struct {
}

func (o *DeleteCertificateSpacesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/{baseSpaceId}/certificates/{id}][%d] deleteCertificateSpacesOK ", 200)
}

func (o *DeleteCertificateSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCertificateSpacesBadRequest creates a DeleteCertificateSpacesBadRequest with default headers values
func NewDeleteCertificateSpacesBadRequest() *DeleteCertificateSpacesBadRequest {
	return &DeleteCertificateSpacesBadRequest{}
}

/*DeleteCertificateSpacesBadRequest handles this case with default header values.

BadRequest
No id parameter was provided.
This resource is read-only and cannot be deleted.
*/
type DeleteCertificateSpacesBadRequest struct {
}

func (o *DeleteCertificateSpacesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/{baseSpaceId}/certificates/{id}][%d] deleteCertificateSpacesBadRequest ", 400)
}

func (o *DeleteCertificateSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
