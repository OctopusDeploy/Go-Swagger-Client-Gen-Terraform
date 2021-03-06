// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetTenantLogoReader is a Reader for the GetTenantLogo structure.
type GetTenantLogoReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantLogoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantLogoOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetTenantLogoNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetTenantLogoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTenantLogoOK creates a GetTenantLogoOK with default headers values
func NewGetTenantLogoOK(writer io.Writer) *GetTenantLogoOK {
	return &GetTenantLogoOK{
		Payload: writer,
	}
}

/*GetTenantLogoOK handles this case with default header values.

OK
*/
type GetTenantLogoOK struct {
	Payload io.Writer
}

func (o *GetTenantLogoOK) Error() string {
	return fmt.Sprintf("[GET /api/tenants/{id}/logo][%d] getTenantLogoOK  %+v", 200, o.Payload)
}

func (o *GetTenantLogoOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetTenantLogoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantLogoNotModified creates a GetTenantLogoNotModified with default headers values
func NewGetTenantLogoNotModified() *GetTenantLogoNotModified {
	return &GetTenantLogoNotModified{}
}

/*GetTenantLogoNotModified handles this case with default header values.

NotModified
*/
type GetTenantLogoNotModified struct {
}

func (o *GetTenantLogoNotModified) Error() string {
	return fmt.Sprintf("[GET /api/tenants/{id}/logo][%d] getTenantLogoNotModified ", 304)
}

func (o *GetTenantLogoNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTenantLogoBadRequest creates a GetTenantLogoBadRequest with default headers values
func NewGetTenantLogoBadRequest() *GetTenantLogoBadRequest {
	return &GetTenantLogoBadRequest{}
}

/*GetTenantLogoBadRequest handles this case with default header values.

No id parameter was provided.
*/
type GetTenantLogoBadRequest struct {
}

func (o *GetTenantLogoBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/tenants/{id}/logo][%d] getTenantLogoBadRequest ", 400)
}

func (o *GetTenantLogoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
