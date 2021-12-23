//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/ScriptExecutions_List.json
func ExampleScriptExecutionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewScriptExecutionsClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<private-cloud-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("ScriptExecution.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/ScriptExecutions_Get.json
func ExampleScriptExecutionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewScriptExecutionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<script-execution-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ScriptExecution.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/ScriptExecutions_CreateOrUpdate.json
func ExampleScriptExecutionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewScriptExecutionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<script-execution-name>",
		armavs.ScriptExecution{
			Properties: &armavs.ScriptExecutionProperties{
				HiddenParameters: []armavs.ScriptExecutionParameterClassification{
					&armavs.ScriptSecureStringExecutionParameter{
						ScriptExecutionParameter: armavs.ScriptExecutionParameter{
							Name: to.StringPtr("<name>"),
							Type: armavs.ScriptExecutionParameterTypeSecureValue.ToPtr(),
						},
						SecureValue: to.StringPtr("<secure-value>"),
					}},
				Parameters: []armavs.ScriptExecutionParameterClassification{
					&armavs.ScriptStringExecutionParameter{
						ScriptExecutionParameter: armavs.ScriptExecutionParameter{
							Name: to.StringPtr("<name>"),
							Type: armavs.ScriptExecutionParameterTypeValue.ToPtr(),
						},
						Value: to.StringPtr("<value>"),
					},
					&armavs.ScriptStringExecutionParameter{
						ScriptExecutionParameter: armavs.ScriptExecutionParameter{
							Name: to.StringPtr("<name>"),
							Type: armavs.ScriptExecutionParameterTypeValue.ToPtr(),
						},
						Value: to.StringPtr("<value>"),
					}},
				Retention:      to.StringPtr("<retention>"),
				ScriptCmdletID: to.StringPtr("<script-cmdlet-id>"),
				Timeout:        to.StringPtr("<timeout>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ScriptExecution.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/ScriptExecutions_Delete.json
func ExampleScriptExecutionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewScriptExecutionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<script-execution-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/ScriptExecutions_GetExecutionLogs.json
func ExampleScriptExecutionsClient_GetExecutionLogs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewScriptExecutionsClient("<subscription-id>", cred, nil)
	res, err := client.GetExecutionLogs(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<script-execution-name>",
		&armavs.ScriptExecutionsGetExecutionLogsOptions{ScriptOutputStreamType: []*armavs.ScriptOutputStreamType{
			armavs.ScriptOutputStreamTypeInformation.ToPtr(),
			armavs.ScriptOutputStreamTypeInformation.ToPtr(),
			armavs.ScriptOutputStreamTypeInformation.ToPtr(),
			armavs.ScriptOutputStreamTypeOutput.ToPtr()},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ScriptExecution.ID: %s\n", *res.ID)
}
