// Copyright 2022 Palantir Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// generated by amalgomate; DO NOT EDIT
package amalgomatedcheck

import (
	"fmt"
	"sort"

	outparamcheck "github.com/palantir/godel-okgo-asset-outparamcheck/internal/generated_src/internal/github.com/palantir/outparamcheck"
)

var programs = map[string]func(){"outparamcheck": func() {
	outparamcheck.AmalgomatedMain()
},
}

func Instance() Amalgomated {
	return &amalgomated{}
}

type Amalgomated interface {
	Run(cmd string)
	Cmds() []string
}

type amalgomated struct{}

func (a *amalgomated) Run(cmd string) {
	if _, ok := programs[cmd]; !ok {
		panic(fmt.Sprintf("Unknown command: \"%v\". Valid values: %v", cmd, a.Cmds()))
	}
	programs[cmd]()
}

func (a *amalgomated) Cmds() []string {
	var cmds []string
	for key := range programs {
		cmds = append(cmds, key)
	}
	sort.Strings(cmds)
	return cmds
}