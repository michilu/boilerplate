package config

const (
	ApplicationEventEnable = "application.event.enable"

	GcpAppengineHostnameFmt  = "gcp.appengine.hostname.fmt"
	GcpAppengineVersionId    = "gcp.appengine.version.id"
	GcpLoggingId             = "gcp.logging.id"
	GcpLoggingIdAlias        = "gcp.logging.id.alias"
	GcpPubsubPushEndpointFmt = "gcp.pubsub.push.endpoint.fmt"

	GithubToken = "github.token"

	GoogleApplicationCredentials = "google.application.credentials"
	GoogleProjectId              = "google.project.id"
	GoogleServicesId             = "google.services.id"

	InfraNutsdbEventPath            = "infra.nutsdb.event.path"
	InfraNutsdbKeystoreAutoRecovery = "infra.nutsdb.keystore.auto-recovery"
	InfraNutsdbKeystorePath         = "infra.nutsdb.keystore.path"

	ServiceConfigFile                  = "service.config.file"
	ServicePprofDuration               = "service.pprof.duration"
	ServiceProfilePprofAddr            = "service.profile.pprof.addr"
	ServiceProfilePprofDuration        = "service.profile.pprof.duration"
	ServiceProfilePprofEnable          = "service.profile.pprof.enable"
	ServiceProfileProfilerDebugLogging = "service.profile.profiler.debug-logging"
	ServiceProfileProfilerEnable       = "service.profile.profiler.enable"
	ServiceSemaphoreParallel           = "service.semaphore.parallel"
	ServiceSlogConsole                 = "service.slog.console"
	ServiceSlogDebug                   = "service.slog.debug"
	ServiceSlogSentryDsn               = "service.slog.sentry.dsn"
	ServiceSlogSentryServerName        = "service.slog.sentry.server-name"
	ServiceSlogSentryServerNameAlias   = "service.slog.sentry.server-name.alias"
	ServiceSlogVerbose                 = "service.slog.verbose"
	ServiceTraceEnable                 = "service.trace.enable"
	ServiceUpdateChannel               = "service.update.channel"
	ServiceUpdateEnable                = "service.update.enable"
	ServiceUpdateForce                 = "service.update.force"
	ServiceUpdateUrl                   = "service.update.url"
)
