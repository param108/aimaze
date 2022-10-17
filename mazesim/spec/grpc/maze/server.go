package maze

import (
	"fmt"

	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
)

type ServerIntf interface {
	CreateSimulation(
		ctx context.Context, req *CreateSimulationRequest,
	) (*Simulation, error)
	Simulate(
		ctx context.Context, req *SimulationAction,
	) (*Simulation, error)
	GetFeaturesV2(
		ctx context.Context, req *Simulation,
	) (*FeaturesV2, error)
}

type Server struct {
	intf ServerIntf
}

func NewServer(intf ServerIntf) *Server {
	return &Server {
		intf: intf,
	}
}
// CreateSimulation - returns a new Simulation State
func (s *Server) CreateSimulation(
	ctx context.Context,
	req *CreateSimulationRequest,
) (*Simulation, error) {
	return s.intf.CreateSimulation(ctx, req)
	//return NewSim()
}

// Simulate - Request an action on a Simulation
//            Returns the new Simulation State
func (s *Server) Simulate(
	ctx context.Context,
	req *SimulationAction,
) (*Simulation, error) {
	return s.intf.Simulate(ctx, req)
	/*x, y, valid := req.Sim.DryMove(req.Action)
	if valid {
		req.Sim.Hero.X = x
		req.Sim.Hero.Y = y
	}
	return req.Sim, nil*/
}

// GetFeaturesV2 - Given a simulation return v2 features
func (s *Server) GetFeaturesV2(
	ctx context.Context,
	req *Simulation,
) (*FeaturesV2, error) {
	return s.intf.GetFeaturesV2(ctx, req)
	/*ret := &FeaturesV2{}
	ret.Features = getInputV2(req)
	return ret, nil*/
}

func (s *Server) mustEmbedUnimplementedSimulatorServer() {

}

// startV2Server - starts the grpc server and never returns
func StartServer(port int, intf ServerIntf) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	server := NewServer(intf)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterSimulatorServer(grpcServer, server)
	grpcServer.Serve(lis)
	return nil
}
