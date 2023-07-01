package main

import (
	pb "authorservice/proto"
	"authorservice/repository"
	"authorservice/service"
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"sync"
)

const (
	defaultHttpPort   = "8180"
	defaultGrpcPort   = "8181"
	AuthorServiceName = "AuthorService"
	defaultMongoHost  = "mongodb://localhost:27017"
	defaultDatabase   = "mosha"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l *log.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		switch lvl {
		case logging.LevelDebug:
			l.Debugf(msg, fields)
		case logging.LevelInfo:
			l.Infof(msg, fields)
		case logging.LevelWarn:
			l.Warnf(msg, fields)
		case logging.LevelError:
			l.Errorf(msg, fields)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}
func main() {
	log.Printf("Starting %s", AuthorServiceName)
	port := getEnv("COMPONENT_PORT", defaultHttpPort)
	mongoHost := getEnv("MONGO_DB_HOST", defaultMongoHost)
	grpcPort := getEnv("GRPC_PORT", defaultGrpcPort)

	database := repository.NewMongoDatabase(mongoHost, defaultDatabase)
	repo := repository.New(database)
	s := service.New(repo)

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		log.Infof("Starting %s http on %s", AuthorServiceName, port)
		// Create a new HttpRouter.
		router := service.NewHttpRouter(s)
		if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router.MakeHandler()); err != nil {
			log.Fatalf("Unable to start service %q: %s", AuthorServiceName, err)
		}
		wg.Done()
	}()

	go func() {
		logger := log.New(os.Stderr)
		loggerOpts := []logging.Option{
			logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		}
		// Create a new GrpcRouter.
		log.Infof("Starting %s grpc on %s", AuthorServiceName, grpcPort)
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				logging.UnaryServerInterceptor(InterceptorLogger(logger), loggerOpts...),
				// Add logging interceptor to grpc server.
			),
		)
		pb.RegisterAuthorServiceServer(grpcServer, service.NewGrpcRouter(s))
		grpcServer.Serve(lis)
		wg.Done()
	}()

	wg.Wait()
}
