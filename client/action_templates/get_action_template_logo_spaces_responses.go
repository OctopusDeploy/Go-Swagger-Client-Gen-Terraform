// Code generated by go-swagger; DO NOT EDIT.

package action_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetActionTemplateLogoSpacesReader is a Reader for the GetActionTemplateLogoSpaces structure.
type GetActionTemplateLogoSpacesReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetActionTemplateLogoSpacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetActionTemplateLogoSpacesOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetActionTemplateLogoSpacesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetActionTemplateLogoSpacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetActionTemplateLogoSpacesOK creates a GetActionTemplateLogoSpacesOK with default headers values
func NewGetActionTemplateLogoSpacesOK(writer io.Writer) *GetActionTemplateLogoSpacesOK {
	return &GetActionTemplateLogoSpacesOK{
		Payload: writer,
	}
}

/*GetActionTemplateLogoSpacesOK handles this case with default header values.

OK
*/
type GetActionTemplateLogoSpacesOK struct {
	Payload io.Writer
}

func (o *GetActionTemplateLogoSpacesOK) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/actiontemplates/{id}/logo][%d] getActionTemplateLogoSpacesOK  %+v", 200, o.Payload)
}

func (o *GetActionTemplateLogoSpacesOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetActionTemplateLogoSpacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetActionTemplateLogoSpacesNotModified creates a GetActionTemplateLogoSpacesNotModified with default headers values
func NewGetActionTemplateLogoSpacesNotModified() *GetActionTemplateLogoSpacesNotModified {
	return &GetActionTemplateLogoSpacesNotModified{}
}

/*GetActionTemplateLogoSpacesNotModified handles this case with default header values.

NotModified
*/
type GetActionTemplateLogoSpacesNotModified struct {
}

func (o *GetActionTemplateLogoSpacesNotModified) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/actiontemplates/{id}/logo][%d] getActionTemplateLogoSpacesNotModified ", 304)
}

func (o *GetActionTemplateLogoSpacesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetActionTemplateLogoSpacesBadRequest creates a GetActionTemplateLogoSpacesBadRequest with default headers values
func NewGetActionTemplateLogoSpacesBadRequest() *GetActionTemplateLogoSpacesBadRequest {
	return &GetActionTemplateLogoSpacesBadRequest{}
}

/*GetActionTemplateLogoSpacesBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetActionTemplateLogoSpacesBadRequest struct {
}

func (o *GetActionTemplateLogoSpacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/{baseSpaceId}/actiontemplates/{id}/logo][%d] getActionTemplateLogoSpacesBadRequest ", 400)
}

func (o *GetActionTemplateLogoSpacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
