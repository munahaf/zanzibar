// Code generated by zanzibar
// @generated
// Checksum : Ong2TfIs3e/Q4Xf7XfIEsQ==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package multi

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceCBackHello(in *jlexer.Lexer, out *ServiceCBack_Hello_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(string)
				}
				*out.Success = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceCBackHello(out *jwriter.Writer, in ServiceCBack_Hello_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.Success))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceCBack_Hello_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceCBackHello(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceCBack_Hello_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceCBackHello(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceCBack_Hello_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceCBackHello(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceCBack_Hello_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceCBackHello(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceCBackHello1(in *jlexer.Lexer, out *ServiceCBack_Hello_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceCBackHello1(out *jwriter.Writer, in ServiceCBack_Hello_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceCBack_Hello_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceCBackHello1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceCBack_Hello_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceCBackHello1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceCBack_Hello_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceCBackHello1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceCBack_Hello_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceCBackHello1(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceBBackHello(in *jlexer.Lexer, out *ServiceBBack_Hello_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(string)
				}
				*out.Success = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceBBackHello(out *jwriter.Writer, in ServiceBBack_Hello_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.Success))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceBBack_Hello_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceBBackHello(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceBBack_Hello_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceBBackHello(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceBBack_Hello_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceBBackHello(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceBBack_Hello_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceBBackHello(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceBBackHello1(in *jlexer.Lexer, out *ServiceBBack_Hello_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceBBackHello1(out *jwriter.Writer, in ServiceBBack_Hello_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceBBack_Hello_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceBBackHello1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceBBack_Hello_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceBBackHello1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceBBack_Hello_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceBBackHello1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceBBack_Hello_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceBBackHello1(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceABackHello(in *jlexer.Lexer, out *ServiceABack_Hello_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(string)
				}
				*out.Success = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceABackHello(out *jwriter.Writer, in ServiceABack_Hello_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.Success))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceABack_Hello_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceABackHello(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceABack_Hello_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceABackHello(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceABack_Hello_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceABackHello(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceABack_Hello_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceABackHello(l, v)
}
func easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceABackHello1(in *jlexer.Lexer, out *ServiceABack_Hello_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceABackHello1(out *jwriter.Writer, in ServiceABack_Hello_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ServiceABack_Hello_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceABackHello1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ServiceABack_Hello_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson536f2a59EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceABackHello1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ServiceABack_Hello_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceABackHello1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ServiceABack_Hello_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson536f2a59DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsIdlClientsMultiMultiServiceABackHello1(l, v)
}
