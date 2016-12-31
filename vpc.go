/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "github.com/aws/aws-sdk-go/service/ec2"

func findVPCs() ([]*ec2.Vpc, error) {
	svc := ec2Session()

	resp, err := svc.DescribeVpcs(nil)
	if err != nil {
		return []*ec2.Vpc{}, err
	}

	vpcs := resp.Vpcs

	// Filter added items
	for i := len(vpcs) - 1; i >= 0; i-- {
		if getEC2TagValue(vpcs[i].Tags, "ernest.service") != "" {
			vpcs = append(vpcs[:i], vpcs[i+1:]...)
		}
	}

	return vpcs, nil
}
