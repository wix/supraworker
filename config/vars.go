package config

var (
	//CFG_PREFIX_COMMUNICATOR defines parameter in the config for Communicators
	CFG_PREFIX_COMMUNICATOR  = "communicator"
	CFG_PREFIX_COMMUNICATORS = "communicators"
	// HTTP Communicator tuning
	// User for allowed response codes definmition.
	CFG_PREFIX_ALLOWED_RESPONSE_CODES = "codes"
	// Defines backoff prefixes
	// More information at
	//   https://github.com/cenkalti/backoff/blob/v4.0.2/exponential.go#L9
	CFG_PREFIX_BACKOFF = "backoff"

	// MaxInterval caps the RetryInterval and not the randomized interval.
	CFG_PREFIX_BACKOFF_MAXINTERVAL = "maxinterval"
	// After MaxElapsedTime the ExponentialBackOff returns Stop.
	// It never stops if MaxElapsedTime == 0.
	CFG_PREFIX_BACKOFF_MAXELAPSEDTIME  = "maxelapsedtime"
	CFG_PREFIX_BACKOFF_INITIALINTERVAL = "initialinterval"

	CFG_COMMUNICATOR_PARAMS_KEY = "params"
	CFG_INTERVAL_PARAMETER      = "interval"
)