// Code generated by Kitex v0.9.1. DO NOT EDIT.

package studentservice

import (
	student_demo "code.byted.org/pku/kitex_demo/kitex_gen/student_demo"
	server "github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler student_demo.StudentService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
