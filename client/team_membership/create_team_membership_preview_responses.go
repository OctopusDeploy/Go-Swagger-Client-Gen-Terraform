// Code generated by go-swagger; DO NOT EDIT.

package team_membership

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// CreateTeamMembershipPreviewReader is a Reader for the CreateTeamMembershipPreview structure.
type CreateTeamMembershipPreviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTeamMembershipPreviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTeamMembershipPreviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTeamMembershipPreviewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTeamMembershipPreviewOK creates a CreateTeamMembershipPreviewOK with default headers values
func NewCreateTeamMembershipPreviewOK() *CreateTeamMembershipPreviewOK {
	return &CreateTeamMembershipPreviewOK{}
}

/*CreateTeamMembershipPreviewOK handles this case with default header values.

TeamMembership[] resource returned
*/
type CreateTeamMembershipPreviewOK struct {
	Payload []*models.TeamMembership
}

func (o *CreateTeamMembershipPreviewOK) Error() string {
	return fmt.Sprintf("[POST /api/teammembership/previewteam][%d] createTeamMembershipPreviewOK  %+v", 200, o.Payload)
}

func (o *CreateTeamMembershipPreviewOK) GetPayload() []*models.TeamMembership {
	return o.Payload
}

func (o *CreateTeamMembershipPreviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamMembershipPreviewBadRequest creates a CreateTeamMembershipPreviewBadRequest with default headers values
func NewCreateTeamMembershipPreviewBadRequest() *CreateTeamMembershipPreviewBadRequest {
	return &CreateTeamMembershipPreviewBadRequest{}
}

/*CreateTeamMembershipPreviewBadRequest handles this case with default header values.

No team resource provided.
*/
type CreateTeamMembershipPreviewBadRequest struct {
}

func (o *CreateTeamMembershipPreviewBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/teammembership/previewteam][%d] createTeamMembershipPreviewBadRequest ", 400)
}

func (o *CreateTeamMembershipPreviewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
