/*
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package apigee_broker

import (
	"omg-cli/config"
)

var tile = config.Tile{
	config.PivnetMetadata{
		"apigee-edge-for-pcf-service-broker",
		"7131",
		"46028",
		"93ae37caa882706ad7dcacf6f485ca84cb5f8afd4ca3850f68bb911533f38592",
	},
	config.OpsManagerMetadata{
		"apigee-edge-for-pcf-service-broker",
		"3.0.0",
	},
	&config.StemcellMetadata{
		config.PivnetMetadata{"stemcells",
			"36314",
			"67872",
			"6c966883018e34edc8c0c61a48b7aa07582571e39e37f9065ff58eff4f4b4423"},
		"light-bosh-stemcell-3468.21-google-kvm-ubuntu-trusty-go_agent",
	},
}

type Tile struct{}

func (*Tile) Definition(*config.EnvConfig) config.Tile {
	return tile
}

func (*Tile) BuiltIn() bool {
	return false
}
