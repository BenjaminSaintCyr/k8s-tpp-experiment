package lttng

/*
#cgo LDFLAGS: -ldl -llttng-ust

#define TRACEPOINT_DEFINE
#include "k8s-tp.h"

void traceStartSpan(uint64_t s_id, uint64_t s_p_id, char* o_name, char* o_ctx) {
	tracepoint(k8s_ust, start_span, s_id, s_p_id, o_name, o_ctx);
}

void traceEndSpan(uint64_t s_id, char* o_ctx) {
	tracepoint(k8s_ust, end_span, s_id, o_ctx);
}

void traceEvent(char* o_name, char* o_ctx) {
	tracepoint(k8s_ust, event, o_name, o_ctx);
}

void traceUnstructEvent(char* o_ctx) {
	tracepoint(k8s_ust, unstruct_event, o_ctx);
}

void traceK8sEvent(char* source_arg, char* type_arg, char* reason_arg, char* message_arg, char* uid_arg, char* name_arg) {
	tracepoint(k8s_ust, k8s_event, source_arg, type_arg, reason_arg, message_arg, uid_arg, name_arg);
}
*/
import "C"

func ReportStartSpan(spanID uint64, parentID uint64, operationName string, context string) {
	C.traceStartSpan(
		C.uint64_t(spanID),
		C.uint64_t(parentID),
		C.CString(operationName),
		C.CString(context),
	)
}

func ReportEndSpan(spanID uint64, context string) {
	C.traceEndSpan(
		C.uint64_t(spanID),
		C.CString(context),
	)
}

func ReportEvent(operationName, context string) {
	C.traceEvent(
		C.CString(operationName),
		C.CString(context),
	)
}

func ReportUnstructEvent(context string) {
	C.traceUnstructEvent(
		C.CString(context),
	)
}

func ReportK8sEvent(source, event_type, reason, message, uid, name string) {
	C.traceK8sEvent(
		C.CString(source),
		C.CString(event_type),
		C.CString(reason),
		C.CString(message),
		C.CString(uid),
		C.CString(name),
	)
}
