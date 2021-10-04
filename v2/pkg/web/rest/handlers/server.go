package handlers

type Server struct {
	options *Options
}

type Options struct {
}

func New(options *Options) *Server {
	return &Server{options: options}
}
