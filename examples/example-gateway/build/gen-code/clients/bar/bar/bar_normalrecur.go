// Code generated by thriftrw v1.19.0. DO NOT EDIT.
// @generated

package bar

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

// Bar_NormalRecur_Args represents the arguments for the Bar.normalRecur function.
//
// The arguments for normalRecur are sent and received over the wire as this struct.
type Bar_NormalRecur_Args struct {
	Request *BarRequestRecur `json:"request,required"`
}

// ToWire translates a Bar_NormalRecur_Args struct into a Thrift-level intermediate
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
func (v *Bar_NormalRecur_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Request == nil {
		return w, errors.New("field Request of Bar_NormalRecur_Args is required")
	}
	w, err = v.Request.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Bar_NormalRecur_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_NormalRecur_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_NormalRecur_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_NormalRecur_Args) FromWire(w wire.Value) error {
	var err error

	requestIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _BarRequestRecur_Read(field.Value)
				if err != nil {
					return err
				}
				requestIsSet = true
			}
		}
	}

	if !requestIsSet {
		return errors.New("field Request of Bar_NormalRecur_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Bar_NormalRecur_Args
// struct.
func (v *Bar_NormalRecur_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Request: %v", v.Request)
	i++

	return fmt.Sprintf("Bar_NormalRecur_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_NormalRecur_Args match the
// provided Bar_NormalRecur_Args.
//
// This function performs a deep comparison.
func (v *Bar_NormalRecur_Args) Equals(rhs *Bar_NormalRecur_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !v.Request.Equals(rhs.Request) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_NormalRecur_Args.
func (v *Bar_NormalRecur_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	err = multierr.Append(err, enc.AddObject("request", v.Request))
	return err
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *Bar_NormalRecur_Args) GetRequest() (o *BarRequestRecur) {
	if v != nil {
		o = v.Request
	}
	return
}

// IsSetRequest returns true if Request is not nil.
func (v *Bar_NormalRecur_Args) IsSetRequest() bool {
	return v != nil && v.Request != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "normalRecur" for this struct.
func (v *Bar_NormalRecur_Args) MethodName() string {
	return "normalRecur"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Bar_NormalRecur_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Bar_NormalRecur_Helper provides functions that aid in handling the
// parameters and return values of the Bar.normalRecur
// function.
var Bar_NormalRecur_Helper = struct {
	// Args accepts the parameters of normalRecur in-order and returns
	// the arguments struct for the function.
	Args func(
		request *BarRequestRecur,
	) *Bar_NormalRecur_Args

	// IsException returns true if the given error can be thrown
	// by normalRecur.
	//
	// An error can be thrown by normalRecur only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for normalRecur
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// normalRecur into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by normalRecur
	//
	//   value, err := normalRecur(args)
	//   result, err := Bar_NormalRecur_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from normalRecur: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*BarResponseRecur, error) (*Bar_NormalRecur_Result, error)

	// UnwrapResponse takes the result struct for normalRecur
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if normalRecur threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Bar_NormalRecur_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Bar_NormalRecur_Result) (*BarResponseRecur, error)
}{}

func init() {
	Bar_NormalRecur_Helper.Args = func(
		request *BarRequestRecur,
	) *Bar_NormalRecur_Args {
		return &Bar_NormalRecur_Args{
			Request: request,
		}
	}

	Bar_NormalRecur_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *BarException:
			return true
		default:
			return false
		}
	}

	Bar_NormalRecur_Helper.WrapResponse = func(success *BarResponseRecur, err error) (*Bar_NormalRecur_Result, error) {
		if err == nil {
			return &Bar_NormalRecur_Result{Success: success}, nil
		}

		switch e := err.(type) {
		case *BarException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Bar_NormalRecur_Result.BarException")
			}
			return &Bar_NormalRecur_Result{BarException: e}, nil
		}

		return nil, err
	}
	Bar_NormalRecur_Helper.UnwrapResponse = func(result *Bar_NormalRecur_Result) (success *BarResponseRecur, err error) {
		if result.BarException != nil {
			err = result.BarException
			return
		}

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Bar_NormalRecur_Result represents the result of a Bar.normalRecur function call.
//
// The result of a normalRecur execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Bar_NormalRecur_Result struct {
	// Value returned by normalRecur after a successful execution.
	Success      *BarResponseRecur `json:"success,omitempty"`
	BarException *BarException     `json:"barException,omitempty"`
}

// ToWire translates a Bar_NormalRecur_Result struct into a Thrift-level intermediate
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
func (v *Bar_NormalRecur_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.BarException != nil {
		w, err = v.BarException.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Bar_NormalRecur_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _BarResponseRecur_Read(w wire.Value) (*BarResponseRecur, error) {
	var v BarResponseRecur
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Bar_NormalRecur_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar_NormalRecur_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar_NormalRecur_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar_NormalRecur_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _BarResponseRecur_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.BarException, err = _BarException_Read(field.Value)
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
	if v.BarException != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Bar_NormalRecur_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Bar_NormalRecur_Result
// struct.
func (v *Bar_NormalRecur_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.BarException != nil {
		fields[i] = fmt.Sprintf("BarException: %v", v.BarException)
		i++
	}

	return fmt.Sprintf("Bar_NormalRecur_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Bar_NormalRecur_Result match the
// provided Bar_NormalRecur_Result.
//
// This function performs a deep comparison.
func (v *Bar_NormalRecur_Result) Equals(rhs *Bar_NormalRecur_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.BarException == nil && rhs.BarException == nil) || (v.BarException != nil && rhs.BarException != nil && v.BarException.Equals(rhs.BarException))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar_NormalRecur_Result.
func (v *Bar_NormalRecur_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	if v.BarException != nil {
		err = multierr.Append(err, enc.AddObject("barException", v.BarException))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Bar_NormalRecur_Result) GetSuccess() (o *BarResponseRecur) {
	if v != nil && v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Bar_NormalRecur_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// GetBarException returns the value of BarException if it is set or its
// zero value if it is unset.
func (v *Bar_NormalRecur_Result) GetBarException() (o *BarException) {
	if v != nil && v.BarException != nil {
		return v.BarException
	}

	return
}

// IsSetBarException returns true if BarException is not nil.
func (v *Bar_NormalRecur_Result) IsSetBarException() bool {
	return v != nil && v.BarException != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "normalRecur" for this struct.
func (v *Bar_NormalRecur_Result) MethodName() string {
	return "normalRecur"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Bar_NormalRecur_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
