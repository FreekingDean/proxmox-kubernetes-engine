package interceptors

import "google.golang.org/grpc"

type streamHandlerFunc func(m any) error

type recvWrapper struct {
	grpc.ServerStream
	f streamHandlerFunc
}

func (s *recvWrapper) RecvMsg(m any) error {
	if err := s.ServerStream.RecvMsg(m); err != nil {
		return err
	}
	if err := s.f(m); err != nil {
		return err
	}
	return nil
}
