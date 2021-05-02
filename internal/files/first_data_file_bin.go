package files

type FirstDataBINFile struct {
	//ffparser is one indexed, column starts at 1
	BinCardLenVal          int    `flatfile:"36,2"`
	BinLowVal              string `flatfile:"2,16"`
	BinHighVal             string `flatfile:"18,16"`
	BinRngeLenVal          int    `flatfile:"34,2"`
	CardIssuerBankName     string `flatfile:"38,60"`
	IssuerCntyCd           string `flatfile:"98,3"`
	BankCardTypInd         string `flatfile:"101,1"`
	BankCardUseInd         string `flatfile:"102,1"`
	PrepaidInd             string `flatfile:"195,1"`
	RegulatorBankIssuedInd string `flatfile:"199,1"`
	ProductId              string `flatfile:"196,3"`
	CardClass              string `flatfile:"205,1"`
	BinUpdate              string `flatfile:"104,2"`
	BinEbtState            string `flatfile:"172,2"`
	BinFsaInd              string `flatfile:"176,1"`
	BinRegulatorInd        string `flatfile:"199,1"`
	BinProductSubtype      string `flatfile:"200,2"`
	BinLargeTicketInd      string `flatfile:"202,1"`
	BinAlpInd              string `flatfile:"203,1"`
	BinAcctFundSrc         string `flatfile:"204,1"`
}