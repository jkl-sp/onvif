// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/juju/errors"
	"github.com/jkl-sp/onvif"
	"github.com/jkl-sp/onvif/sdk"
	"github.com/jkl-sp/onvif/ptz"
)

// Call_SetConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetConfigurationResponse.
func Call_SetConfiguration(ctx context.Context, dev *onvif.Device, request ptz.SetConfiguration) (ptz.SetConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetConfigurationResponse ptz.SetConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetConfiguration")
		return reply.Body.SetConfigurationResponse, errors.Annotate(err, "reply")
	}
}
