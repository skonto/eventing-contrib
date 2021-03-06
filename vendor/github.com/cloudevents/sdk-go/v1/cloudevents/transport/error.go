package transport

import "fmt"

// ErrTransportMessageConversion is an error produced when the transport
// message can not be converted.
type ErrTransportMessageConversion struct {
	fatal     bool
	handled   bool
	transport string
	message   string
}

// NewErrMessageEncodingUnknown makes a new ErrMessageEncodingUnknown.
func NewErrTransportMessageConversion(transport, message string, handled, fatal bool) *ErrTransportMessageConversion {
	return &ErrTransportMessageConversion{
		transport: transport,
		message:   message,
		handled:   handled,
		fatal:     fatal,
	}
}

// IsFatal reports if this error should be considered fatal.
func (e *ErrTransportMessageConversion) IsFatal() bool {
	return e.fatal
}

// Handled reports if this error should be considered accepted and no further action.
func (e *ErrTransportMessageConversion) Handled() bool {
	return e.handled
}

// Error implements error.Error
func (e *ErrTransportMessageConversion) Error() string {
	return fmt.Sprintf("transport %s failed to convert message: %s", e.transport, e.message)
}
