// file: invoices/invoices.go
package invoices

import (
	"errors"
)

/* Invoices have three essential properties, a value (currency and amount), and a status,
which can be one of pending, paid, or canceled. The third crucial value is the recipient details */

// Invoice represents a basic invoice structure
type Invoice struct {
	ID       string
	Amount   int
	Currency string
	Status   string
}

// ErrInvoiceNotFound is returned when an invoice cannot be found
var ErrInvoiceNotFound = errors.New("invoice not found")

// Invoices defines the interface for invoice management
type Invoices interface {
	Create(inv Invoice) (Invoice, error)
	FindOne(id string) (Invoice, error)
	FindMany(ids []string) ([]Invoice, error)
	Cancel(id string) error
}

// mockInvoices implements the Invoices interface for testing and simple usage
type mockInvoices struct {
	invoices map[string]Invoice
}

// New creates a new instance of mockInvoices
func New() Invoices {
	return &mockInvoices{
		invoices: make(map[string]Invoice),
	}
}

// Create adds a new Invoice to the mock store
func (m *mockInvoices) Create(inv Invoice) (Invoice, error) {
	m.invoices[inv.ID] = inv
	return inv, nil
}

// FindOne retrieves a single Invoice by ID
func (m *mockInvoices) FindOne(id string) (Invoice, error) {
	if inv, ok := m.invoices[id]; ok {
		return inv, nil
	}
	return Invoice{}, ErrInvoiceNotFound
}

// FindMany retrieves multiple Invoices by their IDs
func (m *mockInvoices) FindMany(ids []string) ([]Invoice, error) {
	results := []Invoice{}
	for _, id := range ids {
		if inv, ok := m.invoices[id]; ok {
			results = append(results, inv)
		}
	}
	return results, nil
}

// Cancel marks an Invoice as canceled
func (m *mockInvoices) Cancel(id string) error {
	if inv, ok := m.invoices[id]; ok {
		inv.Status = "canceled"
		m.invoices[id] = inv
		return nil
	}
	return ErrInvoiceNotFound
}
