package response

import "go.uber.org/zap/zapcore"

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// MarshalLogObject implements [zapcore.ObjectMarshaler].
func (u UserResponse) MarshalLogObject(zapcore.ObjectEncoder) error {
	panic("unimplemented")
}
