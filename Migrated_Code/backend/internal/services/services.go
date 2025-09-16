package services

type BankService struct{}

func NewBankService() *BankService {
	return &BankService{}
}

type AccountService struct{}

func NewAccountService() *AccountService {
	return &AccountService{}
}

type CustomerService struct{}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

type TransactionService struct{}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}
