package api

import (
	"context"

	"github.com/fanaticscripter/EggOrganizer/ei"
)

func RequestFirstContact(payload *ei.EggIncFirstContactRequest) (*ei.EggIncFirstContactResponse, error) {
	return RequestFirstContactWithContext(context.Background(), payload)
}

func RequestFirstContactWithContext(ctx context.Context, payload *ei.EggIncFirstContactRequest) (*ei.EggIncFirstContactResponse, error) {
	botName := "EggOrganizer"
	payload.DeviceId = &botName // This is actually bot_name for /ei/bot_first_contact.
	if payload.ClientVersion == nil {
		version := ClientVersion
		payload.ClientVersion = &version
	}
	resp := &ei.EggIncFirstContactResponse{}
	err := RequestWithContext(ctx, "/ei/bot_first_contact", payload, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
