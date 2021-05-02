package db

import "time"

type GlobalBinRange struct {
	BinCardLenVal          int       `gorm:"column:bin_card_length_value;not null"`
	BinLowVal              string    `gorm:"column:bin_low_value;type:varchar(16);not null"`
	BinHighVal             string    `gorm:"column:bin_high_value;type:varchar(16);not null"`
	BinRngeLenVal          int       `gorm:"column:bin_range_length_value;not null"`
	CardIssuerBankName     string    `gorm:"column:card_issuer_bank_name;type:varchar(60);not null"`
	IssuerCntyCd           string    `gorm:"column:issuer_country_code;type:varchar(3)"`
	BankCardTypInd         string    `gorm:"column:bank_card_type_indicator;type:varchar(1);not null"`
	BankCardUseInd         string    `gorm:"column:bank_card_use_indicator;type:varchar(1);not null"`
	PrepaidInd             string    `gorm:"column:prepaid_indicator;type:varchar(1)"`
	RegulatorBankIssuedInd string    `gorm:"column:regulator_bank_issued_indicator;type:varchar(1)"`
	ProductId              string    `gorm:"column:product_id;type:varchar(3);not null"`
	CardClass              string    `gorm:"column:card_class;type:varchar(1);not null"`
	CreditHostId           string    `gorm:"column:credit_host_id;type:varchar(2)"`
	NetworkId              string    `gorm:"column:network_id;type:varchar(2);not null"`
	BinUpdate              string    `gorm:"column:bin_update;type:varchar(8);not null"`
	BinEbtState            string    `gorm:"column:bin_electronic_balance_transfer_state;type:varchar(2);not null"`
	BinFsaInd              string    `gorm:"column:bin_flexible_spending_account_indicator;type:varchar(1);not null"`
	BinRegulatorInd        string    `gorm:"column:bin_regulator_indicator;type:varchar(1);not null"`
	BinProductSubtype      string    `gorm:"column:bin_product_subtype;type:varchar(2);not null"`
	BinLargeTicketInd      string    `gorm:"column:bin_large_ticket_indicator;type:varchar(1);not null"`
	BinAlpInd              string    `gorm:"column:bin_account_level_processing_indicator;type:varchar(1);not null"`
	BinAcctFundSrc         string    `gorm:"column:bin_account_funding_source;type:varchar(1);not null"`
	CrtTs                  time.Time `gorm:"column:crt_ts;not null"`
	LastUpdTs              time.Time `gorm:"column:last_upd_ts;not null"`
}
