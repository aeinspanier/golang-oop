package iterator

type Collection interface {
    CreateIterator() Iterator
}