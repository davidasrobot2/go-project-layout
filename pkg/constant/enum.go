package constant

// SourceOfFund - START
type SourceOfFund string

const (
	SourceOfFundQR           SourceOfFund = "QR"
	SourceOfFundCash         SourceOfFund = "CS"
	SourceOfFundBankTransfer SourceOfFund = "BT"
	SourceOfFundDebitCard    SourceOfFund = "DC"
	SourceOfFundCreditCard   SourceOfFund = "CC"
)

func (s *SourceOfFund) Scan(value interface{}) error {
	if val, ok := value.([]byte); ok {
		*s = SourceOfFund(val)
	} else if val, ok := value.(string); ok {
		*s = SourceOfFund(val)
	}
	return nil
}

func (s *SourceOfFund) Value() (string, error) {
	return string(*s), nil
}

// SourceOfFund - END

// TransactionType - START

type TransactionType string

const (
	TransactionTypeSale     TransactionType = "SALE"
	TransactionTypeVoid     TransactionType = "VOID"
	TransactionTypeRefund   TransactionType = "REFUND"
	TransactionTypeTopup    TransactionType = "TOPUP"
	TransactionTypeTransfer TransactionType = "TRANSFER"
)

func (s *TransactionType) Scan(value interface{}) error {
	if val, ok := value.([]byte); ok {
		*s = TransactionType(val)
	} else if val, ok := value.(string); ok {
		*s = TransactionType(val)
	}
	return nil
}

func (s *TransactionType) Value() (string, error) {
	return string(*s), nil
}

// TransactionType - END

// TransactionStatus - START

type TransactionStatus int

const (
	TransactionStatusWaitingForPayment TransactionStatus = 1
	TransactionStatusPending           TransactionStatus = 2
	TransactionStatusSuccess           TransactionStatus = 3
	TransactionStatusCanceled          TransactionStatus = 4
	TransactionStatusFailed            TransactionStatus = 99
)

func (s *TransactionStatus) Scan(value interface{}) error {
	if val, ok := value.(int); ok {
		*s = TransactionStatus(val)
	}
	return nil
}

func (s *TransactionStatus) Value() (int, error) {
	return int(*s), nil
}

// TransactionStatus - END

// MerchantType - START

type MerchantType string

const (
	MerchantTypeFnB MerchantType = "F&B"
	MerchantTypePOS MerchantType = "POS"
	MerchantTypeATM MerchantType = "ATM"
)

func (s *MerchantType) Scan(value interface{}) error {
	if val, ok := value.([]byte); ok {
		*s = MerchantType(val)
	} else if val, ok := value.(string); ok {
		*s = MerchantType(val)
	}
	return nil
}

func (s *MerchantType) Value() (string, error) {
	return string(*s), nil
}

// MerchantType - END

// AccountType - START

type AccountType string

const (
	AccountTypeSavingAccount AccountType = "SA"
	AccountTypeGiro          AccountType = "CA"
	AccountTypeInvestment    AccountType = "IN"
	AccountTypeFixedDeposit  AccountType = "FD"
)

// AccountType - END

// CardType - START

type CardType string

const (
	CardTypeDebit  CardType = "DC"
	CardTypeCredit CardType = "CC"
)

func (s *CardType) Scan(value interface{}) error {
	if val, ok := value.([]byte); ok {
		*s = CardType(val)
	} else if val, ok := value.(string); ok {
		*s = CardType(val)
	}
	return nil
}

func (s *CardType) Value() (string, error) {
	return string(*s), nil
}

// CardType - END

// AdministratorLevel - START

type AdministratorLevel int

const (
	AdministratorLevel1 AdministratorLevel = iota + 1
	AdministratorLevel2 AdministratorLevel = iota + 1
	AdministratorLevel3 AdministratorLevel = iota + 1
)

func (s *AdministratorLevel) Scan(value interface{}) error {
	if val, ok := value.(int); ok {
		*s = AdministratorLevel(val)
	}
	return nil
}

func (s *AdministratorLevel) Value() (int, error) {
	return int(*s), nil
}

// AdministratorLevel - END
