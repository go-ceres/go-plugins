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
	"github.com/go-ceres/go-ceres/server"
	"google.golang.org/grpc"
	"net"
)

type Grpc struct {
	listener           net.Listener
	config             *Config
	server             *grpc.Server
	serverOptions      []grpc.ServerOption
	streamInterceptors []grpc.StreamServerInterceptor
	unaryInterceptors  []grpc.UnaryServerInterceptor
}

func newServer(conf *Config) *Grpc {
	listener, err := net.Listen(conf.Network, conf.Address())
	if err != nil {
		conf.Logger.Panicd()
	}
	return
}

func (s *Grpc) AddUnaryInterceptor(opt ...grpc.UnaryServerInterceptor) {

}

func (s *Grpc) AddStreamInterceptor(opt ...grpc.StreamServerInterceptor) {

}

func (s *Grpc) Info() *server.Info {

	return nil
}

func (s *Grpc) Start() error {
	return nil
}

func (s *Grpc) Stop() error {
	return nil
}

func (s *Grpc) GracefulStop() error {
	return nil
}
