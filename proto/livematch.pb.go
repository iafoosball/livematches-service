// Code generated by protoc-gen-go. DO NOT EDIT.
// source: livematch.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Mode int32

const (
	Mode_defaultMode           Mode = 0
	Mode_oneOnOne              Mode = 1
	Mode_twoOnOne              Mode = 2
	Mode_twoOneTwo             Mode = 3
	Mode_tournamentModeTwoRed  Mode = 4
	Mode_tournamentModeTwoBlue Mode = 5
)

var Mode_name = map[int32]string{
	0: "defaultMode",
	1: "oneOnOne",
	2: "twoOnOne",
	3: "twoOneTwo",
	4: "tournamentModeTwoRed",
	5: "tournamentModeTwoBlue",
}
var Mode_value = map[string]int32{
	"defaultMode":           0,
	"oneOnOne":              1,
	"twoOnOne":              2,
	"twoOneTwo":             3,
	"tournamentModeTwoRed":  4,
	"tournamentModeTwoBlue": 5,
}

func (x Mode) String() string {
	return proto.EnumName(Mode_name, int32(x))
}
func (Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_livematch_e2e59194b3f622d4, []int{0}
}

type Position int32

const (
	Position_defaultPosition Position = 0
	Position_blueAttack      Position = 1
	Position_blueDefense     Position = 2
	Position_redAttach       Position = 3
	Position_redDefense      Position = 4
	Position_spectator       Position = 5
)

var Position_name = map[int32]string{
	0: "defaultPosition",
	1: "blueAttack",
	2: "blueDefense",
	3: "redAttach",
	4: "redDefense",
	5: "spectator",
}
var Position_value = map[string]int32{
	"defaultPosition": 0,
	"blueAttack":      1,
	"blueDefense":     2,
	"redAttach":       3,
	"redDefense":      4,
	"spectator":       5,
}

func (x Position) String() string {
	return proto.EnumName(Position_name, int32(x))
}
func (Position) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_livematch_e2e59194b3f622d4, []int{1}
}

type Match struct {
	XId                  string            `protobuf:"bytes,1,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	XKey                 string            `protobuf:"bytes,2,opt,name=_key,json=Key,proto3" json:"_key,omitempty"`
	TableID              string            `protobuf:"bytes,3,opt,name=tableID,proto3" json:"tableID,omitempty"`
	Started              bool              `protobuf:"varint,4,opt,name=started,proto3" json:"started,omitempty"`
	Users                []*User           `protobuf:"bytes,5,rep,name=users,proto3" json:"users,omitempty"`
	Settings             *Settings         `protobuf:"bytes,6,opt,name=settings,proto3" json:"settings,omitempty"`
	StartTime            string            `protobuf:"bytes,7,opt,name=startTime,proto3" json:"startTime,omitempty"`
	ScoreBlue            int32             `protobuf:"varint,8,opt,name=scoreBlue,proto3" json:"scoreBlue,omitempty"`
	ScoreRed             int32             `protobuf:"varint,9,opt,name=scoreRed,proto3" json:"scoreRed,omitempty"`
	EndTime              string            `protobuf:"bytes,10,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Completed            bool              `protobuf:"varint,11,opt,name=completed,proto3" json:"completed,omitempty"`
	Winner               string            `protobuf:"bytes,12,opt,name=winner,proto3" json:"winner,omitempty"`
	Labels               map[string]string `protobuf:"bytes,14,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Match) Reset()         { *m = Match{} }
func (m *Match) String() string { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()    {}
func (*Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_livematch_e2e59194b3f622d4, []int{0}
}
func (m *Match) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Match.Unmarshal(m, b)
}
func (m *Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Match.Marshal(b, m, deterministic)
}
func (dst *Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Match.Merge(dst, src)
}
func (m *Match) XXX_Size() int {
	return xxx_messageInfo_Match.Size(m)
}
func (m *Match) XXX_DiscardUnknown() {
	xxx_messageInfo_Match.DiscardUnknown(m)
}

var xxx_messageInfo_Match proto.InternalMessageInfo

func (m *Match) GetXId() string {
	if m != nil {
		return m.XId
	}
	return ""
}

func (m *Match) GetXKey() string {
	if m != nil {
		return m.XKey
	}
	return ""
}

func (m *Match) GetTableID() string {
	if m != nil {
		return m.TableID
	}
	return ""
}

func (m *Match) GetStarted() bool {
	if m != nil {
		return m.Started
	}
	return false
}

func (m *Match) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Match) GetSettings() *Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *Match) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Match) GetScoreBlue() int32 {
	if m != nil {
		return m.ScoreBlue
	}
	return 0
}

func (m *Match) GetScoreRed() int32 {
	if m != nil {
		return m.ScoreRed
	}
	return 0
}

func (m *Match) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Match) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *Match) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

func (m *Match) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Admin                bool     `protobuf:"varint,3,opt,name=admin,proto3" json:"admin,omitempty"`
	Ready                bool     `protobuf:"varint,4,opt,name=ready,proto3" json:"ready,omitempty"`
	Position             Position `protobuf:"varint,5,opt,name=position,proto3,enum=proto.Position" json:"position,omitempty"`
	Bet                  int64    `protobuf:"varint,6,opt,name=bet,proto3" json:"bet,omitempty"`
	Color                string   `protobuf:"bytes,7,opt,name=color,proto3" json:"color,omitempty"`
	CurrentTable         string   `protobuf:"bytes,8,opt,name=currentTable,proto3" json:"currentTable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_livematch_e2e59194b3f622d4, []int{1}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetAdmin() bool {
	if m != nil {
		return m.Admin
	}
	return false
}

