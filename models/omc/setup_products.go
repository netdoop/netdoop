package omc

import (
	"github.com/netdoop/netdoop/utils"

	"github.com/heypkg/store/jsontype"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupProducts(db *gorm.DB, schema string) {
	logger := utils.GetLogger()
	pm := &jsontype.Tags{
		"Device.WEB_GUI.Statistics.TotalUpload":                              "TotalUpload",
		"Device.WEB_GUI.Statistics.TotalDownload":                            "TotalDownload",
		"Device.WEB_GUI.Overview.WANStatus.MaxULThroughput":                  "MaxULThroughput",
		"Device.WEB_GUI.Overview.WANStatus.MaxDLThroughput":                  "MaxDLThroughput",
		"Device.WEB_GUI.Network.NR-LTE.Status.NR.SSB_RSRP":                   "SSB_RSRP",
		"Device.WEB_GUI.Network.NR-LTE.Status.NR.SSB_RSRQ":                   "SSB_RSRQ",
		"Device.WEB_GUI.Network.NR-LTE.Status.NR.SSB_SINR":                   "SSB_SINR",
		"Device.WEB_GUI.Network.NR-LTE.Status.NR.SSB_RSSI":                   "SSB_RSSI",
		"Device.WEB_GUI.Network.NR-LTE.Status.NR.NR_TXPower":                 "NR_TXPower",
		"Device.WEB_GUI.Network.NR-LTE.Status.NR.NR_ULMCS":                   "NR_ULMCS",
		"Device.WEB_GUI.Network.NR-LTE.Status.NR.NR_DLMCS":                   "NR_DLMCS",
		"Device.WEB_GUI.Network.NR-LTE.Status.NR.NR_CQI":                     "NR_CQI",
		"Device.WEB_GUI.Network.NR-LTE.Status.NR.NR_PCI":                     "NR_PCI",
		"Device.WEB_GUI.Network.NR-LTE.Status.LTE.RSRP":                      "RSRP",
		"Device.WEB_GUI.Network.NR-LTE.Status.LTE.RSRQ":                      "RSRQ",
		"Device.WEB_GUI.Network.NR-LTE.Status.LTE.SINR":                      "SINR",
		"Device.WEB_GUI.Network.NR-LTE.Status.LTE.RSSI":                      "RSSI",
		"Device.WEB_GUI.Network.NR-LTE.Status.LTE.TXPower":                   "TXPower",
		"Device.WEB_GUI.Network.NR-LTE.Status.LTE.ULMCS":                     "ULMCS",
		"Device.WEB_GUI.Network.NR-LTE.Status.LTE.DLMCS":                     "DLMCS",
		"Device.WEB_GUI.Network.NR-LTE.Status.LTE.CQI":                       "CQI",
		"Device.WEB_GUI.Network.NR-LTE.Status.LTE.PCI":                       "PCI",
		"Device.WEB_GUI.Statistics.ThroughputStatisticsList.1.SentTotal":     "LAN_SentTotal",
		"Device.WEB_GUI.Statistics.ThroughputStatisticsList.2.SentTotal":     "APN1_SentTotal",
		"Device.WEB_GUI.Statistics.ThroughputStatisticsList.3.SentTotal":     "APN2_SentTotal",
		"Device.WEB_GUI.Statistics.ThroughputStatisticsList.4.SentTotal":     "APN3_SentTotal",
		"Device.WEB_GUI.Statistics.ThroughputStatisticsList.5.SentTotal":     "APN3_SentTotal",
		"Device.WEB_GUI.Statistics.ThroughputStatisticsList.1.ReceivedTotal": "LAN_ReceivedTotal",
		"Device.WEB_GUI.Statistics.ThroughputStatisticsList.2.ReceivedTotal": "APN1_ReceivedTotal",
		"Device.WEB_GUI.Statistics.ThroughputStatisticsList.3.ReceivedTotal": "APN2_ReceivedTotal",
		"Device.WEB_GUI.Statistics.ThroughputStatisticsList.4.ReceivedTotal": "APN3_ReceivedTotal",
		"Device.WEB_GUI.Statistics.ThroughputStatisticsList.5.ReceivedTotal": "APN3_ReceivedTotal",
	}

	if _, err := CreateProduct(db, schema, 2, "enb", "LC968", "000000", "LC968", "Normal", "InternetGatewayDevice.", true); err != nil {
		logger.Error("create product", zap.Error(err), zap.String("OUI", "000000"), zap.String("ProductClass", "LC968"))
	}

	if _, err := CreateProduct(db, schema, 3, "enb", "T00000", "T00000", "T00000", "T00000", "Device.", true); err != nil {
		logger.Error("create product", zap.Error(err), zap.String("OUI", "T00000"), zap.String("ProductClass", "T00000"))
	}

	if product, err := CreateProduct(db, schema, 1, "cpe", "LC5200", "D05F64", "LC5200", "5G", "InternetGatewayDevice.", true); err == nil {
		if err := SetProductPerformanceValueDefines(db, product, pm); err != nil {
			logger.Error("set performance value defines of product", zap.Error(err))
		}
	} else {
		logger.Error("create product", zap.Error(err), zap.String("OUI", "D05F64"))
	}

	if product, err := CreateProduct(db, schema, 4, "cpe", "SRW620-a", "88123D", "SRW620-a", "LTE", "Device.", true); err == nil {
		if err := SetProductPerformanceValueDefines(db, product, pm); err != nil {
			logger.Error("set performance value defines of product", zap.Error(err))
		}
	} else {
		logger.Error("create product", zap.Error(err), zap.String("OUI", "88123D"))
	}

	if _, err := CreateProduct(db, schema, 5, "enb", "CTBU5240", "20AC9C", "CTBU5240", "CTBU5240", "Device.", true); err != nil {
		logger.Error("create product", zap.Error(err), zap.String("OUI", "20AC9C"), zap.String("ProductClass", "CTBU5240"))
	}

	if _, err := CreateProduct(db, schema, 6, "enb", "LC908", "000000", "LC908", "Normal", "InternetGatewayDevice.", true); err != nil {
		logger.Error("create product", zap.Error(err), zap.String("OUI", "000000"), zap.String("ProductClass", "LC908"))
	}
}
