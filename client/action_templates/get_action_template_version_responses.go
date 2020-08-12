// Code generated by go-swagger; DO NOT EDIT.

package action_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetActionTemplateVersionReader is a Reader for the GetActionTemplateVersion structure.
type GetActionTemplateVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetActionTemplateVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetActionTemplateVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetActionTemplateVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetActionTemplateVersionOK creates a GetActionTemplateVersionOK with default headers values
func NewGetActionTemplateVersionOK() *GetActionTemplateVersionOK {
	return &GetActionTemplateVersionOK{}
}

/*GetActionTemplateVersionOK handles this case with default header values.

ActionTemplateResource resource returned
ActionTemplateResource[] resource returned
*/
type GetActionTemplateVersionOK struct {
}

func (o *GetActionTemplateVersionOK) Error() string {
	return fmt.Sprintf("[GET /api/actiontemplates/{id}/versions][%d] getActionTemplateVersionOK ", 200)
}

func (o *GetActionTemplateVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActionTemplateVersionBadRequest creates a GetActionTemplateVersionBadRequest with default headers values
func NewGetActionTemplateVersionBadRequest() *GetActionTemplateVersionBadRequest {
	return &GetActionTemplateVersionBadRequest{}
}

/*GetActionTemplateVersionBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetActionTemplateVersionBadRequest struct {
}

func (o *GetActionTemplateVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/actiontemplates/{id}/versions][%d] getActionTemplateVersionBadRequest ", 400)
}

func (o *GetActionTemplateVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}