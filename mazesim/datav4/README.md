# V4

This time we generate the complete path to the Exit when we see the maze.

The output will be an array of 1s and 0s. Each group of 4 values will correspond
to one direction to take.
The 0th index direction (0,1,2,3) will be the first step to be taken, the next index (4,5,6,7) will 
be the next and so on.
There will always be 150 such tuples in the output, many may be all 0s, which implies a noop.
