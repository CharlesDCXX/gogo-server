package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

/*
*
encoding/gob 是 Go 语言标准库中的一个包，用于将 Go 值序列化（编码）和反序列化（解码）。它提供了一种将数据结构转换为字节流的方法，以便在不同的系统之间传输或存储数据。

具体来说，encoding/gob 提供了以下功能：

结构体序列化和反序列化：您可以使用 gob.Register 注册自定义的结构体类型，然后使用 gob.Encoder 将结构体转换为字节流，并使用 gob.Decoder 将字节流转换回结构体。

值类型序列化和反序列化：除了结构体，encoding/gob 还支持对其他基本类型（如整数、浮点数、字符串等）和复合类型（如数组、切片、映射等）的序列化和反序列化。

数据流编码：encoding/gob 提供了 gob.NewEncoder 和 gob.NewDecoder 函数，用于将数据流（如文件、网络连接）与 gob.Encoder 和 gob.Decoder 关联，实现对数据的编码和解码。
*
*/
type GobCodec struct {
	conn io.ReadWriteCloser //实现基本的输入输出方法接口
	buf  *bufio.Writer      //对io.writer实现缓冲
	dec  *gob.Decoder       // gob解码
	enc  *gob.Encoder
}

var _ Codec = (*GobCodec)(nil)

// ReadHeader 方法使用 gob.Decoder 对象从 GobCodec 的 dec 字段中解码字节流，并将解码后的结果存储在 Header 类型的变量中。
func (c *GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

// 同上
func (c *GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

func (c *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()
	if err = c.enc.Encode(h); err != nil {
		log.Println("rpc: gob error encoding header:", err)
		return
	}
	if err = c.enc.Encode(body); err != nil {
		log.Println("rpc: gob error encoding body:", err)
		return
	}
	return
}

func (c *GobCodec) Close() error {
	return c.conn.Close()
}
