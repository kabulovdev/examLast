package repo

import (
	pb "examLast/post_serves/genproto/post"
)

type PostInfoI interface{
	Create(*pb.PostForCreate) (*pb.PostInfo, error)
    GetPost(*pb.Id) (*pb.PostInfo, error)
    GetPosterInfo(*pb.Id) (*pb.PostInfo, error)
    Update(*pb.PostInfo) (*pb.PostInfo, error)
    Delet(*pb.Id) (*pb.EmptyPost, error)
    GetByOwnerId(*pb.Id) (*pb.Posts, error)
}