/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "github.com/aws/aws-sdk-go/service/ec2"

func findNetworks() ([]*ec2.Subnet, error) {
	svc := ec2Session()

	resp, err := svc.DescribeSubnets(nil)
	if err != nil {
		return []*ec2.Subnet{}, err
	}

	networks := resp.Subnets

	// Filter added items
	for i := len(networks) - 1; i >= 0; i-- {
		if getEC2TagValue(networks[i].Tags, "ernest.service") != "" {
			networks = append(networks[:i], networks[i+1:]...)
		}
	}

	return networks, nil
}
