# User-Based Filtering

This directory holds the go code for the first part of Chapter 2 which focuses on user-based filtering techniques.

As discussed in the book, user-based filtering is rather inefficient because it relies on rerunning the algorithm everytime a user makes a new action.

## Usage

```make``` will compile and run the program.

## Structure

The main directory contains ```recommendo.go``` the main executable and a Makefile.

```/reco``` contains the ```recommend(..)``` func which uses all the other packages to do create a list of recommended items to a given user.

```/algo``` contains ```algo.go``` with the generic simFunc function which abstracts euclidean and pearson algorithms. It also contains euclidean and pearson algorithms in there respective ```*.go``` files.

```/sort``` contains map sorting algorithm.

```/data``` contains data schemas and initializer functions.