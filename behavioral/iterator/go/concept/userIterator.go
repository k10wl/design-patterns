package main

type UserIterator interface {
	getNext() *User
	hasMore() bool
}

type UserArrayIterator struct {
	iterationIndex int
	collection     UserCollection
}

func newUserArrayIterator(collection UserCollection) *UserArrayIterator {
	return &UserArrayIterator{
		iterationIndex: 0,
		collection:     collection,
	}
}

func (iterator *UserArrayIterator) getNext() *User {
	if !iterator.hasMore() {
		return nil
	}
	user := &iterator.collection.users[iterator.iterationIndex]
	iterator.iterationIndex++
	return user
}

func (iterator *UserArrayIterator) hasMore() bool {
	return iterator.iterationIndex < len(iterator.collection.users)
}
