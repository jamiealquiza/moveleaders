# moveleaders

Moveleaders takes an input Kafka assignment map and a comma-separated list of broker IDs. Any IDs specified will be demoted from leadership positions.

```
% ./moveleaders -file ../map.json -brokers-to-demote 40002,40005,40008,40011

Partition map changes:
process p0: [40001 40009] -> [40001 40009] no-op
process p1: [40006 40005] -> [40006 40005] no-op
process p2: [40010 40002] -> [40010 40002] no-op
process p3: [40008 40004] -> [40004 40008] preferred leader
process p4: [40010 40011] -> [40010 40011] no-op
process p5: [40002 40007] -> [40007 40002] preferred leader
process p6: [40003 40002] -> [40003 40002] no-op
process p7: [40005 40003] -> [40003 40005] preferred leader
process p8: [40007 40005] -> [40007 40005] no-op
process p9: [40009 40011] -> [40009 40011] no-op
process p10: [40000 40011] -> [40000 40011] no-op
process p11: [40011 40010] -> [40010 40011] preferred leader
process p12: [40009 40007] -> [40009 40007] no-op
process p13: [40004 40006] -> [40004 40006] no-op
process p14: [40002 40001] -> [40001 40002] preferred leader
process p15: [40009 40011] -> [40009 40011] no-op
process p16: [40005 40009] -> [40009 40005] preferred leader
process p17: [40007 40006] -> [40007 40006] no-op
process p18: [40001 40000] -> [40001 40000] no-op
process p19: [40006 40010] -> [40006 40010] no-op
process p20: [40011 40001] -> [40001 40011] preferred leader
process p21: [40010 40006] -> [40010 40006] no-op
process p22: [40009 40002] -> [40009 40002] no-op
process p23: [40011 40010] -> [40010 40011] preferred leader
process p24: [40003 40005] -> [40003 40005] no-op
process p25: [40004 40009] -> [40004 40009] no-op
process p26: [40007 40011] -> [40007 40011] no-op
process p27: [40001 40002] -> [40001 40002] no-op
process p28: [40006 40008] -> [40006 40008] no-op
process p29: [40010 40005] -> [40010 40005] no-op
process p30: [40011 40000] -> [40000 40011] preferred leader
process p31: [40010 40009] -> [40010 40009] no-op
process p32: [40002 40007] -> [40007 40002] preferred leader

wrote move-leaders.json
```
