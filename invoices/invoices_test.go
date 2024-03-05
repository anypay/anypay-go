// file: invoices/invoices_test.go
package invoices

import (
	"testing"
)

func TestInvoiceLifeCycle(t *testing.T) {
	invService := New()
	inv := Invoice{ID: "1", Amount: 100, Currency: "USD", Status: "pending"}

	// Test Create
	createdInv, err := invService.Create(inv)
	if err != nil || createdInv.ID != inv.ID {
		t.Fatalf("Expected invoice to be created, got error: %v", err)
	}

	// Test FindOne
	foundInv, err := invService.FindOne(inv.ID)
	if err != nil || foundInv.ID != inv.ID {
		t.Fatalf("Expected to find invoice with ID %s, got error: %v", inv.ID, err)
	}

	// Test FindMany
	foundInvs, err := invService.FindMany([]string{inv.ID})
	if err != nil || len(foundInvs) != 1 || foundInvs[0].ID != inv.ID {
		t.Fatalf("Expected to find one invoice in FindMany, got error: %v", err)
	}

	// Test Cancel
	err = invService.Cancel(inv.ID)
	if err != nil {
		t.Fatalf("Expected to cancel invoice, got error: %v", err)
	}
	canceledInv, _ := invService.FindOne(inv.ID)
	if canceledInv.Status != "canceled" {
		t.Fatalf("Expected invoice status to be 'canceled', got %s", canceledInv.Status)
	}
}

// Returns a new instance of Invoices
func TestNewReturnsNewInstance(t *testing.T) {
	invoices := New()
	if invoices == nil {
		t.Errorf("Expected a new instance of Invoices, but got nil")
	}
}
