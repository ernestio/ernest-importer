/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
)

func outputVpcs(vpcs []*ec2.Vpc) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "ID", "Name"})

	for i, vpc := range vpcs {
		name := getEC2TagValue(vpc.Tags, "Name")
		table.Append([]string{strconv.Itoa(i + 1), *vpc.VpcId, name})
	}

	table.Render()
}

func outputNetworks(networks []*ec2.Subnet) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "ID", "Name", "Subnet", "AZ", "VPC"})

	for i, network := range networks {
		name := getEC2TagValue(network.Tags, "Name")
		table.Append([]string{strconv.Itoa(i + 1), *network.SubnetId, name, *network.CidrBlock, *network.AvailabilityZone, *network.VpcId})
	}

	table.Render()
}

func outputInstances(instances []*ec2.Instance) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "ID", "Name", "Image", "Type", "IP", "Subnet", "VPC"})

	for i, instance := range instances {
		name := getEC2TagValue(instance.Tags, "Name")
		table.Append([]string{strconv.Itoa(i + 1), *instance.InstanceId, name, *instance.ImageId, *instance.InstanceType, *instance.PrivateIpAddress, *instance.SubnetId, *instance.VpcId})
	}

	table.Render()
}
