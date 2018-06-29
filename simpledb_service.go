package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/simpledb"
)

var (
	trueValue     = true
	titleName     = "title"
	completedName = "completed"
	createdAtName = "created_at"
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

// Create -
func (s *SimpleDBService) Create(item Item) error {
	item.CreatedAt = time.Now()
	completedValue := strconv.FormatBool(item.Completed)
	createdAtValue := fmt.Sprint(item.CreatedAt.Unix())
	in := &simpledb.PutAttributesInput{
		ItemName:   &item.ID,
		DomainName: &s.domain,
		Attributes: []*simpledb.ReplaceableAttribute{
			{
				Name:  &titleName,
				Value: &item.Title,
			},
			{
				Name:  &completedName,
				Value: &completedValue,
			},
			{
				Name:  &createdAtName,
				Value: &createdAtValue,
			},
		},
	}
	_, err := s.db.PutAttributes(in)
	return err
}

// Delete -
func (s *SimpleDBService) Delete(id string) error {
	in := &simpledb.DeleteAttributesInput{
		DomainName: &s.domain,
		ItemName:   &id,
	}
	_, err := s.db.DeleteAttributes(in)
	return err
}

// Update -
func (s *SimpleDBService) Update(item Item) error {
	completedValue := strconv.FormatBool(item.Completed)
	in := &simpledb.PutAttributesInput{
		ItemName:   &item.ID,
		DomainName: &s.domain,
		Attributes: []*simpledb.ReplaceableAttribute{
			{
				Name:    &titleName,
				Replace: &trueValue,
				Value:   &item.Title,
			},
			{
				Name:    &completedName,
				Replace: &trueValue,
				Value:   &completedValue,
			},
		},
	}
	_, err := s.db.PutAttributes(in)
	return err
}

// List -
func (s *SimpleDBService) List() ([]Item, error) {
	q := "SELECT * FROM `todo_20180629` WHERE created_at != '' ORDER BY created_at"
	in := &simpledb.SelectInput{
		SelectExpression: &q,
	}
	out, err := s.db.Select(in)
	if err != nil {
		return nil, err
	}
	var result []Item
	for _, v := range out.Items {
		item := Item{
			ID: *v.Name,
		}
		for _, attr := range v.Attributes {
			switch *attr.Name {
			case "title":
				item.Title = *attr.Value
			case "completed":
				if *attr.Value == "true" {
					item.Completed = true
				}
			}
		}
		result = append(result, item)
	}
	return result, nil
}
