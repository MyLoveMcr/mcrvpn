package main

import (
	_ "github.com/MyLoveMcr/mcrvpn/hiddify_extension"

	"github.com/hiddify/hiddify-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}
