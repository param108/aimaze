from simulation_pb2_grpc import SimulatorStub
from simulation_pb2 import CreateSimulationRequest, SimulationAction
from models import simulate_ai,get_action
import grpc
import os
import numpy
import time

# Maximum tries allowed per maze
max_tries = 150

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

try_num = 1

while (not (sim.maze.exit.y == sim.hero.y and sim.maze.exit.x == sim.hero.x)) and try_num < max_tries :
    features = stub.GetFeaturesV2(sim)

    actionArr = simulate_ai(features.features)[0]
    next_action = get_action(sim, actionArr)
    print(actionArr, next_action.action)

    printSim(sim)
    print("moved:", next_action.action,"try:",try_num)

    sim = stub.Simulate(next_action)
    try_num+=1

if try_num == 150:
    print ("failed, ran out of tries:", try_num)
else:
    print("success made it in:", try_num, "/", max_tries)
