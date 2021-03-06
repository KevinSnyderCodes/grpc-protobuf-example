// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plane/plane.proto

package plane

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Coordinates struct {
	X                    int64    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int64    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coordinates) Reset()         { *m = Coordinates{} }
func (m *Coordinates) String() string { return proto.CompactTextString(m) }
func (*Coordinates) ProtoMessage()    {}
func (*Coordinates) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1f3dea44b14fc9, []int{0}
}

func (m *Coordinates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinates.Unmarshal(m, b)
}
func (m *Coordinates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinates.Marshal(b, m, deterministic)
}
func (m *Coordinates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinates.Merge(m, src)
}
func (m *Coordinates) XXX_Size() int {
	return xxx_messageInfo_Coordinates.Size(m)
}
func (m *Coordinates) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinates.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinates proto.InternalMessageInfo

func (m *Coordinates) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Coordinates) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Point struct {
	Coordinates          *Coordinates `protobuf:"bytes,1,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	Count                uint64       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1f3dea44b14fc9, []int{1}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetCoordinates() *Coordinates {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *Point) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Bounds struct {
	Min                  *Coordinates `protobuf:"bytes,1,opt,name=min,proto3" json:"min,omitempty"`
	Max                  *Coordinates `protobuf:"bytes,2,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Bounds) Reset()         { *m = Bounds{} }
func (m *Bounds) String() string { return proto.CompactTextString(m) }
func (*Bounds) ProtoMessage()    {}
func (*Bounds) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea1f3dea44b14fc9, []int{2}
}

func (m *Bounds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bounds.Unmarshal(m, b)
}
func (m *Bounds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bounds.Marshal(b, m, deterministic)
}
func (m *Bounds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bounds.Merge(m, src)
}
func (m *Bounds) XXX_Size() int {
	return xxx_messageInfo_Bounds.Size(m)
}
func (m *Bounds) XXX_DiscardUnknown() {
	xxx_messageInfo_Bounds.DiscardUnknown(m)
}

var xxx_messageInfo_Bounds proto.InternalMessageInfo

func (m *Bounds) GetMin() *Coordinates {
	if m != nil {
		return m.Min
	}
	return nil
}

func (m *Bounds) GetMax() *Coordinates {
	if m != nil {
		return m.Max
	}
	return nil
}

func init() {
	proto.RegisterType((*Coordinates)(nil), "Coordinates")
	proto.RegisterType((*Point)(nil), "Point")
	proto.RegisterType((*Bounds)(nil), "Bounds")
}

func init() { proto.RegisterFile("plane/plane.proto", fileDescriptor_ea1f3dea44b14fc9) }

var fileDescriptor_ea1f3dea44b14fc9 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xc8, 0x49, 0xcc,
	0x4b, 0xd5, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x4a, 0x9a, 0x5c, 0xdc, 0xce, 0xf9,
	0xf9, 0x45, 0x29, 0x99, 0x79, 0x89, 0x25, 0xa9, 0xc5, 0x42, 0x3c, 0x5c, 0x8c, 0x15, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x8c, 0x15, 0x20, 0x5e, 0xa5, 0x04, 0x13, 0x84, 0x57, 0xa9, 0xe4,
	0xcb, 0xc5, 0x1a, 0x90, 0x9f, 0x99, 0x57, 0x22, 0xa4, 0xc7, 0xc5, 0x9d, 0x8c, 0xd0, 0x03, 0x56,
	0xce, 0x6d, 0xc4, 0xa3, 0x87, 0x64, 0x4e, 0x10, 0xb2, 0x02, 0x21, 0x11, 0x2e, 0xd6, 0xe4, 0xfc,
	0xd2, 0xbc, 0x12, 0xb0, 0x51, 0x2c, 0x41, 0x10, 0x8e, 0x92, 0x07, 0x17, 0x9b, 0x53, 0x7e, 0x69,
	0x5e, 0x4a, 0xb1, 0x90, 0x1c, 0x17, 0x73, 0x6e, 0x66, 0x1e, 0x56, 0x73, 0x40, 0x12, 0x60, 0xf9,
	0xc4, 0x0a, 0xb0, 0x6e, 0x4c, 0xf9, 0xc4, 0x0a, 0xa3, 0x15, 0x8c, 0x5c, 0xac, 0x01, 0x20, 0x3f,
	0x09, 0x29, 0x71, 0x71, 0xb8, 0xa7, 0x96, 0x40, 0x5c, 0x89, 0xa2, 0x50, 0x8a, 0x4d, 0x0f, 0x2c,
	0xaa, 0xc4, 0x20, 0xa4, 0xc6, 0xc5, 0x19, 0x50, 0x0a, 0x51, 0x53, 0x8c, 0xa6, 0x88, 0x5d, 0x0f,
	0xe2, 0x22, 0x25, 0x06, 0x0d, 0x46, 0x21, 0x4d, 0x2e, 0x21, 0x9f, 0xcc, 0x62, 0xa8, 0x42, 0xa7,
	0x4a, 0xa8, 0x5b, 0x61, 0x4a, 0x10, 0x06, 0x1a, 0x30, 0x0a, 0x69, 0x70, 0x71, 0x21, 0x94, 0xe2,
	0xb2, 0x58, 0x83, 0xd1, 0x80, 0x31, 0x89, 0x0d, 0x1c, 0xea, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa5, 0x26, 0x2a, 0x4d, 0x8a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlaneClient is the client API for Plane service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlaneClient interface {
	GetPoint(ctx context.Context, in *Coordinates, opts ...grpc.CallOption) (*Point, error)
	PutPoints(ctx context.Context, opts ...grpc.CallOption) (Plane_PutPointsClient, error)
	ListPointsByBounds(ctx context.Context, in *Bounds, opts ...grpc.CallOption) (Plane_ListPointsByBoundsClient, error)
	ListPoints(ctx context.Context, opts ...grpc.CallOption) (Plane_ListPointsClient, error)
}

