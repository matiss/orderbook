package orderbook

import (
	"testing"
)

// TestFindUpdateOrAddDesc node
func TestUpdateOrAddAsc(t *testing.T) {
	list := List{
		len:  0,
		head: nil,
	}

	// Operations
	list.UpdateOrAddAsc(124005, 100)
	list.UpdateOrAddAsc(124000, 100)
	list.UpdateOrAddAsc(123000, 200)
	list.UpdateOrAddAsc(123500, 3300)
	list.UpdateOrAddAsc(124000, 200)

	if list.len != 4 {
		t.Errorf("Invalid node count! Expected: %d, got: %d", 4, list.len)
		return
	}

	if list.head.Price != 123000 {
		t.Errorf("Invalid head node! Expected: %d, got: %d", 123000, list.head.Price)
	}

	second := list.head.next
	if second.Price != 123500 {
		t.Errorf("Invalid second node! Expected: %d, got: %d", 123500, second.Price)
	}

	third := second.next
	if third.Price != 124000 || third.Size != 200 {
		t.Errorf("Invalid third node! Expected: %d, got: %d", 124000, third.Price)
	}

	fourth := third.next
	if fourth.Price != 124005 {
		t.Errorf("Invalid fourth node! Expected: %d, got: %d", 124005, fourth.Price)
	}
}

// TestFindUpdateOrAddDesc node
func TestUpdateOrAddDesc(t *testing.T) {
	list := List{
		len:  0,
		head: nil,
	}

	// Operations
	list.UpdateOrAddDesc(124005, 100)
	list.UpdateOrAddDesc(124000, 100)
	list.UpdateOrAddDesc(123500, 3300)
	list.UpdateOrAddDesc(123000, 200)
	list.UpdateOrAddDesc(123500, 78012)

	if list.len != 4 {
		t.Errorf("Invalid node count! Expected: %d, got: %d", 4, list.len)
		return
	}

	if list.head.Price != 124005 {
		t.Errorf("Invalid head node! Expected: %d, got: %d", 124005, list.head.Price)
	}

	second := list.head.next
	if second.Price != 124000 {
		t.Errorf("Invalid second node! Expected: %d, got: %d", 124000, second.Price)
	}

	third := second.next
	if third.Price != 123500 || third.Size != 78012 {
		t.Errorf("Invalid third node! Expected: %d, got: %d", 123500, third.Price)
	}

	fourth := third.next
	if fourth.Price != 123000 {
		t.Errorf("Invalid fourth node! Expected: %d, got: %d", 123000, fourth.Price)
	}
}

// TestPrune nodes
func TestPrune(t *testing.T) {
	list := List{
		len: 4,
		head: &ListNode{
			Price: 100000,
			Size:  200,
			next: &ListNode{
				Price: 100001,
				Size:  100,
				next: &ListNode{
					Price: 100002,
					Size:  130,
					next: &ListNode{
						Price: 100010,
						Size:  50,
						next:  nil,
					},
				},
			},
		},
	}

	// Prune
	list.Prune(3)

	// Validate length
	if list.len != 3 {
		t.Errorf("Invalid list length! Expected: %d, got: %d", 3, list.len)
	}

	// Validate head node
	if list.head.Price != 100000 {
		t.Errorf("Invalid head node! Expected: %d, got: %d", 100000, list.head.Price)
	}

	// Validate last node
	current := list.head
	last := current
	for current != nil {
		if current.next == nil {
			last = current
		}

		current = current.next
	}

	if last.Price != 100002 {
		t.Errorf("Invalid last node! Expected: %d, got: %d", 100002, last.Price)
	}
}

// TestRemove node
func TestRemove(t *testing.T) {
	list := List{
		len: 4,
		head: &ListNode{
			Price: 100000,
			Size:  200,
			next: &ListNode{
				Price: 100001,
				Size:  100,
				next: &ListNode{
					Price: 100002,
					Size:  130,
					next: &ListNode{
						Price: 100010,
						Size:  50,
						next:  nil,
					},
				},
			},
		},
	}

	// Remove head node
	err := list.Remove(100000)
	if err != nil {
		t.Error(err)
	}

	// Validate head node
	if list.head.Price != 100001 || list.len != 3 {
		t.Errorf("Invalid head node! Expected: %d, got: %d", 100001, list.head.Price)
	}

	// Try to remove non-existing node
	err = list.Remove(0)
	if err == nil || list.len != 3 {
		t.Errorf("Expected an error")
	}

	// Remove node in the middle
	err = list.Remove(100002)
	if err != nil {
		t.Error(err)
	}

	// Validate list
	if list.len != 2 {
		t.Errorf("Invalid list length. Expected: %d, got: %d", 2, list.len)
	}

	if list.head.Price != 100001 {
		t.Errorf("Invalid head node! Expected: %d, got: %d", 100001, list.head.Price)
	}

	if list.head.next.Price != 100010 {
		t.Errorf("Invalid last node! Expected: %d, got: %d", 100010, list.head.next.Price)
	}
}
