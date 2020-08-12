// Code generated by go-swagger; DO NOT EDIT.

package packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetPackageListNotesSpacesReader is a Reader for the GetPackageListNotesSpaces structure.
type GetPackageListNotesSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackageListNotesSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPackageListNotesSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPackageListNotesSpacesOK creates a GetPackageListNotesSpacesOK with default headers values
func NewGetPackageListNotesSpacesOK() *GetPackageListNotesSpacesOK {
	return &GetPackageListNotesSpacesOK{}
}

/*GetPackageListNotesSpacesOK handles this case with default header values.

PackageNoteListResource resource returned
*/
type GetPackageListNotesSpacesOK struct {
	Payload *models.PackageNoteListResource
}

func (o *GetPackageListNotesSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/packages/notes][%d] getPackageListNotesSpacesOK  %+v", 200, o.Payload)
}

func (o *GetPackageListNotesSpacesOK) GetPayload() *models.PackageNoteListResource {
	return o.Payload
}

func (o *GetPackageListNotesSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PackageNoteListResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}