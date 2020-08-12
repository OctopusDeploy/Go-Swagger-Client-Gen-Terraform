// Code generated by go-swagger; DO NOT EDIT.

package server_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateServerConfigurationReader is a Reader for the UpdateServerConfiguration structure.
type UpdateServerConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServerConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateServerConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateServerConfigurationOK creates a UpdateServerConfigurationOK with default headers values
func NewUpdateServerConfigurationOK() *UpdateServerConfigurationOK {
	return &UpdateServerConfigurationOK{}
}

/*UpdateServerConfigurationOK handles this case with default header values.

OK - Default
*/
type UpdateServerConfigurationOK struct {
}

func (o *UpdateServerConfigurationOK) Error() string {
	return fmt.Sprintf("[PUT /api/serverconfiguration][%d] updateServerConfigurationOK ", 200)
}

func (o *UpdateServerConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
