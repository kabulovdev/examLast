package repo

import (
	pb "examLast/custumer_service/genproto/custum"
)

type CustumerInfoI interface {
	Create(*pb.CustumerForCreate) (*pb.CustumerInfo, error)
	GetByCustumId(*pb.GetId) (*pb.CustumerInfo, error)
	Update(*pb.CustumerInfo) (*pb.CustumerInfo, error)
	DeletCustum(*pb.GetId) (*pb.Empty, error)
	ListAllCustum(*pb.Empty) (*pb.CustumerAll, error)
}
