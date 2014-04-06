# Recommendo

Recommendation system for bookAPI.

## Usage

```make``` will compile and run the program.

## Structure

The main directory contains ```userbased.go``` and ```itembased.go``` the main executables and a Makefile.

```/reco``` contains ```userbasedreco.go``` and ```itembasedreco.go``` which are controllers using all the other packages to create a list of recommended items to a given user.

```/algo``` contains ```algo.go``` with the generic simFunc function which abstracts euclidean and pearson algorithms. It also contains euclidean and pearson algorithms in their respective ```*.go``` files.

```/sort``` contains map sorting algorithm.

```/data``` contains data schemas and initializer functions.