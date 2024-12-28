-- sureshkrishnanv.a_mf_allot_redem_resp definition

CREATE TABLE `a_mf_allot_redem_resp` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Batch_Id` varchar(20) DEFAULT NULL,
  `AllottedAmt` varchar(14) DEFAULT NULL,
  `AllottedUnit` varchar(14) DEFAULT NULL,
  `AllottedNav` varchar(14) DEFAULT NULL,
  `Redemption_Nav` varchar(14) DEFAULT NULL,
  `Amount` varchar(14) DEFAULT NULL,
  `RedeemAmt` varchar(14) DEFAULT NULL,
  `BeneficiaryId` varchar(18) DEFAULT NULL,
  `BranchCode` varchar(14) DEFAULT NULL,
  `BuySell` varchar(1) DEFAULT NULL,
  `ClientCode` varchar(10) DEFAULT NULL,
  `ClientName` varchar(50) DEFAULT NULL,
  `DPCFlag` varchar(2) DEFAULT NULL,
  `DPTrans` varchar(2) DEFAULT NULL,
  `EUIN` varchar(10) DEFAULT NULL,
  `EUINDecl` varchar(2) DEFAULT NULL,
  `FolioNo` varchar(14) DEFAULT NULL,
  `ISIN` varchar(20) DEFAULT NULL,
  `InternalRefNo` varchar(20) DEFAULT NULL,
  `KYCFlag` varchar(6) DEFAULT NULL,
  `MemberCode` varchar(8) DEFAULT NULL,
  `OrderDate` date DEFAULT NULL,
  `OrderNo` varchar(19) DEFAULT NULL,
  `OrderType` varchar(6) DEFAULT NULL,
  `Qty` varchar(14) DEFAULT NULL,
  `RTASchemeCode` varchar(10) DEFAULT NULL,
  `RTATransNo` varchar(20) DEFAULT NULL,
  `Remarks` varchar(1000) DEFAULT NULL,
  `ReportDate` date DEFAULT NULL,
  `SIPRegnDate` varchar(10) DEFAULT NULL,
  `SIPRegnNo` varchar(20) DEFAULT NULL,
  `STT` varchar(8) DEFAULT NULL,
  `SchemeCode` varchar(20) DEFAULT NULL,
  `SettNo` varchar(10) DEFAULT NULL,
  `SettType` varchar(8) DEFAULT NULL,
  `SubBrCode` varchar(10) DEFAULT NULL,
  `SubOrderType` varchar(8) DEFAULT NULL,
  `UserId` varchar(10) DEFAULT NULL,
  `ValidFlag` varchar(2) DEFAULT NULL,
  `Created_by` varchar(100) DEFAULT NULL,
  `Created_date` datetime DEFAULT NULL,
  `Updated_by` varchar(100) DEFAULT NULL,
  `Updated_date` datetime DEFAULT NULL,
  `TP_Request_Id` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb3;


-- sureshkrishnanv.a_mf_amc_master_mnt definition

CREATE TABLE `a_mf_amc_master_mnt` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `AMC_Code` varchar(100) NOT NULL,
  `Amc_Name` varchar(100) DEFAULT NULL,
  `Dark_Theme_URL` mediumtext DEFAULT NULL,
  `Light_Theme_URL` mediumtext DEFAULT NULL,
  `Status` varchar(1) NOT NULL,
  `Created_by` varchar(50) DEFAULT NULL,
  `Created_Date` datetime DEFAULT NULL,
  `Updated_by` varchar(50) DEFAULT NULL,
  `Updated_date` datetime DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb3;


-- sureshkrishnanv.a_mf_clients_rec definition

CREATE TABLE `a_mf_clients_rec` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Occupation` varchar(255) DEFAULT NULL,
  `DateOfJoining` date DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;


-- sureshkrishnanv.a_mf_disclaimer_description definition

