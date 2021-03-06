// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateProjectLogoReader is a Reader for the UpdateProjectLogo structure.
type UpdateProjectLogoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectLogoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectLogoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProjectLogoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProjectLogoOK creates a UpdateProjectLogoOK with default headers values
func NewUpdateProjectLogoOK() *UpdateProjectLogoOK {
	return &UpdateProjectLogoOK{}
}

/*UpdateProjectLogoOK handles this case with default header values.

OK
*/
type UpdateProjectLogoOK struct {
}

func (o *UpdateProjectLogoOK) Error() string {
	return fmt.Sprintf("[PUT /api/projects/{id}/logo][%d] updateProjectLogoOK ", 200)
}

func (o *UpdateProjectLogoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProjectLogoBadRequest creates a UpdateProjectLogoBadRequest with default headers values
func NewUpdateProjectLogoBadRequest() *UpdateProjectLogoBadRequest {
	return &UpdateProjectLogoBadRequest{}
}

/*UpdateProjectLogoBadRequest handles this case with default header values.

Action rejected.
Invalid image provided.
No id parameter was provided.
*/
type UpdateProjectLogoBadRequest struct {
}

func (o *UpdateProjectLogoBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/projects/{id}/logo][%d] updateProjectLogoBadRequest ", 400)
}

func (o *UpdateProjectLogoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
