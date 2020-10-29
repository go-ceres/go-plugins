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
	"errors"
	"github.com/go-ceres/cli/v2"
	"github.com/go-ceres/go-ceres/config"
	"github.com/go-ceres/go-ceres/helper"
	"github.com/go-ceres/go-ceres/plugin"
)

func init() {
	_ = plugin.Register(plugin.NewPlugin(
		plugin.WithName("source"),
		plugin.WithFlags(&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "设置文件路径",
			EnvVars: []string{"CERES_CONFIG"},
		}, &cli.StringFlag{
			Name:    "decode",
			Aliases: []string{"d"},
			Usage:   "设置解码格式",
			EnvVars: []string{"CERES_CONFIG_DECODE"},
		}, &cli.BoolFlag{
			Name:    "watch",
			Aliases: []string{"w"},
			Usage:   "是否监听配置文件变化",
			EnvVars: []string{"CERES_CONFIG_WATCH"},
		}),
		plugin.WithAction(func(ctx *cli.Context) error {
			path := ctx.String("config")
			decode := ctx.String("decode")
			watch := ctx.Bool("watch")
			opts := make([]Option, 0)
			if path == "" {
				return helper.MissingCommand(ctx)
			}
			if decode != "" {
				if config.Unmarshals[decode] == nil {
					opts = append(opts, Unmarshal(decode))
				} else {
					return errors.New("没有配置该解码格式")
				}
			}
			if err := config.Load(NewSource(path, opts...)); err != nil {
				return err
			}
			if watch {
				config.Watch()
			}
			return nil
		}),
	))
}
