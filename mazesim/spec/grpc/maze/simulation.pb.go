// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: simulation.proto

package maze

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Simulation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maze *Maze  `protobuf:"bytes,1,opt,name=maze,proto3" json:"maze,omitempty"`
	Hero *Point `protobuf:"bytes,2,opt,name=hero,proto3" json:"hero,omitempty"`
	// prev - the previous point where the hero stood
	// initially it is set to the starting point of the hero
	Prev *Point `protobuf:"bytes,3,opt,name=prev,proto3" json:"prev,omitempty"`
	// how many actions run so far
	Step int32 `protobuf:"varint,4,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *Simulation) Reset() {
	*x = Simulation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simulation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Simulation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Simulation) ProtoMessage() {}

func (x *Simulation) ProtoReflect() protoreflect.Message {
	mi := &file_simulation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Simulation.ProtoReflect.Descriptor instead.
func (*Simulation) Descriptor() ([]byte, []int) {
	return file_simulation_proto_rawDescGZIP(), []int{0}
}

func (x *Simulation) GetMaze() *Maze {
	if x != nil {
		return x.Maze
	}
	return nil
}

func (x *Simulation) GetHero() *Point {
	if x != nil {
		return x.Hero
	}
	return nil
}

func (x *Simulation) GetPrev() *Point {
	if x != nil {
		return x.Prev
	}
	return nil
}

func (x *Simulation) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

type SimulationAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sim    *Simulation `protobuf:"bytes,1,opt,name=sim,proto3" json:"sim,omitempty"`
	Action string      `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *SimulationAction) Reset() {
	*x = SimulationAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simulation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulationAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulationAction) ProtoMessage() {}

func (x *SimulationAction) ProtoReflect() protoreflect.Message {
	mi := &file_simulation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulationAction.ProtoReflect.Descriptor instead.
func (*SimulationAction) Descriptor() ([]byte, []int) {
	return file_simulation_proto_rawDescGZIP(), []int{1}
}

func (x *SimulationAction) GetSim() *Simulation {
	if x != nil {
		return x.Sim
	}
	return nil
}

func (x *SimulationAction) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type Maze struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size         *Size  `protobuf:"bytes,1,opt,name=size,proto3" json:"size,omitempty"`
	Maze         string `protobuf:"bytes,2,opt,name=maze,proto3" json:"maze,omitempty"`
	Exit         *Point `protobuf:"bytes,3,opt,name=exit,proto3" json:"exit,omitempty"`
	DoorsPerWall int32  `protobuf:"varint,4,opt,name=doors_per_wall,json=doorsPerWall,proto3" json:"doors_per_wall,omitempty"`
}

func (x *Maze) Reset() {
	*x = Maze{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simulation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Maze) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Maze) ProtoMessage() {}

func (x *Maze) ProtoReflect() protoreflect.Message {
	mi := &file_simulation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Maze.ProtoReflect.Descriptor instead.
func (*Maze) Descriptor() ([]byte, []int) {
	return file_simulation_proto_rawDescGZIP(), []int{2}
}

func (x *Maze) GetSize() *Size {
	if x != nil {
		return x.Size
	}
	return nil
}

func (x *Maze) GetMaze() string {
	if x != nil {
		return x.Maze
	}
	return ""
}

func (x *Maze) GetExit() *Point {
	if x != nil {
		return x.Exit
	}
	return nil
}

func (x *Maze) GetDoorsPerWall() int32 {
	if x != nil {
		return x.DoorsPerWall
	}
	return 0
}

type Size struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  int32 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height int32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Size) Reset() {
	*x = Size{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simulation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Size) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Size) ProtoMessage() {}

func (x *Size) ProtoReflect() protoreflect.Message {
	mi := &file_simulation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Size.ProtoReflect.Descriptor instead.
func (*Size) Descriptor() ([]byte, []int) {
	return file_simulation_proto_rawDescGZIP(), []int{3}
}

func (x *Size) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Size) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simulation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_simulation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_simulation_proto_rawDescGZIP(), []int{4}
}

func (x *Point) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Point) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type CreateSimulationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSimulationRequest) Reset() {
	*x = CreateSimulationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simulation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSimulationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSimulationRequest) ProtoMessage() {}

func (x *CreateSimulationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simulation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSimulationRequest.ProtoReflect.Descriptor instead.
func (*CreateSimulationRequest) Descriptor() ([]byte, []int) {
	return file_simulation_proto_rawDescGZIP(), []int{5}
}

type FeaturesV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Features []float64 `protobuf:"fixed64,1,rep,packed,name=features,proto3" json:"features,omitempty"`
}

