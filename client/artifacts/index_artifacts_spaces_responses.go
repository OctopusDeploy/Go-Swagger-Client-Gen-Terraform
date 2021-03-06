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

// IndexArtifactsSpacesReader is a Reader for the IndexArtifactsSpaces structure.
type IndexArtifactsSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexArtifactsSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexArtifactsSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexArtifactsSpacesOK creates a IndexArtifactsSpacesOK with default headers values
func NewIndexArtifactsSpacesOK() *IndexArtifactsSpacesOK {
	return &IndexArtifactsSpacesOK{}
}

/*IndexArtifactsSpacesOK handles this case with default header values.

ResourceCollection_of_ArtifactResource resource returned
*/
type IndexArtifactsSpacesOK struct {
	Payload *models.ResourceCollectionOfArtifactResource
}

func (o *IndexArtifactsSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/artifacts][%d] indexArtifactsSpacesOK  %+v", 200, o.Payload)
}

func (o *IndexArtifactsSpacesOK) GetPayload() *models.ResourceCollectionOfArtifactResource {
	return o.Payload
}

func (o *IndexArtifactsSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfArtifactResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
