package main

type Iterator interface {
	getNext() *User
	hasMore() bool
}
