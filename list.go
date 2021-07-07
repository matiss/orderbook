package orderbook

import (
	"errors"
)

const (
	// OrderBookSideBids Bid side
	OrderBookSideBids = iota
	// OrderBookSideAsks Ask side
	OrderBookSideAsks
)

// ListNode struct
type ListNode struct {
	Price int64
	Size  int64

	next *ListNode
}

// List struct
type List struct {
	len  int
	head *ListNode
}

// AddFront node
func (l *List) AddFront(price, size int64) {
	node := &ListNode{
		Price: price,
		Size:  size,
	}

	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}

	l.len++
}

// AddBack node
func (l *List) AddBack(price, size int64) {
	node := &ListNode{
		Price: price,
		Size:  size,
	}

	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}

	l.len++
}

// UpdateOrAddAsc node
func (l *List) UpdateOrAddAsc(price, size int64) {
	node := &ListNode{
		Price: price,
		Size:  size,
	}

	// Empty list
	if l.head == nil {
		// Insert
		l.head = node
		l.len = 1
		return
	}

	// Replace head node if needed
	if l.head.Price > price {
		node.next = l.head
		l.head = node
		l.len++
		return
	}

	// Traverse
	current := l.head

	for current != nil {
		if current.Price == price {
			// Found node! Update current node.
			current.Price = price
			current.Size = size
			break
		} else if price > current.Price {
			// Validate next
			if (current.next != nil) && (price >= current.next.Price) {
				current = current.next
				continue
			}

			// Insert before current
			node.next = current.next
			current.next = node

			l.len++

			break
		}

		current = current.next
	}
}

// UpdateOrAddDesc node
func (l *List) UpdateOrAddDesc(price, size int64) {
	node := &ListNode{
		Price: price,
		Size:  size,
	}

	// Empty list
	if l.head == nil {
		// Insert
		l.head = node
		l.len = 1
		return
	}

	// Replace head node if needed
	if l.head.Price < price {
		node.next = l.head
		l.head = node
		l.len++
		return
	}

	// Traverse
	current := l.head

	for current != nil {
		if current.Price == price {
			// Found node! Update current node.
			current.Price = price
			current.Size = size
			break
		} else if price < current.Price {
			// Validate next
			if (current.next != nil) && (price <= current.next.Price) {
				current = current.next
				continue
			}

			// Insert before current
			node.next = current.next
			current.next = node

			l.len++

			break
		}

		current = current.next
	}
}

// Prune nodes
func (l *List) Prune(length int) {
	i := 1
	current := l.head
	for current != nil {
		if i == length {
			current.next = nil
			l.len = i
			break
		}

		i++
		current = current.next
	}
}

// Remove node
func (l *List) Remove(price int64) error {

	if l.head == nil {
		return errors.New("Remove: List is empty")
	}

	removed := false

	// Traverse
	var prev *ListNode
	current := l.head

	for current != nil {
		if current.Price == price {
			// Found node! Remove!
			if prev == nil {
				// Remove head node
				l.head = current.next
			} else {
				// Remove current node
				prev.next = current.next
				current = current.next
			}

			l.len--
			removed = true
			break
		}

		prev = current
		current = current.next
	}

	if !removed {
		return errors.New("Remove: node not found")
	}

	return nil
}

// RemoveFront node
func (l *List) RemoveFront() error {
	if l.head == nil {
		return errors.New("RemoveFront: List is empty")
	}

	l.head = l.head.next
	l.len--

	return nil
}

// RemoveBack node
func (l *List) RemoveBack() error {
	if l.head == nil {
		return errors.New("RemoveBack: List is empty")
	}
	var prev *ListNode
	current := l.head

	for current.next != nil {
		prev = current
		current = current.next
	}

	if prev != nil {
		prev.next = nil
	} else {
		l.head = nil
	}

	l.len--

	return nil
}

// Front node
func (l *List) Front() (*ListNode, error) {
	if l.head == nil {
		return nil, errors.New("Front: List is empty")
	}
	return l.head, nil
}

// Last node
func (l *List) Last() (*ListNode, error) {
	if l.head == nil {
		return nil, errors.New("Front: List is empty")
	}

	current := l.head
	for current != nil {
		if current.next == nil {
			// Found last
			return current, nil
		}

		current = current.next
	}

	return nil, nil
}

// Size of list
func (l *List) Size() int {
	return l.len
}
