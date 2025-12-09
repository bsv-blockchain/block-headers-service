package metrics

const appName = "block-header-service"

const (
	requestMetricBaseName  = "http_request"
	requestCounterName     = requestMetricBaseName + "_total"
	requestDurationSecName = requestMetricBaseName + "_duration_seconds"
)

const domainPrefix = "bsv_"

const (
	latestBlockBaseName      = domainPrefix + "latest_block"
	latestBlockHeightName    = latestBlockBaseName + "_height"
	latestBlockTimestampName = latestBlockBaseName + "_timestamp"
)
