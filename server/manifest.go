// This file is automatically generated. Do not modify it manually.

package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "cloudgeex.mattermost.ccr",
  "name": "Channel Moderation",
  "description": "This plugin serves as a starting point for writing a Mattermost plugin.",
  "version": "0.1.0",
  "min_server_version": "5.12.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "webapp": {
    "bundle_path": "webapp/dist/main.js"
  },
  "settings_schema": {
    "header": "",
    "footer": "",
    "settings": [
      {
        "key": "UserFilterKeysEndpoint",
        "display_name": "User Filter Groups API Endpoint",
        "type": "text",
        "help_text": "Endpoint to request for user filter groups.",
        "placeholder": "https://api.service.com/\u003cteam_slug\u003e/filter/user/groups",
        "default": null
      },
      {
        "key": "UserFilterEndpoint",
        "display_name": "User Filter API Endpoint",
        "type": "text",
        "help_text": "Endpoint to request filter users by group.",
        "placeholder": "https://api.service.com/\u003cteam_slug\u003e/filter/user/\u003ckey\u003e",
        "default": null
      },
      {
        "key": "TeamChannelRolesEndpoint",
        "display_name": "Team Channel Roles API Endpoint",
        "type": "text",
        "help_text": "Endpoint to request team channel roles",
        "placeholder": "https://api.service.com/\u003cteam_slug\u003e/channel/roles",
        "default": null
      }
    ]
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}
