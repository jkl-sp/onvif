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

// Call_GetCertificatesStatus forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCertificatesStatusResponse.
func Call_GetCertificatesStatus(ctx context.Context, dev *onvif.Device, request device.GetCertificatesStatus) (device.GetCertificatesStatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCertificatesStatusResponse device.GetCertificatesStatusResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCertificatesStatusResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCertificatesStatus")
		return reply.Body.GetCertificatesStatusResponse, errors.Annotate(err, "reply")
	}
}
