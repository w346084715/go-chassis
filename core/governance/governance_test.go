/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package governance_test

import (
	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/go-chassis/v2/core/governance"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstallProcess(t *testing.T) {
	archaius.Init(archaius.WithMemorySource())
	archaius.Set("servicecomb.customResource.aResource", "test")
	var value string
	governance.InstallProcessor("servicecomb.customResource", func(k, v string) error {
		t.Log("process:" + k)
		value = v
		return nil
	})
	governance.Init()
	assert.Equal(t, "test", value)
}
