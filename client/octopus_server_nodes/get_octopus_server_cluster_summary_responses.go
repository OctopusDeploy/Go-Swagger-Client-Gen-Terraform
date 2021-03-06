// Code generated by go-swagger; DO NOT EDIT.

package octopus_server_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOctopusServerClusterSummaryReader is a Reader for the GetOctopusServerClusterSummary structure.
type GetOctopusServerClusterSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOctopusServerClusterSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOctopusServerClusterSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOctopusServerClusterSummaryOK creates a GetOctopusServerClusterSummaryOK with default headers values
func NewGetOctopusServerClusterSummaryOK() *GetOctopusServerClusterSummaryOK {
	return &GetOctopusServerClusterSummaryOK{}
}

/*GetOctopusServerClusterSummaryOK handles this case with default header values.

OK
*/
type GetOctopusServerClusterSummaryOK struct {
}

func (o *GetOctopusServerClusterSummaryOK) Error() string {
	return fmt.Sprintf("[GET /api/octopusservernodes/summary][%d] getOctopusServerClusterSummaryOK ", 200)
}

func (o *GetOctopusServerClusterSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
