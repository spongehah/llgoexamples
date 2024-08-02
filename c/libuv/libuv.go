package libuv

import (
	"unsafe"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/libuv"
)

const (
	RUN_DEFAULT = libuv.RUN_DEFAULT
	RUN_ONCE    = libuv.RUN_ONCE
	RUN_NOWAIT  = libuv.RUN_NOWAIT
)

const (
	LOOP_BLOCK_SIGNAL = libuv.LOOP_BLOCK_SIGNAL
	METRICS_IDLE_TIME = libuv.METRICS_IDLE_TIME
)

const (
	UV_LEAVE_GROUP = libuv.UV_LEAVE_GROUP
	UV_JOIN_GROUP  = libuv.UV_JOIN_GROUP
)

const (
	UNKNOWN_HANDLE  = libuv.UNKNOWN_HANDLE
	ASYNC           = libuv.ASYNC
	CHECK           = libuv.CHECK
	FS_EVENT        = libuv.FS_EVENT
	FS_POLL         = libuv.FS_POLL
	HANDLE          = libuv.HANDLE
	IDLE            = libuv.IDLE
	NAMED_PIPE      = libuv.NAMED_PIPE
	POLL            = libuv.POLL
	PREPARE         = libuv.PREPARE
	PROCESS         = libuv.PROCESS
	STREAM          = libuv.STREAM
	TCP             = libuv.TCP
	TIMER           = libuv.TIMER
	TTY             = libuv.TTY
	UDP             = libuv.UDP
	SIGNAL          = libuv.SIGNAL
	FILE            = libuv.FILE
	HANDLE_TYPE_MAX = libuv.HANDLE_TYPE_MAX
)

const (
	UNKNOWN_REQ      = libuv.UNKNOWN_REQ
	REQ              = libuv.REQ
	CONNECT          = libuv.CONNECT
	WRITE            = libuv.WRITE
	SHUTDOWN         = libuv.SHUTDOWN
	UDP_SEND         = libuv.UDP_SEND
	FS               = libuv.FS
	WORK             = libuv.WORK
	GETADDRINFO      = libuv.GETADDRINFO
	GETNAMEINFO      = libuv.GETNAMEINFO
	RANDOM           = libuv.RANDOM
	REQ_TYPE_PRIVATE = libuv.REQ_TYPE_PRIVATE
	REQ_TYPE_MAX     = libuv.REQ_TYPE_MAX
)

const (
	READABLE       = libuv.READABLE
	WRITABLE       = libuv.WRITABLE
	DISCONNECT     = libuv.DISCONNECT
	PRIPRIORITIZED = libuv.PRIPRIORITIZED
)

type Loop struct {
	*libuv.Loop
}

type Handle struct {
	*libuv.Handle
	WalkCb WalkCb
}

type Stream struct {
	*libuv.Stream
	ConnectionCb func(server *Stream, status c.Int)
	ReadCb       func(stream *Stream, nread c.Long, buf *Buf)
}

type Poll struct {
	*libuv.Poll
}

type Req struct {
	*libuv.Req
}

type GetAddrInfo struct {
	*libuv.GetAddrInfo
}

type GetNameInfo struct {
	*libuv.GetNameInfo
}

type Shutdown struct {
	*libuv.Shutdown
}

type Write struct {
	_Write  *libuv.Write
	WriteCb func(req *Write, status c.Int)
}

type Connect struct {
	*libuv.Connect
}

type Buf struct {
	*libuv.Buf
}

type WalkCb func(handle *Handle, arg c.Pointer)

type ConnectionCb func(server *Stream, status c.Int)

type ReadCb func(stream *Stream, nread c.Long, buf *Buf)

type WriteCb func(req *Write, status c.Int)

func convertWalkCb(callback WalkCb) func(handle *libuv.Handle, arg c.Pointer) {
	return func(handle *libuv.Handle, arg c.Pointer) {
		hand := &Handle{Handle: handle}
		callback(hand, arg)
	}
}

// DefaultLoop returns the default loop.
func DefaultLoop() *Loop {
	return &Loop{Loop: libuv.DefaultLoop()}
}

// Size returns the size of the loop.
func (l *Loop) Size() uintptr {
	return libuv.LoopSize()
}

// Init initializes the loop.
func (l *Loop) Init() int {
	return int(libuv.LoopInit(l.Loop))
}

// Run runs the loop.
func (l *Loop) Run(mode libuv.RunMode) int {
	return int(libuv.Run(l.Loop, mode))
}

