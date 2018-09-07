/*
Copyright 2018 The Knative Authors

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

package controller

import (
	"fmt"
)

func GetK8sServiceFullname(name string, namespace string) string {
	return fmt.Sprintf("%s.%s.svc.cluster.local", name, namespace)
}

func GetServingK8SServiceNameForObj(name string) string {
	return name + "-service"
}

func GetRevisionHeaderName() string {
	return "knative-serving-revision"
}

func GetRevisionHeaderNamespace() string {
	return "knative-serving-namespace"
}
