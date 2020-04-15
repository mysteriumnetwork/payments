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

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/go-common/pkg/eth"
)

// DeployLinkedMystToken deploys the myst token with the provided links.
// Links should come in the map, via a format "Name":address.
func DeployLinkedMystToken(opts *bind.TransactOpts, backend bind.ContractBackend, links map[string]common.Address) (common.Address, *types.Transaction, *bind.BoundContract, error) {
	return eth.DeployContractWithLinks(opts, backend, MystTokenABI, MystTokenBin, links)
}
