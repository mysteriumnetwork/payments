/* Mysterium network payment library.
 *
 * Copyright (C) 2020 BlockDev AG
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package bindings

/*
 * Tool for binding smart contract abis/bytecode from github repo to go lang
 * List of smart contracts are json files which are looked up in artifacts attached to github release tag
 * Exact contracts which are needed (and their deployment order) can be
 * looked up in migration script here: https://github.com/mysteriumnetwork/payments-smart-contracts/blob/master/migrations/2_deploy_contracts.js
 */

//go:generate go run bindings/abi/abigen.go --githubrepo=mysteriumnetwork/payments-smart-contracts -githubrelease=v2.3 --contracts=MystToken.json,ChannelImplementation.json,Registry.json,HermesImplementation.json --out=bindings --pkg=bindings
