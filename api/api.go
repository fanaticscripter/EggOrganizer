package api

import (
	"context"

	"github.com/google/uuid"

	"github.com/fanaticscripter/EggOrganizer/ei"
)

func RequestFirstContact(payload *ei.EggIncFirstContactRequest) (*ei.EggIncFirstContactResponse, error) {
	return RequestFirstContactWithContext(context.Background(), payload)
}

func RequestFirstContactWithContext(ctx context.Context, payload *ei.EggIncFirstContactRequest) (*ei.EggIncFirstContactResponse, error) {
	if payload.ClientVersion == nil {
		version := ClientVersion
		payload.ClientVersion = &version
	}
	if payload.DeviceId == nil {
		deviceId := uuid.New().String()
		payload.DeviceId = &deviceId
	}
	resp := &ei.EggIncFirstContactResponse{}
	err := RequestAuthenticatedWithContext(ctx, "/ei/first_contact", payload, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
