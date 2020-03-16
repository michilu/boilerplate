// https://github.com/googleapis/google-cloud-go/blob/57a019f/logging/logging.go#L887-L942
func (l *AppengineLoggingWriter) toLogEntry(e Entry) (*logpb.LogEntry, error) {
	if e.LogName != "" {
		return nil, errors.New("logging: Entry.LogName should be not be set when writing")
	}
	t := e.Timestamp
	if t.IsZero() {
		t = now()
	}
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	if e.Trace == "" && e.HTTPRequest != nil && e.HTTPRequest.Request != nil {
		traceHeader := e.HTTPRequest.Request.Header.Get("X-Cloud-Trace-Context")
		if traceHeader != "" {
			// Set to a relative resource name, as described at
			// https://cloud.google.com/appengine/docs/flexible/go/writing-application-logs.
			traceID, spanID, traceSampled := deconstructXCloudTraceContext(traceHeader)
			if traceID != "" {
				e.Trace = fmt.Sprintf("%s/traces/%s", l.client.parent, traceID)
			}
			if e.SpanID == "" {
				e.SpanID = spanID
			}

			// If we previously hadn't set TraceSampled, let's retrieve it
			// from the HTTP request's header, as per:
			//   https://cloud.google.com/trace/docs/troubleshooting#force-trace
			e.TraceSampled = e.TraceSampled || traceSampled
		}
	}
	ent := &logpb.LogEntry{
		Timestamp:      ts,
		Severity:       logtypepb.LogSeverity(e.Severity),
		InsertId:       e.InsertID,
		HttpRequest:    fromHTTPRequest(e.HTTPRequest),
		Operation:      e.Operation,
		Labels:         e.Labels,
		Trace:          e.Trace,
		SpanId:         e.SpanID,
		Resource:       e.Resource,
		SourceLocation: e.SourceLocation,
		TraceSampled:   e.TraceSampled,
	}
	switch p := e.Payload.(type) {
	case string:
		ent.Payload = &logpb.LogEntry_TextPayload{TextPayload: p}
	default:
		s, err := toProtoStruct(p)
		if err != nil {
			return nil, err
		}
		ent.Payload = &logpb.LogEntry_JsonPayload{JsonPayload: s}
	}
	return ent, nil
}
