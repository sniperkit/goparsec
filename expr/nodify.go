/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright (c) 2013 Goparsec AUTHORS. All rights reserved.
// Use of this source code is governed by LICENSE file.

package expr

import (
	"github.com/sniperkit/snk.fork.goparsec"
)

func one2one(ns []parsec.ParsecNode) parsec.ParsecNode {
	if ns == nil || len(ns) == 0 {
		return nil
	}
	return ns[0]
}

func many2many(ns []parsec.ParsecNode) parsec.ParsecNode {
	if ns == nil || len(ns) == 0 {
		return nil
	}
	return ns
}
