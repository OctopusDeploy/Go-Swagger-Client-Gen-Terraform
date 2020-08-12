// Code generated by go-swagger; DO NOT EDIT.

package workers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetDiscoverWorkerSpacesReader is a Reader for the GetDiscoverWorkerSpaces structure.
type GetDiscoverWorkerSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDiscoverWorkerSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDiscoverWorkerSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDiscoverWorkerSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDiscoverWorkerSpacesOK creates a GetDiscoverWorkerSpacesOK with default headers values
func NewGetDiscoverWorkerSpacesOK() *GetDiscoverWorkerSpacesOK {
	return &GetDiscoverWorkerSpacesOK{}
}

/*GetDiscoverWorkerSpacesOK handles this case with default header values.

MachineResource resource returned
*/
type GetDiscoverWorkerSpacesOK struct {
	Payload *models.MachineResource
}

func (o *GetDiscoverWorkerSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/workers/discover][%d] getDiscoverWorkerSpacesOK  %+v", 200, o.Payload)
}

func (o *GetDiscoverWorkerSpacesOK) GetPayload() *models.MachineResource {
	return o.Payload
}

func (o *GetDiscoverWorkerSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoverWorkerSpacesBadRequest creates a GetDiscoverWorkerSpacesBadRequest with default headers values
func NewGetDiscoverWorkerSpacesBadRequest() *GetDiscoverWorkerSpacesBadRequest {
	return &GetDiscoverWorkerSpacesBadRequest{}
}

/*GetDiscoverWorkerSpacesBadRequest handles this case with default header values.

The hostname of the machine to discover must be supplied as URI parameter named 'host'.
There was a controlled failure.
*/
type GetDiscoverWorkerSpacesBadRequest struct {
}

func (o *GetDiscoverWorkerSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/workers/discover][%d] getDiscoverWorkerSpacesBadRequest ", 400)
}

func (o *GetDiscoverWorkerSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}