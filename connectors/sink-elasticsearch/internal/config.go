// Copyright 2022 Linkall Inc.
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

package internal

import (
	cdkgo "github.com/linkall-labs/cdk-go"
	"github.com/pkg/errors"
)

type InsertMode string

const (
	Insert InsertMode = "insert"
	Upsert InsertMode = "upsert"
)

type esConfig struct {
	cdkgo.SinkConfig
	Address   string `json:"address" yaml:"address" validate:"required"`
	IndexName string `json:"index_name" yaml:"index_name" validate:"required"`

	Timeout    int        `json:"timeout" yaml:"timeout"`
	PrimaryKey string     `json:"primary_key" yaml:"primary_key"`
	InsertMode InsertMode `json:"insert_mode" yaml:"insert_mode"`

	Secret Secret
}

func Config() cdkgo.SinkConfigAccessor {
	return &esConfig{}
}

func (cfg *esConfig) GetSecret() cdkgo.SecretAccessor {
	return &cfg.Secret
}

type Secret struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}

func (cfg *esConfig) Validate() error {
	if cfg.InsertMode == Upsert && cfg.PrimaryKey == "" {
		return errors.New("insert mode is upsert but primary key is empty")
	}
	return nil
}
