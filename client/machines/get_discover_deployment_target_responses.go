// Code generated by go-swagger; DO NOT EDIT.

package machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetDiscoverDeploymentTargetReader is a Reader for the GetDiscoverDeploymentTarget structure.
type GetDiscoverDeploymentTargetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDiscoverDeploymentTargetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDiscoverDeploymentTargetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDiscoverDeploymentTargetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDiscoverDeploymentTargetOK creates a GetDiscoverDeploymentTargetOK with default headers values
func NewGetDiscoverDeploymentTargetOK() *GetDiscoverDeploymentTargetOK {
	return &GetDiscoverDeploymentTargetOK{}
}

/*GetDiscoverDeploymentTargetOK handles this case with default header values.

MachineResource resource returned
*/
type GetDiscoverDeploymentTargetOK struct {
	Payload *models.MachineResource
}

func (o *GetDiscoverDeploymentTargetOK) Error() string {
	return fmt.Sprintf("[GET /api/machines/discover][%d] getDiscoverDeploymentTargetOK  %+v", 200, o.Payload)
}

func (o *GetDiscoverDeploymentTargetOK) GetPayload() *models.MachineResource {
	return o.Payload
}

func (o *GetDiscoverDeploymentTargetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoverDeploymentTargetBadRequest creates a GetDiscoverDeploymentTargetBadRequest with default headers values
func NewGetDiscoverDeploymentTargetBadRequest() *GetDiscoverDeploymentTargetBadRequest {
	return &GetDiscoverDeploymentTargetBadRequest{}
}

/*GetDiscoverDeploymentTargetBadRequest handles this case with default header values.

The hostname of the machine to discover must be supplied as URI parameter named 'host'.
There was a controlled failure.
*/
type GetDiscoverDeploymentTargetBadRequest struct {
}

func (o *GetDiscoverDeploymentTargetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/machines/discover][%d] getDiscoverDeploymentTargetBadRequest ", 400)
}

func (o *GetDiscoverDeploymentTargetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}