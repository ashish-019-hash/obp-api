package repositories

type BankRepository interface {
}

type AccountRepository interface {
}

type CustomerRepository interface {
}

type TransactionRepository interface {
}

type InMemoryBankRepository struct{}
type InMemoryAccountRepository struct{}
type InMemoryCustomerRepository struct{}
type InMemoryTransactionRepository struct{}
