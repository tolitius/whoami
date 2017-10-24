#!/bin/bash

: ${TENSOR_FLOW_HOME?"could not find TENSOR_FLOW_HOME, export it to /your-path-to/tenserflow"}

$TENSOR_FLOW_HOME/bazel-bin/tensorflow/examples/image_retraining/retrain $@
