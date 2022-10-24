package models

import (
	pc "examLast/exam_api/genproto/custumer_proto"
	pp "examLast/exam_api/genproto/post_proto"
	pr "examLast/exam_api/genproto/reating_proto"
)


type CustumerAllInfo struct {
	Custumer pc.CustumerInfo
	Posts []Postc

}
type Postc struct {
	Post pp.PostInfo
	Reatings pr.Reatings
}
type Result struct {
	Custum pc.CustumerInfo
	Post []Posts
}

type PostIfos struct {
	poster pc.CustumerInfo
	posth pp.Posts
}

type Posts struct {
	
	Pos pp.PostInfo
	Reating []*pr.ReatingInfo
}

type Error struct {
	Code        int
	Error       error
	Description string
}

