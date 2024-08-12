/*
 * NFT-HOOK-SDK-GO: The NFT-HOOK-SDK-GO is an integration and encapsulation package specifically designed for handling operations related to contract deployment for NFT-HOOK.
 * Copyright (C) 2024  nfthook.xyz
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package config

import (
	"fmt"
)

type NetworkConfig struct {
	ChainID         int64
	RPCURL          string
	ExplorerBaseURL string
}

// Networks 包含所有支持的区块链网络配置
var Networks = map[string]NetworkConfig{
	"ethereum": {
		ChainID:         1,
		RPCURL:          "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID",
		ExplorerBaseURL: "https://etherscan.io",
	},
	"ropsten": {
		ChainID:         3,
		RPCURL:          "https://ropsten.infura.io/v3/YOUR_INFURA_PROJECT_ID",
		ExplorerBaseURL: "https://ropsten.etherscan.io",
	},
	"rinkeby": {
		ChainID:         4,
		RPCURL:          "https://rinkeby.infura.io/v3/YOUR_INFURA_PROJECT_ID",
		ExplorerBaseURL: "https://rinkeby.etherscan.io",
	},
	"goerli": {
		ChainID:         5,
		RPCURL:          "https://goerli.infura.io/v3/YOUR_INFURA_PROJECT_ID",
		ExplorerBaseURL: "https://goerli.etherscan.io",
	},
	"kovan": {
		ChainID:         42,
		RPCURL:          "https://kovan.infura.io/v3/YOUR_INFURA_PROJECT_ID",
		ExplorerBaseURL: "https://kovan.etherscan.io",
	},
	"bsc": {
		ChainID:         56,
		RPCURL:          "https://bsc-dataseed.binance.org/",
		ExplorerBaseURL: "https://bscscan.com",
	},
	"polygon": {
		ChainID:         137,
		RPCURL:          "https://rpc-mainnet.maticvigil.com/",
		ExplorerBaseURL: "https://polygonscan.com",
	},
	"fantom": {
		ChainID:         250,
		RPCURL:          "https://rpcapi.fantom.network",
		ExplorerBaseURL: "https://ftmscan.com",
	},
	"arbitrum": {
		ChainID:         42161,
		RPCURL:          "https://arb1.arbitrum.io/rpc",
		ExplorerBaseURL: "https://arbiscan.io",
	},
	"avalanche": {
		ChainID:         43114,
		RPCURL:          "https://api.avax.network/ext/bc/C/rpc",
		ExplorerBaseURL: "https://snowtrace.io",
	},
}

// GetNetworkConfig 返回指定网络的配置
func GetNetworkConfig(network string) (NetworkConfig, error) {
	config, exists := Networks[network]
	if !exists {
		return NetworkConfig{}, fmt.Errorf("network configuration for %s not found", network)
	}
	return config, nil
}
