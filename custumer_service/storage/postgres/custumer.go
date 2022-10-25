package postgres

import (
	pb "examLast/custumer_service/genproto/custum"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type custumRepo struct {
	db *sqlx.DB
}

func NewCustumRepo(db *sqlx.DB) *custumRepo {
	return &custumRepo{db: db}
}

func (r *custumRepo) DeletCustum(req *pb.GetId) (*pb.Empty, error) {
	_, err := r.db.Query(`UPDATE custumer_base SET deleted_at = NOW() WHERE id =$1`, req.Id)
	if err != nil {
		return &pb.Empty{}, err
	}
	return &pb.Empty{}, nil
}
func (r *custumRepo) Update(req *pb.CustumerInfo) (*pb.CustumerInfo, error) {
	custumResp := pb.CustumerInfo{}

	err := r.db.QueryRow(`
	UPDATE custumer_base
	SET
	first_name = $1,
	last_name = $2,
	email = $3,
	phonenumber = $4
	WHERE id = $5 
	returning id, first_name, last_name, email, phonenumber`,
		req.FirstName, req.LastName, req.Email, req.PhoneNumber, req.Id).
		Scan(&custumResp.Id, &custumResp.FirstName, &custumResp.LastName, &custumResp.Email, &custumResp.PhoneNumber)
	if err != nil {
		return &pb.CustumerInfo{}, err
	}
	fmt.Println(custumResp)
	err = r.db.QueryRow(`UPDATE custumer_bio SET bio=$1 WHERE custumer_id = $2 returning bio`, req.Bio, req.Id).Scan(&custumResp.Bio)
	if err != nil {
		return &pb.CustumerInfo{}, err
	}
	adder := []*pb.CustumAddress{}
	for _, adress := range req.Addresses {
		address := pb.CustumAddress{}
		err = r.db.QueryRow(`UPDATE custumer_address
	SET 
	street=$1,
    home_address=$2 WHERE id = $3 returning id, street, home_address`, adress.Street, adress.HomeAdress, adress.Id).Scan(
			&address.Id,
			&address.Street,
			&address.HomeAdress)

		adder = append(adder, &address)
		fmt.Println(adder)
	}
	if err != nil {
		return &pb.CustumerInfo{}, err
	}
	custumResp.Addresses = adder
	return &custumResp, nil
}
func (r *custumRepo) ListAllCustum(req *pb.Empty) (*pb.CustumerAll, error) {
	result12 := pb.CustumerAll{}
	rows, err := r.db.Query(`
	select 
	id,
	first_name,
	last_name,
	email,
	phonenumber
	from custumer_base where deleted_at is NULL`)
	if err != nil {
		return &pb.CustumerAll{}, err
	}
	resul := []*pb.CustumerInfo{}
	for rows.Next() {
		temp := &pb.CustumerInfo{}
		err = rows.Scan(
			&temp.Id,
			&temp.FirstName,
			&temp.LastName,
			&temp.Email,
			&temp.PhoneNumber)
		if err != nil {
			return &pb.CustumerAll{}, err
		}
		rowlar, err := r.db.Query(`
		select 
		id, 
		street,
		home_address
		from custumer_address where custumer_id=$1
		`, temp.Id)
		if err != nil {
			return &pb.CustumerAll{}, err
		}
		address := []*pb.CustumAddress{}
		for rowlar.Next() {
			med := pb.CustumAddress{}
			err = rowlar.Scan(
				&med.Id,
				&med.Street,
				&med.HomeAdress)
			if err != nil {
				return &pb.CustumerAll{}, err
			}
			address = append(address, &med)
		}
		fmt.Println(temp)
		temp.Addresses = address
		resul = append(resul, temp)
	}
	result12.AllCustum = resul
	return &result12, nil
}
func (r *custumRepo) GetByCustumId(req *pb.GetId) (*pb.CustumerInfo, error) {
	result := pb.CustumerInfo{}
	fmt.Println(req)
	err := r.db.QueryRow(`select 
	id,
	first_name,
	last_name,
	email,
	phonenumber
	from custumer_base where id = $1 and deleted_at is null`, req.Id).Scan(
		&result.Id,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.PhoneNumber,
	)
	if err != nil {
		return &pb.CustumerInfo{}, err
	}
	fmt.Println(result)
	err = r.db.QueryRow(`select 
	bio
	from custumer_bio where custumer_id = $1`, req.Id).Scan(
		&result.Bio)
	if err != nil {
		return &pb.CustumerInfo{}, err
	}
	fmt.Println(result)
	rows, err := r.db.Query(`
	select  
	custumer_id,
    street,
    home_address
	from custumer_address where custumer_id=$1
	`, req.Id)
	addresses := []*pb.CustumAddress{}
	for rows.Next() {
		address := pb.CustumAddress{}
		err = rows.Scan(
			&address.Id,
			&address.Street,
			&address.HomeAdress)
		if err != nil {
			return &pb.CustumerInfo{}, err
		}
		addresses = append(addresses, &address)
	}
	result.Addresses = addresses
	return &result, nil
}

func (r *custumRepo) Create(req *pb.CustumerForCreate) (*pb.CustumerInfo, error) {
	custumResp := pb.CustumerInfo{}
	err := r.db.QueryRow(`
	insert into custumer_base (
		first_name,
		last_name,
		email,
		phonenumber
		) values ($1, $2, $3, $4) returning id, first_name, last_name, email, phonenumber`,
		req.FirstName,
		req.LastName,
		req.Email,
		req.PhoneNumber).Scan(
		&custumResp.Id, &custumResp.FirstName, &custumResp.LastName, &custumResp.Email, &custumResp.PhoneNumber)
	if err != nil {
		return &pb.CustumerInfo{}, err
	}
	fmt.Println(custumResp)
	err = r.db.QueryRow(`
	insert into custumer_bio (
		custumer_id,
		bio
		) values($1, $2) returning bio`,
		custumResp.Id, req.Bio).Scan(&custumResp.Bio)
	if err != nil {
		return &pb.CustumerInfo{}, err
	}
	fmt.Println("bio yes", custumResp)
	addresses := []*pb.CustumAddress{}
	for _, address := range req.Addresses {
		addressResp := pb.CustumAddress{}
		err = r.db.QueryRow(`
		insert into custumer_address ( 
		custumer_id,
		street,
		home_address) values($1,$2,$3)
		returning id, street, home_address`, custumResp.Id, address.Street, address.HomeAdress).Scan(
			&addressResp.Id, &addressResp.Street, &addressResp.HomeAdress)

		addresses = append(addresses, &addressResp)
	}
	fmt.Println(addresses)
	custumResp.Addresses = addresses
	return &custumResp, nil
}
