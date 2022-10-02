# Generate Data V1

# Overview

We want to generate data for a neural network to learn how to navigate the maze.

## Input

The Maze (Hero, Exit, wall positions)

## Label

The direction chosen to move the hero can be one of [up, down, left, right]

The output of the Neural network should be which direction to tell the hero to move.

# V1 data schema

V1 is the brain-dead method. The data is just a single large array.

The Input length will be size of maze e.g. 50x50 + Position of Exit (x,y) + Position of Hero (x,y)

= 2500 + 2 + 2

The output length will be size of action

= 4

## Maze encoding

0 - empty space

1 - wall

The position of Exit and Hero will not be shown

# V1 output schema
action: array [ [1|0], [1|0], [1|0], [1|0] ]
    
    Each element of the array corresponds to one direction.
    
    0 index - up
    
    1 index - down
    
    2 index - left
    
    3 index - right

# Output normalization

Whichever one of the outputs has the greater value will be chosen as the result.

For example: If the output is `[ 0.1, 0.5, 0.3, 0.7 ]` Then `right` will be chosen.

# Implementation

Using a single maze, we will move the hero to various points and calculate the best direction for the Hero to take. We will decide the direction based on which move makes the Hero closer to the exit and is valid.

We will repeat the above for different randomly generated mazes, so that we don't overfit for one maze.

For each trial we will store the Input vector as a csv row in the file `input.csv` and the corresponding output in `labels.csv`
