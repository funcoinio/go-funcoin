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

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// FUNCOIN Foundation Go Bootnodes
	"enode://16dbec786774329f6b861529a01f7c6ccbbc5712cabb3ab380959c6abb03b04e214b37dfb41cc238aefa25c6825a74286758631985367536afbfbd613ce278fd@146.0.32.244:50505",
	"enode://394c83730f3dd0b2cca2f4da4e43e66c6d54f7b570d8af1ae268d55a2796eaf248e4dd383c3b6ec83dbd25bf7ea61f813516cef77b102a264a5ecab1488aa073@43.249.39.234:50505",
	"enode://78d8066d28dac4db56088379e050835f1085a9ada33a8a16b5c60d315da4924e3524c3c6f56f638966f1c34f291e3c14eee7b60c8a454aa7fecdacc4e19a10c2@104.149.133.98:50505",
	"enode://1f280130654b4aadf8beb7b17906699a9572b53cde6a28e8f7359f93d3a8f5607766636e9395495b50ed577ddf5640d111c78f31689d809a0ffebe32158b7ef0@178.172.224.230:50505",
	"enode://e7a8e522cc6a458819bbf971487b7041648a36312dd32b899efc9af5937ef65a17349cef9232eef90d0072d4b3d2267749c65df741df5386fbe2fa72fb4256a1@178.172.224.229:50505",
	"enode://996cf3cf639f7a3353f8b82fca1af6c4102671a5a9481fe605250e2b4bc51704680377acbaa05e59d3c5a2e9b9b1c09e096523219555557505033d6d4f22b18b@5.79.124.129:50505",
	"enode://a47a2ff2b4d4c8238688cf9754c66bf2bf2886dbf0675bc46ec7a53c605c21d58dc09b967402a5070601deec54e90f6375d45ab02f5500d3d93507c431c4fa46@162.210.197.194:50505",
	"enode://74f31a99b0096391a5d06ef2f188827d83eee05291cec667411f1601d79d63a7023d66e07767d9b2bb99406e505ee93bc317f105f32eaba2b957eb10497f71c6@209.58.176.176:50505",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
