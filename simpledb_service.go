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

// Create -
func (s *SimpleDBService) Create(item Item) error {
	titleName := "title"
	completed := "completed"
	completedValue := strconv.FormatBool(item.Completed)
	createdAtName := "created_at"
	createdAtValue := time.Now().Format(time.RFC3339)

	in := &simpledb.PutAttributesInput{
		ItemName:   &item.ID,
		DomainName: &s.domain,
		Attributes: []*simpledb.ReplaceableAttribute{
			{
				Name:  &titleName,
				Value: &item.Title,
			},
			{
				Name:  &completed,
				Value: &completedValue,
			},
			{
				Name:  &createdAtName,
				Value: &createdAtValue,
			},
		},
	}
	s.db.PutAttributes(in)
	return nil
}

// Delete -
func (s *SimpleDBService) Delete(id string) error {
	in := &simpledb.DeleteAttributesInput{
		ItemName:   &id,
		DomainName: &s.domain,
	}
	s.db.DeleteAttributes(in)
	return nil
}

// Update -
func (s *SimpleDBService) Update(item Item) error {
	titleName := "title"
	completed := "completed"
	completedValue := strconv.FormatBool(item.Completed)

	in := &simpledb.GetAttributesInput{
		ItemName:   &item.ID,
		DomainName: &s.domain,
	}
	out, _ := s.db.GetAttributes(in)
	fmt.Println("Out: ===", out)

	inToUpdate := &simpledb.PutAttributesInput{
		ItemName:   &item.ID,
		DomainName: &s.domain,
		Attributes: []*simpledb.ReplaceableAttribute{
			{
				Name:  &titleName,
				Value: &item.Title,
			},
			{
				Name:  &completed,
				Value: &completedValue,
			},
		},
	}
	s.db.PutAttributes(inToUpdate)
	return nil
}

// List -
func (s *SimpleDBService) List() ([]Item, error) {
	q := "select * from `chloe`"
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
