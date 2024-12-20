package nex

import "sync"

// Numeric is a constraint that matches any numeric type.
type Numeric interface {
	int | int32 | int64 | float32 | float64
}

// Incrementer is a generic type that increments a numeric value.
type Incrementer[T Numeric] struct {
	value T
	mu    sync.Mutex
}

// NewIncrementer creates a new Incrementer with the initial value.
func NewIncrementer[T Numeric](initialValue T) *Incrementer[T] {
	return &Incrementer[T]{value: initialValue}
}

// Increment increments the value and returns the new value.
func (i *Incrementer[T]) Increment() T {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.value++
	return i.value
}

// Value returns the current value.
func (i *Incrementer[T]) Value() T {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.value
}
