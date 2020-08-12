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

// CreateReleaseSpacesReader is a Reader for the CreateReleaseSpaces structure.
type CreateReleaseSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReleaseSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateReleaseSpacesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateReleaseSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateReleaseSpacesCreated creates a CreateReleaseSpacesCreated with default headers values
func NewCreateReleaseSpacesCreated() *CreateReleaseSpacesCreated {
	return &CreateReleaseSpacesCreated{}
}

/*CreateReleaseSpacesCreated handles this case with default header values.

ReleaseResource Created
*/
type CreateReleaseSpacesCreated struct {
	Payload *models.ReleaseResource
}

func (o *CreateReleaseSpacesCreated) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/releases][%d] createReleaseSpacesCreated  %+v", 201, o.Payload)
}

func (o *CreateReleaseSpacesCreated) GetPayload() *models.ReleaseResource {
	return o.Payload
}

func (o *CreateReleaseSpacesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReleaseResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReleaseSpacesBadRequest creates a CreateReleaseSpacesBadRequest with default headers values
func NewCreateReleaseSpacesBadRequest() *CreateReleaseSpacesBadRequest {
	return &CreateReleaseSpacesBadRequest{}
}

/*CreateReleaseSpacesBadRequest handles this case with default header values.

A channelId must be provided for a release when more than one exists for that project and no default channel has been specified.
A version must be specified for every included package.
Model validation failed.
No project ID was specified.
No request body was supplied.
Release '{release.Version}' already exists for this project. Please use a different version, or look at using a mask to auto-increment the number.
The channel {channel.Name} specified for this release belongs to the project {channel.ProjectId}. It does not belong to the same project as the release.
The project is disabled, so releases cannot be created
The release number '{release.Version}' does not appear to be a valid semantic version number. The version must consist of between 1 and 4 number only parts (separated by '.') before the optional pre-release part (after the '-'). e.g. 2, 2.1, 2.4.0.23, 2.4-beta and 1-beta.2
The requested package versions for the following steps violate the channel version rules: {0}. Please specify a different package version, channel or provide the `ignoreChannelRules` parameter to override this check.
*/
type CreateReleaseSpacesBadRequest struct {
}

func (o *CreateReleaseSpacesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/{baseSpaceId}/releases][%d] createReleaseSpacesBadRequest ", 400)
}

func (o *CreateReleaseSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
