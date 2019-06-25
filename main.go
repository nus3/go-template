package main

import "fmt"

// Agrigator is
type Agrigator interface {
	CreateIterator() Iterator
}

// Iterator is
type Iterator interface {
	HasNext() bool
	Next() Object
}

// Object is
type Object interface {
	GetName() string
}

// BookShelf is
type BookShelf struct {
	books []*Book
	last  int
}

// CreateIterator returns
func (b *BookShelf) CreateIterator() Iterator {
	var i Iterator
	i = &BookShelfIterator{bookShelf: b, index: 0}

	return i
}

// Append returns
func (b *BookShelf) Append(book *Book) {
	b.books = append(b.books, book)
	b.last++
}

// GetLength returns
func (b BookShelf) GetLength() int {
	return b.last
}

// GetBookAt returns
func (b BookShelf) GetBookAt(index int) *Book {
	return b.books[index]
}

// BookShelfIterator is
type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

// HasNext returns
func (b BookShelfIterator) HasNext() bool {
	if b.index < b.bookShelf.GetLength() {
		return true
	}
	return false
}

// Next returns
func (b *BookShelfIterator) Next() Object {
	book := b.bookShelf.GetBookAt(b.index)
	b.index++

	return book
}

// Book is
type Book struct {
	name string
}

// GetName returns
func (b Book) GetName() string {
	return b.name
}

func main() {
	bookShelf := &BookShelf{}
	bookShelf.Append(&Book{"夜は短し歩けよ乙女"})
	bookShelf.Append(&Book{"有頂天家族"})
	bookShelf.Append(&Book{"ペンギンハイウェイ"})

	var a Agrigator = bookShelf
	iterator := a.CreateIterator()
	for iterator.HasNext() {
		book := iterator.Next()
		fmt.Println(book.GetName())
	}
}
