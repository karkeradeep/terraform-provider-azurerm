package snapshots

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SnapshotsClient struct {
	Client *resourcemanager.Client
}

func NewSnapshotsClientWithBaseURI(api environments.Api) (*SnapshotsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "snapshots", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SnapshotsClient: %+v", err)
	}

	return &SnapshotsClient{
		Client: client,
	}, nil
}
