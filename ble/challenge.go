package ble

import (
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts2"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts3"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts4"
	"github.com/digital-dream-labs/vector-bluetooth/ble/rts5"
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

func handleRTS2ChallengeMessage(v *VectorBLE, msg *rts.RtsConnection_2) ([]byte, bool, error) {
	b, err := rts2.BuildChallengeResponse(msg.GetRtsChallengeMessage().Number)
	if err != nil {
		return nil, false, err
	}

	if err := v.ble.Send(b); err != nil {
		return nil, false, err
	}
	return nil, true, nil
}

func handleRTS3ChallengeMessage(v *VectorBLE, msg *rts.RtsConnection_3) ([]byte, bool, error) {
	b, err := rts3.BuildChallengeResponse(msg.GetRtsChallengeMessage().Number)
	if err != nil {
		return nil, false, err
	}

	if err := v.ble.Send(b); err != nil {
		return nil, false, err
	}
	return nil, true, nil
}

func handleRTS4ChallengeMessage(v *VectorBLE, msg *rts.RtsConnection_4) ([]byte, bool, error) {
	b, err := rts4.BuildChallengeResponse(msg.GetRtsChallengeMessage().Number)
	if err != nil {
		return nil, false, err
	}

	if err := v.ble.Send(b); err != nil {
		return nil, false, err
	}
	return nil, true, nil
}

func handleRTS5ChallengeMessage(v *VectorBLE, msg *rts.RtsConnection_5) ([]byte, bool, error) {
	b, err := rts5.BuildChallengeResponse(msg.GetRtsChallengeMessage().Number)
	if err != nil {
		return nil, false, err
	}

	if err := v.ble.Send(b); err != nil {
		return nil, false, err
	}
	return nil, true, nil
}

func handleRTS2ChallengeSuccessMessage(v *VectorBLE, msg *rts.RtsConnection_2) ([]byte, bool, error) {
	v.state.authorized = true
	return nil, false, nil
}

func handleRTS3ChallengeSuccessMessage(v *VectorBLE, msg *rts.RtsConnection_3) ([]byte, bool, error) {
	v.state.authorized = true
	return nil, false, nil
}

func handleRTS4ChallengeSuccessMessage(v *VectorBLE, msg *rts.RtsConnection_4) ([]byte, bool, error) {
	v.state.authorized = true
	return nil, false, nil
}

func handleRTS5ChallengeSuccessMessage(v *VectorBLE, msg *rts.RtsConnection_5) ([]byte, bool, error) {
	v.state.authorized = true
	return nil, false, nil
}
