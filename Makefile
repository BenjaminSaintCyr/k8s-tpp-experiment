GO_BENCH := go test --run= -bench=. -benchtime=2s
NB_SAMPLES := 20

all: tpp
	go build

k8s-tpp.a: k8s-tpp.c k8s-tp.h
	gcc -I. -c k8s-tpp.c -o k8s-tpp.o
	ar -rc k8s-tpp.a k8s-tpp.o

tpp: k8s-tpp.a

clean:
	rm -f *.o *.a *.out *.stat

benchmark: tpp no-recording-benchmark.stat recording-benchmark.stat
	benchstat no-recording-benchmark.stat recording-benchmark.stat

no-recording-benchmark.stat:
	@echo "[INFO] Without recording"
	for i in $$(seq 1 $(NB_SAMPLES)); do \
		$(GO_BENCH) | tee -a $@; \
	done

trace:
	-lttng destroy
	lttng create benchmark-tracing --output /tmp
	lttng enable-channel --userspace benchmark-channel --num-subbuf=4 --subbuf-size=32M
	lttng enable-event -u -a -c benchmark-channel
	lttng add-context -u -t vpid -t vtid -t procname
	lttng start

recording-benchmark.stat: trace
	@echo "[INFO] With recording"
	for i in $$(seq 1 $(NB_SAMPLES)); do \
		$(GO_BENCH) | tee -a $@; \
	done
	lttng destroy

profile: tpp
	go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
