// Code generated by go-swagger; DO NOT EDIT.

package external_security_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"models"
)

// GetListProvidersThatSupportExternalSecurityGroupsReader is a Reader for the GetListProvidersThatSupportExternalSecurityGroups structure.
type GetListProvidersThatSupportExternalSecurityGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListProvidersThatSupportExternalSecurityGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetListProvidersThatSupportExternalSecurityGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetListProvidersThatSupportExternalSecurityGroupsOK creates a GetListProvidersThatSupportExternalSecurityGroupsOK with default headers values
func NewGetListProvidersThatSupportExternalSecurityGroupsOK() *GetListProvidersThatSupportExternalSecurityGroupsOK {
	return &GetListProvidersThatSupportExternalSecurityGroupsOK{}
}

/*GetListProvidersThatSupportExternalSecurityGroupsOK handles this case with default header values.

IEnumerable_of_AuthenticationProviderThatSupportsGroups resource returned
*/
type GetListProvidersThatSupportExternalSecurityGroupsOK struct {
	Payload []*models.AuthenticationProviderThatSupportsGroups
}

func (o *GetListProvidersThatSupportExternalSecurityGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/externalsecuritygroupproviders][%d] getListProvidersThatSupportExternalSecurityGroupsOK  %+v", 200, o.Payload)
}

func (o *GetListProvidersThatSupportExternalSecurityGroupsOK) GetPayload() []*models.AuthenticationProviderThatSupportsGroups {
	return o.Payload
}

func (o *GetListProvidersThatSupportExternalSecurityGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
