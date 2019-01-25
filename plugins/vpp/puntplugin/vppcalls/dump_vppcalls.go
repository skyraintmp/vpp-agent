// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vppcalls

import (
	"github.com/ligato/vpp-agent/plugins/vpp/model/punt"
)

// PuntDetails includes proto-modelled punt object and its socket path
type PuntDetails struct {
	PuntData   *punt.Punt
	SocketPath []byte
}

// DumpPuntRegisteredSockets returns punt to host via registered socket entries
// TODO since the binary API is not available, all data are read from local cache for now
func (h *PuntVppHandler) DumpPuntRegisteredSockets() (punts []*PuntDetails) {
	for _, puntName := range h.mapping.GetMapping().ListNames() {
		_, puntMeta, found := h.mapping.LookupIdx(puntName)
		if !found {
			continue
		}
		punts = append(punts, &PuntDetails{
			PuntData:   puntMeta.Punt,
			SocketPath: puntMeta.SocketPath,
		})
	}

	if len(punts) > 0 {
		h.log.Warnf("Dump punt socket register: all entries were read from local cache")
	}

	return punts
}