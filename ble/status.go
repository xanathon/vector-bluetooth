package ble

import (
	"encoding/hex"
	"encoding/json"

	"github.com/digital-dream-labs/vector-bluetooth/ble/rts2"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts5"
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

// StatusResponse is the unified response for status messages
type StatusResponse struct {
	WifiSSID      string `json:"wifi_ssid,omitempty"`
	Version       string `json:"version,omitempty"`
	ESN           string `json:"esn,omitempty"`
	WifiState     int    `json:"wifi_state,omitempty"`
	AccessPoint   bool   `json:"access_point,omitempty"`
	OtaInProgress bool   `json:"ota_in_progress,omitempty"`
	HasOwner      bool   `json:"has_owner,omitempty"`
	CloudAuthed   bool   `json:"cloud_authed,omitempty"`
}

// Marshal converts a StatusResponse message to bytes
func (sr *StatusResponse) Marshal() ([]byte, error) {
	return json.Marshal(sr)
}

// Unmarshal converts a StatusResponse byte slice to a StatusResponse
func (sr *StatusResponse) Unmarshal(b []byte) error {
	return json.Unmarshal(b, sr)
}

// GetStatus sends a GetStatus message to the vector robot
func (v *VectorBLE) GetStatus() (*StatusResponse, error) {
	var (
		msg []byte
		err error
	)

	switch v.ble.Version {
	case rtsV2:
		msg, err = rts2.BuildStatusMessage()
	case rtsV3:
	case rtsV4:
	case rtsV5:
		msg, err = rts5.BuildStatusMessage()
	}
	if err != nil {
		return nil, err
	}

	if err := v.ble.Send(msg); err != nil {
		return nil, err
	}

	b, err := v.watch()

	resp := StatusResponse{}
	if err := resp.Unmarshal(b); err != nil {
		return nil, err
	}
	return &resp, err
}

func handleRST2StatusResponse(v *VectorBLE, msg *rts.RtsConnection_2) ([]byte, bool, error) {
	r := msg.GetRtsStatusResponse2()

	ssid, _ := hex.DecodeString(r.WifiSsidHex)

	sr := StatusResponse{
		WifiSSID:      string(ssid),
		Version:       r.Version,
		WifiState:     int(r.WifiState),
		AccessPoint:   r.AccessPoint,
		OtaInProgress: r.OtaInProgress,
	}

	b, err := sr.Marshal()
	return b, false, err
}

func handleRST5StatusResponse(v *VectorBLE, msg *rts.RtsConnection_5) ([]byte, bool, error) {
	r := msg.GetRtsStatusResponse5()

	ssid, _ := hex.DecodeString(r.WifiSsidHex)

	sr := StatusResponse{
		WifiSSID:      string(ssid),
		Version:       r.Version,
		ESN:           r.Esn,
		WifiState:     int(r.WifiState),
		AccessPoint:   r.AccessPoint,
		OtaInProgress: r.OtaInProgress,
		HasOwner:      r.HasOwner,
		CloudAuthed:   r.IsCloudAuthed,
	}

	b, err := sr.Marshal()
	return b, false, err
}