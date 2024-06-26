// Copyright 2024 The Witness Contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package archivista

import (
	"context"

	"github.com/in-toto/go-witness/archivista"
	"github.com/in-toto/go-witness/dsse"
)

type aClient struct {
	url    string
	client Clienter
}

// Define Client Interface for Archivista
type Clienter interface {
	Download(ctx context.Context, gitoid string) (dsse.Envelope, error)
	Store(ctx context.Context, env dsse.Envelope) (string, error)
	SearchGitoids(ctx context.Context, vars archivista.SearchGitoidVariables) ([]string, error)
}

func NewArchivistaClient(url string, client *archivista.Client) Clienter {

	if client == nil {
		return nil
	}

	return &aClient{
		url:    url,
		client: client,
	}
}

func (ac *aClient) Download(ctx context.Context, gitoid string) (dsse.Envelope, error) {
	return ac.client.Download(ctx, gitoid)
}

func (ac *aClient) Store(ctx context.Context, env dsse.Envelope) (string, error) {
	return ac.client.Store(ctx, env)
}

func (ac *aClient) SearchGitoids(ctx context.Context, vars archivista.SearchGitoidVariables) ([]string, error) {
	return ac.client.SearchGitoids(ctx, vars)
}
