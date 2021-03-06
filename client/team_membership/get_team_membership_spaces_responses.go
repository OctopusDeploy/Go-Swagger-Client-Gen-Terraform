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

// GetTeamMembershipSpacesReader is a Reader for the GetTeamMembershipSpaces structure.
type GetTeamMembershipSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamMembershipSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamMembershipSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTeamMembershipSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTeamMembershipSpacesOK creates a GetTeamMembershipSpacesOK with default headers values
func NewGetTeamMembershipSpacesOK() *GetTeamMembershipSpacesOK {
	return &GetTeamMembershipSpacesOK{}
}

/*GetTeamMembershipSpacesOK handles this case with default header values.

IEnumerable_of_TeamMembership resource returned
*/
type GetTeamMembershipSpacesOK struct {
	Payload []*models.TeamMembership
}

func (o *GetTeamMembershipSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/teammembership][%d] getTeamMembershipSpacesOK  %+v", 200, o.Payload)
}

func (o *GetTeamMembershipSpacesOK) GetPayload() []*models.TeamMembership {
	return o.Payload
}

func (o *GetTeamMembershipSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembershipSpacesBadRequest creates a GetTeamMembershipSpacesBadRequest with default headers values
func NewGetTeamMembershipSpacesBadRequest() *GetTeamMembershipSpacesBadRequest {
	return &GetTeamMembershipSpacesBadRequest{}
}

/*GetTeamMembershipSpacesBadRequest handles this case with default header values.

No userId parameter was provided.
*/
type GetTeamMembershipSpacesBadRequest struct {
}

func (o *GetTeamMembershipSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/teammembership][%d] getTeamMembershipSpacesBadRequest ", 400)
}

func (o *GetTeamMembershipSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
