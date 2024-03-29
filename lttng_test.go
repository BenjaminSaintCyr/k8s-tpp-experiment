package lttng

import (
	"fmt"
	"testing"
)

func BenchmarkCGO(b *testing.B) {
	CallCgo(b.N)
}

// BenchmarkGo must be called with `-gcflags -l` to avoid inlining.

func BenchmarkGo(b *testing.B) {
	CallGo(b.N)
}

func BenchmarkReportStartSpan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReportStartSpan(1, 0, "Event", fmt.Sprintf("Source: %s, Type: %s, Reason: %s, Message: %s, UID: %s, Name: %s", "{deployment-controller }", "Normal", "ScalingReplicaSet", "Scaled up replica set stress-deployment-6b4477b8fd to 1", "3cca339f-1449-4fe3-b018-ee72c4dd4c26", "stress-deployment"))
	}
}

func BenchmarkReportEndSpan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReportEndSpan(1, "testing")
	}
}

func BenchmarkReportEvent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReportEvent("Event", fmt.Sprintf("Source: %s, Type: %s, Reason: %s, Message: %s, UID: %s, Name: %s", "{deployment-controller }", "Normal", "ScalingReplicaSet", "Scaled up replica set stress-deployment-6b4477b8fd to 1", "3cca339f-1449-4fe3-b018-ee72c4dd4c26", "stress-deployment"))
	}
}

func BenchmarkReportUnstructEvent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReportUnstructEvent(fmt.Sprintf("Source: %s, Type: %s, Reason: %s, Message: %s, UID: %s, Name: %s", "{deployment-controller }", "Normal", "ScalingReplicaSet", "Scaled up replica set stress-deployment-6b4477b8fd to 1", "3cca339f-1449-4fe3-b018-ee72c4dd4c26", "stress-deployment"))
	}
}

func BenchmarkReportK8sEvent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReportK8sEvent("{deployment-controller }", "Normal", "ScalingReplicaSet", "Scaled up replica set stress-deployment-6b4477b8fd to 1", "3cca339f-1449-4fe3-b018-ee72c4dd4c26", "stress-deployment")
	}
}

func BenchmarkByteEvent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReportByteEvent(fmt.Sprintf("Source: %s, Type: %s, Reason: %s, Message: %s, UID: %s, Name: %s", "{deployment-controller }", "Normal", "ScalingReplicaSet", "Scaled up replica set stress-deployment-6b4477b8fd to 1", "3cca339f-1449-4fe3-b018-ee72c4dd4c26", "stress-deployment"))
	}
}
