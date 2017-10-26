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

package server

import (
	"sync/atomic"
)

type xSession struct {
	xsql                   *xSQL
	crud                   *xCrud
	sessionID              uint32
	sendWarnings           bool
	sendXPluginDeprecation bool
}

func (xcc *mysqlXClientConn) createXSession() *xSession {
	return &xSession{
		xsql:                   createXSQL(xcc),
		crud:                   createCrud(xcc),
		sessionID:              atomic.AddUint32(&baseSessionID, 1),
		sendWarnings:           true,
		sendXPluginDeprecation: true,
	}
}

func (xs *xSession) setSendWarnings(flag bool) {
	xs.sendWarnings = flag
}

func (xs *xSession) getSendWarnings() bool {
	return xs.sendWarnings
}

func (xs *xSession) setXPluginDeprecation(flag bool) {
	xs.sendXPluginDeprecation = flag
}

func (xs *xSession) getXPluginDeprecation() bool {
	return xs.sendXPluginDeprecation
}
