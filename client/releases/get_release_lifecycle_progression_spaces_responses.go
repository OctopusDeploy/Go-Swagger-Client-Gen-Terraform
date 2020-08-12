// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetReleaseLifecycleProgressionSpacesReader is a Reader for the GetReleaseLifecycleProgressionSpaces structure.
type GetReleaseLifecycleProgressionSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleaseLifecycleProgressionSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReleaseLifecycleProgressionSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReleaseLifecycleProgressionSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReleaseLifecycleProgressionSpacesOK creates a GetReleaseLifecycleProgressionSpacesOK with default headers values
func NewGetReleaseLifecycleProgressionSpacesOK() *GetReleaseLifecycleProgressionSpacesOK {
	return &GetReleaseLifecycleProgressionSpacesOK{}
}

/*GetReleaseLifecycleProgressionSpacesOK handles this case with default header values.

LifecycleProgressionResource resource returned
*/
type GetReleaseLifecycleProgressionSpacesOK struct {
	Payload *models.LifecycleProgressionResource
}

func (o *GetReleaseLifecycleProgressionSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/releases/{id}/progression][%d] getReleaseLifecycleProgressionSpacesOK  %+v", 200, o.Payload)
}

func (o *GetReleaseLifecycleProgressionSpacesOK) GetPayload() *models.LifecycleProgressionResource {
	return o.Payload
}

func (o *GetReleaseLifecycleProgressionSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LifecycleProgressionResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseLifecycleProgressionSpacesBadRequest creates a GetReleaseLifecycleProgressionSpacesBadRequest with default headers values
func NewGetReleaseLifecycleProgressionSpacesBadRequest() *GetReleaseLifecycleProgressionSpacesBadRequest {
	return &GetReleaseLifecycleProgressionSpacesBadRequest{}
}

/*GetReleaseLifecycleProgressionSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetReleaseLifecycleProgressionSpacesBadRequest struct {
}

func (o *GetReleaseLifecycleProgressionSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/releases/{id}/progression][%d] getReleaseLifecycleProgressionSpacesBadRequest ", 400)
}

func (o *GetReleaseLifecycleProgressionSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}