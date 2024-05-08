package main

// Config конфигурация приложения
type Config struct {
	LogLevel          string `long:"log-level" description:"Log level: panic, fatal, warn or warning, info, debug" env:"CL_LOG_LEVEL" required:"true"`
	LogJSON           bool   `long:"log-json" description:"Enable force log format JSON" env:"CL_LOG_JSON"`
	HttpPrivateListen string `long:"http-private-listen" description:"Listening host:port for private http-server" env:"CL_HTTP_PRIVATE_LISTEN" required:"true"`
	HttpPublicListen  string `long:"http-public-listen" description:"Listening host:port for public http-server" env:"CL_HTTP_PUBLIC_LISTEN" required:"true"`
	GrpcListen        string `long:"grpc-listen" description:"Listening host:port for grpc-server" env:"CL_GRPC_LISTEN" required:"true"`
	StoreEndpoint     string `long:"store-endpoint" description:"Store Grpc Endpoint" env:"CL_STORE_ENDPOINT" required:"true"`

	// EnablePprof включение отладки при помощи pprof
	EnablePprof bool `long:"enable-pprof" description:"Enable pprof server" env:"CL_ENABLE_PPROF"`
}
