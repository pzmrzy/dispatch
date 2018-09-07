///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package cmd

import (
	"context"

	"github.com/vmware/dispatch/pkg/api/v1"
	"github.com/vmware/dispatch/pkg/client"
	"strings"
	"fmt"
)

// CallUpdateFunction makes the API call to update a function
func CallUpdateFunction(c client.FunctionsClient) ModelAction {
	return func(input interface{}) error {
		function := input.(*v1.Function)

		_, err := c.UpdateFunction(context.TODO(), "", function)
		if err != nil {
			if create && strings.HasPrefix(fmt.Sprint(err), "[Code: 404] ") {
				_, err := c.CreateFunction(context.TODO(), "", function)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}

		return nil
	}
}
