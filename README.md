# sector_penalty 

#### Estimated Filecoin network penalties for miner faulted sector

- [x] windowPoSt  handleProvingDeadline penalty
- [x] sector InitialPledge calculate
- [x] sector terminate penalty
- [ ]  depositToBurn
- [ ] todo

#### usege:

```bash
$ go run main.go windowPost         
Current network penalty for windowPoSt failure of 32GB sector: 0.003561825011436986 FIL = 3561825011436986 attoFIL
```

or: designated miner

```bash
$ go run main.go windowPost f0397332        
Current network penalty for windowPoSt failure of 32GB sector: 0.003561831381557233 FIL = 3561831381557233 attoFIL 
miner f0397332 fault sectors 232 , recoveries sectors 0 
miner f0397332 Estimated 24-hour windowPoSt penalty for failed sectors: 0.8263448805212781 FIL = 826344880521278056 attoFIL
```

```bash
go run main.go init-pledge
InitialPledge for 32 GiB:  0.17011580252903266
InitialPledge for 1 TiB:  5.443705680929045  
```

```bash
go run main.go sector terminate f01132416 16781
miner: f01132416 sector ID: 16781 , start height: 1451551  expiration height: 3000401,  ExpectedDayReward: 651360883065113 
penalty for terminating sector : 58320649085073059 attoFIL about 0.0583206491 FIL 
```




![Alt](https://repobeats.axiom.co/api/embed/7e33e91436ecc8340f7b2e2988047e1f4a2c016e.svg "Repobeats analytics image")
