// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudinstancesGetReader is a Reader for the PcloudCloudinstancesGet structure.
type PcloudCloudinstancesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudCloudinstancesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudCloudinstancesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudCloudinstancesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudCloudinstancesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudCloudinstancesGetOK creates a PcloudCloudinstancesGetOK with default headers values
func NewPcloudCloudinstancesGetOK() *PcloudCloudinstancesGetOK {
	return &PcloudCloudinstancesGetOK{}
}

/*PcloudCloudinstancesGetOK handles this case with default header values.

OK
*/
type PcloudCloudinstancesGetOK struct {
	Payload *models.CloudInstance
}

func (o *PcloudCloudinstancesGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesGetBadRequest creates a PcloudCloudinstancesGetBadRequest with default headers values
func NewPcloudCloudinstancesGetBadRequest() *PcloudCloudinstancesGetBadRequest {
	return &PcloudCloudinstancesGetBadRequest{}
}

/*PcloudCloudinstancesGetBadRequest handles this case with default header values.

Bad Request
*/
type PcloudCloudinstancesGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesGetNotFound creates a PcloudCloudinstancesGetNotFound with default headers values
func NewPcloudCloudinstancesGetNotFound() *PcloudCloudinstancesGetNotFound {
	return &PcloudCloudinstancesGetNotFound{}
}

/*PcloudCloudinstancesGetNotFound handles this case with default header values.

Not Found
*/
type PcloudCloudinstancesGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesGetInternalServerError creates a PcloudCloudinstancesGetInternalServerError with default headers values
func NewPcloudCloudinstancesGetInternalServerError() *PcloudCloudinstancesGetInternalServerError {
	return &PcloudCloudinstancesGetInternalServerError{}
}

/*PcloudCloudinstancesGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudCloudinstancesGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}][%d] pcloudCloudinstancesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}