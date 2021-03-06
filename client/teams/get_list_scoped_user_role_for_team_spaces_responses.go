// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetListScopedUserRoleForTeamSpacesReader is a Reader for the GetListScopedUserRoleForTeamSpaces structure.
type GetListScopedUserRoleForTeamSpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListScopedUserRoleForTeamSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetListScopedUserRoleForTeamSpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetListScopedUserRoleForTeamSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetListScopedUserRoleForTeamSpacesOK creates a GetListScopedUserRoleForTeamSpacesOK with default headers values
func NewGetListScopedUserRoleForTeamSpacesOK() *GetListScopedUserRoleForTeamSpacesOK {
	return &GetListScopedUserRoleForTeamSpacesOK{}
}

/*GetListScopedUserRoleForTeamSpacesOK handles this case with default header values.

ResourceCollection_of_ScopedUserRoleResource resource returned
*/
type GetListScopedUserRoleForTeamSpacesOK struct {
	Payload *models.ResourceCollectionOfScopedUserRoleResource
}

func (o *GetListScopedUserRoleForTeamSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/teams/{id}/scopeduserroles][%d] getListScopedUserRoleForTeamSpacesOK  %+v", 200, o.Payload)
}

func (o *GetListScopedUserRoleForTeamSpacesOK) GetPayload() *models.ResourceCollectionOfScopedUserRoleResource {
	return o.Payload
}

func (o *GetListScopedUserRoleForTeamSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionOfScopedUserRoleResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListScopedUserRoleForTeamSpacesBadRequest creates a GetListScopedUserRoleForTeamSpacesBadRequest with default headers values
func NewGetListScopedUserRoleForTeamSpacesBadRequest() *GetListScopedUserRoleForTeamSpacesBadRequest {
	return &GetListScopedUserRoleForTeamSpacesBadRequest{}
}

/*GetListScopedUserRoleForTeamSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetListScopedUserRoleForTeamSpacesBadRequest struct {
}

func (o *GetListScopedUserRoleForTeamSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/teams/{id}/scopeduserroles][%d] getListScopedUserRoleForTeamSpacesBadRequest ", 400)
}

func (o *GetListScopedUserRoleForTeamSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
