# Tracepoint Benchmark

This repository contains a benchmarking tool designed to evaluate the performance of tracepoints using the LTTng (Linux Trace Toolkit Next Generation) framework.

## Requirements

- lttng-tools
- lttng-ust

## Getting Started

To get started with the tracepoint benchmark, follow these steps:

1. Clone this repository:

   ```shell
   git clone https://github.com/BenjaminSaintCyr/k8s-tpp-experiment
   ```

2. Install the required dependencies:

   ```shell
   sudo apt-get install lttng-tools liblttng-ust1 liblttng-ust-dev
   ```

3. Run the benchmark

   ```shell
   make benchmark
   ```
