// Code generated by go-swagger; DO NOT EDIT.

package server_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetServerStatusHealthReader is a Reader for the GetServerStatusHealth structure.
type GetServerStatusHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerStatusHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerStatusHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 418:
		result := NewGetServerStatusHealthIMATeapot()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetServerStatusHealthOK creates a GetServerStatusHealthOK with default headers values
func NewGetServerStatusHealthOK() *GetServerStatusHealthOK {
	return &GetServerStatusHealthOK{}
}

/*GetServerStatusHealthOK handles this case with default header values.

ServerStatusHealthResource resource returned
*/
type GetServerStatusHealthOK struct {
	Payload *models.ServerStatusHealthResource
}

func (o *GetServerStatusHealthOK) Error() string {
	return fmt.Sprintf("[GET /api/serverstatus/health][%d] getServerStatusHealthOK  %+v", 200, o.Payload)
}

func (o *GetServerStatusHealthOK) GetPayload() *models.ServerStatusHealthResource {
	return o.Payload
}

func (o *GetServerStatusHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerStatusHealthResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServerStatusHealthIMATeapot creates a GetServerStatusHealthIMATeapot with default headers values
func NewGetServerStatusHealthIMATeapot() *GetServerStatusHealthIMATeapot {
	return &GetServerStatusHealthIMATeapot{}
}

/*GetServerStatusHealthIMATeapot handles this case with default header values.

ServerStatusHealthResource resource returned
*/
type GetServerStatusHealthIMATeapot struct {
	Payload *models.ServerStatusHealthResource
}

func (o *GetServerStatusHealthIMATeapot) Error() string {
	return fmt.Sprintf("[GET /api/serverstatus/health][%d] getServerStatusHealthIMATeapot  %+v", 418, o.Payload)
}

func (o *GetServerStatusHealthIMATeapot) GetPayload() *models.ServerStatusHealthResource {
	return o.Payload
}

func (o *GetServerStatusHealthIMATeapot) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerStatusHealthResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}