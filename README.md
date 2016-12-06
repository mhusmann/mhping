mhping
======

Description
___________

Parallel ping my home hosts in my network.

The package is installed by go install 

Documentation
_____________

I wrote this kind of tool in several languages: Python, Ruby, elixer and also
in Google Go, because I was looking for a tool that checks my own network.

When running such a tool sequentially then I must wait for all time outs and
that takes time. So the idea was: Let them run in parallel. Doing so I only have
to wait for one time out and not all the others.

Since the fping tool is available I found no need to invent it by myself.

Added flag retry to adjust fping with more retries if necessary

Slice