func (m *User) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

func (m *User) GetPosition() Position {
	if m != nil {
		return m.Position
	}
	return Position_defaultPosition
}

func (m *User) GetBet() int64 {
	if m != nil {
		return m.Bet
	}
	return 0
}

func (m *User) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *User) GetCurrentTable() string {
	if m != nil {
		return m.CurrentTable
	}
	return ""
}

type Settings struct {
	Tournament           bool     `protobuf:"varint,1,opt,name=tournament,proto3" json:"tournament,omitempty"`
	Drunk                bool     `protobuf:"varint,2,opt,name=drunk,proto3" json:"drunk,omitempty"`
	Payed                bool     `protobuf:"varint,3,opt,name=payed,proto3" json:"payed,omitempty"`
	Bet                  bool     `protobuf:"varint,4,opt,name=bet,proto3" json:"bet,omitempty"`
	MaxGoals             int32    `protobuf:"varint,5,opt,name=maxGoals,proto3" json:"maxGoals,omitempty"`
	MaxTime              int32    `protobuf:"varint,6,opt,name=maxTime,proto3" json:"maxTime,omitempty"`
	Rated                bool     `protobuf:"varint,7,opt,name=rated,proto3" json:"rated,omitempty"`
	SwitchPositions      bool     `protobuf:"varint,8,opt,name=switchPositions,proto3" json:"switchPositions,omitempty"`
	Mode                 Mode     `protobuf:"varint,9,opt,name=mode,proto3,enum=proto.Mode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_livematch_e2e59194b3f622d4, []int{2}
}
func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (dst *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(dst, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetTournament() bool {
	if m != nil {
		return m.Tournament
	}
	return false
}

func (m *Settings) GetDrunk() bool {
	if m != nil {
		return m.Drunk
	}
	return false
}

func (m *Settings) GetPayed() bool {
	if m != nil {
		return m.Payed
	}
	return false
}

func (m *Settings) GetBet() bool {
	if m != nil {
		return m.Bet
	}
	return false
}

func (m *Settings) GetMaxGoals() int32 {
	if m != nil {
		return m.MaxGoals
	}
	return 0
}

func (m *Settings) GetMaxTime() int32 {
	if m != nil {
		return m.MaxTime
	}
	return 0
}

func (m *Settings) GetRated() bool {
	if m != nil {
		return m.Rated
	}
	return false
}

func (m *Settings) GetSwitchPositions() bool {
	if m != nil {
		return m.SwitchPositions
	}
	return false
}

func (m *Settings) GetMode() Mode {
	if m != nil {
		return m.Mode
	}
	return Mode_defaultMode
}

type Command struct {
	Settings             *Settings `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	Position             Position  `protobuf:"varint,2,opt,name=position,proto3,enum=proto.Position" json:"position,omitempty"`
	User                 *User     `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	AddGoalBlue          bool      `protobuf:"varint,4,opt,name=addGoalBlue,proto3" json:"addGoalBlue,omitempty"`
	AddGoalRed           bool      `protobuf:"varint,5,opt,name=addGoalRed,proto3" json:"addGoalRed,omitempty"`
	KickUser             string    `protobuf:"bytes,6,opt,name=kickUser,proto3" json:"kickUser,omitempty"`
	Join                 string    `protobuf:"bytes,7,opt,name=join,proto3" json:"join,omitempty"`
	Table                *Table    `protobuf:"bytes,8,opt,name=table,proto3" json:"table,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_livematch_e2e59194b3f622d4, []int{3}
}
func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (dst *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(dst, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetSettings() *Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *Command) GetPosition() Position {
	if m != nil {
		return m.Position
	}
	return Position_defaultPosition
}

func (m *Command) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Command) GetAddGoalBlue() bool {
	if m != nil {
		return m.AddGoalBlue
	}
	return false
}

