// Code generated by thriftrw v1.27.0. DO NOT EDIT.
// @generated

package abc

import (
	errors "errors"
	fmt "fmt"
	strings "strings"

	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
)

// AppDemoService_Call_Args represents the arguments for the AppDemoService.Call function.
//
// The arguments for Call are sent and received over the wire as this struct.
type AppDemoService_Call_Args struct {
}

// ToWire translates a AppDemoService_Call_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *AppDemoService_Call_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a AppDemoService_Call_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a AppDemoService_Call_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v AppDemoService_Call_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *AppDemoService_Call_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a AppDemoService_Call_Args
// struct.
func (v *AppDemoService_Call_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("AppDemoService_Call_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this AppDemoService_Call_Args match the
// provided AppDemoService_Call_Args.
//
// This function performs a deep comparison.
func (v *AppDemoService_Call_Args) Equals(rhs *AppDemoService_Call_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of AppDemoService_Call_Args.
func (v *AppDemoService_Call_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "Call" for this struct.
func (v *AppDemoService_Call_Args) MethodName() string {
	return "Call"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *AppDemoService_Call_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// AppDemoService_Call_Helper provides functions that aid in handling the
// parameters and return values of the AppDemoService.Call
// function.
var AppDemoService_Call_Helper = struct {
	// Args accepts the parameters of Call in-order and returns
	// the arguments struct for the function.
	Args func() *AppDemoService_Call_Args

	// IsException returns true if the given error can be thrown
	// by Call.
	//
	// An error can be thrown by Call only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for Call
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// Call into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by Call
	//
	//   value, err := Call(args)
	//   result, err := AppDemoService_Call_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from Call: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(int32, error) (*AppDemoService_Call_Result, error)

	// UnwrapResponse takes the result struct for Call
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if Call threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := AppDemoService_Call_Helper.UnwrapResponse(result)
	UnwrapResponse func(*AppDemoService_Call_Result) (int32, error)
}{}

func init() {
	AppDemoService_Call_Helper.Args = func() *AppDemoService_Call_Args {
		return &AppDemoService_Call_Args{}
	}

	AppDemoService_Call_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	AppDemoService_Call_Helper.WrapResponse = func(success int32, err error) (*AppDemoService_Call_Result, error) {
		if err == nil {
			return &AppDemoService_Call_Result{Success: &success}, nil
		}

		return nil, err
	}
	AppDemoService_Call_Helper.UnwrapResponse = func(result *AppDemoService_Call_Result) (success int32, err error) {

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// AppDemoService_Call_Result represents the result of a AppDemoService.Call function call.
//
// The result of a Call execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type AppDemoService_Call_Result struct {
	// Value returned by Call after a successful execution.
	Success *int32 `json:"success,omitempty"`
}

// ToWire translates a AppDemoService_Call_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *AppDemoService_Call_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueI32(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("AppDemoService_Call_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a AppDemoService_Call_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a AppDemoService_Call_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v AppDemoService_Call_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *AppDemoService_Call_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TI32 {
				var x int32
				x, err = field.Value.GetI32(), error(nil)
				v.Success = &x
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("AppDemoService_Call_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a AppDemoService_Call_Result
// struct.
func (v *AppDemoService_Call_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}

	return fmt.Sprintf("AppDemoService_Call_Result{%v}", strings.Join(fields[:i], ", "))
}

func _I32_EqualsPtr(lhs, rhs *int32) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this AppDemoService_Call_Result match the
// provided AppDemoService_Call_Result.
//
// This function performs a deep comparison.
func (v *AppDemoService_Call_Result) Equals(rhs *AppDemoService_Call_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_I32_EqualsPtr(v.Success, rhs.Success) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of AppDemoService_Call_Result.
func (v *AppDemoService_Call_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		enc.AddInt32("success", *v.Success)
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *AppDemoService_Call_Result) GetSuccess() (o int32) {
	if v != nil && v.Success != nil {
		return *v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *AppDemoService_Call_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "Call" for this struct.
func (v *AppDemoService_Call_Result) MethodName() string {
	return "Call"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *AppDemoService_Call_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
