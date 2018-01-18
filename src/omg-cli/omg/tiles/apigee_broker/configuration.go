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
	"encoding/json"
	"omg-cli/config"
	"omg-cli/omg/tiles"
	"omg-cli/ops_manager"
)

type Properties struct {
	ApigeeEdgeConfigurations ApigeeEdgeConfigurationValue `json:".properties.apigee_configurations`
}

type ApigeeEdgeConfigurationValue struct {
	Configs []ApigeeEdgeConfiguration `json:"value"`
}

type ApigeeEdgeConfiguration struct {
	ConfigurationName       string `json:"configuration_name"`
	Org                     string `json:"org"`
	Env                     string `json:"env"`
	ApigeeDashboardUrl      string `json:"apigee_dashboard_url"`
	ApigeeMgmtApiUrl        string `json:"apigee_mgmt_api_url"`
	ApigeeProxyDomain       string `json:"apigee_proxy_domain"`
	ApigeeProxyHostTemplate string `json:"apigee_proxy_host_template"`
	ApigeeProxyNameTemplate string `json:"APIGEE_PROXY_NAME_TEMPLATE"`
}

func (*Tile) Configure(envConfig *config.EnvConfig, cfg *config.Config, om *ops_manager.Sdk) error {
	if err := om.StageProduct(tile.Product); err != nil {
		return err
	}

	network := tiles.NetworkConfig(cfg.ServicesSubnetName, cfg)

	networkBytes, err := json.Marshal(&network)
	if err != nil {
		return err
	}

	properties := Properties{
		ApigeeEdgeConfigurations: ApigeeEdgeConfigurationValue{[]ApigeeEdgeConfiguration{
			{
				ConfigurationName: "Default Configuration",
				ApigeeDashboardUrl: "https://enterprise.apigee.com/platform/#/",
				ApigeeMgmtApiUrl: "https://api.enterprise.apigee.com/v1",
				ApigeeProxyDomain: "apigee.net",
				ApigeeProxyHostTemplate: "${org}-${env}.${domain}",
				ApigeeProxyNameTemplate: "cf-${route}",
			},
		}},
	}

	propertiesBytes, err := json.Marshal(&properties)
	if err != nil {
		return err
	}

	resoruces := "{}"

	return om.ConfigureProduct(tile.Product.Name, string(networkBytes), string(propertiesBytes), resoruces)
}
