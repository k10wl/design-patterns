package main

type Component interface {
	Search(string) string
	GetPath(int) string
}
