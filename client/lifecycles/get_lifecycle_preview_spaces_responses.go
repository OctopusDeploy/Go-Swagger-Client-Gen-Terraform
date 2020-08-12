// Code generated by go-swagger; DO NOT EDIT.

package lifecycles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetLifecyclePreviewSpacesReader is a Reader for the GetLifecyclePreviewSpaces structure.
type GetLifecyclePreviewSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLifecyclePreviewSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLifecyclePreviewSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLifecyclePreviewSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLifecyclePreviewSpacesOK creates a GetLifecyclePreviewSpacesOK with default headers values
func NewGetLifecyclePreviewSpacesOK() *GetLifecyclePreviewSpacesOK {
	return &GetLifecyclePreviewSpacesOK{}
}

/*GetLifecyclePreviewSpacesOK handles this case with default header values.

LifecycleResource resource returned
*/
type GetLifecyclePreviewSpacesOK struct {
	Payload *models.LifecycleResource
}

func (o *GetLifecyclePreviewSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/lifecycles/{id}/preview][%d] getLifecyclePreviewSpacesOK  %+v", 200, o.Payload)
}

func (o *GetLifecyclePreviewSpacesOK) GetPayload() *models.LifecycleResource {
	return o.Payload
}

func (o *GetLifecyclePreviewSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LifecycleResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLifecyclePreviewSpacesBadRequest creates a GetLifecyclePreviewSpacesBadRequest with default headers values
func NewGetLifecyclePreviewSpacesBadRequest() *GetLifecyclePreviewSpacesBadRequest {
	return &GetLifecyclePreviewSpacesBadRequest{}
}

/*GetLifecyclePreviewSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetLifecyclePreviewSpacesBadRequest struct {
}

func (o *GetLifecyclePreviewSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/lifecycles/{id}/preview][%d] getLifecyclePreviewSpacesBadRequest ", 400)
}

func (o *GetLifecyclePreviewSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}