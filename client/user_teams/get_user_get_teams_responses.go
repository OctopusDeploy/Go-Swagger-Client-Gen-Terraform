// Code generated by go-swagger; DO NOT EDIT.

package user_teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetUserGetTeamsReader is a Reader for the GetUserGetTeams structure.
type GetUserGetTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserGetTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserGetTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserGetTeamsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserGetTeamsOK creates a GetUserGetTeamsOK with default headers values
func NewGetUserGetTeamsOK() *GetUserGetTeamsOK {
	return &GetUserGetTeamsOK{}
}

/*GetUserGetTeamsOK handles this case with default header values.

IEnumerable_of_ProjectedTeamReferenceDataItem resource returned
*/
type GetUserGetTeamsOK struct {
	Payload []*models.ProjectedTeamReferenceDataItem
}

func (o *GetUserGetTeamsOK) Error() string {
	return fmt.Sprintf("[GET /api/users/{id}/teams][%d] getUserGetTeamsOK  %+v", 200, o.Payload)
}

func (o *GetUserGetTeamsOK) GetPayload() []*models.ProjectedTeamReferenceDataItem {
	return o.Payload
}

func (o *GetUserGetTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGetTeamsBadRequest creates a GetUserGetTeamsBadRequest with default headers values
func NewGetUserGetTeamsBadRequest() *GetUserGetTeamsBadRequest {
	return &GetUserGetTeamsBadRequest{}
}

/*GetUserGetTeamsBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetUserGetTeamsBadRequest struct {
}

func (o *GetUserGetTeamsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/users/{id}/teams][%d] getUserGetTeamsBadRequest ", 400)
}

func (o *GetUserGetTeamsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}