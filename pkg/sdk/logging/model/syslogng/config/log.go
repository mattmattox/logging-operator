// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"github.com/kube-logging/logging-operator/pkg/sdk/logging/model/syslogng/config/render"
	"github.com/siliconbrain/go-seqs/seqs"
)

func logDefStmt(sourceRefs []string, transforms []render.Renderer, destRefs []string) render.Renderer {
	return braceDefStmt("log", "", render.AllOf(
		render.AllFrom(seqs.Map(seqs.FromSlice(sourceRefs), sourceRefStmt)),
		render.AllOf(transforms...),
		render.AllFrom(seqs.Map(seqs.FromSlice(destRefs), destinationRefStmt)),
	))
}

func sourceRefStmt(name string) render.Renderer {
	return parenDefStmt("source", render.Literal(name))
}

func filterRefStmt(name string) render.Renderer {
	return parenDefStmt("filter", render.Literal(name))
}

func destinationRefStmt(name string) render.Renderer {
	return parenDefStmt("destination", render.Literal(name))
}
