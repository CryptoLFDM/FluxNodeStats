# FluxNodeStats
This little api written in go is designed to calcul node rentability on flux.

We harvest date from 

**https://api.runonflux.io/daemon/getzelnodecount**

**https://explorer.runonflux.io/api/statistics/total**


## Config


````yaml
---
api_port: 8080
api_log_file: ""
api_adress: localhost

````

## Usage

We wrote it to be able to run as api, little change is needed if you only want reporting

``````bash
go build
./FluxNodeStats -config path/to_your_config
``````
Then a server will start based on the config file.

Actually four routes are availables

````log
[GIN-debug] GET    /health                      --> FluxNodeStats/routes.Health (3 handlers)
[GIN-debug] GET    /flux_nodes_data             --> FluxNodeStats/routes.HarvestNodesInfo (3 handlers)
[GIN-debug] GET    /flux_blocs_data             --> FluxNodeStats/routes.HarvestBlocksInfo (3 handlers)
[GIN-debug] GET    /calcul_nodes_rentability    --> FluxNodeStats/routes.CalculNodesRentability (3 handlers)

````
### routes

All route about data return Json from server call, log json directly on console actually.

#### health
is just the health check


``````json
{"status": "We are Alive"}

``````

####  flux_nodes_data

Retrieves nodes stats

````json
{
  "status": "success",
  "data": {
    "total": 4805,
    "stable": 4805,
    "basic-enabled": 3413,
    "super-enabled": 843,
    "bamf-enabled": 549,
    "cumulus-enabled": 3413,
    "nimbus-enabled": 843,
    "stratus-enabled": 549,
    "ipv4": 4805,
    "ipv6": 0,
    "onion": 0
  }
}
````

#### flux_blocs_data

Retrieves blocs stats

````json
{
  "n_blocks_mined": 720,
  "time_between_blocks": 119.35744089012518,
  "mined_currency_amount": 5400000000000,
  "transaction_fees": 10506720,
  "number_of_transactions": 31157,
  "outputs_volume": 4524888089462817,
  "difficulty": "29229.53813487229",
  "network_hash_ps": 1950978.9888888889,
  "blocks_by_pool": [
    {
      "address": "t1ZSUsQv2dHmAdsHPTcfd3UAigyreaYTxjq",
      "poolName": "MinerPool",
      "url": "https://flux.minerpool.org",
      "blocks_found": 383,
      "percent_total": "53.19"
    },
    {
      "address": "t1J7WrTnM4tWr5mSjxVNECxKrMmQzggNqJy",
      "poolName": "Flux Labs",
      "url": "https://fluxpools.net/coins/flux",
      "blocks_found": 163,
      "percent_total": "22.64"
    },
    {
      "address": "t1gomWuf4C5qpgHoTQjLUEP9PUdMcCcJXjb",
      "poolName": "Hero Miners",
      "url": "https://flux.herominers.com",
      "blocks_found": 71,
      "percent_total": "9.86"
    },
    {
      "address": "t1JKRwXGfKTGfPV1z48rvoLyabk31z3xwHa",
      "poolName": "2Miners PPLNS",
      "url": "https://2miners.com/zel-mining-pool",
      "blocks_found": 44,
      "percent_total": "6.11"
    },
    {
      "address": "t1Um6mMrMWAa72gh15e1cZ299pK7azCx5YF",
      "poolName": "MinerPool SOLO",
      "url": "https://solo-flux.minerpool.org",
      "blocks_found": 25,
      "percent_total": "3.47"
    },
    {
      "address": "t1exvkXQb3RZFGaobUs87foqw93K3fVq5wb",
      "poolName": "ZPOOL",
      "url": "https://zpool.ca",
      "blocks_found": 7,
      "percent_total": "0.97"
    },
    {
      "address": "t1ZqpNxeMtPQVGtZ7hEkWJJv74v3FQHCEMj",
      "poolName": "2Miners SOLO",
      "url": "https://2miners.com/solo-zel-mining-pool",
      "blocks_found": 7,
      "percent_total": "0.97"
    },
    {
      "address": "t1RpbpUWos9fNbAFPXy9fAm1jVqYYBUi1ou",
      "poolName": "CoinBlockers",
      "url": "https://zel.coinblockers.com/",
      "blocks_found": 6,
      "percent_total": "0.83"
    },
    {
      "address": "t1P8Woa26PQvePZSurvH51yrqgLayywghJe",
      "poolName": "EnigmaPool",
      "url": "https://enigmapool.com",
      "blocks_found": 5,
      "percent_total": "0.69"
    },
    {
      "address": "t1ZVgGSRzm8ofu31vNkvtKhVunzabWKzGyV",
      "poolName": "SoloPool.org",
      "url": "https://zel.solopool.org/",
      "blocks_found": 3,
      "percent_total": "0.42"
    },
    {
      "address": "t1YiZpbVfzyWYHZwgGVaqqaLShtcHrF1mmy",
      "poolName": "Unknown",
      "url": "",
      "blocks_found": 2,
      "percent_total": "0.28"
    },
    {
      "address": "t1RDHkTgTnVordf5fbyuuUWHH8khPSxK8Xf",
      "poolName": "Flux Labs Solo",
      "url": "https://fluxpools.net/coins/flux-solo",
      "blocks_found": 2,
      "percent_total": "0.28"
    },
    {
      "address": "t1QMb3WFmm1KLHPCGyB3NAZtZviKfezxui4",
      "poolName": "Unknown",
      "url": "",
      "blocks_found": 1,
      "percent_total": "0.14"
    },
    {
      "address": "t1WKGsEWiqCwZXeiBY2VTvmVzLgMYpk6YEK",
      "poolName": "Unknown",
      "url": "",
      "blocks_found": 1,
      "percent_total": "0.14"
    }
  ]
}
````

