This is ASCII-ART-WEB-DOCKERIZE project 

# Ascii-art-stylize consists on making your site :

For this project you must create at least :

* one Dockerfile
* one image
* one container
It must apply metadata to Docker objects.

It have to take caution of unused object (often referred to as "garbage collection").

# Authors 

atemerzh and asundeto

# Usage: how to run

To run the Web server with DockerFile write this to console:

1) make run

# What is run?

run command consists docker build and run commands

run:
	docker build -t ascii .
	docker run -dp 8081:8081 ascii:latest
	$(info running on http://localhost:8081)

* To stop docker do next steps:

1) write to console - docker ps

2) copy CONTAINER ID

3) write - docker kill || docker stop [CONTAINER ID]

# Description

This project must be written in Go.

The code must respect the good practices.

It is recommended to have test files for unit testing.

Some banner files with a specific graphical template representation using ASCII will be given. The files are formatted in a way that is not necessary to change them.

* shadow
* standard
* thinkertoy

# Implementation details: algorithm of ascii-art

We save the entire txt file into a variable, and then by algorithms save the runes of each desired character into an array. This array is concatenated into a single line string via the connect function
