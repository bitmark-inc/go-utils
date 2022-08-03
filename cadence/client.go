// SPDX-License-Identifier: BSD-2-Clause
// Copyright (c) 2022-2022 Bitmark Inc.
// Use of this source code is governed by an BSD 2 Clause
// license that can be found in the LICENSE file.
package cadence

import (
	"context"

	"go.uber.org/cadence/client"
	"go.uber.org/cadence/workflow"
)

// CadenceWorkerClient manages multiple cadence worker service clients
type CadenceWorkerClient struct {
	service  string
	hostPort string
	domain   string
	clients  map[string]client.Client
}

// NewWorkerClient creates a new client instance
func NewWorkerClient(
	service string, //  e.g.: "cadence-frontend"
	hostPort string, // e.g.: "localhost:7933"
	domain string, //   e.g.: "some-domain"
) *CadenceWorkerClient {
	return &CadenceWorkerClient{
		service:  service,
		hostPort: hostPort,
		domain:   domain,
		clients:  map[string]client.Client{},
	}
}

// AddService register a service client
func (c *CadenceWorkerClient) AddService(clientName string) {

	serviceClient := BuildCadenceServiceClient(c.hostPort, clientName, c.service)

	cadenceWorker := client.NewClient(
		serviceClient,
		c.Domain,
		&client.Options{},
	)

	c.clients[clientName] = cadenceWorker
}

// StartWorkflow triggers a workflow in a specific client
func (c *CadenceWorkerClient) StartWorkflow(
	ctx context.Context,
	clientName string,
	options client.StartWorkflowOptions,
	workflowFunc interface{},
	args ...interface{},
) (*workflow.Execution, error) {

	return c.clients[clientName].StartWorkflow(ctx, options, workflowFunc, args...)
}
