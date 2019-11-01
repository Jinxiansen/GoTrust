# GoTrust

After downloading the project, unzip `TrustWalletCore.zip` and run `main_test.go` .


![](images/image.jpg)


## Support the following coins：


```json

[
    {
        "id": "nebulas",
        "name": "Nebulas",
        "symbol": "NAS",
        "decimals": 18,
        "blockchain": "Nebulas",
        "derivationPath": "m/44'/2718'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://explorer.nebulas.io/#/tx/",
        "info": {
            "url": "https://nebulas.io",
            "client": "https://github.com/nebulasio/go-nebulas",
            "clientPublic": "https://mainnet.nebulas.io",
            "clientDocs": "https://wiki.nebulas.io/en/latest/dapp-development/rpc/rpc.html"
        }
    },
    {
        "id": "ethereum",
        "name": "Ethereum",
        "symbol": "ETH",
        "decimals": 18,
        "blockchain": "Ethereum",
        "derivationPath": "m/44'/60'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://etherscan.io/tx/",
        "info": {
            "url": "https://ethereum.org",
            "client": "https://github.com/ethereum/go-ethereum",
            "clientPublic": "https://mainnet.infura.io",
            "clientDocs": "https://github.com/ethereum/wiki/wiki/JSON-RPC"
        }
    },
    {
        "id": "bitcoin",
        "name": "Bitcoin",
        "symbol": "BTC",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/84'/0'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "p2pkhPrefix": 0,
        "p2shPrefix": 5,
        "hrp": "bc",
        "publicKeyHasher": "sha256ripemd",
        "base58Hasher": "sha256d",
        "xpub": "zpub",
        "xprv": "zprv",
        "explorer": "https://blockchair.com/bitcoin/transaction/",
        "info": {
            "url": "https://bitcoin.org",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "bitcoincash",
        "name": "Bitcoin Cash",
        "symbol": "BCH",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/44'/145'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "p2pkhPrefix": 0,
        "p2shPrefix": 5,
        "hrp": "bitcoincash",
        "publicKeyHasher": "sha256ripemd",
        "base58Hasher": "sha256d",
        "xpub": "xpub",
        "xprv": "xprv",
        "explorer": "https://blockchair.com/bitcoin-cash/transaction/",
        "info": {
            "url": "https://bitcoincash.org",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "callisto",
        "name": "Callisto",
        "symbol": "CLO",
        "decimals": 18,
        "blockchain": "Ethereum",
        "derivationPath": "m/44'/820'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://explorer2.callisto.network/tx/",
        "info": {
            "url": "https://callisto.network",
            "client": "https://github.com/EthereumCommonwealth/go-callisto",
            "clientPublic": "https://clo-geth.0xinfra.com",
            "clientDocs": "https://github.com/ethereum/wiki/wiki/JSON-RPC"
        }
    },
    {
        "id": "dash",
        "name": "Dash",
        "symbol": "DASH",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/44'/5'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "p2pkhPrefix": 76,
        "p2shPrefix": 16,
        "publicKeyHasher": "sha256ripemd",
        "base58Hasher": "sha256d",
        "xpub": "xpub",
        "xprv": "xprv",
        "explorer": "https://blockchair.com/dash/transaction/",
        "info": {
            "url": "https://dash.org",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "decred",
        "name": "Decred",
        "symbol": "DCR",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/44'/42'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "staticPrefix": 7,
        "p2pkhPrefix": 63,
        "p2shPrefix": 26,
        "publicKeyHasher": "blake256ripemd",
        "base58Hasher": "blake256d",
        "xpub": "dpub",
        "xprv": "dprv",
        "explorer": "https://dcrdata.decred.org/tx/",
        "info": {
            "url": "https://decred.org",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "digibyte",
        "name": "DigiByte",
        "symbol": "DGB",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/84'/20'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "p2pkhPrefix": 30,
        "p2shPrefix": 63,
        "hrp": "dgb",
        "publicKeyHasher": "sha256ripemd",
        "base58Hasher": "sha256d",
        "xpub": "zpub",
        "xprv": "zprv",
        "explorer": "https://digiexplorer.info/tx/",
        "info": {
            "url": "https://www.digibyte.io",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "doge",
        "name": "Dogecoin",
        "symbol": "DOGE",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/44'/3'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "p2pkhPrefix": 30,
        "p2shPrefix": 22,
        "publicKeyHasher": "sha256ripemd",
        "base58Hasher": "sha256d",
        "xpub": "dgub",
        "xprv": "dgpv",
        "explorer": "https://blockchair.com/dogecoin/transaction/",
        "info": {
            "url": "https://dogecoin.com",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "classic",
        "name": "Ethereum Classic",
        "symbol": "ETC",
        "decimals": 18,
        "blockchain": "Ethereum",
        "derivationPath": "m/44'/61'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://gastracker.io/tx/",
        "info": {
            "url": "https://ethereumclassic.org",
            "client": "https://github.com/ethereumclassic/go-ethereum",
            "clientPublic": "https://www.ethercluster.com/etc",
            "clientDocs": "https://github.com/ethereum/wiki/wiki/JSON-RPC"
        }
    },
    {
        "id": "gochain",
        "name": "GoChain",
        "symbol": "GO",
        "decimals": 18,
        "blockchain": "Ethereum",
        "derivationPath": "m/44'/6060'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://explorer.gochain.io/tx/",
        "info": {
            "url": "https://gochain.io",
            "client": "https://github.com/gochain-io/gochain",
            "clientPublic": "https://rpc.gochain.io",
            "clientDocs": "https://github.com/ethereum/wiki/wiki/JSON-RPC"
        }
    },
    {
        "id": "groestlcoin",
        "name": "Groestlcoin",
        "symbol": "GRS",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/84'/17'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "p2pkhPrefix": 36,
        "p2shPrefix": 5,
        "hrp": "grs",
        "publicKeyHasher": "sha256ripemd",
        "base58Hasher": "groestl512d",
        "xpub": "zpub",
        "xprv": "zprv",
        "explorer": "https://blockchair.com/groestlcoin/transaction/",
        "info": {
            "url": "https://www.groestlcoin.org",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "icon",
        "name": "ICON",
        "symbol": "ICX",
        "decimals": 18,
        "blockchain": "Icon",
        "derivationPath": "m/44'/74'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://tracker.icon.foundation/transaction/",
        "info": {
            "url": "https://icon.foundation",
            "client": "https://github.com/icon-project/icon-rpc-server",
            "clientPublic": "http://ctz.icxstation.com:9000/api/v3",
            "clientDocs": "https://www.icondev.io/docs/icon-json-rpc-v3"
        }
    },
    {
        "id": "iost",
        "name": "IOST",
        "symbol": "IOST",
        "decimals": 8,
        "blockchain": "IOST",
        "derivationPath": "m/44'/291'/0'/0'/0'",
        "curve": "ed25519",
        "publicKeyType": "ed25519",
        "explorer": "https://www.iostabc.com/tx/",
        "info": {
            "url": "https://iost.io",
            "client": "https://github.com/iost-official/go-iost",
            "clientPublic": "",
            "clientDocs": "https://developers.iost.io/docs/en/6-reference/API.html"
        }
    },
    {
        "id": "litecoin",
        "name": "Litecoin",
        "symbol": "LTC",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/84'/2'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "p2pkhPrefix": 48,
        "p2shPrefix": 50,
        "hrp": "ltc",
        "publicKeyHasher": "sha256ripemd",
        "base58Hasher": "sha256d",
        "xpub": "zpub",
        "xprv": "zprv",
        "explorer": "https://blockchair.com/litecoin/transaction/",
        "info": {
            "url": "https://litecoin.org",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "ontology",
        "name": "Ontology",
        "symbol": "ONT",
        "decimals": 0,
        "blockchain": "Ontology",
        "derivationPath": "m/44'/1024'/0'/0/0",
        "curve": "nist256p1",
        "publicKeyType": "nist256p1",
        "explorer": "https://explorer.ont.io/transaction/",
        "info": {
            "url": "https://ont.io",
            "client": "https://github.com/ontio/ontology",
            "clientPublic": "http://dappnode1.ont.io:20336",
            "clientDocs": "https://github.com/ontio/ontology/blob/master/docs/specifications/rpc_api.md"
        }
    },
    {
        "id": "viacoin",
        "name": "Viacoin",
        "symbol": "VIA",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/84'/14'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "p2pkhPrefix": 71,
        "p2shPrefix": 33,
        "hrp": "via",
        "publicKeyHasher": "sha256ripemd",
        "base58Hasher": "sha256d",
        "xpub": "zpub",
        "xprv": "zprv",
        "explorer": "https://explorer.viacoin.org/tx/",
        "info": {
            "url": "https://viacoin.org",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "poa",
        "name": "POA Network",
        "symbol": "POA",
        "decimals": 18,
        "blockchain": "Ethereum",
        "derivationPath": "m/44'/178'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://poaexplorer.com/txid/search/",
        "info": {
            "url": "https://poa.network",
            "client": "https://github.com/poanetwork/parity-ethereum",
            "clientPublic": "https://core.poa.network",
            "clientDocs": "https://github.com/ethereum/wiki/wiki/JSON-RPC"
        }
    },
    {
        "id": "thundertoken",
        "name": "Thunder Token",
        "symbol": "TT",
        "decimals": 18,
        "blockchain": "Ethereum",
        "derivationPath": "m/44'/1001'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://scan.thundercore.com/transactions/",
        "info": {
            "url": "https://thundercore.com",
            "client": "https://github.com/thundercore/pala",
            "clientPublic": "https://mainnet-rpc.thundercore.com",
            "clientDocs": "https://github.com/ethereum/wiki/wiki/JSON-RPC"
        }
    },
    {
        "id": "tomochain",
        "name": "TomoChain",
        "symbol": "TOMO",
        "decimals": 18,
        "blockchain": "Ethereum",
        "derivationPath": "m/44'/889'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://scan.tomochain.com/txs/",
        "info": {
            "url": "https://tomochain.com",
            "client": "https://github.com/tomochain/tomochain",
            "clientPublic": "https://rpc.tomochain.com",
            "clientDocs": "https://github.com/ethereum/wiki/wiki/JSON-RPC"
        }
    },
    {
        "id": "tron",
        "name": "Tron",
        "symbol": "TRX",
        "decimals": 6,
        "blockchain": "Tron",
        "derivationPath": "m/44'/195'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://tronscan.org/#/transaction/",
        "info": {
            "url": "https://tron.network",
            "client": "https://github.com/tronprotocol/java-tron",
            "clientPublic": "https://api.trongrid.io",
            "clientDocs": "https://developers.tron.network/docs/tron-wallet-rpc-api"
        }
    },
    {
        "id": "vechain",
        "name": "VeChain",
        "symbol": "VET",
        "decimals": 18,
        "blockchain": "Vechain",
        "derivationPath": "m/44'/818'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://insight.vecha.in/#/txs/",
        "info": {
            "url": "https://vechain.org",
            "client": "https://github.com/vechain/thor",
            "clientPublic": "",
            "clientDocs": "https://doc.vechainworld.io/docs"
        }
    },
    {
        "id": "wanchain",
        "name": "Wanchain",
        "symbol": "WAN",
        "decimals": 18,
        "blockchain": "Wanchain",
        "derivationPath": "m/44'/5718350'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://www.wanscan.org/tx/",
        "info": {
            "url": "https://wanchain.org",
            "client": "https://github.com/wanchain/go-wanchain",
            "clientPublic": "",
            "clientDocs": "https://github.com/ethereum/wiki/wiki/JSON-RPC"
        }
    },
    {
        "id": "zcoin",
        "name": "Zcoin",
        "symbol": "XZC",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/44'/136'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "p2pkhPrefix": 82,
        "p2shPrefix": 7,
        "publicKeyHasher": "sha256ripemd",
        "base58Hasher": "sha256d",
        "xpub": "xpub",
        "xprv": "xprv",
        "explorer": "https://explorer.zcoin.io/tx/",
        "info": {
            "url": "https://zcoin.io",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "zcash",
        "name": "Zcash",
        "symbol": "ZEC",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/44'/133'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "staticPrefix": 28,
        "p2pkhPrefix": 184,
        "p2shPrefix": 189,
        "publicKeyHasher": "sha256ripemd",
        "base58Hasher": "sha256d",
        "xpub": "xpub",
        "xprv": "xprv",
        "explorer": "https://chain.so/tx/ZEC/",
        "info": {
            "url": "https://z.cash",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "binance",
        "name": "Binance",
        "displayName": "BNB",
        "symbol": "BNB",
        "decimals": 8,
        "blockchain": "Binance",
        "derivationPath": "m/44'/714'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "hrp": "bnb",
        "explorer": "https://explorer.binance.org/tx/",
        "info": {
            "url": "https://binance.org",
            "client": "https://github.com/binance-chain/node-binary",
            "clientPublic": "https://dex.binance.org",
            "clientDocs": "https://docs.binance.org/api-reference/dex-api/paths.html"
        }
    },
    {
        "id": "ripple",
        "name": "XRP",
        "symbol": "XRP",
        "decimals": 6,
        "blockchain": "Ripple",
        "derivationPath": "m/44'/144'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "explorer": "https://bithomp.com/explorer/",
        "info": {
            "url": "https://ripple.com/xrp",
            "client": "https://github.com/ripple/rippled",
            "clientPublic": "https://s2.ripple.com:51234",
            "clientDocs": "https://xrpl.org/rippled-api.html"
        }
    },
    {
        "id": "tezos",
        "name": "Tezos",
        "symbol": "XTZ",
        "decimals": 6,
        "blockchain": "Tezos",
        "derivationPath": "m/44'/1729'/0'/0'",
        "curve": "ed25519",
        "publicKeyType": "ed25519",
        "explorer": "https://tezos.id/",
        "info": {
            "url": "https://tezos.com",
            "client": "https://gitlab.com/tezos/tezos",
            "clientPublic": "https://rpc.tulip.tools/mainnet",
            "clientDocs": "https://tezos.gitlab.io/tezos/api/rpc.html"
        }
    },
    {
        "id": "nimiq",
        "name": "Nimiq",
        "symbol": "NIM",
        "decimals": 5,
        "blockchain": "Nimiq",
        "derivationPath": "m/44'/242'/0'/0'",
        "curve": "ed25519",
        "publicKeyType": "ed25519",
        "explorer": "https://nimiq.watch/#",
        "info": {
            "url": "https://nimiq.com",
            "client": "https://github.com/nimiq/core-rs",
            "clientPublic": "",
            "clientDocs": "https://github.com/nimiq/core-js/wiki/JSON-RPC-API"
        }
    },
    {
        "id": "stellar",
        "name": "Stellar",
        "symbol": "XLM",
        "decimals": 7,
        "blockchain": "Stellar",
        "derivationPath": "m/44'/148'/0'",
        "curve": "ed25519",
        "publicKeyType": "ed25519",
        "explorer": "https://stellarscan.io/transaction/",
        "info": {
            "url": "https://stellar.org",
            "client": "https://github.com/stellar/go",
            "clientPublic": "https://horizon.stellar.org",
            "clientDocs": "https://www.stellar.org/developers/horizon/reference"
        }
    },
    {
        "id": "aion",
        "name": "Aion",
        "symbol": "AION",
        "decimals": 18,
        "blockchain": "Aion",
        "derivationPath": "m/44'/425'/0'/0'/0'",
        "curve": "ed25519",
        "publicKeyType": "ed25519",
        "explorer": "https://mainnet.aion.network/#/transaction/",
        "info": {
            "url": "https://aion.network",
            "client": "https://github.com/aionnetwork/aion",
            "clientPublic": "",
            "clientDocs": "https://github.com/aionnetwork/aion/wiki/JSON-RPC-API-Docs"
        }
    },
    {
        "id": "cosmos",
        "name": "Cosmos",
        "symbol": "ATOM",
        "decimals": 6,
        "blockchain": "Cosmos",
        "derivationPath": "m/44'/118'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "hrp": "cosmos",
        "explorer": "https://www.mintscan.io/txs/",
        "info": {
            "url": "https://cosmos.network",
            "client": "https://github.com/cosmos/cosmos-sdk",
            "clientPublic": "https://stargate.cosmos.network",
            "clientDocs": "https://cosmos.network/rpc"
        }
    },
    {
        "id": "neo",
        "name": "NEO",
        "symbol": "NEO",
        "decimals": 8,
        "blockchain": "NEO",
        "derivationPath": "m/44'/888'/0'/0'/0'",
        "curve": "nist256p1",
        "publicKeyType": "nist256p1",
        "explorer": "https://neoscan.io/transaction/",
        "info": {
            "url": "https://neo.org",
            "client": "https://github.com/neo-project/neo",
            "clientPublic": "",
            "clientDocs": "https://www.stellar.org/developers/horizon/reference"
        }
    },
    {
        "id": "kin",
        "name": "Kin",
        "symbol": "KIN",
        "decimals": 5,
        "blockchain": "Stellar",
        "derivationPath": "m/44'/2017'/0'",
        "curve": "ed25519",
        "publicKeyType": "ed25519",
        "explorer": "https://kinexplorer.com/tx/",
        "info": {
            "url": "https://www.kin.org",
            "client": "https://github.com/kinecosystem/go",
            "clientPublic": "https://horizon.kinfederation.com",
            "clientDocs": "https://www.stellar.org/developers/horizon/reference"
        }
    },
    {
        "id": "nuls",
        "name": "NULS",
        "symbol": "NULS",
        "decimals": 8,
        "blockchain": "NULS",
        "derivationPath": "m/44'/8964'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "explorer": "https://nulscan.io/",
        "info": {
            "url": "https://nuls.io",
            "client": "https://github.com/nuls-io/nuls-v2",
            "clientPublic": "https://public1.nuls.io/",
            "clientDocs": "https://docs.nuls.io/"
        }
    },
    {
        "id": "theta",
        "name": "Theta",
        "symbol": "THETA",
        "decimals": 18,
        "blockchain": "Theta",
        "derivationPath": "m/44'/500'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://explorer.thetatoken.org/txs/",
        "info": {
            "url": "https://www.thetatoken.org",
            "client": "https://github.com/thetatoken/theta-protocol-ledger",
            "clientPublic": "",
            "clientDocs": "https://github.com/thetatoken/theta-mainnet-integration-guide/blob/master/docs/api.md#api-reference"
        }
    },
    {
        "id": "qtum",
        "name": "Qtum",
        "symbol": "QTUM",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/44'/2301'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "p2pkhPrefix": 58,
        "p2shPrefix": 50,
        "hrp": "qc",
        "publicKeyHasher": "sha256ripemd",
        "base58Hasher": "sha256d",
        "xpub": "xpub",
        "xprv": "xprv",
        "explorer": "https://qtum.info/tx/",
        "info": {
            "url": "https://qtum.org",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "bravocoin",
        "name": "BravoCoin",
        "symbol": "BRAVO",
        "decimals": 3,
        "blockchain": "EOS",
        "derivationPath": "m/44'/282'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "explorer": "https://explorer.bravocoin.com/txid/",
        "info": {
            "url": "https://bravocoin.com",
            "client": "https://github.com/bravo-project/bravo",
            "clientPublic": "",
            "clientDocs": "https://github.com/bravo-project/bravo"
        }
    },
    {
        "id": "steem",
        "name": "Steem",
        "symbol": "STEEM",
        "decimals": 3,
        "blockchain": "Steem",
        "derivationPath": "m/44'/135'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "explorer": "https://steemblockexplorer.com/tx/",
        "info": {
            "url": "http://steem.io",
            "client": "https://github.com/steemit/steem",
            "clientPublic": "",
            "clientDocs": "https://github.com/steemit/steem"
        }
    },
    {
        "id": "eos",
        "name": "EOS",
        "symbol": "EOS",
        "decimals": 4,
        "blockchain": "EOS",
        "derivationPath": "m/44'/194'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "explorer": "https://bloks.io/transaction/",
        "info": {
            "url": "http://eos.io",
            "client": "https://github.com/eosio/eos",
            "clientPublic": "",
            "clientDocs": "https://developers.eos.io/eosio-nodeos/reference"
        }
    },
    {
        "id": "nano",
        "name": "Nano",
        "symbol": "NANO",
        "decimals": 30,
        "blockchain": "Nano",
        "derivationPath": "m/44'/165'/0'",
        "curve": "ed25519Blake2bNano",
        "publicKeyType": "ed25519Blake2b",
        "explorer": "https://www.nanode.co/block/",
        "url": "https://nano.org",
        "rpcNodeInfo": "https://github.com/nanocurrency/nano-node",
        "info": {
            "url": "https://nano.org",
            "client": "https://github.com/nanocurrency/nano-node",
            "clientPublic": "",
            "clientDocs": "https://docs.nano.org/commands/rpc-protocol/"
        }
    },
    {
        "id": "iotex",
        "name": "IoTeX",
        "symbol": "IOTX",
        "decimals": 18,
        "blockchain": "IoTeX",
        "derivationPath": "m/44'/304'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://iotexscan.io/action/",
        "info": {
            "url": "https://iotex.io",
            "client": "https://github.com/iotexproject/iotex-core",
            "clientPublic": "",
            "clientDocs": "https://docs.iotex.io/#api"
        }
    },
    {
        "id": "zilliqa",
        "name": "Zilliqa",
        "symbol": "ZIL",
        "decimals": 12,
        "blockchain": "Zilliqa",
        "derivationPath": "m/44'/313'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "hrp": "zil",
        "explorer": "https://viewblock.io/zilliqa/tx/",
        "info": {
            "url": "https://zilliqa.com",
            "client": "https://github.com/Zilliqa/Zilliqa",
            "clientPublic": "https://api.zilliqa.com",
            "clientDocs": "https://apidocs.zilliqa.com"
        }
    },
    {
        "id": "semux",
        "name": "Semux",
        "symbol": "SEM",
        "decimals": 9,
        "blockchain": "Semux",
        "derivationPath": "m/44'/7562605'/0'/0'/0'",
        "curve": "ed25519",
        "publicKeyType": "ed25519",
        "explorer": "https://semux.info/explorer/transaction/",
        "info": {
            "url": "https://www.semux.org",
            "client": "https://github.com/semuxproject/semux-core",
            "clientPublic": "",
            "clientDocs": "https://www.semux.org"
        }
    },
    {
        "id": "zelcash",
        "name": "Zelcash",
        "symbol": "ZEL",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/44'/19167'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "staticPrefix": 28,
        "p2pkhPrefix": 184,
        "p2shPrefix": 189,
        "publicKeyHasher": "sha256ripemd",
        "base58Hasher": "sha256d",
        "xpub": "xpub",
        "xprv": "xprv",
        "explorer": "https://explorer.zel.cash/tx/",
        "info": {
            "url": "https://zel.cash",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "https://blockbook.zel.cash",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "ark",
        "name": "ARK",
        "symbol": "ARK",
        "decimals": 8,
        "blockchain": "Ark",
        "derivationPath": "m/44'/111'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "explorer": "https://explorer.ark.io/transaction/",
        "info": {
            "url": "http://ark.io",
            "client": "https://github.com/ArkEcosystem/core",
            "clientPublic": "",
            "clientDocs": "https://docs.ark.io/api/json-rpc/"
        }
    },
    {
        "id": "ravencoin",
        "name": "Ravencoin",
        "symbol": "RVN",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/44'/175'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "p2pkhPrefix": 60,
        "p2shPrefix": 122,
        "publicKeyHasher": "sha256ripemd",
        "base58Hasher": "sha256d",
        "xpub": "xpub",
        "xprv": "xprv",
        "explorer": "https://ravencoin.network/tx/",
        "info": {
            "url": "https://ravencoin.org",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "waves",
        "name": "Waves",
        "symbol": "WAVES",
        "decimals": 8,
        "blockchain": "Waves",
        "derivationPath": "m/44'/5741564'/0'/0'/0'",
        "curve": "ed25519",
        "publicKeyType": "curve25519",
        "explorer": "https://wavesexplorer.com/tx/",
        "info": {
            "url": "https://wavesplatform.com",
            "client": "https://github.com/wavesplatform/Waves",
            "clientPublic": "https://nodes.wavesnodes.com",
            "clientDocs": "https://nodes.wavesnodes.com/api-docs/index.html"
        }
    },
    {
        "id": "aeternity",
        "name": "Aeternity",
        "symbol": "AE",
        "decimals": 18,
        "blockchain": "Aeternity",
        "derivationPath": "m/44'/457'/0'/0'/0'",
        "curve": "ed25519",
        "publicKeyType": "ed25519",
        "explorer": "https://explorer.aepps.com/#/tx/",
        "info": {
            "url": "https://aeternity.com",
            "client": "https://github.com/aeternity/aeternity",
            "clientPublic": "https://sdk-mainnet.aepps.com",
            "clientDocs": "http://aeternity.com/api-docs/"
        }
    },
    {
        "id": "terra",
        "name": "Terra",
        "symbol": "LUNA",
        "decimals": 6,
        "blockchain": "Cosmos",
        "derivationPath": "m/44'/330'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "hrp": "terra",
        "explorer": "https://terra.stake.id/?#/tx/",
        "info": {
            "url": "https://terra.money",
            "client": "https://github.com/terra-project/core",
            "clientPublic": "",
            "clientDocs": "https://docs.terra.money"
        }
    },
    {
        "id": "monacoin",
        "name": "Monacoin",
        "symbol": "MONA",
        "decimals": 8,
        "blockchain": "Bitcoin",
        "derivationPath": "m/44'/22'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "p2pkhPrefix": 50,
        "p2shPrefix": 55,
        "hrp": "mona",
        "publicKeyHasher": "sha256ripemd",
        "base58Hasher": "sha256d",
        "xpub": "xpub",
        "xprv": "xprv",
        "explorer": "https://blockbook.electrum-mona.org/tx/",
        "info": {
            "url": "https://monacoin.org",
            "client": "https://github.com/trezor/blockbook",
            "clientPublic": "",
            "clientDocs": "https://github.com/trezor/blockbook/blob/master/docs/api.md"
        }
    },
    {
        "id": "fio",
        "name": "FIO",
        "symbol": "FIO",
        "decimals": 9,
        "blockchain": "FIO",
        "derivationPath": "m/44'/235'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1",
        "explorer": "https://fio.foundation",
        "url": "https://fio.foundation",
        "rpcNodeInfo": "https://fio.foundation",
        "info": {
            "url": "https://fio.foundation",
            "client": "https://fio.foundation",
            "clientPublic": "",
            "clientDocs": "https://fio.foundation"
        }
    },
    {
        "id": "harmony",
        "name": "Harmony",
        "symbol": "ONE",
        "decimals": 18,
        "blockchain": "Harmony",
        "hrp": "one",
        "derivationPath": "m/44'/1023'/0'/0/0",
        "curve": "secp256k1",
        "publicKeyType": "secp256k1Extended",
        "explorer": "https://explorer.harmony.one/#/tx/",
        "info": {
            "url": "https://harmony.one",
            "client": "https://github.com/harmony-one/go-sdk",
            "clientPublic": "",
            "clientDocs": "https://app.gitbook.com/@harmony-one/s/sdk-wiki"
        }
    },
    {
        "id": "solana",
        "name": "Solana",
        "symbol": "SOL",
        "decimals": 13,
        "blockchain": "Solana",
        "derivationPath": "m/44'/501'/0'",
        "curve": "ed25519",
        "publicKeyType": "ed25519",
        "explorer": "https://explorer.solana.com/tx/",
        "info": {
            "url": "https://solana.com",
            "client": "https://github.com/solana-labs/solana",
            "clientPublic": "http://api.mainnet.solana.com",
            "clientDocs": "https://solana-labs.github.io/book/"
        }
    },
    {
        "id": "near",
        "name": "NEAR",
        "symbol": "NEAR",
        "decimals": 18,
        "blockchain": "NEAR",
        "derivationPath": "m/44'/397'/0'",
        "curve": "ed25519",
        "publicKeyType": "ed25519",
        "explorer": "https://explorer.nearprotocol.com",
        "info": {
            "url": "https://nearprotocol.com",
            "client": "https://github.com/nearprotocol/nearcore",
            "clientPublic": "https://rpc.nearprotocol.com",
            "clientDocs": "https://docs.nearprotocol.com"
        }
    },
    {
        "id": "algorand",
        "name": "Algorand",
        "symbol": "ALGO",
        "decimals": 6,
        "blockchain": "Algorand",
        "derivationPath": "m/44'/283'/0'/0'/0'",
        "curve": "ed25519",
        "publicKeyType": "ed25519",
        "explorer": "https://algoexplorer.io/tx/",
        "info": {
            "url": "https://www.algorand.com/",
            "client": "https://github.com/algorand/go-algorand",
            "clientPublic": "https://indexer.algorand.network",
            "clientDocs": "https://developer.algorand.org/docs/algod-rest-paths"
        }
    }
]

```



**Thanks [Wallet-Core](https://github.com/trustwallet/wallet-core)** ！






