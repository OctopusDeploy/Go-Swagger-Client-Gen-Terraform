// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetProjectLogoSpacesReader is a Reader for the GetProjectLogoSpaces structure.
type GetProjectLogoSpacesReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectLogoSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectLogoSpacesOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetProjectLogoSpacesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetProjectLogoSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectLogoSpacesOK creates a GetProjectLogoSpacesOK with default headers values
func NewGetProjectLogoSpacesOK(writer io.Writer) *GetProjectLogoSpacesOK {
	return &GetProjectLogoSpacesOK{
		Payload: writer,
	}
}

/*GetProjectLogoSpacesOK handles this case with default header values.

OK
*/
type GetProjectLogoSpacesOK struct {
	Payload io.Writer
}

func (o *GetProjectLogoSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/projects/{id}/logo][%d] getProjectLogoSpacesOK  %+v", 200, o.Payload)
}

func (o *GetProjectLogoSpacesOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetProjectLogoSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectLogoSpacesNotModified creates a GetProjectLogoSpacesNotModified with default headers values
func NewGetProjectLogoSpacesNotModified() *GetProjectLogoSpacesNotModified {
	return &GetProjectLogoSpacesNotModified{}
}

/*GetProjectLogoSpacesNotModified handles this case with default header values.

NotModified
*/
type GetProjectLogoSpacesNotModified struct {
}

func (o *GetProjectLogoSpacesNotModified) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/projects/{id}/logo][%d] getProjectLogoSpacesNotModified ", 304)
}

func (o *GetProjectLogoSpacesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectLogoSpacesBadRequest creates a GetProjectLogoSpacesBadRequest with default headers values
func NewGetProjectLogoSpacesBadRequest() *GetProjectLogoSpacesBadRequest {
	return &GetProjectLogoSpacesBadRequest{}
}

/*GetProjectLogoSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetProjectLogoSpacesBadRequest struct {
}

func (o *GetProjectLogoSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/projects/{id}/logo][%d] getProjectLogoSpacesBadRequest ", 400)
}

func (o *GetProjectLogoSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
