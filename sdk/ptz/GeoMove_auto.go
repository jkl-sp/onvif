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

// Call_GeoMove forwards the call to dev.CallMethod() then parses the payload of the reply as a GeoMoveResponse.
func Call_GeoMove(ctx context.Context, dev *onvif.Device, request ptz.GeoMove) (ptz.GeoMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GeoMoveResponse ptz.GeoMoveResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GeoMoveResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GeoMove")
		return reply.Body.GeoMoveResponse, errors.Annotate(err, "reply")
	}
}
