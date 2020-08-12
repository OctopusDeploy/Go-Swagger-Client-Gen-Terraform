// Code generated by go-swagger; DO NOT EDIT.

package artifacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetArtifactContentReader is a Reader for the GetArtifactContent structure.
type GetArtifactContentReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetArtifactContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArtifactContentOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetArtifactContentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArtifactContentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetArtifactContentMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArtifactContentOK creates a GetArtifactContentOK with default headers values
func NewGetArtifactContentOK(writer io.Writer) *GetArtifactContentOK {
	return &GetArtifactContentOK{
		Payload: writer,
	}
}

/*GetArtifactContentOK handles this case with default header values.

OK
*/
type GetArtifactContentOK struct {
	Payload io.Writer
}

func (o *GetArtifactContentOK) Error() string {
	return fmt.Sprintf("[GET /api/artifacts/{id}/content][%d] getArtifactContentOK  %+v", 200, o.Payload)
}

func (o *GetArtifactContentOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetArtifactContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactContentNoContent creates a GetArtifactContentNoContent with default headers values
func NewGetArtifactContentNoContent() *GetArtifactContentNoContent {
	return &GetArtifactContentNoContent{}
}

/*GetArtifactContentNoContent handles this case with default header values.

NoContent
*/
type GetArtifactContentNoContent struct {
}

func (o *GetArtifactContentNoContent) Error() string {
	return fmt.Sprintf("[GET /api/artifacts/{id}/content][%d] getArtifactContentNoContent ", 204)
}

func (o *GetArtifactContentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetArtifactContentBadRequest creates a GetArtifactContentBadRequest with default headers values
func NewGetArtifactContentBadRequest() *GetArtifactContentBadRequest {
	return &GetArtifactContentBadRequest{}
}

/*GetArtifactContentBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetArtifactContentBadRequest struct {
}

func (o *GetArtifactContentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/artifacts/{id}/content][%d] getArtifactContentBadRequest ", 400)
}

func (o *GetArtifactContentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetArtifactContentMethodNotAllowed creates a GetArtifactContentMethodNotAllowed with default headers values
func NewGetArtifactContentMethodNotAllowed() *GetArtifactContentMethodNotAllowed {
	return &GetArtifactContentMethodNotAllowed{}
}

/*GetArtifactContentMethodNotAllowed handles this case with default header values.

Artifact content supports GET and PUT methods only.
*/
type GetArtifactContentMethodNotAllowed struct {
}

func (o *GetArtifactContentMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/artifacts/{id}/content][%d] getArtifactContentMethodNotAllowed ", 405)
}

func (o *GetArtifactContentMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
