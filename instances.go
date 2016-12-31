/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "github.com/aws/aws-sdk-go/service/ec2"

func findInstances() ([]*ec2.Instance, error) {
	svc := ec2Session()

	resp, err := svc.DescribeInstances(nil)
	if err != nil || len(resp.Reservations) < 1 {
		return []*ec2.Instance{}, err
	}

	instances := resp.Reservations[0].Instances

	// Filter added items
	for i := len(instances) - 1; i >= 0; i-- {
		if getEC2TagValue(instances[i].Tags, "ernest.service") != "" {
			instances = append(instances[:i], instances[i+1:]...)
		}
	}

	return instances, nil
}
