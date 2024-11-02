<p align="center"><img style="width: 300px" src="./img/SimpleMR.png"></img></p>
<h1 align="center">SimpleMR<br>A Simple Implementation of Google MapReduce</h1>

This application is a from scratch project to help me (KevinZonda) familiar with
Google's MapReduce.

## Basic Concept

In original paper, the Worker refers to a singe CPU core.
In this implementation, a Worker refers to a go routine.

RPC over different workers is not done as the same as paper one.
In this implementation, WaitGroup is used to wait for all workers
instead of working with RPC to do communication.
