//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/ExpressRouteLinkGet.json
func ExampleExpressRouteLinksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRouteLinksClient().Get(ctx, "rg1", "portName", "linkName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRouteLink = armnetwork.ExpressRouteLink{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/linkName"),
	// 	Name: to.Ptr("linkName"),
	// 	Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
	// 		AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
	// 		ColoLocation: to.Ptr("coloLocationName"),
	// 		ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
	// 		InterfaceName: to.Ptr("Ethernet 0/0"),
	// 		PatchPanelID: to.Ptr("patchPanelId1"),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RackID: to.Ptr("rackId1"),
	// 		RouterName: to.Ptr("router1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4883fa5dbf6f2c9093fac8ce334547e9dfac68fa/specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/ExpressRouteLinkList.json
func ExampleExpressRouteLinksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExpressRouteLinksClient().NewListPager("rg1", "portName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ExpressRouteLinkListResult = armnetwork.ExpressRouteLinkListResult{
		// 	Value: []*armnetwork.ExpressRouteLink{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link1"),
		// 			Name: to.Ptr("link1"),
		// 			Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
		// 				AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
		// 				ColoLocation: to.Ptr("coloLocation1"),
		// 				ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 				InterfaceName: to.Ptr("Ethernet 0/0"),
		// 				PatchPanelID: to.Ptr("patchPanelId1"),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				RackID: to.Ptr("rackId1"),
		// 				RouterName: to.Ptr("router1"),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRoutePorts/portName/links/link2"),
		// 			Name: to.Ptr("link2"),
		// 			Properties: &armnetwork.ExpressRouteLinkPropertiesFormat{
		// 				AdminState: to.Ptr(armnetwork.ExpressRouteLinkAdminStateDisabled),
		// 				ColoLocation: to.Ptr("coloLocation2"),
		// 				ConnectorType: to.Ptr(armnetwork.ExpressRouteLinkConnectorTypeLC),
		// 				InterfaceName: to.Ptr("Ethernet 0/0"),
		// 				PatchPanelID: to.Ptr("patchPanelId2"),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				RackID: to.Ptr("rackId2"),
		// 				RouterName: to.Ptr("router2"),
		// 			},
		// 	}},
		// }
	}
}
