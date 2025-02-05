/**
* @license
* Copyright 2020 Dynatrace LLC
* 
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
* 
*     http://www.apache.org/licenses/LICENSE-2.0
* 
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/


package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dtcookie/dynatrace/api/config/dashboards"
	"github.com/dtcookie/dynatrace/api/config/maintenancewindows"
	"github.com/dtcookie/dynatrace/api/config/managementzones"
	"github.com/dtcookie/dynatrace/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func createSchema(args []string) bool {
	if len(args) == 1 {
		return false
	}
	if strings.TrimSpace(args[1]) != "schema" {
		return false
	}
	var res interface{}
	if strings.TrimSpace(args[2]) == "dashboard" {
		res = new(dashboards.Dashboard)
	} else if strings.TrimSpace(args[2]) == "managementzone" {
		res = new(managementzones.ManagementZone)
	} else if strings.TrimSpace(args[2]) == "maintenancewindow" {
		res = new(maintenancewindows.MaintenanceWindow)
	}
	resource := terraform.ResourceFor(res)

	var err error
	var w *os.File

	os.Remove("resource.go")

	if w, err = os.Create(strings.TrimSpace("resource.go")); err != nil {
		panic(err)
	}
	defer w.Close()

	fmt.Fprintln(w, "package main")
	fmt.Fprintln(w)
	fmt.Fprintln(w, `import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"`)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "func asdf() *schema.Resource {")
	fmt.Fprint(w, "\treturn ")
	terraform.DumpResource(w, resource, "")
	fmt.Fprintln(w, "}")

	return true
}

func download(args []string) bool {
	if len(args) == 1 {
		return false
	}
	if strings.TrimSpace(args[1]) != "download" {
		return false
	}

	if len(args) < 3 {
		fmt.Println("Usage: terraform-provider-dynatrace download <environment-url> <api-token> [<target-folder>")
		os.Exit(0)
	}
	targetFolder := "configuration"
	environmentURL := args[2]
	apiToken := args[3]
	if len(args) > 4 {
		targetFolder = args[4]
	}

	if err := os.RemoveAll(targetFolder); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	if err := importDashboards(targetFolder+"/dashboards", environmentURL, apiToken); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	if err := importManagementZones(targetFolder+"/management_zones", environmentURL, apiToken); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	if err := importCustomServices(targetFolder+"/custom_services", environmentURL, apiToken); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	if err := importAlertingProfiles(targetFolder+"/alerting_profiles", environmentURL, apiToken); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	if err := importRequestAttributes(targetFolder+"/request_attributes", environmentURL, apiToken); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	return true
}

func main() {
	if createSchema(os.Args) {
		return
	}
	if convert(os.Args) {
		return
	}
	if download(os.Args) {
		return
	}
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return Provider()
		},
	})
}
