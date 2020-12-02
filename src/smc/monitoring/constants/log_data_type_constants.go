package constants

type DataType int

const (
	IPS_LOGS                   DataType = 0
	FW_LOGS                    DataType = 1
	ALERTS                     DataType = 3
	ALERT_EVENTS               DataType = 4
	IPS_RECORDINGS             DataType = 5
	COUNTERS                   DataType = 8
	AUDIT                      DataType = 9
	SSL_VPN                    DataType = 28
	THIRD_PARTY                DataType = 29
	BLACKLIST_LOG              DataType = 35
	L2FW_LOGS                  DataType = 36
	SSL_VPN_SESSION_MONITORING DataType = 38
)
