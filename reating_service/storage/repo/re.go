package repo

import (
	pb "examLast/reating_service/genproto/reating"
)

type ReatingInfoI interface {
	Create(*pb.ReatingForCreate) (*pb.ReatingInfo, error)
	GetReating(*pb.Id) (*pb.ReatingInfo, error)
	Update(*pb.ReatingInfo) (*pb.ReatingInfo, error)
	Delet(*pb.Id) (*pb.EmptyReating, error)
	GetPostReating(*pb.Id) (*pb.Reatings, error)
}
