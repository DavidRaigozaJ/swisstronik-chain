// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: swisstronik/vesting/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreateMonthlyVestingAccount struct {
	Creator   string                                   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ToAddress string                                   `protobuf:"bytes,2,opt,name=toAddress,proto3" json:"toAddress,omitempty"`
	StartTime int64                                    `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Amount    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	Month     int64                                    `protobuf:"varint,5,opt,name=month,proto3" json:"month,omitempty"`
}

func (m *MsgCreateMonthlyVestingAccount) Reset()         { *m = MsgCreateMonthlyVestingAccount{} }
func (m *MsgCreateMonthlyVestingAccount) String() string { return proto.CompactTextString(m) }
func (*MsgCreateMonthlyVestingAccount) ProtoMessage()    {}
func (*MsgCreateMonthlyVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_899431b613aba5b5, []int{0}
}
func (m *MsgCreateMonthlyVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateMonthlyVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateMonthlyVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateMonthlyVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateMonthlyVestingAccount.Merge(m, src)
}
func (m *MsgCreateMonthlyVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateMonthlyVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateMonthlyVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateMonthlyVestingAccount proto.InternalMessageInfo

func (m *MsgCreateMonthlyVestingAccount) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateMonthlyVestingAccount) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *MsgCreateMonthlyVestingAccount) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *MsgCreateMonthlyVestingAccount) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *MsgCreateMonthlyVestingAccount) GetMonth() int64 {
	if m != nil {
		return m.Month
	}
	return 0
}

type MsgCreateMonthlyVestingAccountResponse struct {
}

func (m *MsgCreateMonthlyVestingAccountResponse) Reset() {
	*m = MsgCreateMonthlyVestingAccountResponse{}
}
func (m *MsgCreateMonthlyVestingAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateMonthlyVestingAccountResponse) ProtoMessage()    {}
func (*MsgCreateMonthlyVestingAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_899431b613aba5b5, []int{1}
}
func (m *MsgCreateMonthlyVestingAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateMonthlyVestingAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateMonthlyVestingAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateMonthlyVestingAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateMonthlyVestingAccountResponse.Merge(m, src)
}
func (m *MsgCreateMonthlyVestingAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateMonthlyVestingAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateMonthlyVestingAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateMonthlyVestingAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateMonthlyVestingAccount)(nil), "swisstronik.vesting.MsgCreateMonthlyVestingAccount")
	proto.RegisterType((*MsgCreateMonthlyVestingAccountResponse)(nil), "swisstronik.vesting.MsgCreateMonthlyVestingAccountResponse")
}

func init() { proto.RegisterFile("swisstronik/vesting/tx.proto", fileDescriptor_899431b613aba5b5) }

var fileDescriptor_899431b613aba5b5 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbd, 0x6e, 0xe2, 0x40,
	0x10, 0xc7, 0xbd, 0xe7, 0x83, 0x13, 0x7b, 0x9d, 0x8f, 0xc2, 0x07, 0xdc, 0x82, 0x28, 0x4e, 0x6e,
	0xb2, 0x1b, 0x40, 0xa9, 0x52, 0x01, 0x35, 0x8d, 0x15, 0xa5, 0x48, 0x13, 0xd9, 0x66, 0x65, 0x56,
	0xc4, 0x1e, 0xe4, 0x59, 0x08, 0x3c, 0x45, 0xd2, 0xe5, 0x1d, 0xf2, 0x24, 0x94, 0x94, 0xa9, 0x92,
	0x08, 0xde, 0x20, 0x4f, 0x10, 0xf9, 0x03, 0x41, 0x11, 0x51, 0xa4, 0xb2, 0x67, 0x7e, 0x3b, 0x33,
	0xff, 0xf9, 0xa0, 0x0d, 0xbc, 0x57, 0x88, 0x3a, 0x81, 0x58, 0x4d, 0xc5, 0x42, 0xa2, 0x56, 0x71,
	0x28, 0xf4, 0x92, 0xcf, 0x12, 0xd0, 0x60, 0xfd, 0x39, 0xa2, 0xbc, 0xa0, 0xb5, 0x6a, 0x08, 0x21,
	0x64, 0x5c, 0xa4, 0x7f, 0xf9, 0xd3, 0x1a, 0x0b, 0x00, 0x23, 0x40, 0xe1, 0x7b, 0x28, 0xc5, 0xa2,
	0xe3, 0x4b, 0xed, 0x75, 0x44, 0x00, 0x2a, 0xce, 0x79, 0xfb, 0x83, 0x50, 0x36, 0xc2, 0x70, 0x98,
	0x48, 0x4f, 0xcb, 0x11, 0xc4, 0x7a, 0x72, 0xb7, 0xba, 0xce, 0x33, 0xf6, 0x83, 0x00, 0xe6, 0xb1,
	0xb6, 0x6c, 0xfa, 0x2b, 0x48, 0x31, 0x24, 0x36, 0x69, 0x11, 0xa7, 0xe2, 0xee, 0x4d, 0xab, 0x41,
	0x2b, 0x1a, 0xfa, 0xe3, 0x71, 0x22, 0x11, 0xed, 0x1f, 0x19, 0x3b, 0x38, 0xac, 0x7f, 0x94, 0xa2,
	0xf6, 0x12, 0x7d, 0xab, 0x55, 0x24, 0x6d, 0xb3, 0x45, 0x1c, 0xd3, 0xad, 0x64, 0x9e, 0x2b, 0x15,
	0x49, 0x2b, 0xa0, 0x65, 0x2f, 0x4a, 0x0b, 0xd8, 0x3f, 0x5b, 0xa6, 0xf3, 0xbb, 0xfb, 0x97, 0xe7,
	0x52, 0x79, 0x2a, 0x95, 0x17, 0x52, 0xf9, 0x10, 0x54, 0x3c, 0x38, 0x5f, 0xbf, 0x36, 0x8d, 0xe7,
	0xb7, 0xa6, 0x13, 0x2a, 0x3d, 0x99, 0xfb, 0x3c, 0x80, 0x48, 0x14, 0x7d, 0xe5, 0x9f, 0x33, 0x1c,
	0x4f, 0x85, 0x5e, 0xcd, 0x24, 0x66, 0x01, 0xe8, 0x16, 0xa9, 0xad, 0x2a, 0x2d, 0x45, 0x69, 0x53,
	0x76, 0x29, 0x2b, 0x9f, 0x1b, 0x6d, 0x87, 0xfe, 0x3f, 0xdd, 0xb3, 0x2b, 0x71, 0x06, 0x31, 0xca,
	0xee, 0x13, 0xa1, 0xe6, 0x08, 0x43, 0xeb, 0x81, 0xd0, 0xfa, 0xa9, 0x19, 0xf5, 0xf8, 0x17, 0x2b,
	0xe1, 0xa7, 0x8b, 0xd4, 0x2e, 0xbf, 0x11, 0xb4, 0x57, 0x36, 0xb8, 0x58, 0x6f, 0x19, 0xd9, 0x6c,
	0x19, 0x79, 0xdf, 0x32, 0xf2, 0xb8, 0x63, 0xc6, 0x66, 0xc7, 0x8c, 0x97, 0x1d, 0x33, 0x6e, 0xea,
	0xc7, 0xb7, 0xb3, 0x3c, 0x5c, 0x4f, 0x3a, 0x1e, 0xbf, 0x9c, 0xad, 0xbd, 0xf7, 0x19, 0x00, 0x00,
	0xff, 0xff, 0x50, 0x99, 0x96, 0x74, 0x61, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateMonthlyVestingAccount(ctx context.Context, in *MsgCreateMonthlyVestingAccount, opts ...grpc.CallOption) (*MsgCreateMonthlyVestingAccountResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateMonthlyVestingAccount(ctx context.Context, in *MsgCreateMonthlyVestingAccount, opts ...grpc.CallOption) (*MsgCreateMonthlyVestingAccountResponse, error) {
	out := new(MsgCreateMonthlyVestingAccountResponse)
	err := c.cc.Invoke(ctx, "/swisstronik.vesting.Msg/CreateMonthlyVestingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateMonthlyVestingAccount(context.Context, *MsgCreateMonthlyVestingAccount) (*MsgCreateMonthlyVestingAccountResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateMonthlyVestingAccount(ctx context.Context, req *MsgCreateMonthlyVestingAccount) (*MsgCreateMonthlyVestingAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMonthlyVestingAccount not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateMonthlyVestingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateMonthlyVestingAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateMonthlyVestingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swisstronik.vesting.Msg/CreateMonthlyVestingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateMonthlyVestingAccount(ctx, req.(*MsgCreateMonthlyVestingAccount))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "swisstronik.vesting.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMonthlyVestingAccount",
			Handler:    _Msg_CreateMonthlyVestingAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "swisstronik/vesting/tx.proto",
}

func (m *MsgCreateMonthlyVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateMonthlyVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateMonthlyVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Month != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Month))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.StartTime != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateMonthlyVestingAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateMonthlyVestingAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateMonthlyVestingAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateMonthlyVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.StartTime != 0 {
		n += 1 + sovTx(uint64(m.StartTime))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.Month != 0 {
		n += 1 + sovTx(uint64(m.Month))
	}
	return n
}

func (m *MsgCreateMonthlyVestingAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateMonthlyVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateMonthlyVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateMonthlyVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Month", wireType)
			}
			m.Month = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Month |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateMonthlyVestingAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateMonthlyVestingAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateMonthlyVestingAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
