# Guesstimate Groth16 proving time

[Here is the heuristics from libsnark](https://github.com/scipr-lab/libsnark/blob/master/libsnark/zk_proof_systems/ppzksnark/README.md#asymptotic-performance).
```bash
(time go test ./... -v -bench=. -run=Bench -benchtime=1x -timeout 150m)
```
