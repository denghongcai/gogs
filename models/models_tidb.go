// +build tidb

// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	_ "github.com/go-xorm/tidb"
	_ "github.com/pingcap/tidb"
)

func init() {
	EnableTidb = true
}
