CREATE TABLE IF NOT EXISTS global_bin_range (
	bin_card_length_value 			INT		NOT NULL,
	bin_low_value 				VARCHAR(16) 		NOT NULL,
	bin_high_value 				VARCHAR(16) 		NOT NULL,
	bin_range_length_value 			bigint 		NOT NULL,
	card_issuer_bank_name 			VARCHAR(60) 		NOT NULL,
	issuer_country_code 			VARCHAR(3),
	bank_card_type_indicator 			VARCHAR(1) 		NOT NULL,
	bank_card_use_indicator 			VARCHAR(1) 		NOT NULL,
	prepaid_indicator 				VARCHAR(1),
	regulator_bank_issued_indicator		VARCHAR(1),
	credit_host_id 				VARCHAR(2)		NOT NULL,
	product_id 				VARCHAR(3) 		NOT NULL,
	card_class 				VARCHAR(1) 		NOT NULL,
	network_id 				VARCHAR(2) 		NOT NULL,
	bin_update 				VARCHAR(8)		NOT NULL,
	bin_electronic_balance_transfer_state 		VARCHAR(2) 		NOT NULL,
	bin_flexible_spending_account_indicator 	VARCHAR(1) 		NOT NULL,
	bin_regulator_indicator 			VARCHAR(1) 		NOT NULL,
	bin_product_subtype 			VARCHAR(2) 		NOT NULL,
	bin_large_ticket_indicator 			VARCHAR(1) 		NOT NULL,
	bin_account_level_processing_indicator 	VARCHAR(1) 		NOT NULL,
	bin_account_funding_source 			VARCHAR(1) 		NOT NULL,
	crt_ts 					TIMESTAMPTZ 	NOT NULL,
	last_upd_ts 				TIMESTAMPTZ 	NOT NULL,
    PRIMARY KEY (bin_card_length_value, bin_low_value, bin_high_value)
);

--                                                             ,
-- INDEX postgres.public.global_bin_range (bin_card_length_value, bin_low_value, bin_high_value) STORING (bin_range_length_value, card_issuer_bank_name, issuer_country_code, bank_card_type_indicator, bank_card_use_indicator, prepaid_indicator, regulator_bank_issued_indicator, credit_host_id, product_id, card_class, network_id, bin_update, bin_electronic_balance_transfer_state, bin_flexible_spending_account_indicator, bin_regulator_indicator, bin_product_subtype, bin_large_ticket_indicator, bin_account_level_processing_indicator, bin_account_funding_source),
-- FAMILY "primary" (bin_card_length_value, bin_low_value, bin_high_value)