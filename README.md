# binarysearch
Generic binary search module for golang.

It can search any index space, it doesn't have to be a slice, and it isn't limited to a set of pre-defined types.
The user provides a function that indicates whether a particular index is too low or too high, and the module finds the right index.
