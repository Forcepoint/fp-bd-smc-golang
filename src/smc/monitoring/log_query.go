package monitoring

import "github.cicd.cloud.fpdev.io/BD/fp-smc-golang/src/smc/monitoring/constants"

var DefaultLogFields = []constants.LogField{
	constants.TIMESTAMP,
	constants.ALERTSEVERITY,
	constants.ACTION,
	constants.NODEID,
	constants.SRC,
	constants.SPORT,
	constants.DST,
	constants.DPORT,
	constants.PROTOCOL,
	constants.EVENT,
	constants.INFOMSG,
}

func QueryLogs() {
}
