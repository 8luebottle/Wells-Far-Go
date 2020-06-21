package repository

type BankStorer interface {
}

type bankRepository struct{}

func ParseBankStorer() BankStorer {
	return &bankRepository{}
}
