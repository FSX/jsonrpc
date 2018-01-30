// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package jsonrpc

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

func easyjson4086215fDecodeAutogitJsonrpc(in *jlexer.Lexer, out *Response) {
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
		case "jsonrpc":
			out.JSONRPC = string(in.String())
		case "result":
			if in.IsNull() {
				in.Skip()
				out.Result = nil
			} else {
				if out.Result == nil {
					out.Result = new(json.RawMessage)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Result).UnmarshalJSON(data))
				}
			}
		case "error":
			if in.IsNull() {
				in.Skip()
				out.Error = nil
			} else {
				if out.Error == nil {
					out.Error = new(Error)
				}
				(*out.Error).UnmarshalEasyJSON(in)
			}
		case "id":
			out.Id = Id(in.String())
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
func easyjson4086215fEncodeAutogitJsonrpc(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"jsonrpc\":")
	out.String(string(in.JSONRPC))
	if in.Result != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"result\":")
		if in.Result == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Result).MarshalJSON())
		}
	}
	if in.Error != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"error\":")
		if in.Error == nil {
			out.RawString("null")
		} else {
			(*in.Error).MarshalEasyJSON(out)
		}
	}
	if in.Id != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.String(string(in.Id))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4086215fEncodeAutogitJsonrpc(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4086215fEncodeAutogitJsonrpc(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4086215fDecodeAutogitJsonrpc(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4086215fDecodeAutogitJsonrpc(l, v)
}
func easyjson4086215fDecodeAutogitJsonrpc1(in *jlexer.Lexer, out *Request) {
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
		case "jsonrpc":
			out.JSONRPC = string(in.String())
		case "method":
			out.Method = string(in.String())
		case "params":
			if in.IsNull() {
				in.Skip()
				out.Params = nil
			} else {
				if out.Params == nil {
					out.Params = new(json.RawMessage)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Params).UnmarshalJSON(data))
				}
			}
		case "id":
			out.Id = Id(in.String())
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
func easyjson4086215fEncodeAutogitJsonrpc1(out *jwriter.Writer, in Request) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"jsonrpc\":")
	out.String(string(in.JSONRPC))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"method\":")
	out.String(string(in.Method))
	if in.Params != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"params\":")
		if in.Params == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Params).MarshalJSON())
		}
	}
	if in.Id != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"id\":")
		out.String(string(in.Id))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Request) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4086215fEncodeAutogitJsonrpc1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Request) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4086215fEncodeAutogitJsonrpc1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Request) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4086215fDecodeAutogitJsonrpc1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Request) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4086215fDecodeAutogitJsonrpc1(l, v)
}
func easyjson4086215fDecodeAutogitJsonrpc2(in *jlexer.Lexer, out *Error) {
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
		case "code":
			out.Code = int64(in.Int64())
		case "message":
			out.Message = string(in.String())
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				if out.Data == nil {
					out.Data = new(json.RawMessage)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Data).UnmarshalJSON(data))
				}
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
func easyjson4086215fEncodeAutogitJsonrpc2(out *jwriter.Writer, in Error) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"code\":")
	out.Int64(int64(in.Code))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"message\":")
	out.String(string(in.Message))
	if in.Data != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"data\":")
		if in.Data == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.Data).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Error) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4086215fEncodeAutogitJsonrpc2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Error) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4086215fEncodeAutogitJsonrpc2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Error) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4086215fDecodeAutogitJsonrpc2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Error) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4086215fDecodeAutogitJsonrpc2(l, v)
}
func easyjson4086215fDecodeAutogitJsonrpc3(in *jlexer.Lexer, out *BatchResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(BatchResponse, 0, 8)
			} else {
				*out = BatchResponse{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *Response
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(Response)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4086215fEncodeAutogitJsonrpc3(out *jwriter.Writer, in BatchResponse) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v BatchResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4086215fEncodeAutogitJsonrpc3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BatchResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4086215fEncodeAutogitJsonrpc3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BatchResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4086215fDecodeAutogitJsonrpc3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BatchResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4086215fDecodeAutogitJsonrpc3(l, v)
}
func easyjson4086215fDecodeAutogitJsonrpc4(in *jlexer.Lexer, out *BatchRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(BatchRequest, 0, 8)
			} else {
				*out = BatchRequest{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 *Request
			if in.IsNull() {
				in.Skip()
				v4 = nil
			} else {
				if v4 == nil {
					v4 = new(Request)
				}
				(*v4).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4086215fEncodeAutogitJsonrpc4(out *jwriter.Writer, in BatchRequest) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			if v6 == nil {
				out.RawString("null")
			} else {
				(*v6).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v BatchRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4086215fEncodeAutogitJsonrpc4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BatchRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4086215fEncodeAutogitJsonrpc4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BatchRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4086215fDecodeAutogitJsonrpc4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BatchRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4086215fDecodeAutogitJsonrpc4(l, v)
}
