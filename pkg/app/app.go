// Copyright 2018-2019 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package app

import "context"
import "github.com/cs3org/reva/pkg/storage"

// Registry is the interface that application registries implement
// for discovering application providers
type Registry interface {
	FindProvider(ctx context.Context, mimeType string) (*ProviderInfo, error)
	ListProviders(ctx context.Context) ([]*ProviderInfo, error)
}

// ProviderInfo contains the information
// about a Application Provider
type ProviderInfo struct {
	Location string
}

// Provider is the interface that application providers implement
// for providing the iframe location to a iframe UI Provider
type Provider interface {
	GetIFrame(ctx context.Context, resID *storage.ResourceID, token string) (string, error)
}
