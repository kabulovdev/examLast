package main

import (
	"net"

	"examLast/custumer_service/config"
	pb "examLast/custumer_service/genproto/custumer_proto"
	"examLast/custumer_service/pkg/db"
	"examLast/custumer_service/pkg/logger"
	"examLast/custumer_service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "")
	defer logger.Cleanup(log)

	log.Info("main:sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("datbase", cfg.PostgresDatabase))
	connDb, err := db.ConnectToDb(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}
	storeService := service.NewCustumService(connDb, log)
	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterCustomServiceServer(s, storeService)
	log.Info("main: server runing",
		logger.String("port", cfg.RPCPort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
