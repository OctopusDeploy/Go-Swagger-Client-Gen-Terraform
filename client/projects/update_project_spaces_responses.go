// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// UpdateProjectSpacesReader is a Reader for the UpdateProjectSpaces structure.
type UpdateProjectSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProjectSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProjectSpacesOK creates a UpdateProjectSpacesOK with default headers values
func NewUpdateProjectSpacesOK() *UpdateProjectSpacesOK {
	return &UpdateProjectSpacesOK{}
}

/*UpdateProjectSpacesOK handles this case with default header values.

ProjectResource Modified
*/
type UpdateProjectSpacesOK struct {
	Payload *models.ProjectResource
}

func (o *UpdateProjectSpacesOK) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/projects/{id}][%d] updateProjectSpacesOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectSpacesOK) GetPayload() *models.ProjectResource {
	return o.Payload
}

func (o *UpdateProjectSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectSpacesBadRequest creates a UpdateProjectSpacesBadRequest with default headers values
func NewUpdateProjectSpacesBadRequest() *UpdateProjectSpacesBadRequest {
	return &UpdateProjectSpacesBadRequest{}
}

/*UpdateProjectSpacesBadRequest handles this case with default header values.

A channel ID must be provided to enable Automatic Release Creation.
A package step must be provided to enable Automatic Release Creation.
A project cannot be named '{0}'
Model validation failed.
No id parameter was provided.
No request body was supplied.
Only one of either a donor package step or a version template may be provided.
Please re-enter the password when changing the repository URL or username.
The Automatic Release Creation (ARC) package must be sourced from the Octopus built-in feed. ARC settings can be found at Project > Triggers https://g.octopushq.com/AutoReleaseCreation
The Automatic Release Creation package could not be found (see Automatic Release Creation settings, via Project > Triggers)
The package step provided could not be found in the project's deployment process (see Automatic Release Creation settings, via Project > Triggers).
The package step provided is disabled (see Automatic Release Creation settings, via Project > Triggers).
The specified Automatic Release Creation channel belongs to a different project.
This resource is read-only and cannot be modified.
Unable to connect to the specified Git repository. Please use the test button to validate the connection details.
Unable to convert project to version control: project type not supported
*/
type UpdateProjectSpacesBadRequest struct {
}

func (o *UpdateProjectSpacesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/{baseSpaceId}/projects/{id}][%d] updateProjectSpacesBadRequest ", 400)
}

func (o *UpdateProjectSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