#### calcul_nodes_rentability

Based on the two handlers flux_blocs_data & flux_nodes_data, calcul rentabilty of each node

``````json
[
  {
    "nodes_default": {
      "name": "Cumulus",
      "disk_type": "SSD/NVME",
      "cpu": 2,
      "threads": 4,
      "gb_ram": 8,
      "collateral": 1000,
      "gb_disk_size": 220,
      "mb_disk_speed": 180,
      "mb_bandwidth": 25,
      "mb_eps_min": 250,
      "reward": 7.5
    },
    "flux_reward": 5.625,
    "flux_instant_pa_reward": 2.8125,
    "flux_later_pa_reward": 2.8125,
    "delay_reward_minutes": 6926.533988764044,
    "delay_reward_day": 4.810093047752808,
    "flux_reward_7_day": 8.185912332484985,
    "flux_instant_pa_reward_7_day": 4.092956166242493,
    "flux_later_pa_reward_7_day": 4.092956166242493,
    "flux_total_reward_7_day": 16.37182466496997,
    "flux_reward_15_day": 17.541240712467825,
    "flux_instant_pa_reward_15_day": 8.770620356233913,
    "flux_later_pa_reward_15_day": 8.770620356233913,
    "flux_total_reward_15_day": 35.08248142493565,
    "flux_reward_30_day": 35.08248142493565,
    "flux_instant_pa_reward_30_day": 17.541240712467825,
    "flux_later_pa_reward_30_day": 17.541240712467825,
    "flux_total_reward_30_day": 70.1649628498713,
    "flux_reward_180_day": 210.4948885496139,
    "flux_instant_pa_reward_180_day": 105.24744427480695,
    "flux_later_pa_reward_180_day": 105.24744427480695,
    "flux_total_reward_180_day": 420.9897770992278,
    "flux_reward_365_day": 426.83685733671706,
    "flux_instant_pa_reward_365_day": 213.41842866835853,
    "flux_later_pa_reward_365_day": 213.41842866835853,
    "flux_total_reward_365_day": 853.6737146734341
  },
  {
    "nodes_default": {
      "name": "Nimbus",
      "disk_type": "SSD/NVME",
      "cpu": 4,
      "threads": 8,
      "gb_ram": 32,
      "collateral": 12500,
      "gb_disk_size": 440,
      "mb_disk_speed": 180,
      "mb_bandwidth": 50,
      "mb_eps_min": 640,
      "reward": 12.5
    },
    "flux_reward": 9.375,
    "flux_instant_pa_reward": 4.6875,
    "flux_later_pa_reward": 4.6875,
    "delay_reward_minutes": 1705.3660112359548,
    "delay_reward_day": 1.1842819522471908,
    "flux_reward_7_day": 55.413324399207205,
    "flux_instant_pa_reward_7_day": 27.706662199603603,
    "flux_later_pa_reward_7_day": 27.706662199603603,
    "flux_total_reward_7_day": 110.82664879841441,
    "flux_reward_15_day": 118.74283799830116,
    "flux_instant_pa_reward_15_day": 59.37141899915058,
    "flux_later_pa_reward_15_day": 59.37141899915058,
    "flux_total_reward_15_day": 237.4856759966023,
    "flux_reward_30_day": 237.4856759966023,
    "flux_instant_pa_reward_30_day": 118.74283799830116,
    "flux_later_pa_reward_30_day": 118.74283799830116,
    "flux_total_reward_30_day": 474.9713519932046,
    "flux_reward_180_day": 1424.9140559796137,
    "flux_instant_pa_reward_180_day": 712.4570279898069,
    "flux_later_pa_reward_180_day": 712.4570279898069,
    "flux_total_reward_180_day": 2849.8281119592275,
    "flux_reward_365_day": 2889.4090579586614,
    "flux_instant_pa_reward_365_day": 1444.7045289793307,
    "flux_later_pa_reward_365_day": 1444.7045289793307,
    "flux_total_reward_365_day": 5778.818115917323
  },
  {
    "nodes_default": {
      "name": "Stratus",
      "disk_type": "SSD/NVME",
      "cpu": 8,
      "threads": 16,
      "gb_ram": 64,
      "collateral": 40000,
      "gb_disk_size": 880,
      "mb_disk_speed": 400,
      "mb_bandwidth": 100,
      "mb_eps_min": 1520,
      "reward": 30
    },
    "flux_reward": 22.5,
    "flux_instant_pa_reward": 11.25,
    "flux_later_pa_reward": 11.25,
    "delay_reward_minutes": 1109.2961376404494,
    "delay_reward_day": 0.7703445400280898,
    "flux_reward_7_day": 204.45397067947923,
    "flux_instant_pa_reward_7_day": 102.22698533973961,
    "flux_later_pa_reward_7_day": 102.22698533973961,
    "flux_total_reward_7_day": 408.90794135895845,
    "flux_reward_15_day": 438.1156514560269,
    "flux_instant_pa_reward_15_day": 219.05782572801346,
    "flux_later_pa_reward_15_day": 219.05782572801346,
    "flux_total_reward_15_day": 876.2313029120538,
    "flux_reward_30_day": 876.2313029120538,
    "flux_instant_pa_reward_30_day": 438.1156514560269,
    "flux_later_pa_reward_30_day": 438.1156514560269,
    "flux_total_reward_30_day": 1752.4626058241076,
    "flux_reward_180_day": 5257.387817472323,
    "flux_instant_pa_reward_180_day": 2628.6939087361616,
    "flux_later_pa_reward_180_day": 2628.6939087361616,
    "flux_total_reward_180_day": 10514.775634944646,
    "flux_reward_365_day": 10660.814185429988,
    "flux_instant_pa_reward_365_day": 5330.407092714994,
    "flux_later_pa_reward_365_day": 5330.407092714994,
    "flux_total_reward_365_day": 21321.628370859977
  }
]
``````