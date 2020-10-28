// Copyright 2020 Go-Ceres
// Author https://github.com/go-ceres/go-ceres
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

package file

import (
	"github.com/go-ceres/cli/v2"
	"github.com/go-ceres/go-ceres/config"
	"github.com/go-ceres/go-ceres/helper"
	"github.com/go-ceres/go-ceres/plugin"
)

func init() {
	plugin.Register(plugin.NewPlugin(
		plugin.WithName("source"),
		plugin.WithFlags(&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "This is the configuration path of the configuration component",
			EnvVars: []string{"CERES_CONFIG"},
		}),
		plugin.WithAction(func(ctx *cli.Context) error {
			path := ctx.String("config")
			if path == "" {
				return helper.MissingCommand(ctx)
			}
			if err := config.Load(NewSource(path)); err != nil {
				return err
			}
			return nil
		}),
	))
}
