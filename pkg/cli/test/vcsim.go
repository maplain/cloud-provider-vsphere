/*
Copyright 2016 The Kubernetes Authors.

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

package test

import (
	"log"

	"github.com/vmware/govmomi/simulator"
)

// NewServiceInstance returns a new vCenter simulator's model and
// server object used to access the simulator's data and control its
// lifecycle.
func NewServiceInstance() (*simulator.Model, *simulator.Server, error) {
	model := simulator.VPX()
	err := model.Create()
	if err != nil {
		log.Fatal(err)
	}
	return model, model.Service.NewServer(), nil
}