CREATE TABLE `a_mf_disclaimer_description` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Title` varchar(30) DEFAULT NULL,
  `Content` varchar(1000) DEFAULT NULL,
  `Description` varchar(1000) DEFAULT NULL,
  `Link` varchar(200) DEFAULT NULL,
  `Created_by` varchar(100) DEFAULT NULL,
  `Created_date` datetime DEFAULT NULL,
  `Updated_by` varchar(100) DEFAULT NULL,
  `Updated_date` datetime DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb3;


-- sureshkrishnanv.a_mf_order_status_resp definition

CREATE TABLE `a_mf_order_status_resp` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Amount` varchar(20) DEFAULT NULL,
  `BuySell` varchar(1) DEFAULT NULL,
  `BuySellType` varchar(15) DEFAULT NULL,
  `ClientCode` varchar(20) DEFAULT NULL,
  `DPC` varchar(5) DEFAULT NULL,
  `DPFolioNo` varchar(20) DEFAULT NULL,
  `DPTrans` varchar(20) DEFAULT NULL,
  `Date` varchar(20) DEFAULT NULL,
  `EUIN` varchar(20) DEFAULT NULL,
  `EUINDecl` varchar(20) DEFAULT NULL,
  `EntryBy` varchar(20) DEFAULT NULL,
  `FolioNo` varchar(20) DEFAULT NULL,
  `ISIN` varchar(20) DEFAULT NULL,
  `MINRedeemFlag` varchar(20) DEFAULT NULL,
  `MemberRemarks` varchar(250) DEFAULT NULL,
  `OrderNumber` varchar(20) DEFAULT NULL,
  `OrderRemarks` varchar(250) DEFAULT NULL,
  `OrderStatus` varchar(20) DEFAULT NULL,
  `OrderType` varchar(20) DEFAULT NULL,
  `SIPRegnDate` varchar(20) DEFAULT NULL,
  `SIPRegnNo` varchar(20) DEFAULT NULL,
  `SchemeCode` varchar(20) DEFAULT NULL,
  `SettNo` varchar(20) DEFAULT NULL,
  `SettType` varchar(20) DEFAULT NULL,
  `SubBrCode` varchar(20) DEFAULT NULL,
  `SubBrokerARNCode` varchar(20) DEFAULT NULL,
  `SubOrderType` varchar(20) DEFAULT NULL,
  `Time` varchar(20) DEFAULT NULL,
  `Units` varchar(20) DEFAULT NULL,
  `TP_Request_Id` int(11) DEFAULT NULL,
  `Created_by` varchar(100) DEFAULT NULL,
  `Created_date` datetime DEFAULT NULL,
  `Updated_by` varchar(100) DEFAULT NULL,
  `Updated_date` datetime DEFAULT NULL,
  `Batch_Id` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=70 DEFAULT CHARSET=utf8mb3;


-- sureshkrishnanv.a_mf_orders definition

