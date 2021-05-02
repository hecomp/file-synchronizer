package db

import (
	"time"

	"github.com/hecomp/file-synchronizer/internal/files"
)

type IGlobalBinRange interface {}

func NewFirstDataBinRange(binData *files.FirstDataBINFile) IGlobalBinRange{
	return &GlobalBinRange{
		BinCardLenVal: binData.BinCardLenVal,
		BinLowVal: binData.BinLowVal,
		BinHighVal: binData.BinHighVal,
		BinRngeLenVal: binData.BinRngeLenVal,
		CardIssuerBankName: binData.CardIssuerBankName,
		IssuerCntyCd: binData.IssuerCntyCd,
		BankCardTypInd: binData.BankCardTypInd,
		PrepaidInd: binData.PrepaidInd,
		RegulatorBankIssuedInd: binData.RegulatorBankIssuedInd,
		ProductId: binData.ProductId,
		CardClass: binData.CardClass,
		CreditHostId: "",
		NetworkId: "",
		BinUpdate: binData.BinUpdate,
		BinEbtState: binData.BinEbtState,
		BinFsaInd: binData.BinFsaInd,
		BinRegulatorInd: binData.BinRegulatorInd,
		BinProductSubtype: binData.BinProductSubtype,
		BinLargeTicketInd: binData.BinLargeTicketInd,
		BinAlpInd: binData.BinAlpInd,
		BinAcctFundSrc: binData.BinAcctFundSrc,
		CrtTs: time.Now(),
		LastUpdTs: time.Now(),
	}
}
