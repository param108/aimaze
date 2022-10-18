In this we switch back to sending the complete maze as input along with the coordinates of Hero and Exit.
THat is 
2500 1s or 0s (1 for wall, 0 otherwise)
Hero X (Normalized by dividing by width of maze)
Hero Y (Normalized by dividing by height of maze)
Exit X (Normalized by dividing by width of maze)
Exit Y (Normalized by dividing by height of maze)

In this version we actually solve the maze using recursion and
then use the solution path to generate the data.

Output is still the same
[Up, Down, Left, Right]
Where only one of these is 1 signifying the correct direction
to move.
