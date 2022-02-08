package httpserver

func (s *Server) routerInit() {
	s.routerAddExpress()
	s.routerAddOpenBlog()
}

// ExpressRouterAdd 快递柜相关路由
func (s *Server) routerAddExpress() {
	s.engine.POST("/express-box/get-package", s.express.GetPackage.Handle) // 注释1
}

// ExpressRouterAdd 快递柜相关路由
func (s *Server) routerAddOpenBlog() {
	s.engine.GET("/openblog/api/v1/user/:profile", s.user.GetUserInfo.Handle)
}
