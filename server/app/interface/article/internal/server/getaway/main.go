package getaway

import (
	"context"
	"log"
	"net/http"
	v1 "yuumi/api/interface/article/v1"
	"yuumi/middleware"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Server struct {
	CorsEnable          bool
	CorsWhiteListRegexp string
	Addr                string
	GrpcAddr            string

	server  *http.Server
	mux     *runtime.ServeMux
	context context.Context
	cancel  context.CancelFunc
}

func (s Server) Start() {
	defer s.cancel()
	log.Printf("http server listening at: %s \n", s.Addr)
	s.server.ListenAndServe()
}

func (s Server) Stop() {
	log.Printf("http server shutdown at: %s \n", s.Addr)
	s.server.Shutdown(s.context)
}

func (s *Server) SetMux() {
	s.mux = runtime.NewServeMux(
		runtime.WithMetadata(func(_ context.Context, request *http.Request) metadata.MD {
			md := metadata.Pairs(
				"grpcgateway-path", request.URL.Path,
				"grpcgateway-method", request.Method,
			)
			return md
		}),
	)
}

func (s *Server) SetContext() {
	ctx, cancel := context.WithCancel(context.Background())
	s.context = ctx
	s.cancel = cancel
}

func (s *Server) SetServer() {
	s.server = &http.Server{
		Addr: s.Addr,
		Handler: middleware.CorsMiddleware(s.mux, &middleware.CorsConfig{
			Enabled:         s.CorsEnable,
			WhiteListRegexp: s.CorsWhiteListRegexp,
		}),
	}
}

func (s *Server) RegisterHandlerFromEndpoint() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// register endpoint...
	v1.RegisterArticleInterfaceHandlerFromEndpoint(s.context, s.mux, s.GrpcAddr, opts)
	v1.RegisterCommentInterfaceHandlerFromEndpoint(s.context, s.mux, s.GrpcAddr, opts)
}
