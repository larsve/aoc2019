package intcodecpu

type (
	// StdIn read one integer from standard input
	StdIn = func() int
	// StdOut write one integer to standard output
	StdOut = func(data int)
)
