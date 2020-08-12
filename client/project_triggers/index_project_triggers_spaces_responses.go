// Code generated by go-swagger; DO NOT EDIT.

package project_triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// IndexProjectTriggersSpacesReader is a Reader for the IndexProjectTriggersSpaces structure.
type IndexProjectTriggersSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexProjectTriggersSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndexProjectTriggersSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIndexProjectTriggersSpacesOK creates a IndexProjectTriggersSpacesOK with default headers values
func NewIndexProjectTriggersSpacesOK() *IndexProjectTriggersSpacesOK {
	return &IndexProjectTriggersSpacesOK{}
}

/*IndexProjectTriggersSpacesOK handles this case with default header values.

ResourceCollection_of_ProjectTriggerResource resource returned
*/
type IndexProjectTriggersSpacesOK struct {
	Payload *models.ResourceCollectionOfProjectTriggerResource
}

func (o *IndexProjectTriggersSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/projecttriggers][%d] indexProjectTriggersSpacesOK  %+v", 200, o.Payload)
}

func (o *IndexProjectTriggersSpacesOK) GetPayload() *models.ResourceCollectionOfProjectTriggerResource {
	return o.Payload
}

func (o *IndexProjectTriggersSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfProjectTriggerResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
