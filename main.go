/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import "os"

var (
	// AWSID ...
	AWSID string
	// AWSKEY ...
	AWSKEY string
	// SERVICENAME ...
	SERVICENAME string
)

func main() {
	AWSID = os.Getenv("ID")
	AWSKEY = os.Getenv("KEY")
	SERVICENAME = os.Args[1]

	tagVPCS()
	tagNetworks()
	tagInstances()
}

func tagVPCS() {
	vpcs, err := findVPCs()
	if err != nil {
		panic(err)
	}

	outputVpcs(vpcs)
	s := selectSingular()

	err = setEC2Tags([]*string{vpcs[s].VpcId}, "ernest.service", SERVICENAME)
	if err != nil {
		exit(err)
	}
}

func tagNetworks() {
	nws, err := findNetworks()
	if err != nil {
		panic(err)
	}

	outputNetworks(nws)

	for _, s := range selectMultiple(len(nws)) {
		err = setEC2Tags([]*string{nws[s].SubnetId}, "ernest.service", SERVICENAME)
		if err != nil {
			exit(err)
		}
	}
}

func tagInstances() {
	ins, err := findInstances()
	if err != nil {
		panic(err)
	}

	outputInstances(ins)

	for group, members := range selectGroups() {
		for _, s := range members {
			err = setEC2Tags([]*string{ins[s].InstanceId}, "ernest.service", SERVICENAME)
			if err != nil {
				exit(err)
			}

			err = setEC2Tags([]*string{ins[s].InstanceId}, "ernest.instance_group", group)
			if err != nil {
				exit(err)
			}
		}
	}
}
