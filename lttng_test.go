package lttng

import "testing"

func BenchmarkReportStartSpan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReportStartSpan(1, 0, "Event", "Source: {deployment-controller }, Type: Normal, Reason: ScalingReplicaSet, Message: Scaled up replica set stress-deployment-6b4477b8fd to 1, UID: 3cca339f-1449-4fe3-b018-ee72c4dd4c26, Name: stress-deployment")
	}
}

func BenchmarkReportEndSpan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReportEndSpan(1, "testing")
	}
}

func BenchmarkReportEvent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReportEvent("Event", "Source: {deployment-controller }, Type: Normal, Reason: ScalingReplicaSet, Message: Scaled up replica set stress-deployment-6b4477b8fd to 1, UID: 3cca339f-1449-4fe3-b018-ee72c4dd4c26, Name: stress-deployment")
	}
}
