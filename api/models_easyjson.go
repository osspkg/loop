// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package api

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

func easyjsonD2b7633eDecodeGoArwosOrgLoopyApi(in *jlexer.Lexer, out *EntityKV) {
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
		case "k":
			out.Key = string(in.String())
		case "v":
			if in.IsNull() {
				in.Skip()
				out.Val = nil
			} else {
				if out.Val == nil {
					out.Val = new(string)
				}
				*out.Val = string(in.String())
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
func easyjsonD2b7633eEncodeGoArwosOrgLoopyApi(out *jwriter.Writer, in EntityKV) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"k\":"
		out.RawString(prefix[1:])
		out.String(string(in.Key))
	}
	{
		const prefix string = ",\"v\":"
		out.RawString(prefix)
		if in.Val == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Val))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EntityKV) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGoArwosOrgLoopyApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EntityKV) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGoArwosOrgLoopyApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EntityKV) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGoArwosOrgLoopyApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EntityKV) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGoArwosOrgLoopyApi(l, v)
}
func easyjsonD2b7633eDecodeGoArwosOrgLoopyApi1(in *jlexer.Lexer, out *EntitiesKV) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(EntitiesKV, 0, 2)
			} else {
				*out = EntitiesKV{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 EntityKV
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGoArwosOrgLoopyApi1(out *jwriter.Writer, in EntitiesKV) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v EntitiesKV) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGoArwosOrgLoopyApi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EntitiesKV) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGoArwosOrgLoopyApi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EntitiesKV) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGoArwosOrgLoopyApi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EntitiesKV) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGoArwosOrgLoopyApi1(l, v)
}
