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

// GetDeploymentTargetConnectionStatusSpacesReader is a Reader for the GetDeploymentTargetConnectionStatusSpaces structure.
type GetDeploymentTargetConnectionStatusSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentTargetConnectionStatusSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentTargetConnectionStatusSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeploymentTargetConnectionStatusSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentTargetConnectionStatusSpacesOK creates a GetDeploymentTargetConnectionStatusSpacesOK with default headers values
func NewGetDeploymentTargetConnectionStatusSpacesOK() *GetDeploymentTargetConnectionStatusSpacesOK {
	return &GetDeploymentTargetConnectionStatusSpacesOK{}
}

/*GetDeploymentTargetConnectionStatusSpacesOK handles this case with default header values.

MachineConnectionStatus resource returned
*/
type GetDeploymentTargetConnectionStatusSpacesOK struct {
	Payload *models.MachineConnectionStatus
}

func (o *GetDeploymentTargetConnectionStatusSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/machines/{id}/connection][%d] getDeploymentTargetConnectionStatusSpacesOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentTargetConnectionStatusSpacesOK) GetPayload() *models.MachineConnectionStatus {
	return o.Payload
}

func (o *GetDeploymentTargetConnectionStatusSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineConnectionStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentTargetConnectionStatusSpacesBadRequest creates a GetDeploymentTargetConnectionStatusSpacesBadRequest with default headers values
func NewGetDeploymentTargetConnectionStatusSpacesBadRequest() *GetDeploymentTargetConnectionStatusSpacesBadRequest {
	return &GetDeploymentTargetConnectionStatusSpacesBadRequest{}
}

/*GetDeploymentTargetConnectionStatusSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetDeploymentTargetConnectionStatusSpacesBadRequest struct {
}

func (o *GetDeploymentTargetConnectionStatusSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/machines/{id}/connection][%d] getDeploymentTargetConnectionStatusSpacesBadRequest ", 400)
}

func (o *GetDeploymentTargetConnectionStatusSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
