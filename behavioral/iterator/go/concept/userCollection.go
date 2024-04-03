package main

type IterableUserCollection interface {
	createUserIterator() UserIterator
}

type UserCollection struct {
	users []User
}

func (collection UserCollection) createUserIterator() UserIterator {
	return newUserArrayIterator(collection)
}
