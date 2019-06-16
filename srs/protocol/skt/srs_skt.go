package skt

import (
	"bufio"
	"io"
	"net"
	"time"
)

type SrsIOReadWriter struct {
	conn     net.Conn
	IOReader *bufio.Reader
	IOWriter *bufio.Writer
}

func NewSrsIOReadWriter(c net.Conn) *SrsIOReadWriter {
	rw := &SrsIOReadWriter{
		conn:c,
	}
	rw.IOReader = bufio.NewReader(rw.conn)
	rw.IOWriter = bufio.NewWriter(rw.conn)
	return rw
}

func (this *SrsIOReadWriter) Read(b []byte) (int, error) {
	return this.IOReader.Read(b)
}

func (this *SrsIOReadWriter) ReadWithTimeout(b []byte, timeoutms uint32) (int, error) {
	this.conn.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(timeoutms)))
	return this.IOReader.Read(b)
}

func (this *SrsIOReadWriter) ReadFully(b []byte, timeoutms uint32) (int, error) {
	return io.ReadFull(this.conn, b)
}

func (this *SrsIOReadWriter) ReadFullyWithTimeout(b []byte, timeoutms uint32) (int, error) {
	this.conn.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(timeoutms)))
	return io.ReadFull(this.conn, b)
}

func (this *SrsIOReadWriter) Write(b []byte) (int, error) {
	return this.IOWriter.Write(b)
}

func (this *SrsIOReadWriter) WriteWithTimeout(b []byte, timeoutms uint32) (int, error) {
	this.conn.SetWriteDeadline(time.Now().Add(time.Millisecond * time.Duration(timeoutms)))
	return this.IOWriter.Write(b)
}