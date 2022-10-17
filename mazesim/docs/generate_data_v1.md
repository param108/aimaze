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

For each trial we will store the Input vector as a csv row in the file `inputs.csv` and the corresponding output in `labels.csv`

# Results

The model failed to train approproately.
After 20 epochs
accuracy
	training         	 (min:    0.461, max:    0.508, cur:    0.508)
	validation       	 (min:    0.448, max:    0.474, cur:    0.458)
Loss
	training         	 (min:    0.407, max:    0.464, cur:    0.407)
	validation       	 (min:    0.414, max:    0.421, cur:    0.417)
Mean Squared Error
	training         	 (min:    0.140, max:    0.148, cur:    0.140)
	validation       	 (min:    0.142, max:    0.145, cur:    0.143)
617/617 [==============================] - 22s 36ms/step - loss: 0.4072 - accuracy: 0.5081 - mse: 0.1397 - val_loss: 0.4174 - val_accuracy: 0.4576 - val_mse: 0.1432
Test loss: 0.41744720935821533
Test accuracy: 0.4576219916343689
Model: "sequential"

