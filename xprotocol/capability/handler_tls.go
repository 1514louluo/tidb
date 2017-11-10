// Copyright 2017 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package capability

import (
	"github.com/pingcap/tipb/go-mysqlx/Connection"
	"github.com/pingcap/tipb/go-mysqlx/Datatypes"
)

// HandlerTLS is read only value handler.
type HandlerTLS struct {
}

// IsSupport implements Handler interface.
func (h *HandlerTLS) IsSupport() bool {
	return false
}

// GetName implements Handler interface.
func (h *HandlerTLS) GetName() string {
	return "tls"
}

// Get implements Handler interface.
func (h *HandlerTLS) Get() *Mysqlx_Connection.Capability {
	return nil
}

// Set implements Handler interface.
func (h *HandlerTLS) Set(any *Mysqlx_Datatypes.Any) bool {
	return false
}