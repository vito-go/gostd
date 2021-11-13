package myrpc

import "github.com/liushihao/gostd/logic/myrpc/teacher"

func (s *Server) register() {
	var receivers = []interface{}{
		s.server.Register(&teacher.Teach{}),
	}
	for i := 0; i < len(receivers); i++ {
		err := s.server.Register(receivers[i])
		if err != nil {
			panic(err)
		}
	}

}
