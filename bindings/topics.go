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

import "github.com/ethereum/go-ethereum/common"

// ChannelImplementationExitRequestTopics the topic for exit request events.
var ChannelImplementationExitRequestTopics = [][]common.Hash{
	{
		common.HexToHash("0xe60f0366d8d61555184ea027447889648bae94ebfb1202a39544b6b6803969db"),
	},
}