func (x *FeaturesV2) Reset() {
	*x = FeaturesV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simulation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeaturesV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeaturesV2) ProtoMessage() {}

func (x *FeaturesV2) ProtoReflect() protoreflect.Message {
	mi := &file_simulation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeaturesV2.ProtoReflect.Descriptor instead.
func (*FeaturesV2) Descriptor() ([]byte, []int) {
	return file_simulation_proto_rawDescGZIP(), []int{6}
}

func (x *FeaturesV2) GetFeatures() []float64 {
	if x != nil {
		return x.Features
	}
	return nil
}

var File_simulation_proto protoreflect.FileDescriptor

var file_simulation_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x73, 0x0a, 0x0a, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x19, 0x0a, 0x04, 0x6d, 0x61, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x4d, 0x61, 0x7a, 0x65, 0x52, 0x04, 0x6d, 0x61, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x04, 0x68,
	0x65, 0x72, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x04, 0x68, 0x65, 0x72, 0x6f, 0x12, 0x1a, 0x0a, 0x04, 0x70, 0x72, 0x65, 0x76, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x04, 0x70,
	0x72, 0x65, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x22, 0x49, 0x0a, 0x10, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x03, 0x73,
	0x69, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x73, 0x69, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x77, 0x0a, 0x04, 0x4d, 0x61, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x04, 0x65, 0x78, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x04, 0x65, 0x78, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x6f, 0x6f, 0x72, 0x73, 0x5f, 0x70,
	0x65, 0x72, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64,
	0x6f, 0x6f, 0x72, 0x73, 0x50, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x22, 0x34, 0x0a, 0x04, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x22, 0x23, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x28, 0x0a, 0x0a, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x56, 0x32, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x01, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x32, 0x9d, 0x01, 0x0a, 0x09,
	0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x39, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x08, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0x11, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x0b, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x29, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x56,
	0x32, 0x12, 0x0b, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0b,
	0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x56, 0x32, 0x42, 0x10, 0x5a, 0x0e, 0x73,
	0x70, 0x65, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x7a, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_simulation_proto_rawDescOnce sync.Once
	file_simulation_proto_rawDescData = file_simulation_proto_rawDesc
)

func file_simulation_proto_rawDescGZIP() []byte {
	file_simulation_proto_rawDescOnce.Do(func() {
		file_simulation_proto_rawDescData = protoimpl.X.CompressGZIP(file_simulation_proto_rawDescData)
	})
	return file_simulation_proto_rawDescData
}

var file_simulation_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_simulation_proto_goTypes = []interface{}{
	(*Simulation)(nil),              // 0: Simulation
	(*SimulationAction)(nil),        // 1: SimulationAction
	(*Maze)(nil),                    // 2: Maze
	(*Size)(nil),                    // 3: Size
	(*Point)(nil),                   // 4: Point
	(*CreateSimulationRequest)(nil), // 5: CreateSimulationRequest
	(*FeaturesV2)(nil),              // 6: FeaturesV2
}
var file_simulation_proto_depIdxs = []int32{
	2, // 0: Simulation.maze:type_name -> Maze
	4, // 1: Simulation.hero:type_name -> Point
	4, // 2: Simulation.prev:type_name -> Point
	0, // 3: SimulationAction.sim:type_name -> Simulation
	3, // 4: Maze.size:type_name -> Size
	4, // 5: Maze.exit:type_name -> Point
	5, // 6: Simulator.CreateSimulation:input_type -> CreateSimulationRequest
	1, // 7: Simulator.Simulate:input_type -> SimulationAction
	0, // 8: Simulator.GetFeaturesV2:input_type -> Simulation
	0, // 9: Simulator.CreateSimulation:output_type -> Simulation
	0, // 10: Simulator.Simulate:output_type -> Simulation
	6, // 11: Simulator.GetFeaturesV2:output_type -> FeaturesV2
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_simulation_proto_init() }
func file_simulation_proto_init() {
	if File_simulation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_simulation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Simulation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_simulation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulationAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_simulation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Maze); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_simulation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Size); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_simulation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_simulation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSimulationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_simulation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeaturesV2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_simulation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_simulation_proto_goTypes,
		DependencyIndexes: file_simulation_proto_depIdxs,
		MessageInfos:      file_simulation_proto_msgTypes,
	}.Build()
	File_simulation_proto = out.File
	file_simulation_proto_rawDesc = nil
	file_simulation_proto_goTypes = nil
	file_simulation_proto_depIdxs = nil
}
