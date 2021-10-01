// Base on https://github.com/fanaticscripter/EggContractor/blob/3ce2cdc9ee767ecc8cbdfa4ae0ac90d248dc8694/api/request.go

package api

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context/ctxhttp"
	"google.golang.org/protobuf/proto"

	"github.com/fanaticscripter/EggOrganizer/ei"
)

const (
	ClientVersion uint32 = 36
	AppVersion    string = "1.22.1"
	AppBuild      string = "1.22.1.1"
	Platform      string = "IOS"
)

const _apiPrefix = "https://www.auxbrain.com"

var _client *http.Client

func init() {
	_client = &http.Client{
		Timeout: 5 * time.Second,
	}
}

func Request(endpoint string, reqMsg proto.Message, respMsg proto.Message) error {
	return RequestWithContext(context.Background(), endpoint, reqMsg, respMsg)
}

func RequestWithContext(ctx context.Context, endpoint string, reqMsg proto.Message, respMsg proto.Message) error {
	apiUrl := _apiPrefix + endpoint
	reqBin, err := proto.Marshal(reqMsg)
	if err != nil {
		return errors.Wrapf(err, "marshaling payload %+v for %s", reqMsg, apiUrl)
	}
	enc := base64.StdEncoding
	reqDataEncoded := enc.EncodeToString(reqBin)
	log.Infof("POST %s: %+v", apiUrl, reqMsg)
	log.Debugf("POST %s data=%s", apiUrl, reqDataEncoded)
	resp, err := ctxhttp.PostForm(ctx, _client, apiUrl, url.Values{"data": {reqDataEncoded}})
	if err != nil {
		return errors.Wrapf(err, "POST %s", apiUrl)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "POST %s", apiUrl)
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return errors.Errorf("POST %s: HTTP %d: %#v", apiUrl, resp.StatusCode, string(body))
	}
	respBinBuf := make([]byte, enc.DecodedLen(len(body)))
	n, err := enc.Decode(respBinBuf, body)
	if err != nil {
		return errors.Wrapf(err, "base64 decoding %s reponse (%#v)", apiUrl, string(body))
	}
	err = proto.Unmarshal(respBinBuf[:n], respMsg)
	if err != nil {
		return errors.Wrapf(err, "unmarshaling %s response (%#v)", apiUrl, string(body))
	}
	return nil
}

func RequestAuthenticated(endpoint string, reqMsg proto.Message, respMsg proto.Message) error {
	return RequestAuthenticatedWithContext(context.Background(), endpoint, reqMsg, respMsg)
}

func RequestAuthenticatedWithContext(ctx context.Context, endpoint string, reqMsg proto.Message, respMsg proto.Message) error {
	apiUrl := _apiPrefix + endpoint
	authenticatedMsg := &ei.AuthenticatedMessage{}
	err := RequestWithContext(ctx, endpoint, reqMsg, authenticatedMsg)
	if err != nil {
		return err
	}
	payload := authenticatedMsg.Message
	err = proto.Unmarshal(payload, respMsg)
	if err != nil {
		return errors.Wrapf(err, "unmarshaling %s authenticated payload (%#v)", apiUrl, base64.RawStdEncoding.EncodeToString(payload))
	}
	return nil
}