func (m *Command) GetAddGoalRed() bool {
	if m != nil {
		return m.AddGoalRed
	}
	return false
}

func (m *Command) GetKickUser() string {
	if m != nil {
		return m.KickUser
	}
	return ""
}

func (m *Command) GetJoin() string {
	if m != nil {
		return m.Join
	}
	return ""
}

func (m *Command) GetTable() *Table {
	if m != nil {
		return m.Table
	}
	return nil
}

type Table struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Location             string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	LastMaintenance      string   `protobuf:"bytes,3,opt,name=lastMaintenance,proto3" json:"lastMaintenance,omitempty"`
	Ready                bool     `protobuf:"varint,4,opt,name=ready,proto3" json:"ready,omitempty"`
	InGame               bool     `protobuf:"varint,5,opt,name=inGame,proto3" json:"inGame,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}
func (*Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_livematch_e2e59194b3f622d4, []int{4}
}
func (m *Table) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Table.Unmarshal(m, b)
}
func (m *Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Table.Marshal(b, m, deterministic)
}
func (dst *Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Table.Merge(dst, src)
}
func (m *Table) XXX_Size() int {
	return xxx_messageInfo_Table.Size(m)
}
func (m *Table) XXX_DiscardUnknown() {
	xxx_messageInfo_Table.DiscardUnknown(m)
}

var xxx_messageInfo_Table proto.InternalMessageInfo

func (m *Table) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Table) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Table) GetLastMaintenance() string {
	if m != nil {
		return m.LastMaintenance
	}
	return ""
}

func (m *Table) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

func (m *Table) GetInGame() bool {
	if m != nil {
		return m.InGame
	}
	return false
}

func init() {
	proto.RegisterType((*Match)(nil), "proto.Match")
	proto.RegisterMapType((map[string]string)(nil), "proto.Match.LabelsEntry")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*Settings)(nil), "proto.Settings")
	proto.RegisterType((*Command)(nil), "proto.Command")
	proto.RegisterType((*Table)(nil), "proto.Table")
	proto.RegisterEnum("proto.Mode", Mode_name, Mode_value)
	proto.RegisterEnum("proto.Position", Position_name, Position_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LivematchClient is the client API for Livematch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LivematchClient interface {
	Send(ctx context.Context, opts ...grpc.CallOption) (Livematch_SendClient, error)
}

type livematchClient struct {
	cc *grpc.ClientConn
}

func NewLivematchClient(cc *grpc.ClientConn) LivematchClient {
	return &livematchClient{cc}
}

func (c *livematchClient) Send(ctx context.Context, opts ...grpc.CallOption) (Livematch_SendClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Livematch_serviceDesc.Streams[0], "/proto.Livematch/Send", opts...)
	if err != nil {
		return nil, err
	}
	x := &livematchSendClient{stream}
	return x, nil
}

type Livematch_SendClient interface {
	Send(*Command) error
	Recv() (*Match, error)
	grpc.ClientStream
}

type livematchSendClient struct {
	grpc.ClientStream
}

func (x *livematchSendClient) Send(m *Command) error {
	return x.ClientStream.SendMsg(m)
}

func (x *livematchSendClient) Recv() (*Match, error) {
	m := new(Match)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LivematchServer is the server API for Livematch service.
type LivematchServer interface {
	Send(Livematch_SendServer) error
}

func RegisterLivematchServer(s *grpc.Server, srv LivematchServer) {
	s.RegisterService(&_Livematch_serviceDesc, srv)
}

func _Livematch_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LivematchServer).Send(&livematchSendServer{stream})
}

type Livematch_SendServer interface {
	Send(*Match) error
	Recv() (*Command, error)
	grpc.ServerStream
}

type livematchSendServer struct {
	grpc.ServerStream
}

func (x *livematchSendServer) Send(m *Match) error {
	return x.ServerStream.SendMsg(m)
}

func (x *livematchSendServer) Recv() (*Command, error) {
	m := new(Command)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Livematch_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Livematch",
	HandlerType: (*LivematchServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Send",
			Handler:       _Livematch_Send_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "livematch.proto",
}

func init() { proto.RegisterFile("livematch.proto", fileDescriptor_livematch_e2e59194b3f622d4) }

var fileDescriptor_livematch_e2e59194b3f622d4 = []byte{
	// 817 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xae, 0x63, 0x3b, 0xeb, 0x1c, 0x87, 0x24, 0x0c, 0x05, 0x99, 0x15, 0x82, 0x25, 0x57, 0xd1,
	0x22, 0xad, 0xaa, 0xe5, 0x82, 0x9f, 0x3b, 0xa0, 0xa8, 0xaa, 0xe8, 0xaa, 0xc8, 0x5d, 0xae, 0xab,
	0x89, 0x3d, 0xa5, 0x26, 0xf6, 0x4c, 0x64, 0x4f, 0x1a, 0xf2, 0x0c, 0xbc, 0x01, 0xcf, 0xc3, 0x3b,
	0xf0, 0x38, 0x70, 0xce, 0x99, 0xb1, 0x93, 0x2e, 0x54, 0x7b, 0x15, 0x7f, 0xdf, 0x99, 0x9f, 0x73,
	0xce, 0xf7, 0x9d, 0x09, 0xcc, 0xeb, 0xea, 0x8d, 0x6a, 0xa4, 0x2d, 0x5e, 0x5f, 0x6d, 0x5b, 0x63,
	0x8d, 0x88, 0xf9, 0x67, 0xf9, 0x57, 0x08, 0xf1, 0x0d, 0xd1, 0x62, 0x0e, 0xe1, 0xcb, 0xaa, 0xcc,
	0x82, 0x8b, 0x60, 0x35, 0xc9, 0x47, 0x4f, 0x4b, 0xf1, 0x3e, 0x44, 0x2f, 0x37, 0xea, 0x90, 0x8d,
	0x98, 0x09, 0x7f, 0x52, 0x07, 0x91, 0xc1, 0x99, 0x95, 0xeb, 0x5a, 0x3d, 0x7d, 0x9c, 0x85, 0xcc,
	0xf6, 0x90, 0x22, 0x9d, 0x95, 0xad, 0x55, 0x65, 0x16, 0x61, 0x24, 0xc9, 0x7b, 0x28, 0x3e, 0x87,
	0x78, 0xd7, 0xa9, 0xb6, 0xcb, 0xe2, 0x8b, 0x70, 0x95, 0x5e, 0xa7, 0xee, 0xfe, 0xab, 0x5f, 0x90,
	0xcb, 0x5d, 0x44, 0x7c, 0x01, 0x49, 0xa7, 0xac, 0xad, 0xf4, 0xaf, 0x5d, 0x36, 0xc6, 0xdd, 0xe9,
	0xf5, 0xdc, 0xaf, 0x7a, 0xe1, 0xe9, 0x7c, 0x58, 0x20, 0x3e, 0x81, 0x09, 0x1f, 0x7d, 0x5b, 0x35,
	0x2a, 0x3b, 0xe3, 0x2c, 0x8e, 0x04, 0x47, 0x0b, 0xd3, 0xaa, 0xef, 0xeb, 0x9d, 0xca, 0x12, 0x8c,
	0xc6, 0xf9, 0x91, 0x10, 0xe7, 0x78, 0x11, 0x81, 0x1c, 0xd3, 0x9c, 0x70, 0x70, 0xc0, 0x54, 0x81,
	0xd2, 0x25, 0x9f, 0x0a, 0xae, 0x36, 0x0f, 0xe9, 0xcc, 0xc2, 0x34, 0xdb, 0x5a, 0x51, 0x75, 0x29,
	0x57, 0x77, 0x24, 0xc4, 0x47, 0x30, 0xde, 0x57, 0x5a, 0xab, 0x36, 0x9b, 0xf2, 0x36, 0x8f, 0xc4,
	0x23, 0x18, 0xd7, 0x72, 0xad, 0xea, 0x2e, 0x9b, 0x71, 0xe1, 0x99, 0x2f, 0x89, 0xbb, 0x7d, 0xf5,
	0x8c, 0x43, 0x3f, 0x6a, 0xdb, 0x1e, 0x72, 0xbf, 0xee, 0xfc, 0x1b, 0x48, 0x4f, 0x68, 0xb1, 0x80,
	0x90, 0xda, 0xef, 0x04, 0xa1, 0x4f, 0xf1, 0x10, 0xe2, 0x37, 0x92, 0x0a, 0x73, 0x92, 0x38, 0xf0,
	0xed, 0xe8, 0xeb, 0x60, 0xf9, 0x77, 0x00, 0x11, 0x75, 0x54, 0xcc, 0x60, 0x74, 0x14, 0xb1, 0x2a,
	0xa9, 0x62, 0xea, 0xb1, 0x96, 0x4d, 0xbf, 0x6b, 0xc0, 0x74, 0x9c, 0x2c, 0x9b, 0x4a, 0xb3, 0x96,
	0x49, 0xee, 0x00, 0xb1, 0xad, 0x92, 0xe5, 0xc1, 0xeb, 0xe8, 0x00, 0x49, 0xb4, 0x35, 0x5d, 0x65,
	0x2b, 0xa3, 0x51, 0xc8, 0x60, 0x35, 0x1b, 0x24, 0xfa, 0xd9, 0xd3, 0xf9, 0xb0, 0x80, 0x32, 0x5f,
	0x2b, 0xcb, 0x52, 0x86, 0x39, 0x7d, 0xd2, 0xa1, 0x85, 0xa9, 0x4d, 0xeb, 0x05, 0x73, 0x40, 0x2c,
	0x61, 0x5a, 0xec, 0xda, 0x56, 0x69, 0x7b, 0x4b, 0x36, 0x62, 0xbd, 0x26, 0xf9, 0x5b, 0xdc, 0xf2,
	0x9f, 0x00, 0x92, 0xde, 0x05, 0xe2, 0x53, 0x00, 0x6b, 0x76, 0x9c, 0xbd, 0xb6, 0x5c, 0x65, 0x92,
	0x9f, 0x30, 0x74, 0x4d, 0xd9, 0xee, 0xf4, 0x86, 0x4b, 0xc5, 0xdc, 0x19, 0x10, 0xbb, 0x95, 0x07,
	0xd4, 0xce, 0xd7, 0xc9, 0xa0, 0x4f, 0xd2, 0x55, 0xc9, 0x49, 0x62, 0xaf, 0x1a, 0xf9, 0xfb, 0x13,
	0x23, 0xeb, 0x8e, 0x6b, 0x44, 0x77, 0xf4, 0x98, 0xdc, 0x81, 0xdf, 0xec, 0x8e, 0x31, 0x87, 0x7a,
	0xc8, 0xfd, 0x92, 0xe4, 0x8c, 0x33, 0xdf, 0x2f, 0x02, 0x62, 0x05, 0xf3, 0x6e, 0x5f, 0xa1, 0xd2,
	0x7d, 0x7b, 0x3a, 0xae, 0x2e, 0xc9, 0xef, 0xd2, 0xe2, 0x33, 0x88, 0x1a, 0x53, 0x2a, 0xf6, 0xe3,
	0x6c, 0x18, 0x8f, 0x1b, 0xa4, 0x72, 0x0e, 0x2c, 0xff, 0x1c, 0xc1, 0xd9, 0x0f, 0xa6, 0x69, 0xa4,
	0x2e, 0xdf, 0x9a, 0x94, 0xe0, 0xbe, 0x49, 0x39, 0xd5, 0x6c, 0x74, 0x9f, 0x66, 0x98, 0x06, 0x19,
	0x83, 0x7b, 0x74, 0x67, 0x4a, 0x39, 0x20, 0x2e, 0x20, 0x95, 0x65, 0x49, 0xdd, 0xe0, 0xd9, 0x72,
	0x7d, 0x3b, 0xa5, 0x48, 0x1d, 0x0f, 0x69, 0xbe, 0x62, 0xa7, 0xce, 0x91, 0xa1, 0xfe, 0x6e, 0xaa,
	0x62, 0x43, 0x67, 0x72, 0x13, 0xd1, 0x8b, 0x3d, 0x16, 0x02, 0xa2, 0xdf, 0x0c, 0x5a, 0xd1, 0xf9,
	0x83, 0xbf, 0xd1, 0x1e, 0xb1, 0x1d, 0x7c, 0x91, 0x5e, 0x4f, 0x7d, 0x4e, 0xec, 0x8b, 0xdc, 0x85,
	0x96, 0x7f, 0x04, 0x10, 0x33, 0xf1, 0x7f, 0xce, 0xaf, 0x4d, 0x21, 0x87, 0xea, 0xf1, 0xb6, 0x1e,
	0x93, 0x3a, 0xb5, 0xec, 0xec, 0x8d, 0xac, 0xb4, 0x55, 0x5a, 0xea, 0x42, 0xf9, 0xf7, 0xec, 0x2e,
	0xfd, 0x8e, 0x69, 0xc0, 0x99, 0xaf, 0xf4, 0x13, 0x9a, 0x29, 0x57, 0xa5, 0x47, 0x97, 0x7b, 0x88,
	0x48, 0x38, 0x7c, 0x4b, 0xd3, 0x52, 0xbd, 0x92, 0xbb, 0xda, 0x12, 0x5c, 0x3c, 0x10, 0x53, 0x48,
	0x8c, 0x56, 0xcf, 0xf5, 0x73, 0xad, 0x16, 0x01, 0x21, 0xbb, 0x37, 0x0e, 0x8d, 0xc4, 0x7b, 0x30,
	0x61, 0xa4, 0x6e, 0xf7, 0x66, 0x11, 0xa2, 0xd3, 0x1e, 0x1e, 0x1d, 0x4d, 0xdb, 0x91, 0xc6, 0xee,
	0x2d, 0x22, 0xf1, 0x31, 0x7c, 0xf8, 0x9f, 0x08, 0x35, 0x7e, 0x11, 0x5f, 0x6a, 0x48, 0x7a, 0x4d,
	0xc5, 0x07, 0x30, 0xf7, 0x97, 0xf7, 0x14, 0x26, 0x30, 0x03, 0x58, 0xe3, 0xd2, 0xef, 0xac, 0x95,
	0xc5, 0x06, 0x53, 0xc0, 0x0c, 0x09, 0x3f, 0x56, 0xaf, 0x94, 0xee, 0x7c, 0x16, 0xad, 0x2a, 0x39,
	0xfe, 0x1a, 0xb3, 0xc0, 0xf5, 0x08, 0xfb, 0x70, 0x44, 0xe1, 0x6e, 0xab, 0x0a, 0x2b, 0xad, 0x69,
	0x17, 0xf1, 0xf5, 0x57, 0x30, 0x79, 0xd6, 0xff, 0xa1, 0x88, 0x4b, 0x88, 0x5e, 0xe0, 0x5b, 0x29,
	0x66, 0x5e, 0x20, 0x6f, 0xd6, 0xf3, 0xe9, 0xe9, 0x8b, 0xb7, 0x7c, 0xb0, 0x0a, 0x1e, 0x05, 0xeb,
	0x31, 0x53, 0x5f, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x17, 0xbf, 0xe5, 0x12, 0x90, 0x06, 0x00,
	0x00,
}
