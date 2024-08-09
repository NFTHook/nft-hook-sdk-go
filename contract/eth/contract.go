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
package eth

import (
	"math/big"
	"nft-hook-sdk-go/util"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractSDK struct {
	client *ethclient.Client
	auth   *bind.TransactOpts
}

func NewContractSDK(rpcURL string, keyfile string, passphrase string, chainId *big.Int) (*ContractSDK, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	key, err := os.ReadFile(keyfile)
	if err != nil {
		return nil, err
	}

	ks := keystore.NewKeyStore(".", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.Import(key, passphrase, passphrase)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, account, chainId)
	if err != nil {
		return nil, err
	}

	return &ContractSDK{
		client: client,
		auth:   auth,
	}, nil
}

func NewContractSDKFromPrivateKey(rpcURL string, privateKey string) (*ContractSDK, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1)) // 更改为你所用的链ID
	if err != nil {
		return nil, err
	}

	return &ContractSDK{
		client: client,
		auth:   auth,
	}, nil
}

func (sdk *ContractSDK) DeployContract(abiPath, binPath string, constructorArgs ...interface{}) (common.Address, *types.Transaction, *bind.BoundContract, error) {
	var abiBytes, binBytes []byte
	var err error

	if strings.HasPrefix(abiPath, "http://") || strings.HasPrefix(abiPath, "https://") {
		abiBytes, err = util.FetchRemoteFile(abiPath)
		if err != nil {
			return common.Address{}, nil, nil, err
		}
	} else {
		abiBytes, err = os.ReadFile(abiPath)
		if err != nil {
			return common.Address{}, nil, nil, err
		}
	}

	if strings.HasPrefix(binPath, "http://") || strings.HasPrefix(binPath, "https://") {
		binBytes, err = util.FetchRemoteFile(binPath)
		if err != nil {
			return common.Address{}, nil, nil, err
		}
	} else {
		binBytes, err = os.ReadFile(binPath)
		if err != nil {
			return common.Address{}, nil, nil, err
		}
	}

	parsedABI, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(sdk.auth, parsedABI, common.FromHex(string(binBytes)), sdk.client, constructorArgs...)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	return address, tx, contract, nil
}
