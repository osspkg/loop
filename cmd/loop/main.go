/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@gmail.com>. All rights reserved.
 *  Use of this source code is governed by a BSD-3-Clause license that can be found in the LICENSE file.
 */

package main

import (
	"go.arwos.org/loopy/internal/pkg/db"
	"go.arwos.org/loopy/internal/server"
	"go.osspkg.com/goppy"
	"go.osspkg.com/goppy/plugins/web"
)

func main() {
	app := goppy.New()
	app.Plugins(
		web.WithHTTP(),
		web.WithHTTPDebug(),
		web.WithWebsocketServer(),
	)
	app.Plugins(
		server.Plugin,
		db.Plugin,
	)
	app.Run()
}