// Stop closes the loop.
func (l *Loop) Stop() int {
	return int(libuv.LoopClose(l.Loop))
}

// Default creates a new loop.
func (l *Loop) Default() *libuv.Loop {
	return libuv.LoopDefault()
}

// New creates a new loop.
func (l *Loop) New() *libuv.Loop {
	return libuv.LoopNew()
}

// Deprecated: use LoopClose instead.
// Delete closes the loop.
func (l *Loop) Delete() int {
	return int(libuv.LoopDelete(l.Loop))
}

// Alive returns the status of the loop.
func (l *Loop) Alive() int {
	return int(libuv.LoopAlive(l.Loop))
}

// Close closes the loop.
func (l *Loop) Close() int {
	return int(libuv.LoopClose(l.Loop))
}

// Configure configures the loop.
func (l *Loop) Configure(loop *Loop, option libuv.LoopOption, arg int) int {
	return int(libuv.LoopConfigure(l.Loop, option, c.Int(arg)))
}

// Walk walks the loop.
func (l *Loop) Walk(walkCb WalkCb, arg c.Pointer) {
	libuv.LoopWalk(l.Loop, func(_handle *libuv.Handle, arg c.Pointer) {
		handle := (*Handle)(unsafe.Pointer(_handle))
		handle.WalkCb(handle, arg)
	}, arg)
}

// Fork forks the loop.
func (l *Loop) Fork(loop *Loop) int {
	return int(libuv.LoopFork(l.Loop))
}

// UpdateTime updates the time of the loop.
func (l *Loop) UpdateTime() {
	libuv.LoopUpdateTime(l.Loop)
}

// Now returns the current time of the loop.
func (l *Loop) Now() uint64 {
	return uint64(libuv.LoopNow(l.Loop))
}

// BackendFd returns the backend file descriptor of the loop.
func (l *Loop) BackendFd() int {
	return int(libuv.LoopBackendFd(l.Loop))
}

// BackendTimeout returns the backend timeout of the loop.
func (l *Loop) BackendTimeout() int {
	return int(libuv.LoopBackendTimeout(l.Loop))
}

// ----------------------------------------------

/* Buf related functions and method. */

// InitBuf initializes a buffer with the given c.Char slice.
func InitBuf(buffer []c.Char) Buf {
	buf := libuv.InitBuf((*c.Char)(unsafe.Pointer(&buffer[0])), c.Uint(unsafe.Sizeof(buffer)))
	return Buf{Buf: &buf}
}

// ----------------------------------------------

/* Handle related functions and method. */

// Ref references the handle.
func (h *Handle) Ref() {
	h.Handle.Ref()
}

// Unref unreferences the handle.
func (h *Handle) Unref() {
	h.Handle.Unref()
}

// HasRef returns true if the handle has a reference.
func (h *Handle) HasRef() int {
	return int(h.Handle.HasRef())
}

// HandleSize returns the size of the handle.
func HandleSize(handleType libuv.HandleType) uintptr {
	return libuv.HandleSize(handleType)
}

// GetType returns the type of the handle.
func (h *Handle) GetType() libuv.HandleType {
	return h.Handle.GetType()
}

// HandleTypeName returns the name of the handle type.
func HandleTypeName(handleType libuv.HandleType) string {
	return c.GoString(libuv.HandleTypeName(handleType))
}

// GetData returns the data of the handle.
func (h *Handle) GetData() c.Pointer {
	return h.Handle.GetData()
}

// GetLoop returns the loop of the handle.
func (h *Handle) GetLoop() *Loop {
	return &Loop{Loop: h.Handle.GetLoop()}
}

// SetData sets the data of the handle.
func (h *Handle) SetData(data c.Pointer) {
	h.Handle.SetData(data)
}

// IsActive returns true if the handle is active.
func (h *Handle) IsActive() int {
	return int(h.Handle.IsActive())
}

// Close closes the handle.
func (h *Handle) Close(closeCb libuv.CloseCb) {
	h.Handle.Close(closeCb)
}

// SendBufferSize returns the send buffer size of the handle.
func (h *Handle) SendBufferSize(value *c.Int) int {
	return int(h.Handle.SendBufferSize(value))
}

// RecvBufferSize returns the receive buffer size of the handle.
func (h *Handle) RecvBufferSize(value *c.Int) int {
	return int(h.Handle.RecvBufferSize(value))
}

// Fileno returns the file number of the handle.
func (h *Handle) Fileno(fd *libuv.OsFd) int {
	return int(h.Handle.Fileno(fd))
}

