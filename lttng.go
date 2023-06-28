package lttng

/*
#cgo LDFLAGS: -ldl -llttng-ust

#define TRACEPOINT_DEFINE
#include "k8s-tp.h"

void foo() { }

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

void traceByteEvent(char* message, unsigned int length) {
	tracepoint(k8s_ust, byte_event, message, length);
}

*/
import "C"
import "unsafe"

func CallCgo(n int) {
	for i := 0; i < n; i++ {
		C.foo()
	}
}

func foo() {}

func CallGo(n int) {
	for i := 0; i < n; i++ {
		foo()
	}
}

func ReportStartSpan(spanID uint64, parentID uint64, operationName string, context string) {
	operationNameC := C.CString(operationName)
	contextC := C.CString(context)
	C.traceStartSpan(
		C.uint64_t(spanID),
		C.uint64_t(parentID),
		operationNameC,
		contextC,
	)
	C.free(unsafe.Pointer(operationNameC))
	C.free(unsafe.Pointer(contextC))
}

func ReportEndSpan(spanID uint64, context string) {
	contextC := C.CString(context)
	C.traceEndSpan(
		C.uint64_t(spanID),
		contextC,
	)
	C.free(unsafe.Pointer(contextC))
}

func ReportEvent(operationName, context string) {
	contextC := C.CString(context)
	C.traceEvent(
		C.CString(operationName),
		contextC,
	)
	C.free(unsafe.Pointer(contextC))
}

func ReportUnstructEvent(context string) {
	contextC := C.CString(context)
	C.traceUnstructEvent(
		contextC,
	)
	C.free(unsafe.Pointer(contextC))
}

func ReportK8sEvent(source, event_type, reason, message, uid, name string) {
	sourceC := C.CString(source)
	event_typeC := C.CString(event_type)
	reasonC := C.CString(reason)
	messageC := C.CString(message)
	uidC := C.CString(uid)
	nameC := C.CString(name)
	C.traceK8sEvent(
		sourceC,
		event_typeC,
		reasonC,
		messageC,
		uidC,
		nameC,
	)
	C.free(unsafe.Pointer(sourceC))
	C.free(unsafe.Pointer(event_typeC))
	C.free(unsafe.Pointer(reasonC))
	C.free(unsafe.Pointer(messageC))
	C.free(unsafe.Pointer(uidC))
	C.free(unsafe.Pointer(nameC))
}

func ReportByteEvent(message string) {
	messageBytes := []byte(message)
	C.traceByteEvent(
		(*C.char)(unsafe.Pointer(&messageBytes[0])),
		C.uint(len(messageBytes)),
	)
}
