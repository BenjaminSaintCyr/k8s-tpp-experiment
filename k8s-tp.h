#undef TRACEPOINT_PROVIDER
#define TRACEPOINT_PROVIDER k8s_ust

#undef TRACEPOINT_INCLUDE
#define TRACEPOINT_INCLUDE "./k8s-tp.h"

#if !defined(_TP_H) || defined(TRACEPOINT_HEADER_MULTI_READ)
#define _TP_H

#include <lttng/tracepoint.h>

TRACEPOINT_EVENT(
    k8s_ust,
    start_span,

    /* Input arguments */
    TP_ARGS(
        uint64_t, s_id,
        uint64_t, s_p_id,
        char*, o_name,
        char*, o_ctx
    ),

    /* Output event fields */
    TP_FIELDS(
        ctf_integer(uint64_t, span_id, s_id)
        ctf_integer(uint64_t, parent_span_id, s_p_id)
        ctf_string(op_name, o_name)
        ctf_string(op_ctx, o_ctx)
    )
)

TRACEPOINT_EVENT(
    k8s_ust,
    end_span,

    /* Input arguments */
    TP_ARGS(
        uint64_t, s_id,
        char*, o_ctx
    ),

    /* Output event fields */
    TP_FIELDS(
        ctf_integer(uint64_t, span_id, s_id)
        ctf_string(op_ctx, o_ctx)
    )
)

TRACEPOINT_EVENT(
    k8s_ust,
    event,

    /* Input arguments */
    TP_ARGS(
        char*, o_name,
        char*, o_ctx
    ),

    /* Output event fields */
    TP_FIELDS(
        ctf_string(op_name, o_name)
        ctf_string(op_ctx, o_ctx)
    )
)

TRACEPOINT_EVENT(
    k8s_ust,
    unstruct_event,

    /* Input arguments */
    TP_ARGS(
        char*, o_ctx
    ),

    /* Output event fields */
    TP_FIELDS(
        ctf_string(op_ctx, o_ctx)
    )
)

#endif /* _TP_H */

#include <lttng/tracepoint-event.h>
