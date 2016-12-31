/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func ec2Session() *ec2.EC2 {
	creds := credentials.NewStaticCredentials(AWSID, AWSKEY, "")
	return ec2.New(session.New(), &aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: creds,
	})
}

func getEC2TagValue(tags []*ec2.Tag, key string) string {
	for _, t := range tags {
		if *t.Key == key {
			return *t.Value
		}
	}

	return ""
}

func setEC2Tags(ids []*string, key, value string) error {
	svc := ec2Session()

	req := &ec2.CreateTagsInput{
		Resources: ids,
	}

	req.Tags = append(req.Tags, &ec2.Tag{
		Key:   &key,
		Value: &value,
	})

	_, err := svc.CreateTags(req)

	return err
}
