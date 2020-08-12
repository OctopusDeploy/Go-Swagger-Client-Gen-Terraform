// Code generated by go-swagger; DO NOT EDIT.

package lifecycles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetLifecycleProjectsReader is a Reader for the GetLifecycleProjects structure.
type GetLifecycleProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLifecycleProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLifecycleProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLifecycleProjectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLifecycleProjectsOK creates a GetLifecycleProjectsOK with default headers values
func NewGetLifecycleProjectsOK() *GetLifecycleProjectsOK {
	return &GetLifecycleProjectsOK{}
}

/*GetLifecycleProjectsOK handles this case with default header values.

IEnumerable_of_ProjectResource resource returned
*/
type GetLifecycleProjectsOK struct {
	Payload []*models.ProjectResource
}

func (o *GetLifecycleProjectsOK) Error() string {
	return fmt.Sprintf("[GET /api/lifecycles/{id}/projects][%d] getLifecycleProjectsOK  %+v", 200, o.Payload)
}

func (o *GetLifecycleProjectsOK) GetPayload() []*models.ProjectResource {
	return o.Payload
}

func (o *GetLifecycleProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLifecycleProjectsBadRequest creates a GetLifecycleProjectsBadRequest with default headers values
func NewGetLifecycleProjectsBadRequest() *GetLifecycleProjectsBadRequest {
	return &GetLifecycleProjectsBadRequest{}
}

/*GetLifecycleProjectsBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetLifecycleProjectsBadRequest struct {
}

func (o *GetLifecycleProjectsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/lifecycles/{id}/projects][%d] getLifecycleProjectsBadRequest ", 400)
}

func (o *GetLifecycleProjectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}