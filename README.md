# gobinsearch
Generic binary search module for golang.

It can search any index space, it doesn't have to be a slice.
The user provides a function that indicates whether a particular index is too low or too high, and the module finds the right index.