CREATE TABLE `a_mf_orders` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Broker_id` int(11) DEFAULT NULL,
  `Isin` varchar(20) DEFAULT NULL,
  `OrderId` varchar(12) DEFAULT NULL,
  `TransCode` varchar(3) NOT NULL,
  `TransNo` varchar(19) NOT NULL,
  `ClientCode` varchar(20) NOT NULL,
  `SchemeCd` varchar(20) NOT NULL,
  `BuySell` varchar(1) NOT NULL,
  `BuySellType` varchar(10) NOT NULL,
  `DPTxn` varchar(10) NOT NULL,
  `OrderVal` varchar(14) DEFAULT NULL,
  `Qty` varchar(12) DEFAULT NULL,
  `NavValue` decimal(10,2) NOT NULL,
  `AllRedeem` varchar(1) DEFAULT NULL,
  `FolioNo` varchar(20) DEFAULT NULL,
  `Remarks` varchar(255) DEFAULT NULL,
  `KYCStatus` varchar(1) NOT NULL,
  `RefNo` varchar(20) DEFAULT NULL,
  `SubBrCode` varchar(15) DEFAULT NULL,
  `EUIN` varchar(20) NOT NULL,
  `EUINVal` varchar(1) NOT NULL,
  `MinRedeem` varchar(1) DEFAULT NULL,
  `DPC` varchar(1) NOT NULL,
  `IPAdd` varchar(20) NOT NULL,
  `Parma1` varchar(20) DEFAULT NULL,
  `Param2` varchar(25) DEFAULT NULL,
  `Param3` varchar(20) DEFAULT NULL,
  `MobileNo` varchar(15) DEFAULT NULL,
  `EmailID` varchar(50) DEFAULT NULL,
  `MandateID` varchar(10) DEFAULT NULL,
  `Filler1` varchar(30) DEFAULT NULL,
  `Filler2` varchar(30) DEFAULT NULL,
  `Filler3` varchar(30) DEFAULT NULL,
  `Filler4` varchar(30) DEFAULT NULL,
  `Filler5` varchar(30) DEFAULT NULL,
  `Filler6` varchar(30) DEFAULT NULL,
  `Client_Ref_Id` varchar(20) DEFAULT NULL,
  `TP_Reference_Id` int(11) DEFAULT NULL,
  `Prov_trade_date` varchar(20) DEFAULT NULL,
  `Offline_Order` varchar(20) DEFAULT NULL,
  `Settlement_type` varchar(20) DEFAULT NULL,
  `Prov_settlement_Date` varchar(20) DEFAULT NULL,
  `Scheme_Ref_Id` int(11) DEFAULT NULL,
  `Nav_Ref_Id` int(11) DEFAULT NULL,
  `Concile_Flag` varchar(1) DEFAULT NULL,
  `Sys_Comments` varchar(120) DEFAULT NULL,
  `Concile_Comments` varchar(120) DEFAULT NULL,
  `Uncon_Comments` varchar(120) DEFAULT NULL,
  `CreatedBy` varchar(100) DEFAULT NULL,
  `CreatedDate` datetime DEFAULT NULL,
  `UpdatedBy` varchar(100) DEFAULT NULL,
  `UpdatedDate` datetime DEFAULT NULL,
  `schemeType` varchar(100) DEFAULT NULL,
  `Nfo` varchar(1) DEFAULT NULL,
  `SettlementNo` varchar(30) DEFAULT NULL,
  `PayinStatus` varchar(3) DEFAULT NULL,
  `Usno` varchar(25) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb3;


-- sureshkrishnanv.a_mf_orders_resp definition

CREATE TABLE `a_mf_orders_resp` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `OrdersID` int(11) DEFAULT NULL,
  `Broker_id` int(11) DEFAULT NULL,
  `TransactionCode` varchar(4) CHARACTER SET utf8mb4 DEFAULT NULL,
  `TransNo` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL,
  `TP_OrderId` varchar(12) CHARACTER SET utf8mb4 DEFAULT NULL,
  `ClientCode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Remarks` varchar(1000) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Successflag` varchar(4) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Ep_Status` varchar(3) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Reconcile_Flag` varchar(2) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Flag_Comment` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Cancel_Flag_Comment` varchar(100) DEFAULT NULL,
  `UnFlagged_Comment` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Ep_Status_Comment` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Ep_Status_Updatedby` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Ep_Completed_Time` datetime DEFAULT NULL,
  `Ep_Rejected_Time` datetime DEFAULT NULL,
  `CreatedBy` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `CreatedDate` datetime DEFAULT NULL,
  `UpdatedBy` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `UpdatedDate` datetime DEFAULT NULL,
  `Usno` varchar(25) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb3;


-- sureshkrishnanv.a_mf_orders_status definition

CREATE TABLE `a_mf_orders_status` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `OrderId` int(11) DEFAULT NULL,
  `StatusCode` int(5) DEFAULT NULL,
  `Response_Id` int(11) DEFAULT NULL,
  `Response_Type` varchar(100) DEFAULT NULL,
  `Status_Date` datetime DEFAULT NULL,
  `API_Type` varchar(2) DEFAULT NULL,
  `CreatedBy` varchar(100) DEFAULT NULL,
  `CreatedDate` datetime DEFAULT NULL,
  `UpdatedBy` varchar(100) DEFAULT NULL,
  `UpdatedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=480 DEFAULT CHARSET=utf8mb3;


