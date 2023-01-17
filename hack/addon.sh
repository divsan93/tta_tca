#!/bin/bash

IMG="${IMG:-tackle2-addon-test-generator:latest}"

cat <<EOF | kubectl apply -f -
kind: Addon
apiVersion: tackle.konveyor.io/v1alpha1
metadata:
  name: testgen
  namespace: konveyor-tackle
spec:
  image: ${IMG}
EOF
