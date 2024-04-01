package main

type CheckChain interface {
	setNext(CheckChain)
	handle(*Guest)
}
