// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main MIX network.
var MainnetBootnodes = []string{
	// Singapore Parity
	"enode://aeb6070deb50efeb41c5e4283a6a3b08ff701fef90e3522161c145f30df2852af3dfc51ba74591f7c9d96b11ca4c3c2b354bf58dd243f2d877f6eecc2373fd1d@139.162.15.124:30313",
	// Singapore Geth
	"enode://cb4b0568607c97a3af4857700152fad9c79908a89319310268c7f49184e3f7e0c88c8d3504696f521c8aafcd00644f9325f76cba9427181b1493ff7ce230df00@139.162.15.124:30303",
	// Freemont Parity
	"enode://e0c926dcdc5c1cf58b2ecba371c577c28c28c91f9b210093178a812389b65e5b53f0e478753b94fceb0b36645b779a915ca57c0c48507fe4d7f786508653656c@74.207.240.177:30313",
	// Freemont Geth
	"enode://37fe8c0f8d667f110316de3109dcc2e9c9247219998e9ac395db39511bef697f9646f73d86a43da491efce35458f53ac3ac0994b85f7086bbf244809d4e7eb7f@74.207.240.177:30303",
	// Newark Parity
	"enode://99fff4ed887d6a6a7b6e03a657c35c06d0eede1909ec289a362bad9d37dd4085886461bbce83aa484ce1327badb3c5958365caa851d71de49dc4530e075b64bc@45.79.128.151:30313",
	// Newark Geth
	"enode://4b4c06445175b77ac07eba03ac624f72e4e86065ee8bbdbca5f882a2a333e2608e6ef2ae5b5779816abd8d07ae6ec08f75888edeab3bc06f2b75871feb14afef@45.79.128.151:30303",
	// London Parity
	"enode://fd80e04c75559cfdd9ed8c08ef2c39c5bc95021f7cbaf31acb601914bc7dac7c34b470b90a05e519bc8a8435a46e1ce51053ae07fac31a83567285c34a79c6bf@139.162.224.203:30313",
	// London Geth
	"enode://d2e678247450c0a7aefdcf03787d4905deda2a6889f02071b241a1309a6bec6ea1c8b1160f69635dee0b4e00cc83656a601b4528cccbd85744b811654ab24b13@139.162.224.203:30303",
	// Frankfurt Parity
	"enode://4742134a153c108855eb16563424887ed3aa5b6b74e4b713c8e93a10c376d954ff3041442716bdf9ee28fab2ea09f04d07e3366f834ea472c19820b7337eb27a@172.104.130.233:30313",
	// Frankfurt Geth
	"enode://9aaeae0129af1bbb2e3590a71cd1ad0c320aa1c03e15c9eb3563cee2d8a7ca43473f43197b6dc0befe5bcef6185196360fa8b4b0d82d8ef11e2ff65553a1efa5@172.104.130.233:30303",
	// Tokyo Parity
	"enode://799d0a8836e17ef7fcc58b3d5ced5bb1fe474b31a09851f938d381f4556fa8954ca308f6a178d22ed56769a8b878ac8f9cc62c889f9cafab45a3bd4f6024bb29@172.104.68.7:30313",
	// Tokyo Geth
	"enode://ab7b1aab2439aafadcb52f8353e60fc1eea55ee5a01b4ddf46ecdeaa2e869c4bf305249757dc74baa78cf05c5d98ffe5c2a008851f08cab6096c78a08dee7c17@172.104.68.7:30303",
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://c1f8b7c2ac4453271fa07d8e9ecf9a2e8285aa0bd0c07df0131f47153306b0736fd3db8924e7a9bf0bed6b1d8d4f87362a71b033dc7c64547728d953e43e59b2@52.64.155.147:30303",
	"enode://f4a9c6ee28586009fb5a96c8af13a58ed6d8315a9eee4772212c1d4d9cebe5a8b8a78ea4434f318726317d04a3f531a1ef0420cf9752605a562cfe858c46e263@213.186.16.82:30303",

	// Ethereum Foundation bootnode
	"enode://a61215641fb8714a373c80edbfa0ea8878243193f57c96eeb44d0bc019ef295abd4e044fd619bfc4c59731a73fb79afe84e9ab6da0c743ceb479cbb6d263fa91@3.11.147.67:30303",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:30303",
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// These DNS names provide bootstrap connectivity for public testnets and the mainnet.
// See https://github.com/ethereum/discv4-dns-lists for more information.
var KnownDNSNetworks = map[common.Hash]string{
	MainnetGenesisHash: dnsPrefix + "all.mainnet.ethdisco.net",
	RopstenGenesisHash: dnsPrefix + "all.ropsten.ethdisco.net",
	RinkebyGenesisHash: dnsPrefix + "all.rinkeby.ethdisco.net",
	GoerliGenesisHash:  dnsPrefix + "all.goerli.ethdisco.net",
}
