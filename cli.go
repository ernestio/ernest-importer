/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func selectSingular() int {
	fmt.Println("Please select which item to add:")

	input := read()

	return getInt(input) - 1
}

func selectMultiple(max int) []int {
	fmt.Println("Please select which item(s) to add:")

	input := read()

	return parseMultipleSelection(input)
}

func selectGroups() map[string][]int {
	m := make(map[string][]int)

	fmt.Println("To group items together, please define the name of the instance group, followed by which instances are to be added")
	fmt.Println("i.e. web:1-8 db:1,2,3")
	fmt.Println("Please select which item(s) to add:")

	input := read()

	selections := strings.Split(input, " ")

	for _, s := range selections {
		x := strings.Split(s, ":")
		m[x[0]] = parseMultipleSelection(x[1])
	}

	return m
}

func parseMultipleSelection(input string) []int {
	var is []int

	selections := strings.Split(input, ",")

	for _, s := range selections {

		values := strings.Split(s, "-")
		if len(values) == 2 {
			for i := getInt(values[0]); i <= getInt(values[1]); i++ {
				is = append(is, i-1)
			}

		} else {
			is = append(is, getInt(s)-1)
			// TODO: check if selection > max
		}

	}

	return is
}

func getInt(val string) int {
	v := strings.TrimSpace(val)

	i, err := strconv.Atoi(v)
	if err != nil {
		exit(err)
	}

	return i
}

func read() string {
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		exit(err)
	}

	return line
}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
