## Description
The repository is the set of tools to collect cosmos node networking performance.
The metrics collection repository contains three key directories: prometheus, gaia, and metrics.

The metrics are being collected each 5s triggered by Prometheus that queries middle layer app written in go and running on `http://127.0.0.1:3000` and underneath operates with a cosmos node running on `http://127.0.0.1:26657`

```Gaia Node(API) <- Go App <- Prometheus```

The application uses only two endpoints: `http://localhost:26657/status`, `http://localhost:26657/net_info` with three properties collected: `n_peers, latest_block_time, latest_block_height` which can be easilly tuned.
Full specification of Cosmos API can be found by the link https://docs.tendermint.com/v0.34/rpc/#/

## Cosmos Installation Manual
Installation of the cosmos network even on a testnet might be tricky, the application has undergone updates and has multiple testnets running. I also find that the documentation is not always correct having wrong or outdated values in readme files.

During the development, the test network on my local machine was launched manually.

Testnet explorer: https://explorer.theta-testnet.polypore.xyz/
Testent id: theta-testnet-001


1. Install and build node with network genesis file downloaded
```
git clone https://github.com/cosmos/testnets/tree/master/public
git checkout v7.0.0

make install
gaiad init CUSTOM_MONIKER --chain-id theta-testnet-001

wget https://github.com/cosmos/testnets/raw/master/public/genesis.json.gz
gunzip genesis.json.gz
cp genesis.json $HOME/.gaia/config/genesis.json

gaiad config chain-id theta-testnet-001
```

2. Set the network parameters as described here, see https://hub.cosmos.network/main/hub-tutorials/join-testnet.html
Or just copy the config files from this repo to Gaia home dir `$HOME/.gaia/config`. 
```
cp gaia/app.toml $HOME/.gaia/config
cp gaia/client.toml $HOME/.gaia/config
cp gaia/config.toml $HOME/.gaia/config
```

3. Launch gaia network
```
gaiad start --x-crisis-skip-assert-invariants
```

## Cosmos Installation Ansible
```
https://github.com/hyphacoop/cosmos-ansible
```

## Metrics - Go Daemon
```
nano /etc/systemd/system/gaiametrics.service

sudo systemctl daemon-reload
sudo systemctl start gaiametrics.service
sudo systemctl enable gaiametrics.service
```

See metrics running with
```
journalctl -fu gaiametrics.service
journalctl -u gaiametrics.service -f
```

![prometeus desync](img/prometheus_desync.png)
![metrics running](img/metrics_go.png)

## Prometheus Daemon
```
nano /etc/systemd/system/prometheus.service

sudo systemctl daemon-reload
sudo systemctl start prometheus.service
sudo systemctl enable prometheus.service
```

## Node app 
```
npm run dev
npm run build
http://127.0.0.1:3000/
```