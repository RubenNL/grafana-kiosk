package webaudio

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// ContextID context's UUID in string.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-ContextId
type ContextID string

// String returns the ContextID as string value.
func (t ContextID) String() string {
	return string(t)
}

// ContextType enum of BaseAudioContext types.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-ContextType
type ContextType string

// String returns the ContextType as string value.
func (t ContextType) String() string {
	return string(t)
}

// ContextType values.
const (
	ContextTypeRealtime ContextType = "realtime"
	ContextTypeOffline  ContextType = "offline"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ContextType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ContextType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ContextType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ContextType(in.String()) {
	case ContextTypeRealtime:
		*t = ContextTypeRealtime
	case ContextTypeOffline:
		*t = ContextTypeOffline

	default:
		in.AddError(errors.New("unknown ContextType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ContextType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ContextState enum of AudioContextState from the spec.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-ContextState
type ContextState string

// String returns the ContextState as string value.
func (t ContextState) String() string {
	return string(t)
}

// ContextState values.
const (
	ContextStateSuspended ContextState = "suspended"
	ContextStateRunning   ContextState = "running"
	ContextStateClosed    ContextState = "closed"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t ContextState) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t ContextState) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *ContextState) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch ContextState(in.String()) {
	case ContextStateSuspended:
		*t = ContextStateSuspended
	case ContextStateRunning:
		*t = ContextStateRunning
	case ContextStateClosed:
		*t = ContextStateClosed

	default:
		in.AddError(errors.New("unknown ContextState value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *ContextState) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// ContextRealtimeData fields in AudioContext that change in real-time. These
// are not updated on OfflineAudioContext.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-ContextRealtimeData
type ContextRealtimeData struct {
	CurrentTime    float64 `json:"currentTime,omitempty"`    // The current context time in second in BaseAudioContext.
	RenderCapacity float64 `json:"renderCapacity,omitempty"` // The time spent on rendering graph divided by render quantum duration, and multiplied by 100. 100 means the audio renderer reached the full capacity and glitch may occur.
}

// BaseAudioContext protocol object for BaseAudioContext.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/WebAudio#type-BaseAudioContext
type BaseAudioContext struct {
	ContextID             ContextID            `json:"contextId"`
	ContextType           ContextType          `json:"contextType"`
	ContextState          ContextState         `json:"contextState"`
	RealtimeData          *ContextRealtimeData `json:"realtimeData,omitempty"`
	CallbackBufferSize    float64              `json:"callbackBufferSize"`    // Platform-dependent callback buffer size.
	MaxOutputChannelCount float64              `json:"maxOutputChannelCount"` // Number of output channels supported by audio hardware in use.
	SampleRate            float64              `json:"sampleRate"`            // Context sample rate.
}
