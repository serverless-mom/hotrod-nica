// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package customer

import (
	"context"
)

// Customer contains data about a customer.
type Customer struct {
	ID       int64
	Name     string
	Location string
}

// Interface exposed by the Customer service.
type Interface interface {
	Get(ctx context.Context, customerID string) (*Customer, error)
	List(ctx context.Context) ([]Customer, error)
	Put(ctx context.Context, customer *Customer) error
}
