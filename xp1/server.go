package xrpc

import (
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"xrpc/codec"
)

type Server struct {
	lis net.Listener
}

func NewServer(l net.Listener) *Server {
	return &Server{lis: l}
}

var defaultCodecType = codec.GobType

func (server *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }()

	//这里默认为gob解析，后面会说 如何根据客户端传过来的类型进行选择
	f := codec.NewCodecFuncMap[defaultCodecType]
	server.serveCodec(f(conn))
}

var invalidRequest = struct{}{}

func (server *Server) serveCodec(cc codec.Codec) {
	for {
		req, err := server.readRequest(cc)
		if err != nil {
			if req == nil {
				fmt.Printf("close服务 %s \n", err)
				break
			}
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, invalidRequest)
			continue
		}
		go server.handleRequest(cc, req)
	}
	cc.Close()
}

type request struct {
	h            *codec.Header
	argv, replyv reflect.Value
}

func (server *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println("rpc server: read header error:", err)
		}
		return nil, err
	}
	return &h, nil
}

func (server *Server) readRequest(cc codec.Codec) (*request, error) {
	h, err := server.readRequestHeader(cc)
	if err != nil {
		return nil, err
	}
	req := &request{h: h}
	// 这里暂时不提，后面会说
	req.argv = reflect.New(reflect.TypeOf(""))
	if err = cc.ReadBody(req.argv.Interface()); err != nil {
		log.Println("rpc server: read argv err:", err)
	}
	return req, nil
}

func (server *Server) sendResponse(cc codec.Codec, h *codec.Header, body interface{}) {
	if err := cc.Write(h, body); err != nil {
		log.Println("rpc server: write response error:", err)
	}
}

func (server *Server) handleRequest(cc codec.Codec, req *request) {
	log.Println(req.h, req.argv.Elem())
	req.replyv = reflect.ValueOf(fmt.Sprintf("xrpc resp %d", req.h.Seq))
	server.sendResponse(cc, req.h, req.replyv.Interface())
}

func (server *Server) Run() {
	for {
		conn, err := server.lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error:", err)
			return
		}
		go server.ServeConn(conn)
	}
}

func (server *Server) Close() {
	server.lis.Close()
}
