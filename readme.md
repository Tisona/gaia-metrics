## Installation

see https://github.com/cosmos/testnets/tree/master/public

theta-testnet-001

git checkout v7.0.0

make install
gaiad init CUSTOM_MONIKER --chain-id theta-testnet-001

wget https://github.com/cosmos/testnets/raw/master/public/genesis.json.gz
gunzip genesis.json.gz
cp genesis.json $HOME/.gaia/config/genesis.json


cd $HOME/.gaia/config
sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0.001uatom"/' app.toml
gsed -i 's/persistent_peers = ""/persistent_peers = ""/' config.toml

gaiad config chain-id theta-testnet-001



gaiad start --x-crisis-skip-assert-invariants


## Graphana
curl -O https://dl.grafana.com/oss/release/grafana-7.1.5.darwin-amd64.tar.gz
./bin/grafana-server web


## Websockets
https://hub.cosmos.network/main/getting-started/quickstart.html
https://explorer.theta-testnet.polypore.xyz/blocks
http://localhost:26657/status?
https://github.com/websockets/wscat

https://docs.tendermint.com/v0.34/tendermint-core/subscription.html