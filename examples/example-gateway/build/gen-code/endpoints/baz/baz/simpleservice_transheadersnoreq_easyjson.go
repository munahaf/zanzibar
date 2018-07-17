// Code generated by zanzibar
// @generated
// Checksum : bcn101rmOtcy7nJt1kIuKQ==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package baz

import (
	json "encoding/json"
	fmt "fmt"
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

func easyjsonEcca83c7DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTransHeadersNoReq(in *jlexer.Lexer, out *SimpleService_TransHeadersNoReq_Result) {
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
		key := in.UnsafeString()
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
					out.Success = new(TransHeader)
				}
				easyjsonEcca83c7DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz(in, &*out.Success)
			}
		case "authErr":
			if in.IsNull() {
				in.Skip()
				out.AuthErr = nil
			} else {
				if out.AuthErr == nil {
					out.AuthErr = new(AuthErr)
				}
				easyjsonEcca83c7DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz1(in, &*out.AuthErr)
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
func easyjsonEcca83c7EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTransHeadersNoReq(out *jwriter.Writer, in SimpleService_TransHeadersNoReq_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		const prefix string = ",\"success\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonEcca83c7EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz(out, *in.Success)
	}
	if in.AuthErr != nil {
		const prefix string = ",\"authErr\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonEcca83c7EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz1(out, *in.AuthErr)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_TransHeadersNoReq_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEcca83c7EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTransHeadersNoReq(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_TransHeadersNoReq_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEcca83c7EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTransHeadersNoReq(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_TransHeadersNoReq_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEcca83c7DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTransHeadersNoReq(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_TransHeadersNoReq_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEcca83c7DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTransHeadersNoReq(l, v)
}
func easyjsonEcca83c7DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz1(in *jlexer.Lexer, out *AuthErr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var MessageSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "message":
			out.Message = string(in.String())
			MessageSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !MessageSet {
		in.AddError(fmt.Errorf("key 'message' is required"))
	}
}
func easyjsonEcca83c7EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz1(out *jwriter.Writer, in AuthErr) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	out.RawByte('}')
}
func easyjsonEcca83c7DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz(in *jlexer.Lexer, out *TransHeader) {
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
		key := in.UnsafeString()
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
func easyjsonEcca83c7EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBaz(out *jwriter.Writer, in TransHeader) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}
func easyjsonEcca83c7DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTransHeadersNoReq1(in *jlexer.Lexer, out *SimpleService_TransHeadersNoReq_Args) {
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
		key := in.UnsafeString()
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
func easyjsonEcca83c7EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTransHeadersNoReq1(out *jwriter.Writer, in SimpleService_TransHeadersNoReq_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_TransHeadersNoReq_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEcca83c7EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTransHeadersNoReq1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_TransHeadersNoReq_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEcca83c7EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTransHeadersNoReq1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_TransHeadersNoReq_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEcca83c7DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTransHeadersNoReq1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_TransHeadersNoReq_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEcca83c7DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazBazSimpleServiceTransHeadersNoReq1(l, v)
}
