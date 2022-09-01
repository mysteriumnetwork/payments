package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mysteriumnetwork/payments/client/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMultiChainAddressKeeper(t *testing.T) {
	defaultReg := common.HexToAddress("0x1")
	defaultMyst := common.HexToAddress("0x2")
	defaultHermes := common.HexToAddress("0x3")
	defaultChImp := common.HexToAddress("0x4")
	defaultHermes2 := common.HexToAddress("0x31")
	fmt.Println(defaultReg.Hex(), defaultMyst.Hex())
	chain1Addrs := SmartContractAddresses{
		Registry:                    defaultReg,
		Myst:                        defaultMyst,
		ActiveHermes:                defaultHermes,
		ActiveChannelImplementation: defaultChImp,
		KnownHermeses:               []common.Address{defaultHermes, defaultHermes2},
	}
	mcak := *NewMultiChainAddressKeeper(map[int64]SmartContractAddresses{
		1: chain1Addrs,
	})
	t.Run("get addresses for chain", func(t *testing.T) {
		addrs, err := mcak.GetAddressesForChain(1)
		assert.NoError(t, err)
		assert.Equal(t, chain1Addrs, addrs)

		_, err = mcak.GetAddressesForChain(2)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnknownChain)
	})

	t.Run("getters", func(t *testing.T) {
		addr, err := mcak.GetRegistryAddress(1)
		assert.NoError(t, err)
		assert.Equal(t, defaultReg, addr)

		_, err = mcak.GetRegistryAddress(2)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnknownChain)

		addr, err = mcak.GetMystAddress(1)
		assert.NoError(t, err)
		assert.Equal(t, defaultMyst, addr)

		_, err = mcak.GetMystAddress(2)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnknownChain)

		addr, err = mcak.GetActiveHermes(1)
		assert.NoError(t, err)
		assert.Equal(t, defaultHermes, addr)

		_, err = mcak.GetActiveHermes(2)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnknownChain)

		addr, err = mcak.GetActiveChannelImplementation(1)
		assert.NoError(t, err)
		assert.Equal(t, defaultChImp, addr)

		_, err = mcak.GetActiveChannelImplementation(2)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnknownChain)

		addrs, err := mcak.GetKnownHermeses(1)
		assert.NoError(t, err)
		assert.Equal(t, []common.Address{defaultHermes, defaultHermes2}, addrs)

		_, err = mcak.GetKnownHermeses(2)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnknownChain)
	})

	t.Run("channel address", func(t *testing.T) {
		channel, err := mcak.GetActiveChannelAddress(1, common.HexToAddress("0x8"))
		assert.NoError(t, err)
		assert.Equal(t, common.HexToAddress("0xed31e71bdaf1e7cfcbaf87206a590f16255be7ce"), channel)

		_, err = mcak.GetActiveChannelAddress(2, common.HexToAddress("0x8"))
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUnknownChain)

		channel, err = mcak.GetArbitraryChannelAddress(common.HexToAddress("0x11"), common.HexToAddress("0x12"), common.HexToAddress("0x13"), common.HexToAddress("0x14"))
		assert.NoError(t, err)
		assert.Equal(t, common.HexToAddress("0x185ae932baf1803c53613bda27a30ecb311bb53a"), channel)
	})

	t.Run("multichain address provider", func(t *testing.T) {
		mcbc := NewMultichainBlockchainClient(map[int64]BC{
			1: NewBlockchain(NewDefaultEthClientGetter(&mocks.EtherClientMock{}), time.Second),
		})
		mcap := NewMultiChainAddressProvider(&mcak, mcbc)
		assert.Equal(t, mcbc, mcap.bc)
		assert.Equal(t, &mcak, mcap.MultiChainAddressKeeper)

		t.Run("cache", func(t *testing.T) {
			chainId := int64(1)
			addr := common.HexToAddress("0x50")
			key := mcap.chCacheKey(chainId, addr)
			assert.Equal(t, fmt.Sprintf("%d|%s", chainId, addr.Hash()), key)

			_, ok := mcap.chCacheGet(chainId, addr)
			assert.False(t, ok)

			chImpl := common.HexToAddress("0x2")
			mcap.chCacheSet(chainId, addr, chImpl)

			got, ok := mcap.chCacheGet(chainId, addr)
			assert.True(t, ok)
			assert.Equal(t, chImpl, got)
		})

		t.Run("channel address", func(t *testing.T) {
			ch, err := mcap.GetHermesChannelAddress(1, common.HexToAddress("0x20"), defaultHermes)
			assert.NoError(t, err)
			assert.Equal(t, common.HexToAddress("0xbe1614cfdd363bb378c86b208debb128c24fc290"), ch)

			hermes2 := common.HexToAddress("0x800")
			chImpl2 := common.HexToAddress("0x250")
			mcap.chCacheSet(1, hermes2, chImpl2)

			ch, err = mcap.GetHermesChannelAddress(1, common.HexToAddress("0x20"), hermes2)
			assert.NoError(t, err)
			assert.Equal(t, common.HexToAddress("0x2c3d47c15d79413eed1a66ae597994794fa3cd19"), ch)
		})
	})
}