type planeClient struct {
	cc *grpc.ClientConn
}

func NewPlaneClient(cc *grpc.ClientConn) PlaneClient {
	return &planeClient{cc}
}

func (c *planeClient) GetPoint(ctx context.Context, in *Coordinates, opts ...grpc.CallOption) (*Point, error) {
	out := new(Point)
	err := c.cc.Invoke(ctx, "/Plane/GetPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planeClient) PutPoints(ctx context.Context, opts ...grpc.CallOption) (Plane_PutPointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Plane_serviceDesc.Streams[0], "/Plane/PutPoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &planePutPointsClient{stream}
	return x, nil
}

type Plane_PutPointsClient interface {
	Send(*Coordinates) error
	CloseAndRecv() (*Bounds, error)
	grpc.ClientStream
}

type planePutPointsClient struct {
	grpc.ClientStream
}

func (x *planePutPointsClient) Send(m *Coordinates) error {
	return x.ClientStream.SendMsg(m)
}

func (x *planePutPointsClient) CloseAndRecv() (*Bounds, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Bounds)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *planeClient) ListPointsByBounds(ctx context.Context, in *Bounds, opts ...grpc.CallOption) (Plane_ListPointsByBoundsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Plane_serviceDesc.Streams[1], "/Plane/ListPointsByBounds", opts...)
	if err != nil {
		return nil, err
	}
	x := &planeListPointsByBoundsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Plane_ListPointsByBoundsClient interface {
	Recv() (*Point, error)
	grpc.ClientStream
}

type planeListPointsByBoundsClient struct {
	grpc.ClientStream
}

func (x *planeListPointsByBoundsClient) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *planeClient) ListPoints(ctx context.Context, opts ...grpc.CallOption) (Plane_ListPointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Plane_serviceDesc.Streams[2], "/Plane/ListPoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &planeListPointsClient{stream}
	return x, nil
}

type Plane_ListPointsClient interface {
	Send(*Coordinates) error
	Recv() (*Point, error)
	grpc.ClientStream
}

type planeListPointsClient struct {
	grpc.ClientStream
}

func (x *planeListPointsClient) Send(m *Coordinates) error {
	return x.ClientStream.SendMsg(m)
}

func (x *planeListPointsClient) Recv() (*Point, error) {
	m := new(Point)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PlaneServer is the server API for Plane service.
type PlaneServer interface {
	GetPoint(context.Context, *Coordinates) (*Point, error)
	PutPoints(Plane_PutPointsServer) error
	ListPointsByBounds(*Bounds, Plane_ListPointsByBoundsServer) error
	ListPoints(Plane_ListPointsServer) error
}

func RegisterPlaneServer(s *grpc.Server, srv PlaneServer) {
	s.RegisterService(&_Plane_serviceDesc, srv)
}

func _Plane_GetPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Coordinates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaneServer).GetPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Plane/GetPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaneServer).GetPoint(ctx, req.(*Coordinates))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plane_PutPoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PlaneServer).PutPoints(&planePutPointsServer{stream})
}

type Plane_PutPointsServer interface {
	SendAndClose(*Bounds) error
	Recv() (*Coordinates, error)
	grpc.ServerStream
}

type planePutPointsServer struct {
	grpc.ServerStream
}

func (x *planePutPointsServer) SendAndClose(m *Bounds) error {
	return x.ServerStream.SendMsg(m)
}

func (x *planePutPointsServer) Recv() (*Coordinates, error) {
	m := new(Coordinates)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Plane_ListPointsByBounds_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Bounds)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PlaneServer).ListPointsByBounds(m, &planeListPointsByBoundsServer{stream})
}

type Plane_ListPointsByBoundsServer interface {
	Send(*Point) error
	grpc.ServerStream
}

type planeListPointsByBoundsServer struct {
	grpc.ServerStream
}

func (x *planeListPointsByBoundsServer) Send(m *Point) error {
	return x.ServerStream.SendMsg(m)
}

func _Plane_ListPoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PlaneServer).ListPoints(&planeListPointsServer{stream})
}

type Plane_ListPointsServer interface {
	Send(*Point) error
	Recv() (*Coordinates, error)
	grpc.ServerStream
}

type planeListPointsServer struct {
	grpc.ServerStream
}

func (x *planeListPointsServer) Send(m *Point) error {
	return x.ServerStream.SendMsg(m)
}

func (x *planeListPointsServer) Recv() (*Coordinates, error) {
	m := new(Coordinates)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Plane_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Plane",
	HandlerType: (*PlaneServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPoint",
			Handler:    _Plane_GetPoint_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PutPoints",
			Handler:       _Plane_PutPoints_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ListPointsByBounds",
			Handler:       _Plane_ListPointsByBounds_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListPoints",
			Handler:       _Plane_ListPoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "plane/plane.proto",
}
