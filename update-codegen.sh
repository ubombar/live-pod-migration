#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

bash ./vendor/k8s.io/code-generator/generate-groups.sh all /home/ubombar/Documents/EDGENET/live-pod-migration/pkg/generated /home/ubombar/Documents/EDGENET/live-pod-migration/pkg/apis live-pod-migration:v1alpha1 --output-base "${GOPATH}/src" --go-header-file "hack/boilerplate.go.txt"

# github.com/ubombar/live-pod-migration/