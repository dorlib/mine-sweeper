package main

type Cell struct {
	IsBomb     bool
	IsVisible  bool
	IsHasFlag  bool
	closeBombs int
}