-- sureshkrishnanv.a_mf_scheme_master definition

CREATE TABLE `a_mf_scheme_master` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `BatchId` varchar(30) CHARACTER SET utf8mb4 DEFAULT NULL,
  `SchemeCurType` varchar(3) CHARACTER SET utf8mb4 DEFAULT NULL,
  `FileEffective_Date` varchar(30) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Pledgeable` varchar(2) CHARACTER SET utf8mb4 DEFAULT NULL,
  `UniqueNo` varchar(30) CHARACTER SET utf8mb4 DEFAULT NULL,
  `SchemeCode` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `MfType` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL,
  `RTASchemeCode` varchar(15) CHARACTER SET utf8mb4 DEFAULT NULL,
  `AMCSchemeCode` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL,
  `ISIN` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL,
  `AMCCode` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `SchemeType` varchar(30) CHARACTER SET utf8mb4 DEFAULT NULL,
  `SchemePlan` varchar(30) CHARACTER SET utf8mb4 DEFAULT NULL,
  `SchemeName` varchar(300) CHARACTER SET utf8mb4 DEFAULT NULL,
  `PurchaseAllowed` varchar(5) CHARACTER SET utf8mb4 DEFAULT NULL,
  `PurchaseTransactionMode` varchar(5) CHARACTER SET utf8mb4 DEFAULT NULL,
  `MinimumPurchaseAmount` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `AdditionalPurchaseAmount` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `MaximumPurchaseAmount` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `PurchaseAmountMultiplier` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `PurchaseCutoffTime` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `RedemptionAllowed` varchar(5) CHARACTER SET utf8mb4 DEFAULT NULL,
  `RedemptionTransactionMode` varchar(5) CHARACTER SET utf8mb4 DEFAULT NULL,
  `MinimumRedemptionQty` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `RedemptionQtyMultiplier` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `MaximumRedemptionQty` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `RedemptionAmountMinimum` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `RedemptionAmountMaximum` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `RedemptionAmountMultiple` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `RedemptionCutoffTime` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `RTAAgentCode` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `AMCActiveFlag` varchar(5) CHARACTER SET utf8mb4 DEFAULT NULL,
  `DividendReinvestmentFlag` varchar(5) CHARACTER SET utf8mb4 DEFAULT NULL,
  `SIPFlag` varchar(5) CHARACTER SET utf8mb4 DEFAULT NULL,
  `STPFlag` varchar(5) CHARACTER SET utf8mb4 DEFAULT NULL,
  `SWPFlag` varchar(5) CHARACTER SET utf8mb4 DEFAULT NULL,
  `SwitchFlag` varchar(5) CHARACTER SET utf8mb4 DEFAULT NULL,
  `SettlementType` varchar(5) CHARACTER SET utf8mb4 DEFAULT NULL,
  `AMCInd` varchar(10) CHARACTER SET utf8mb4 DEFAULT NULL,
  `FaceValue` varchar(30) CHARACTER SET utf8mb4 DEFAULT NULL,
  `StartDate` varchar(30) CHARACTER SET utf8mb4 DEFAULT NULL,
  `EndDate` varchar(30) CHARACTER SET utf8mb4 DEFAULT NULL,
  `ExitLoadFlag` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL,
  `ExitLoad` varchar(300) CHARACTER SET utf8mb4 DEFAULT NULL,
  `LockInPeriodFlag` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL,
  `LockInPeriod` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL,
  `ChannelPartnerCode` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL,
  `ReOpeningDate` varchar(30) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Source_name` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Created_by` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Created_Date` datetime DEFAULT NULL,
  `Updated_By` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Updated_Date` datetime DEFAULT NULL,
  `Soft_Del` varchar(2) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Soft_Del_by` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL,
  `Soft_Del_Date` datetime DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=22699 DEFAULT CHARSET=utf8mb3;