package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/simpledb"
)

var (
	titleName     = "title"
	completedName = "completed"
	createdAtName = "created_at"
	replace       = true
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

func (s *SimpleDBService) Create(item Item) error {
	item.CreatedAt = time.Now()
	compltedValue := strconv.FormatBool(item.Completed)
	createdAtValue := fmt.Sprint(item.CreatedAt.Unix())

	input := &simpledb.PutAttributesInput{
		ItemName:   &item.ID,
		DomainName: &s.domain,
		Attributes: []*simpledb.ReplaceableAttribute{
			{
				Name:  &titleName,
				Value: &item.Title,
			},
			{
				Name:  &completedName,
				Value: &compltedValue,
			},
			{
				Name:  &createdAtName,
				Value: &createdAtValue,
			},
		},
	}

	_, err := s.db.PutAttributes(input)
	return err
}

func (s *SimpleDBService) List() ([]Item, error) {
	var result []Item
	consistentRead := true
	SelectExpression := "select * from `" + s.domain + "` where created_at is not null order by created_at asc"
	input := &simpledb.SelectInput{
		ConsistentRead:   &consistentRead,
		NextToken:        nil,
		SelectExpression: &SelectExpression,
	}
	output, err := s.db.Select(input)

	for _, item := range output.Items {
		var tmpItem Item
		tmpItem.ID = *item.Name

		for _, attr := range item.Attributes {
			switch *attr.Name {
			case "title":
				tmpItem.Title = *attr.Value
			case "completed":
				if *attr.Value == "true" {
					tmpItem.Completed = true
				}
			}
		}
		result = append(result, tmpItem)
	}

	return result, err
}

func (s *SimpleDBService) Update(item Item) error {
	compltedValue := strconv.FormatBool(item.Completed)

	input := &simpledb.PutAttributesInput{
		ItemName:   &item.ID,
		DomainName: &s.domain,
		Attributes: []*simpledb.ReplaceableAttribute{
			{
				Name:    &titleName,
				Value:   &item.Title,
				Replace: &replace,
			},
			{
				Name:    &completedName,
				Value:   &compltedValue,
				Replace: &replace,
			},
		},
	}

	_, err := s.db.PutAttributes(input)
	return err
}

func (s *SimpleDBService) Delete(id string) error {
	input := &simpledb.DeleteAttributesInput{
		ItemName:   &id,
		DomainName: &s.domain,
	}
	_, err := s.db.DeleteAttributes(input)
	return err
}
