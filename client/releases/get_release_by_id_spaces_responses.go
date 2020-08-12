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

// GetReleaseByIDSpacesReader is a Reader for the GetReleaseByIDSpaces structure.
type GetReleaseByIDSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleaseByIDSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReleaseByIDSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReleaseByIDSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReleaseByIDSpacesOK creates a GetReleaseByIDSpacesOK with default headers values
func NewGetReleaseByIDSpacesOK() *GetReleaseByIDSpacesOK {
	return &GetReleaseByIDSpacesOK{}
}

/*GetReleaseByIDSpacesOK handles this case with default header values.

Load ReleaseResource by ID
*/
type GetReleaseByIDSpacesOK struct {
	Payload *models.ReleaseResource
}

func (o *GetReleaseByIDSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/releases/{id}][%d] getReleaseByIdSpacesOK  %+v", 200, o.Payload)
}

func (o *GetReleaseByIDSpacesOK) GetPayload() *models.ReleaseResource {
	return o.Payload
}

func (o *GetReleaseByIDSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReleaseResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleaseByIDSpacesBadRequest creates a GetReleaseByIDSpacesBadRequest with default headers values
func NewGetReleaseByIDSpacesBadRequest() *GetReleaseByIDSpacesBadRequest {
	return &GetReleaseByIDSpacesBadRequest{}
}

/*GetReleaseByIDSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetReleaseByIDSpacesBadRequest struct {
}

func (o *GetReleaseByIDSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/releases/{id}][%d] getReleaseByIdSpacesBadRequest ", 400)
}

func (o *GetReleaseByIDSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}