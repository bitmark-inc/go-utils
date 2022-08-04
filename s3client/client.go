// SPDX-License-Identifier: BSD-2-Clause
// Copyright (c) 2022-2022 Bitmark Inc.
// Use of this source code is governed by an BSD 2 Clause
// license that can be found in the LICENSE file.

package s3client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
)

type Client struct {
	bucket   string
	folder   string
	s3Client *s3.Client
}

// New creates an S3 client
//   folder is optional root folder, if present "<folder>/" is prepended to upload path
func New(region string, bucket string, folder string, access string, secret string) (*Client, error) {

	creds := credentials.NewStaticCredentialsProvider(access, secret, "")

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	c := &Client{
		bucket:   bucket,
		folder:   folder,
		s3Client: s3.NewFromConfig(cfg),
	}
	return c, nil
}

// PathFromTimestamp create a path with a root folder compatible to Athena table with timestamp partitioning
func (c *Client) PathFromTimestamp(prefix string, timestamp time.Time) (string, error) {
	partition := timestamp.Format("2006/01/02/15")    // hourly partitions
	tsPart := timestamp.Format("2006-01-02-15-04-05") // timestamp
	u := uuid.New()
	uText, err := u.MarshalText()
	if err != nil {
		return "", err
	}
	path := fmt.Sprintf("%s/%s-%s-%s", partition, prefix, tsPart, uText)
	if c.folder != "" {
		path = c.folder + "/" + path
	}
	return path, nil
}

// Upload a json record to a path in the bucket
func (c *Client) Upload(path string, data interface{}) error {

	buffer := &bytes.Buffer{}
	enc := json.NewEncoder(buffer)
	if err := enc.Encode(data); err != nil {
		return err
	}

	// not needed ad encoder appears to add a "\n"
	//buffer.WriteString("\n") // ensure proper termination of JSON record for Athena

	uploader := manager.NewUploader(c.s3Client)
	_, err := uploader.Upload(
		context.TODO(),
		&s3.PutObjectInput{
			Bucket: aws.String(c.bucket),
			Key:    aws.String(path),
			Body:   buffer,
		})

	return err
}
