GO_BENCH := go test --run= -bench=. -benchtime=2s
NB_SAMPLES := 6
TRACE_DIR := traces

all: tpp
	go build

k8s-tpp.a: k8s-tpp.c k8s-tp.h
	gcc -I. -c k8s-tpp.c -o k8s-tpp.o
	ar -rc k8s-tpp.a k8s-tpp.o

.PHONY: tpp
tpp: k8s-tpp.a

.PHONY: clean
clean:
	rm -rf *.o *.a *.out *.stat $(TRACE_DIR)

.PHONY: benchmark
benchmark: tpp no-recording-benchmark.stat recording-benchmark.stat
	benchstat no-recording-benchmark.stat recording-benchmark.stat

no-recording-benchmark.stat:
	@echo "[INFO] Without recording"
	for i in $$(seq 1 $(NB_SAMPLES)); do \
		$(GO_BENCH) | tee -a $@; \
	done

.PHONY: trace
trace:
	-lttng destroy
	mkdir -p traces
	lttng create benchmark-tracing --output $(TRACE_DIR)
	lttng enable-channel --userspace benchmark-channel --num-subbuf=4 --subbuf-size=32M
	lttng enable-event -u -a -c benchmark-channel
	lttng add-context -u -t vpid -t vtid -t procname
	lttng start

recording-benchmark.stat: trace
	@echo "[INFO] With recording"
	for i in $$(seq 1 $(NB_SAMPLES)); do \
		$(GO_BENCH) | tee -a $@;\
		lttng clear;\
	done
	lttng destroy

.PHONY: profile
profile: tpp
	go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
