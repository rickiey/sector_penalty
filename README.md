# sector_penalty 

#### Estimated Filecoin network penalties for miner faulted sector

- [x] windowPoSt  handleProvingDeadline penalty  
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


