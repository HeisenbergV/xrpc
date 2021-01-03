package xrpc

import "xrpc/codec"

type Option struct {
	MagicNumber int
	CodecType   codec.Type
}

var DefaultOption = &Option{}
