/*
Copyright (c) 2015 Golang Croatia
All rights reserved.

The MIT License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package users ...
package users

import (
	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/manager"
	"github.com/golanghr/platform/options"
	"github.com/golanghr/platform/service"
	"google.golang.org/grpc"
)

// Service - tbd ...
type Service struct {

	// Options - Is a global service options struct ... tbd ...
	options.Options

	// Servicer - Is a Servicer interface ... tbd ...
	service.Servicer

	// Managerer - Here GRPC server is located. The thing is, we really need just
	// gRPC so we're not going to complicate this struct more than needed
	manager.Managerer

	// Logging -
	logging.Logging
}

// GrpcServer - Will return back actual grpc.Server
// I understand that this looks like a hack but I'd rather have it require to satisfy interface
// than having it require to satisfy nil.
func (s *Service) GrpcServer() *grpc.Server {
	return s.Managerer.Interface().(*manager.GrpcServer).Server
}

// NewService -
func NewService(opts options.Options, logger logging.Logging) (*Service, error) {

	serv, err := service.New(opts)

	if err != nil {
		return nil, err
	}

	grpcServer, err := manager.NewGrpcServer(serv, opts, logger)

	if err != nil {
		return nil, err
	}

	sc := &Service{
		Options:   opts,
		Servicer:  serv,
		Managerer: grpcServer,
		Logging:   logger,
	}

	RegisterUsersServer(sc.GrpcServer(), sc)

	return sc, nil
}
