package coins

// Each coin has a plugin that implements the following interface:

// Assuming other structures like Address, Account, Price, PaymentOption, Confirmation, Transaction, and Payment are defined elsewhere

type GetNewAddress struct {
	Account Account
	Address Address
}

type BroadcastTx struct {
	Txhex string
	Txid  string
	Txkey string
}

type Plugin struct {
	Currency string
	Chain    string
	Decimals int
	Token    string
}

type BuildSignedPayment struct {
	PaymentOption PaymentOption
	Mnemonic      string
}

type BroadcastTxResult struct {
	Txid    string
	Txhex   string
	Success bool
	Result  interface{}
}

type VerifyPayment struct {
	PaymentOption PaymentOption
	Transaction   Transaction
	Protocol      string
}

type PaymentOption struct {
	Currency string
	Amount   float64
	Address  string // Optional: For cryptocurrencies, etc.
	// Add more fields as needed based on your specific requirements
}

type ValidateUnsignedTx struct {
	PaymentOption PaymentOption
	Transactions  []Transaction
}

type Transaction struct {
	TxID   string // Transaction ID
	TxHex  string // Raw transaction hex
	TxKey  string // Transaction key or secret, if applicable
	Amount float64
	From   string // Sender's address or identifier
	To     string // Recipient's address or identifier
	// Include additional fields as needed
}

type AbstractPlugin interface {
	BuildSignedPayment(params BuildSignedPayment) (Transaction, error)
	VerifyPayment(params VerifyPayment) (bool, error)
	ValidateAddress(address string) (bool, error)
	GetTransaction(txid string) (Transaction, error)
	BroadcastTx(params BroadcastTx) (BroadcastTxResult, error)
	GetNewAddress(params GetNewAddress) (string, error)
	TransformAddress(address string) (string, error)
	GetConfirmation(txid string) (*Confirmation, error)
	GetPayments(txid string) ([]Payment, error)
	ParsePayments(transaction Transaction) ([]Payment, error)
	GetPrice() (Price, error)
}

type Account struct {
	ID     int
	Name   string
	Email  string
	Status string // e.g., "active", "suspended"
	// Additional fields as needed
}

type Address struct {
	ID         int
	AccountID  int // Assuming an address is associated with an account
	Street     string
	City       string
	State      string
	PostalCode string
	Country    string
	// Additional fields as needed
}

type Payment struct {
	ID        int
	AccountID int
	Amount    float64
	Currency  string
	Status    string // e.g., "pending", "completed", "failed"
	// Additional fields as needed
}

type Price struct {
	ID       int
	ItemID   int // Assuming prices are associated with specific items or services
	Amount   float64
	Currency string
	// Additional fields as needed
}

type Confirmation struct {
	ID            int
	TransactionID int
	Status        string // e.g., "confirmed", "unconfirmed"
	// Additional fields as needed, like timestamps
}
