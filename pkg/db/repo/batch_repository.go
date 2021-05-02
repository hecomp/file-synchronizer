package repo

import (
	"github.com/go-kit/kit/log"
	"github.com/hecomp/file-synchronizer/pkg/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const GlobalBinRange string = "global_bin_range"

type IBatchRepository interface {
	UploadFirstDataBinRange(records []*db.GlobalBinRange) error
}

type BatchRepository struct {
	db     *gorm.DB
	logger log.Logger
}

func NewBatchRepository(db *gorm.DB, logger log.Logger) IBatchRepository {
	return &BatchRepository{
		db,
		logger,
	}
}

func (b BatchRepository) UploadFirstDataBinRange(records []*db.GlobalBinRange) error {
	b.logger.Log("creating files")

	result := b.db.Scopes(b.GlobalBinTable()).Clauses(clause.OnConflict{DoNothing: true}).Create(&records)
	if err := result.Error; err != nil {
		return err
	}
	return nil
	//
	//insertQuery := `insert into global_bin_range (bin_card_length_value, bin_low_value, bin_high_value, bin_range_length_value, card_issuer_bank_name, issuer_country_code, bank_card_type_indicator, bank_card_use_indicator, prepaid_indicator, regulator_bank_issued_indicator, product_id, card_class, credit_host_id, network_id, bin_update, bin_electronic_balance_transfer_state, bin_flexible_spending_account_indicator, bin_regulator_indicator, bin_product_subtype, bin_large_ticket_indicator, bin_account_level_processing_indicator, bin_account_funding_source, crt_ts, last_upd_ts)
    //                values (:BinCardLenVal, :BinLowVal, :BinHighVal, :BinRngeLenVal, :CardIssuerBankName, :IssuerCntyCd, :BankCardTypInd, :BankCardUseInd, :PrepaidInd, :RegulatorBankIssuedInd, :ProductId, :CardClass, :CreditHostId, :NetworkId, :BinUpdate, :BinEbtState, :BinFsaInd, :BinRegulatorInd, :BinProductSubtype, :BinLargeTicketInd, :BinAlpInd, :BinAcctFundSrc, :CrtTs, :LastUpdTs)`
	//_, err := b.db.Exec(insertQuery, pq.Array(records))
	//return err
}

func (b BatchRepository) GlobalBinTable() func (tx *gorm.DB) *gorm.DB {
	return func (tx *gorm.DB) *gorm.DB {
		return tx.Table(GlobalBinRange)
	}
}