// Pipe creates a new pipe.
func Pipe(fds [2]libuv.File, readFlags int, writeFlags int) int {
	return int(libuv.Pipe(fds, c.Int(readFlags), c.Int(writeFlags)))
}

// Socketpair creates a new socket pair.
func Socketpair(_type int, protocol int, socketVector [2]libuv.OsSock, flag0 int, flag1 int) int {
	return int(libuv.Socketpair(c.Int(_type), c.Int(protocol), socketVector, c.Int(flag0), c.Int(flag1)))
}

// IsClosing returns true if the handle is closing.
func (h *Handle) IsClosing() int {
	return int(h.Handle.IsClosing())
}

// ----------------------------------------------

/* Req related functions and method. */

// ReqSize returns the size of the request.
func ReqSize(reqType libuv.ReqType) uintptr {
	return libuv.ReqSize(reqType)
}

// GetData returns the data of the request.
func (req *Req) GetData() c.Pointer {
	return req.Req.GetData()
}

// SetData sets the data of the request.
func (req *Req) SetData(data c.Pointer) {
	req.Req.SetData(data)
}

// GetType returns the type of the request.
func (req *Req) GetType() libuv.ReqType {
	return req.Req.GetType()
}

// TypeName returns the name of the request type.
func TypeName(reqType libuv.ReqType) string {
	return c.GoString(libuv.TypeName(reqType))
}

// ----------------------------------------------

/* Stream related functions and method. */

// GetWriteQueueSize returns the size of the write queue.
func (s *Stream) GetWriteQueueSize() uintptr {
	return s.Stream.GetWriteQueueSize()
}

// Listen listens to the stream.
func (s *Stream) Listen(backlog int, connectionCb ConnectionCb) int {
	s.ConnectionCb = connectionCb
	return int(s.Stream.Listen(c.Int(backlog), func(_server *libuv.Stream, status c.Int) {
		server := (*Stream)(unsafe.Pointer(_server))
		server.ConnectionCb(server, status)
	}))
}

// Accept accepts the stream.
func (server *Stream) Accept(client *Stream) int {
	return int(server.Stream.Accept(client.Stream))
}

// StartRead starts reading from the stream.
func (s *Stream) StartRead(allocCb libuv.AllocCb, readCb ReadCb) int {
	s.ReadCb = readCb
	return int(s.Stream.StartRead(allocCb, func(_stream *libuv.Stream, nread c.Long, buf *libuv.Buf) {
		stream := (*Stream)(unsafe.Pointer(_stream))
		stream.ReadCb(stream, nread, &Buf{Buf: buf})
	}))
}

// StopRead stops reading from the stream.
func (s *Stream) StopRead() int {
	return int(s.Stream.StopRead())
}

// Write writes to the stream.
func (w *Write) Write(stream *Stream, bufs *Buf, nbufs int, writeCb WriteCb) int {
	w.WriteCb = writeCb
	return int(w._Write.Write(stream.Stream, bufs.Buf, c.Uint(nbufs), func(_req *libuv.Write, status c.Int) {
		req := (*Write)(unsafe.Pointer(_req))
		req.WriteCb(req, status)
	}))
}

// Write2 writes to the stream.
func (w *Write) Write2(stream *Stream, bufs *Buf, nbufs int, sendStream *Stream, writeCb WriteCb) int {
	w.WriteCb = writeCb
	return int(w._Write.Write2(stream.Stream, bufs.Buf, c.Uint(nbufs), sendStream.Stream, func(_req *libuv.Write, status c.Int) {
		req := (*Write)(unsafe.Pointer(_req))
		req.WriteCb(req, status)
	}))
}

// TryWrite tries to write to the stream.
func (s *Stream) TryWrite(bufs *Buf, nbufs uint) int {
	return int(s.Stream.TryWrite(bufs.Buf, c.Uint(nbufs)))
}

// TryWrite2 tries to write to the stream.
func (s *Stream) TryWrite2(bufs *Buf, nbufs uint, sendStream *Stream) int {
	return int(s.Stream.TryWrite2(bufs.Buf, c.Uint(nbufs), sendStream.Stream))
}

// IsReadable returns true if the stream is readable.
func (s *Stream) IsReadable() int {
	return int(s.Stream.IsReadable())
}

// IsWritable returns true if the stream is writable.
func (s *Stream) IsWritable() int {
	return int(s.Stream.IsWritable())
}

// SetBlocking sets the blocking status of the stream.
func (s *Stream) SetBlocking(blocking int) int {
	return int(s.Stream.SetBlocking(c.Int(blocking)))
}

// ----------------------------------------------