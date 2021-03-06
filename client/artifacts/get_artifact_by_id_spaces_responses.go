// Code generated by go-swagger; DO NOT EDIT.

package artifacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetArtifactByIDSpacesReader is a Reader for the GetArtifactByIDSpaces structure.
type GetArtifactByIDSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArtifactByIDSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArtifactByIDSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArtifactByIDSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArtifactByIDSpacesOK creates a GetArtifactByIDSpacesOK with default headers values
func NewGetArtifactByIDSpacesOK() *GetArtifactByIDSpacesOK {
	return &GetArtifactByIDSpacesOK{}
}

/*GetArtifactByIDSpacesOK handles this case with default header values.

Load ArtifactResource by ID
*/
type GetArtifactByIDSpacesOK struct {
	Payload *models.ArtifactResource
}

func (o *GetArtifactByIDSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/artifacts/{id}][%d] getArtifactByIdSpacesOK  %+v", 200, o.Payload)
}

func (o *GetArtifactByIDSpacesOK) GetPayload() *models.ArtifactResource {
	return o.Payload
}

func (o *GetArtifactByIDSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ArtifactResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactByIDSpacesBadRequest creates a GetArtifactByIDSpacesBadRequest with default headers values
func NewGetArtifactByIDSpacesBadRequest() *GetArtifactByIDSpacesBadRequest {
	return &GetArtifactByIDSpacesBadRequest{}
}

/*GetArtifactByIDSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetArtifactByIDSpacesBadRequest struct {
}

func (o *GetArtifactByIDSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/artifacts/{id}][%d] getArtifactByIdSpacesBadRequest ", 400)
}

func (o *GetArtifactByIDSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
