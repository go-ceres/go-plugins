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

package grpc

import (
	"fmt"
	"github.com/go-ceres/go-ceres/config"
	"github.com/go-ceres/go-ceres/logger"
)

type Config struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Network string `json:"network"`
	Logger  logger.Logger
}

func RawConfig(key string) *Config {
	conf := DefaultConfig()
	if err := config.Get(key).Scan(&conf); err != nil {
		panic("配置文件读取失败")
	}
	return &conf
}

func ScanConfig(name string) *Config {
	return RawConfig("ceres.server." + name)
}

func DefaultConfig() Config {
	return Config{
		Host:    "127.0.0.1",
		Port:    9102,
		Network: "tcp4",
		Logger:  logger.FrameLogger.With(logger.String("mod", "")),
	}
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Config) Build() *Grpc {
	return newServer(c)
}
