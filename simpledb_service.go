package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/simpledb"
)

// SimpleDBService provides persistence layer via AWS SimpleDB
type SimpleDBService struct {
	db     *simpledb.SimpleDB
	domain string
}

// NewSimpleDBService -
func NewSimpleDBService(domain string) Service {
	sess := session.Must(session.NewSession())
	db := simpledb.New(sess)
	return &SimpleDBService{
		db:     db,
		domain: domain,
	}
}

// func (s *SimpleDBService) Create(item Item) error {
// 	titleName := "title"
// 	completedName := "completed"
// 	completedValue := strconv.FormatBool(item.Completed)
// 	createdAtName := "created_at"
// 	createdAtValue := fmt.Sprintf("%d", item.CreatedAt.Unix())
// 	in := &simpledb.PutAttributesInput{
// 		ItemName : &item.ID,
// 		DomainName : &s.domain,
// 		Attributes: []*simpledb.ReplaceableAttribute{
// 			{
// 				Name: &titleName,
// 				Value: &item.Title,
// 			},
// 			{
// 				Name: &completedName,
// 				Value: &completedValue,
// 			},
// 			{
// 				Name: &createdAtName,
// 				Value: &createdAtValue,
// 			},
// 		},
// 	}
// 	_, err := s.db.PutAttributes(in)

// 	return err
// }


func (s *SimpleDBService) List() ([]Item, error) {
	var query = "select * from " + s.domain + " where created_at is not null order by created_at"
	in := &simpledb.SelectInput{
		SelectExpression: &query,
	}

	out, err := s.db.Select(in);

	var list = []Item{}

	if err != nil {
		return nil, err
	}

	for _, db_item := range out.Items {
		item := Item {
			ID : *db_item.Name,
		}

		for _, attr := range db_item.Attributes {
			switch(*attr.Name){
			case "title":
				item.Title = *attr.Value
			case "completed":
				if *attr.Value == "true"{
					item.Completed = true
				}
			}
		}

		list = append(list, item)
	}
	return list, nil
}

func (s *SimpleDBService) Upsert(item Item, isUpdate bool) error {

	titleName := "title"
	completedName := "completed"
	completedValue := strconv.FormatBool(item.Completed)
	in := &simpledb.PutAttributesInput{
		ItemName : &item.ID,
		DomainName : &s.domain,
		Attributes: []*simpledb.ReplaceableAttribute{
			{
				Name: &titleName,
				Value: &item.Title,
				Replace: &isUpdate,
			},
			{
				Name: &completedName,
				Value: &completedValue,
				Replace: &isUpdate,
			},
		},
	}

	// create
	if !isUpdate {
		item.CreatedAt = time.Now()
		createdAtName := "created_at"
		createdAtValue := fmt.Sprint(item.CreatedAt.Unix())
		createdAttr := simpledb.ReplaceableAttribute{
				Name: &createdAtName,
				Value: &createdAtValue,
				Replace: &isUpdate,
		}
		in.Attributes = append(in.Attributes, &createdAttr)
	}

	_, err := s.db.PutAttributes(in)

	return err
}

func (s *SimpleDBService) Delete(itemId string) error {
	in := &simpledb.DeleteAttributesInput {
		ItemName : &itemId,
		DomainName: &s.domain,
	}
	_, err := s.db.DeleteAttributes(in)
	return err
}