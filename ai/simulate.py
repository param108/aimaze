from simulation_pb2_grpc import SimulatorStub
from simulation_pb2 import CreateSimulationRequest

import grpc
import os

def printSim(sim):
    os.system('clear')
    x = 0
    y = 0
    idx = 0
    for y in range(sim.maze.size.height):
        for x in range(sim.maze.size.width):
            if x == sim.hero.x and y == sim.hero.y:
                print('H', end='')
            elif x == sim.maze.exit.x and y == sim.maze.exit.y:
                print('E', end='')
            else:
                print(sim.maze.maze[idx],end='')
            idx+=1
        print('')
        x = 0

channel = grpc.insecure_channel('localhost:9999')

stub = SimulatorStub(channel)

sim = stub.CreateSimulation(CreateSimulationRequest())

printSim(sim)
