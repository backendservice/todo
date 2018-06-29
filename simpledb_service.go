package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/simpledb"
)

// SimpleDBService provides persistence layer via AWS SimpleDB
type SimpleDBService struct {
	db     *simpledb.SimpleDB
	domain string
}

// NewSimpleDBService -
func NewSimpleDBService(domain string) *SimpleDBService {
	sess := session.Must(session.NewSession())
	db := simpledb.New(sess)
	return &SimpleDBService{
		db:     db,
		domain: domain,
	}
}
