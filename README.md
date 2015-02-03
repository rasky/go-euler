# go-euler
Project Euler solutions in Go

This repository contains solutions to Project Euler problems written in the Go programming language.

The package "natural" contains common utilities about natural numbers that are useful in many of those problems,
while the rest of the repository is structured as a big testsuite. You can use:

  $ go test

to run all the problems; obviously, the tests will fail if the solution doesn't match the correct one. If you
want to have more details, you can activate verbose output:

  $ go test -test.v=1

This will also show the execution time of each problem; notice that it also activates some logging from each
test (I usually log intermediate solutions), so it can be come a little verbose.

## Performance goals

Project Euler launched 10 years ago, and the stated performance goal for the soultions was "less than one
minute of runtime". My own goal is to take less than 10 seconds for most problems (though I don't know if it's
possible as I've yet to solve most of them), with a second goal of less than 1 second when algorithmically
feasable. 

As a compromise towards this goal, in a few problems, I have reimplemented a specialized versions of
the generic algorithms provided in the "natural" package; I make this only when it makes a noticable performance
difference, and I think it's a good school for how and when abstractions leak and impact performance. On
the other hand, some problems are algorithmically heavy (eg: `O(n^5)`) so they tend to stress the
implementation details more than in normal everyday code.
