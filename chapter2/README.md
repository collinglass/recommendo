# Chapter 2

This directory holds the go code for Chapter 2 which focuses on user based and item based filtering techniques.

As discussed in the book, user-based filtering is rather inefficient because it relies on rerunning the algorithm everytime a user makes a new action.

## Usage

```make``` will compile and run the program.

## Structure

The main directory contains ```userbased.go``` and ```itembased.go``` the main executables and a Makefile.

```/reco``` contains the ```Recommend(..)``` func which uses all the other packages to create a list of recommended items to a given user. It also holds the ```TransformPrefs(prefs)``` function which transforms the prefs map from user-item to item-user ratings, used in Item-Based filtering.

```/algo``` contains ```algo.go``` with the generic simFunc function which abstracts euclidean and pearson algorithms. It also contains euclidean and pearson algorithms in there respective ```*.go``` files.

```/sort``` contains map sorting algorithm.

```/data``` contains data schemas and initializer functions.