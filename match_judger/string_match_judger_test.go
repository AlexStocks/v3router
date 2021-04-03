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

package match_judger

import (
	"github.com/dubbogo/v3router/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStringMatchJudger(t *testing.T) {
	assert.True(t, NewStringMatchJudger(&config.StringMatch{
		Exact: "abc",
	}).Judge("abc"))

	assert.False(t, NewStringMatchJudger(&config.StringMatch{
		Exact: "abcd",
	}).Judge("abc"))

	assert.True(t, NewStringMatchJudger(&config.StringMatch{
		Prefix: "abc",
	}).Judge("abcd"))

	assert.False(t, NewStringMatchJudger(&config.StringMatch{
		Exact: "abcd",
	}).Judge("abdc"))

	assert.True(t, NewStringMatchJudger(&config.StringMatch{
		Empty: "true",
	}).Judge(""))

	assert.False(t, NewStringMatchJudger(&config.StringMatch{
		NoEmpty: "true",
	}).Judge(""))
}
