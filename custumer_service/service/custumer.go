package service

import (
	"context"
	"fmt"

	pb "examLast/custumer_service/genproto/custum"
	l "examLast/custumer_service/pkg/logger"
	//grpcclient "custumer/service/grpc_client"
	"examLast/custumer_service/storage"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustumerService struct {
	//post    *grpcclient.ServiceManager
	storage storage.IStorage
	logger  l.Logger
}

func NewCustumService(db *sqlx.DB, log l.Logger) *CustumerService {
	return &CustumerService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *CustumerService) Update(cxt context.Context, req *pb.CustumerInfo) (*pb.CustumerInfo, error) {
	fmt.Println(req)
	store, err := s.storage.Custum().Update(req)
	if err != nil {
		s.logger.Error("error while updating castumer", l.Any("error updating custumer", err))
		return &pb.CustumerInfo{}, status.Error(codes.Internal, "something went wrong")
	}
	return store, nil
}

func (s *CustumerService) DeletCustum(cxt context.Context, req *pb.GetId) (*pb.Empty, error) {
	fmt.Println(req)
	store, err := s.storage.Custum().DeletCustum(req)
	if err != nil {
		s.logger.Error("error while deleting castumer", l.Any("error deleting custumer", err))
		return &pb.Empty{}, status.Error(codes.Internal, "something went wrong")
	}
	return store, nil
}
func (s *CustumerService) ListAllCustum(cxt context.Context, req *pb.Empty) (*pb.CustumerAll, error) {
	fmt.Println(req)
	store, err := s.storage.Custum().ListAllCustum(req)
	if err != nil {
		s.logger.Error("error while deleting castumer", l.Any("error deleting custumer", err))
		return &pb.CustumerAll{}, status.Error(codes.Internal, "something went wrong")
	}
	return store, nil
}
func (s *CustumerService) GetByCustumId(ctx context.Context, req *pb.GetId) (*pb.CustumerInfo, error) {
	fmt.Println(req)
	store, err := s.storage.Custum().GetByCustumId(req)
	if err != nil {
		s.logger.Error("error while geting castumer", l.Any("error geting store by owner id", err))
		return &pb.CustumerInfo{}, status.Error(codes.Internal, "something went wrong")
	}
	return store, nil
}

func (s *CustumerService) Create(ctx context.Context, req *pb.CustumerForCreate) (*pb.CustumerInfo, error) {
	fmt.Println(req)
	store, err := s.storage.Custum().Create(req)
	if err != nil {
		s.logger.Error("error while creating custumer", l.Any("error creating custumer", err))
		return &pb.CustumerInfo{}, status.Error(codes.Internal, "something went wrong")
	}
	return store, nil
}
