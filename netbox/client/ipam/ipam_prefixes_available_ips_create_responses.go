// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/netbox-community/go-netbox/netbox/models"
)

// IpamPrefixesAvailableIpsCreateReader is a Reader for the IpamPrefixesAvailableIpsCreate structure.
type IpamPrefixesAvailableIpsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesAvailableIpsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamPrefixesAvailableIpsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIpamPrefixesAvailableIpsCreateCreated creates a IpamPrefixesAvailableIpsCreateCreated with default headers values
func NewIpamPrefixesAvailableIpsCreateCreated() *IpamPrefixesAvailableIpsCreateCreated {
	return &IpamPrefixesAvailableIpsCreateCreated{}
}

/*IpamPrefixesAvailableIpsCreateCreated handles this case with default header values.

IpamPrefixesAvailableIpsCreateCreated ipam prefixes available ips create created
*/
type IpamPrefixesAvailableIpsCreateCreated struct {
	Payload *models.IPAddress
}

func (o *IpamPrefixesAvailableIpsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-ips/][%d] ipamPrefixesAvailableIpsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamPrefixesAvailableIpsCreateCreated) GetPayload() *models.IPAddress {
	return o.Payload
}

func (o *IpamPrefixesAvailableIpsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
