// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/jkl-sp/onvif"
	"github.com/jkl-sp/onvif/sdk"
	"github.com/jkl-sp/onvif/device"
)

// Call_GetDynamicDNS forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDynamicDNSResponse.
func Call_GetDynamicDNS(ctx context.Context, dev *onvif.Device, request device.GetDynamicDNS) (device.GetDynamicDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDynamicDNSResponse device.GetDynamicDNSResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDynamicDNSResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetDynamicDNS")
		return reply.Body.GetDynamicDNSResponse, errors.Annotate(err, "reply")
	}
}
