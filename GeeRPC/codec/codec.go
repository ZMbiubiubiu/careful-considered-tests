package codec

import "io"

type Header struct {
	ServiceMethod string // format service.method
	Seq           uint64 // sequence number chosen by client,用来区分不同的请求
	Error         string
}

// Codec 对消息体进行编解码
type Codec interface {
	io.Closer
	ReadHeader(*Header) error         // 解码消息Header中
	ReadBody(interface{}) error       // 解码消息到body中
	Write(*Header, interface{}) error // 编码消息
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
