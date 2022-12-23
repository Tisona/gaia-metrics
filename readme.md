## Ansible installation
```
https://github.com/cosmos/testnets/tree/master/public
https://github.com/hyphacoop/cosmos-ansible
```

## Cosmovisor
```
journalctl -fu cosmovisor
```

## Prometheus
```
nano /etc/systemd/system/prometheus.service

journalctl -fu prometheus.service
journalctl -u prometheus.service -f

sudo systemctl daemon-reload
sudo systemctl start prometheus.service
sudo systemctl enable prometheus.service
```

## Manual Installation Cosmos
```
Testnet https://explorer.theta-testnet.polypore.xyz/

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
```


## Node app 
```
npm run dev
npm run build
http://127.0.0.1:3000/
```

## Go app 
```
go run main.go
go build
http://127.0.0.1:3000/
```

## API used
```
https://docs.tendermint.com/v0.34/rpc/#/
http://localhost:26657/status
http://localhost:26657/net_info
```