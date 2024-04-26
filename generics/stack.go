package main

import (
	"errors"
	"reflect"
)

type Stack[T any] struct {
	store []T
}

func (s *Stack[T]) Push(item T) {
	s.store = append(s.store, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Size() == 0 {
		return reflect.Zero(reflect.TypeOf(s.store[0])).Interface().(T), errors.New("Stack is empty")
	}

	topIndex := len(s.store) - 1
	top := s.store[topIndex]
	s.store = s.store[:topIndex]
	return top, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.Size() == 0 {
		return reflect.Zero(reflect.TypeOf(s.store[0])).Interface().(T), errors.New("Stack is empty")
	}

	topIndex := len(s.store) - 1
	top := s.store[topIndex]
	return top, nil
}

func (s *Stack[T]) Size() int {
	return len(s.store)
}
