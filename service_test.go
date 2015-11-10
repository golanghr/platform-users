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
	"testing"
	"time"

	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/options"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	serviceName        = "Users"
	serviceDescription = "Testing user service"
	serviceVersion     = 0.1
	grpcAddr           = ""
	grpcTLS            = false
	grpcTLSCert        = ""
	grpcTLSKey         = ""
	grpcMaxStreams     = uint32(20)
)

func getOptions() (options.Options, error) {
	return options.New("memo", map[string]interface{}{
		"service-name":                serviceName,
		"service-description":         serviceDescription,
		"service-version":             serviceVersion,
		"grpc-addr":                   grpcAddr,
		"grpc-tls":                    grpcTLS,
		"grpc-tls-cert":               grpcTLSCert,
		"grpc-tls-key":                grpcTLSKey,
		"grpc-max-concurrent-streams": grpcMaxStreams,
	})
}

// TestServerInstance -
func TestServerInstance(t *testing.T) {
	Convey("By creating new service we are getting back service and options are matching", t, func() {
		opts, err := getOptions()
		So(err, ShouldBeNil)

		service, err := NewService(opts, logging.New(opts))

		So(service, ShouldHaveSameTypeAs, &Service{})
		So(err, ShouldBeNil)

		So(service.Name(), ShouldEqual, serviceName)
		So(service.Description(), ShouldEqual, serviceDescription)
		So(service.Version(), ShouldEqual, serviceVersion)
	})
}

func TestServerManager(t *testing.T) {
	Convey("", t, func() {
		opts, err := getOptions()
		So(err, ShouldBeNil)

		service, err := NewService(opts, logging.New(opts))
		So(service, ShouldHaveSameTypeAs, &Service{})
		So(err, ShouldBeNil)

		go service.Start()

		defer func() {
			time.Sleep(1 * time.Second)
			So(service.Stop(), ShouldBeNil)
		}()

	})
}
