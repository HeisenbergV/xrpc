package xrpc

import "xrpc/codec"

type Option struct {
	MagicNumber int
	CodecType   codec.Type
}

const MagicNumber = 0x3bef5c

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}
