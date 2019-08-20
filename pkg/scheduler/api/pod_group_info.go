/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"volcano.sh/volcano/pkg/apis/scheduling"
)

// PodGroupPhase is the phase of a pod group at the current time.
type PodGroupPhase string

// These are the valid phase of podGroups.
const (
	//PodGroupVersionV1Alpha1 represents PodGroupVersion of V1Alpha1
	PodGroupVersionV1Alpha1 string = "v1alpha1"

	//PodGroupVersionV1Alpha2 represents PodGroupVersion of V1Alpha2
	PodGroupVersionV1Alpha2 string = "v1alpha2"
)

// PodGroup is a collection of Pod; used for batch workload.
type PodGroup struct {
	scheduling.PodGroup

	//Version represents the version of PodGroup
	Version string
}
