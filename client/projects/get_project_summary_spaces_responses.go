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

// GetProjectSummarySpacesReader is a Reader for the GetProjectSummarySpaces structure.
type GetProjectSummarySpacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectSummarySpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectSummarySpacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectSummarySpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectSummarySpacesOK creates a GetProjectSummarySpacesOK with default headers values
func NewGetProjectSummarySpacesOK() *GetProjectSummarySpacesOK {
	return &GetProjectSummarySpacesOK{}
}

/*GetProjectSummarySpacesOK handles this case with default header values.

ProjectSummary resource returned
*/
type GetProjectSummarySpacesOK struct {
	Payload *models.ProjectSummary
}

func (o *GetProjectSummarySpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/projects/{id}/summary][%d] getProjectSummarySpacesOK  %+v", 200, o.Payload)
}

func (o *GetProjectSummarySpacesOK) GetPayload() *models.ProjectSummary {
	return o.Payload
}

func (o *GetProjectSummarySpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSummarySpacesBadRequest creates a GetProjectSummarySpacesBadRequest with default headers values
func NewGetProjectSummarySpacesBadRequest() *GetProjectSummarySpacesBadRequest {
	return &GetProjectSummarySpacesBadRequest{}
}

/*GetProjectSummarySpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetProjectSummarySpacesBadRequest struct {
}

func (o *GetProjectSummarySpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/projects/{id}/summary][%d] getProjectSummarySpacesBadRequest ", 400)
}

func (o *GetProjectSummarySpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}