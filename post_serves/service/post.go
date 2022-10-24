package service

import (
	"context"
	"fmt"

	pb "examLast/post_serves/genproto/post"
	l "examLast/post_serves/pkg/logger"
	"examLast/post_serves/storage"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewPostService(db *sqlx.DB, log l.Logger) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *PostService) Create(cxt context.Context, req *pb.PostForCreate) (*pb.PostInfo, error) {
	fmt.Println(req)
	Post, err := s.storage.Post().Create(req)
	if err != nil {
		s.logger.Error("error while creating post", l.Any("error creating post", err))
		return &pb.PostInfo{}, status.Error(codes.Internal, "something went wrong")
	}
	return Post, nil
}
func (s *PostService) GetPost(cxt context.Context, req *pb.Id) (*pb.PostInfo, error) {
	fmt.Println(req)
	Post, err := s.storage.Post().GetPost(req)
	if err != nil {
		s.logger.Error("error while geting post", l.Any("error geting post", err))
		return &pb.PostInfo{}, status.Error(codes.Internal, "something went wrong")
	}
	return Post, nil
}
func (s *PostService) Update(cxt context.Context, req *pb.PostInfo) (*pb.PostInfo, error) {
	fmt.Println(req)
	Post, err := s.storage.Post().Update(req)
	if err != nil {
		s.logger.Error("error while updating post", l.Any("error updating post", err))
		return &pb.PostInfo{}, status.Error(codes.Internal, "something went wrong")
	}
	return Post, nil
}
func (s *PostService) Delet(cxt context.Context, req *pb.Id) (*pb.EmptyPost, error) {
	fmt.Println(req)
	Post, err := s.storage.Post().Delet(req)
	if err != nil {
		s.logger.Error("error while deleting post", l.Any("error deleting post", err))
		return &pb.EmptyPost{}, status.Error(codes.Internal, "something went wrong")
	}
	return Post, nil
}
func (s *PostService) GetByOwnerID(cxt context.Context, req *pb.Id) (*pb.Posts, error) {
	fmt.Println(req)
	Post, err := s.storage.Post().GetByOwnerId(req)
	if err != nil {
		s.logger.Error("error while deleting post", l.Any("error deleting post", err))
		return &pb.Posts{}, status.Error(codes.Internal, "something went wrong")
	}
	return Post, nil
}
