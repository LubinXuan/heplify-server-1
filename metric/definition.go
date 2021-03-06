package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// HEP, SIP Metrics
	packetsByType = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "heplify_packets_total",
		Help: "Total packets by HEP type"},
		[]string{"type"})
	packetsBySize = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_packets_size",
		Help: "Packet size by HEP type"},
		[]string{"type"})
	methodResponses = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "heplify_method_response",
		Help: "SIP method and response counter"},
		[]string{"target_name", "direction", "node_id", "response", "method"})
	reasonCause = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "heplify_reason_isup_total",
		Help: "ISUP Q.850 cause from reason header"},
		[]string{"cause", "node_id"})
	srd = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_kpi_srd",
		Help: "SIP Session Request Delay KPI"},
		[]string{"src_target_name", "dst_target_name"})
	rrd = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_kpi_rrd",
		Help: "SIP Registration Request Delay"},
		[]string{"src_target_name", "dst_target_name"})
	logSeverity = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "heplify_severity_total",
		Help: "Log errors and warnings"},
		[]string{"node_id", "severity", "host"})

	// X-RTP-Stat Metrics
	xrtpCS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_xrtp_cs",
		Help: "XRTP call setup time"},
		[]string{"target_name"})
	xrtpJIR = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_xrtp_jir",
		Help: "XRTP received jitter"},
		[]string{"target_name"})
	xrtpJIS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_xrtp_jis",
		Help: "XRTP sent jitter"},
		[]string{"target_name"})
	xrtpPLR = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_xrtp_plr",
		Help: "XRTP received packets lost"},
		[]string{"target_name"})
	xrtpPLS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_xrtp_pls",
		Help: "XRTP sent packets lost"},
		[]string{"target_name"})
	xrtpDLE = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_xrtp_dle",
		Help: "XRTP mean rtt"},
		[]string{"target_name"})
	xrtpMOS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_xrtp_mos",
		Help: "XRTP mos"},
		[]string{"target_name"})

	// RTCP Metrics
	rtcpFractionLost = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtcp_fraction_lost",
		Help: "RTCP fraction lost"},
		[]string{"node_id"})
	rtcpPacketsLost = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtcp_packets_lost",
		Help: "RTCP packets lost"},
		[]string{"node_id"})
	rtcpJitter = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtcp_jitter",
		Help: "RTCP jitter"},
		[]string{"node_id"})
	rtcpDLSR = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtcp_dlsr",
		Help: "RTCP dlsr"},
		[]string{"node_id"})

	// RTCP-XR Metrics
	rtcpxrFractionLost = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtcpxr_fraction_lost",
		Help: "RTCPXR fraction lost"},
		[]string{"node_id"})
	rtcpxrFractionDiscard = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtcpxr_fraction_discard",
		Help: "RTCPXR fraction discard"},
		[]string{"node_id"})
	rtcpxrBurstDensity = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtcpxr_burst_density",
		Help: "RTCPXR burst density"},
		[]string{"node_id"})
	rtcpxrBurstDuration = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtcpxr_burst_duration",
		Help: "RTCPXR burst duration"},
		[]string{"node_id"})
	rtcpxrGapDensity = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtcpxr_gap_density",
		Help: "RTCPXR gap density"},
		[]string{"node_id"})
	rtcpxrGapDuration = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtcpxr_gap_duration",
		Help: "RTCPXR gap duration"},
		[]string{"node_id"})
	rtcpxrRoundTripDelay = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtcpxr_round_trip_delay",
		Help: "RTCPXR round trip delay"},
		[]string{"node_id"})
	rtcpxrEndSystemDelay = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtcpxr_end_system_delay",
		Help: "RTCPXR end system delay"},
		[]string{"node_id"})

	// VQ-RTCP-XR Metrics
	vqrtcpxrNLR = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_vqrtcpxr_nlr",
		Help: "VQ-RTCPXR network packet loss rate"},
		[]string{"node_id"})
	vqrtcpxrJDR = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_vqrtcpxr_jdr",
		Help: "VQ-RTCPXR jitter buffer discard rate"},
		[]string{"node_id"})
	vqrtcpxrIAJ = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_vqrtcpxr_iaj",
		Help: "VQ-RTCPXR interarrival jitter"},
		[]string{"node_id"})
	vqrtcpxrMOSLQ = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_vqrtcpxr_moslq",
		Help: "VQ-RTCPXR MOS listening voice quality"},
		[]string{"node_id"})
	vqrtcpxrMOSCQ = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_vqrtcpxr_moscq",
		Help: "VQ-RTCPXR MOS conversation voice quality"},
		[]string{"node_id"})

	// RTPAgent Metrics
	rtpagentDelta = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtpagent_delta",
		Help: "RTPAgent delta"},
		[]string{"node_id"})
	rtpagentJitter = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtpagent_jitter",
		Help: "RTPAgent jitter"},
		[]string{"node_id"})
	rtpagentMOS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtpagent_mos",
		Help: "RTPAgent mos"},
		[]string{"node_id"})
	rtpagentPacketsLost = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "heplify_rtpagent_packets_lost",
		Help: "RTPAgent packets lost"},
		[]string{"node_id"})

	// Horaclifix Metrics
	horaclifixRtpMOS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "horaclifix_rtp_mos",
		Help: "Incoming RTP MOS"},
		[]string{"sbc_name", "direction", "inc_realm", "out_realm"})
	horaclifixRtpRVAL = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "horaclifix_rtp_rval",
		Help: "Incoming RTP rVal"},
		[]string{"sbc_name", "direction", "inc_realm", "out_realm"})
	horaclifixRtpPackets = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "horaclifix_rtp_packets",
		Help: "Incoming RTP packets"},
		[]string{"sbc_name", "direction", "inc_realm", "out_realm"})
	horaclifixRtpLostPackets = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "horaclifix_rtp_lost_packets",
		Help: "Incoming RTP lostPackets"},
		[]string{"sbc_name", "direction", "inc_realm", "out_realm"})
	horaclifixRtpAvgJitter = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "horaclifix_rtp_avg_jitter",
		Help: "Incoming RTP avgJitter"},
		[]string{"sbc_name", "direction", "inc_realm", "out_realm"})
	horaclifixRtpMaxJitter = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "horaclifix_rtp_max_jitter",
		Help: "Incoming RTP maxJitter"},
		[]string{"sbc_name", "direction", "inc_realm", "out_realm"})
	horaclifixRtcpPackets = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "horaclifix_rtcp_packets",
		Help: "Incoming RTCP packets"},
		[]string{"sbc_name", "direction", "inc_realm", "out_realm"})
	horaclifixRtcpLostPackets = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "horaclifix_rtcp_lost_packets",
		Help: "Incoming RTCP lostPackets"},
		[]string{"sbc_name", "direction", "inc_realm", "out_realm"})
	horaclifixRtcpAvgJitter = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "horaclifix_rtcp_avg_jitter",
		Help: "Incoming RTCP avgJitter"},
		[]string{"sbc_name", "direction", "inc_realm", "out_realm"})
	horaclifixRtcpMaxJitter = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "horaclifix_rtcp_max_jitter",
		Help: "Incoming RTCP maxJitter"},
		[]string{"sbc_name", "direction", "inc_realm", "out_realm"})
	horaclifixRtcpAvgLAT = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "horaclifix_rtcp_avg_lat",
		Help: "Incoming RTCP avgLat"},
		[]string{"sbc_name", "direction", "inc_realm", "out_realm"})
	horaclifixRtcpMaxLAT = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "horaclifix_rtcp_max_lat",
		Help: "Incoming RTCP maxLat"},
		[]string{"sbc_name", "direction", "inc_realm", "out_realm"})

	// Janus Metrics
	janusRtt                 = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_rtt"})
	janusLost                = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_lost"})
	janusLostByRemote        = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_lost_by_remote"})
	janusJitterLocal         = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_jitter_local"})
	janusJitterRemote        = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_jitter_remote"})
	janusInLinkQuality       = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_in_link_quality"})
	janusInMediaLinkQuality  = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_in_media_link_quality"})
	janusOutLinkQuality      = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_out_link_quality"})
	janusOutMediaLinkQuality = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_out_media_link_quality"})
	janusPacketsReceived     = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_packets_received"})
	janusPacketsSent         = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_packets_sent"})
	janusBytesReceived       = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_bytes_received"})
	janusBytesSent           = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_bytes_sent"})
	janusNacksReceived       = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_nacks_received"})
	janusNacksSent           = promauto.NewGauge(prometheus.GaugeOpts{Name: "janus_nacks_sent"})

	// JSON Paths

	rtcpPaths = [][]string{
		[]string{"report_blocks", "[0]", "fraction_lost"},
		[]string{"report_blocks", "[0]", "packets_lost"},
		[]string{"report_blocks", "[0]", "ia_jitter"},
		[]string{"report_blocks", "[0]", "dlsr"},
		[]string{"report_blocks_xr", "fraction_lost"},
		[]string{"report_blocks_xr", "fraction_discard"},
		[]string{"report_blocks_xr", "burst_density"},
		[]string{"report_blocks_xr", "gap_density"},
		[]string{"report_blocks_xr", "burst_duration"},
		[]string{"report_blocks_xr", "gap_duration"},
		[]string{"report_blocks_xr", "round_trip_delay"},
		[]string{"report_blocks_xr", "end_system_delay"},
	}
	rtpPaths = [][]string{
		[]string{"DELTA"},
		[]string{"JITTER"},
		[]string{"MOS"},
		[]string{"PACKET_LOSS"},
	}
	horaclifixPaths = [][]string{
		[]string{"NAME"},
		[]string{"INC_REALM"},
		[]string{"OUT_REALM"},
		[]string{"INC_MOS"},
		[]string{"INC_RVAL"},
		[]string{"INC_RTP_PK"},
		[]string{"INC_RTP_PK_LOSS"},
		[]string{"INC_RTP_AVG_JITTER"},
		[]string{"INC_RTP_MAX_JITTER"},
		[]string{"INC_RTCP_PK"},
		[]string{"INC_RTCP_PK_LOSS"},
		[]string{"INC_RTCP_AVG_JITTER"},
		[]string{"INC_RTCP_MAX_JITTER"},
		[]string{"INC_RTCP_AVG_LAT"},
		[]string{"INC_RTCP_MAX_LAT"},
		[]string{"OUT_MOS"},
		[]string{"OUT_RVAL"},
		[]string{"OUT_RTP_PK"},
		[]string{"OUT_RTP_PK_LOSS"},
		[]string{"OUT_RTP_AVG_JITTER"},
		[]string{"OUT_RTP_MAX_JITTER"},
		[]string{"OUT_RTCP_PK"},
		[]string{"OUT_RTCP_PK_LOSS"},
		[]string{"OUT_RTCP_AVG_JITTER"},
		[]string{"OUT_RTCP_MAX_JITTER"},
		[]string{"OUT_RTCP_AVG_LAT"},
		[]string{"OUT_RTCP_MAX_LAT"},
	}
	janusPaths = [][]string{
		[]string{"rtt"},
		[]string{"lost"},
		[]string{"lost-by-remote"},
		[]string{"jitter-local"},
		[]string{"jitter-remote"},
		[]string{"in-link-quality"},
		[]string{"in-media-link-quality"},
		[]string{"out-link-quality"},
		[]string{"out-media-link-quality"},
		[]string{"packets-received"},
		[]string{"packets-sent"},
		[]string{"bytes-received"},
		[]string{"bytes-sent"},
		[]string{"nacks-received"},
		[]string{"nacks-sent"},
	}
)
