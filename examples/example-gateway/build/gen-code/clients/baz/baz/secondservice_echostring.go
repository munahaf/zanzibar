// Code generated by thriftrw v1.19.0. DO NOT EDIT.
// @generated

package baz

import (
	errors "errors"
	fmt "fmt"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// SecondService_EchoString_Args represents the arguments for the SecondService.echoString function.
//
// The arguments for echoString are sent and received over the wire as this struct.
type SecondService_EchoString_Args struct {
	Arg string `json:"arg,required"`
}

// ToWire translates a SecondService_EchoString_Args struct into a Thrift-level intermediate
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
func (v *SecondService_EchoString_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Arg), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a SecondService_EchoString_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SecondService_EchoString_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SecondService_EchoString_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SecondService_EchoString_Args) FromWire(w wire.Value) error {
	var err error

	argIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Arg, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				argIsSet = true
			}
		}
	}

	if !argIsSet {
		return errors.New("field Arg of SecondService_EchoString_Args is required")
	}

	return nil
}

// String returns a readable string representation of a SecondService_EchoString_Args
// struct.
func (v *SecondService_EchoString_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Arg: %v", v.Arg)
	i++

	return fmt.Sprintf("SecondService_EchoString_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this SecondService_EchoString_Args match the
// provided SecondService_EchoString_Args.
//
// This function performs a deep comparison.
func (v *SecondService_EchoString_Args) Equals(rhs *SecondService_EchoString_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Arg == rhs.Arg) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of SecondService_EchoString_Args.
func (v *SecondService_EchoString_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("arg", v.Arg)
	return err
}

// GetArg returns the value of Arg if it is set or its
// zero value if it is unset.
func (v *SecondService_EchoString_Args) GetArg() (o string) {
	if v != nil {
		o = v.Arg
	}
	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "echoString" for this struct.
func (v *SecondService_EchoString_Args) MethodName() string {
	return "echoString"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *SecondService_EchoString_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// SecondService_EchoString_Helper provides functions that aid in handling the
// parameters and return values of the SecondService.echoString
// function.
var SecondService_EchoString_Helper = struct {
	// Args accepts the parameters of echoString in-order and returns
	// the arguments struct for the function.
	Args func(
		arg string,
	) *SecondService_EchoString_Args

	// IsException returns true if the given error can be thrown
	// by echoString.
	//
	// An error can be thrown by echoString only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for echoString
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// echoString into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by echoString
	//
	//   value, err := echoString(args)
	//   result, err := SecondService_EchoString_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from echoString: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(string, error) (*SecondService_EchoString_Result, error)

	// UnwrapResponse takes the result struct for echoString
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if echoString threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := SecondService_EchoString_Helper.UnwrapResponse(result)
	UnwrapResponse func(*SecondService_EchoString_Result) (string, error)
}{}

func init() {
	SecondService_EchoString_Helper.Args = func(
		arg string,
	) *SecondService_EchoString_Args {
		return &SecondService_EchoString_Args{
			Arg: arg,
		}
	}

	SecondService_EchoString_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	SecondService_EchoString_Helper.WrapResponse = func(success string, err error) (*SecondService_EchoString_Result, error) {
		if err == nil {
			return &SecondService_EchoString_Result{Success: &success}, nil
		}

		return nil, err
	}
	SecondService_EchoString_Helper.UnwrapResponse = func(result *SecondService_EchoString_Result) (success string, err error) {

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// SecondService_EchoString_Result represents the result of a SecondService.echoString function call.
//
// The result of a echoString execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type SecondService_EchoString_Result struct {
	// Value returned by echoString after a successful execution.
	Success *string `json:"success,omitempty"`
}

// ToWire translates a SecondService_EchoString_Result struct into a Thrift-level intermediate
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
func (v *SecondService_EchoString_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueString(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("SecondService_EchoString_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a SecondService_EchoString_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a SecondService_EchoString_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v SecondService_EchoString_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *SecondService_EchoString_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
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
		return fmt.Errorf("SecondService_EchoString_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a SecondService_EchoString_Result
// struct.
func (v *SecondService_EchoString_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}

	return fmt.Sprintf("SecondService_EchoString_Result{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this SecondService_EchoString_Result match the
// provided SecondService_EchoString_Result.
//
// This function performs a deep comparison.
func (v *SecondService_EchoString_Result) Equals(rhs *SecondService_EchoString_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Success, rhs.Success) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of SecondService_EchoString_Result.
func (v *SecondService_EchoString_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		enc.AddString("success", *v.Success)
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *SecondService_EchoString_Result) GetSuccess() (o string) {
	if v != nil && v.Success != nil {
		return *v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *SecondService_EchoString_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "echoString" for this struct.
func (v *SecondService_EchoString_Result) MethodName() string {
	return "echoString"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *SecondService_EchoString_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
