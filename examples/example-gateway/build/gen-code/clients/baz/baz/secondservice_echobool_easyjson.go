// Code generated by zanzibar
// @generated
// Checksum : NPm0WopGtJDBQK4Uv5WRRw==
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

func easyjsonCefa8dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoBool(in *jlexer.Lexer, out *SecondService_EchoBool_Result) {
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
					out.Success = new(bool)
				}
				*out.Success = bool(in.Bool())
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
func easyjsonCefa8dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoBool(out *jwriter.Writer, in SecondService_EchoBool_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"success\":")
		if in.Success == nil {
			out.RawString("null")
		} else {
			out.Bool(bool(*in.Success))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SecondService_EchoBool_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCefa8dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoBool(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SecondService_EchoBool_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCefa8dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoBool(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SecondService_EchoBool_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCefa8dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoBool(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SecondService_EchoBool_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCefa8dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoBool(l, v)
}
func easyjsonCefa8dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoBool1(in *jlexer.Lexer, out *SecondService_EchoBool_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var ArgSet bool
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
		case "arg":
			out.Arg = bool(in.Bool())
			ArgSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !ArgSet {
		in.AddError(fmt.Errorf("key 'arg' is required"))
	}
}
func easyjsonCefa8dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoBool1(out *jwriter.Writer, in SecondService_EchoBool_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"arg\":")
	out.Bool(bool(in.Arg))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SecondService_EchoBool_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCefa8dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoBool1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SecondService_EchoBool_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCefa8dfbEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoBool1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SecondService_EchoBool_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCefa8dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoBool1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SecondService_EchoBool_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCefa8dfbDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoBool1(l, v)
}
