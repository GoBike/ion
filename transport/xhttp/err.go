package xhttp

import "fmt"

var (
	// PhaseNewRequest is an error during request generation.
	PhaseNewRequest = "New Request"

	// PhaseEncode is an error during request/response encoding.
	PhaseEncode = "Encode"

	// PhaseDo is an error during the execution phase of the request.
	PhaseDo = "Do"

	// PhaseDecode is an error during request/response decoding.
	PhaseDecode = "Decode"
)

// Error is an error occurs at some phase within transport
type Error struct {
	// Phase is the 1/4 stages where error was generated.
	Phase string

	// Err is the concrete error
	Err error
}

// Error implements the error interface.
//"Let's turn Our Error-struct into real error!"
func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Phase, e.Err)
}
