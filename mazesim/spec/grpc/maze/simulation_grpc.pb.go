// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: simulation.proto

package maze

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SimulatorClient is the client API for Simulator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimulatorClient interface {
	// CreateSimulation - returns a new Simulation State
	CreateSimulation(ctx context.Context, in *CreateSimulationRequest, opts ...grpc.CallOption) (*Simulation, error)
	// Simulate - Request an action on a Simulation
	//
	//	Returns the new Simulation State
	Simulate(ctx context.Context, in *SimulationAction, opts ...grpc.CallOption) (*Simulation, error)
	// GetFeaturesV2 - Given a simulation return v2 features
	GetFeaturesV2(ctx context.Context, in *Simulation, opts ...grpc.CallOption) (*FeaturesV2, error)
}

type simulatorClient struct {
	cc grpc.ClientConnInterface
}

func NewSimulatorClient(cc grpc.ClientConnInterface) SimulatorClient {
	return &simulatorClient{cc}
}

func (c *simulatorClient) CreateSimulation(ctx context.Context, in *CreateSimulationRequest, opts ...grpc.CallOption) (*Simulation, error) {
	out := new(Simulation)
	err := c.cc.Invoke(ctx, "/Simulator/CreateSimulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulatorClient) Simulate(ctx context.Context, in *SimulationAction, opts ...grpc.CallOption) (*Simulation, error) {
	out := new(Simulation)
	err := c.cc.Invoke(ctx, "/Simulator/Simulate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simulatorClient) GetFeaturesV2(ctx context.Context, in *Simulation, opts ...grpc.CallOption) (*FeaturesV2, error) {
	out := new(FeaturesV2)
	err := c.cc.Invoke(ctx, "/Simulator/GetFeaturesV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimulatorServer is the server API for Simulator service.
// All implementations must embed UnimplementedSimulatorServer
// for forward compatibility
type SimulatorServer interface {
	// CreateSimulation - returns a new Simulation State
	CreateSimulation(context.Context, *CreateSimulationRequest) (*Simulation, error)
	// Simulate - Request an action on a Simulation
	//
	//	Returns the new Simulation State
	Simulate(context.Context, *SimulationAction) (*Simulation, error)
	// GetFeaturesV2 - Given a simulation return v2 features
	GetFeaturesV2(context.Context, *Simulation) (*FeaturesV2, error)
	mustEmbedUnimplementedSimulatorServer()
}

// UnimplementedSimulatorServer must be embedded to have forward compatible implementations.
type UnimplementedSimulatorServer struct {
}

func (UnimplementedSimulatorServer) CreateSimulation(context.Context, *CreateSimulationRequest) (*Simulation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSimulation not implemented")
}
func (UnimplementedSimulatorServer) Simulate(context.Context, *SimulationAction) (*Simulation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Simulate not implemented")
}
func (UnimplementedSimulatorServer) GetFeaturesV2(context.Context, *Simulation) (*FeaturesV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeaturesV2 not implemented")
}
func (UnimplementedSimulatorServer) mustEmbedUnimplementedSimulatorServer() {}

// UnsafeSimulatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimulatorServer will
// result in compilation errors.
type UnsafeSimulatorServer interface {
	mustEmbedUnimplementedSimulatorServer()
}

func RegisterSimulatorServer(s grpc.ServiceRegistrar, srv SimulatorServer) {
	s.RegisterService(&Simulator_ServiceDesc, srv)
}

func _Simulator_CreateSimulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSimulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulatorServer).CreateSimulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Simulator/CreateSimulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulatorServer).CreateSimulation(ctx, req.(*CreateSimulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Simulator_Simulate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimulationAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulatorServer).Simulate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Simulator/Simulate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulatorServer).Simulate(ctx, req.(*SimulationAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Simulator_GetFeaturesV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Simulation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimulatorServer).GetFeaturesV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Simulator/GetFeaturesV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimulatorServer).GetFeaturesV2(ctx, req.(*Simulation))
	}
	return interceptor(ctx, in, info, handler)
}

// Simulator_ServiceDesc is the grpc.ServiceDesc for Simulator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Simulator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Simulator",
	HandlerType: (*SimulatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSimulation",
			Handler:    _Simulator_CreateSimulation_Handler,
		},
		{
			MethodName: "Simulate",
			Handler:    _Simulator_Simulate_Handler,
		},
		{
			MethodName: "GetFeaturesV2",
			Handler:    _Simulator_GetFeaturesV2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simulation.proto",
}
