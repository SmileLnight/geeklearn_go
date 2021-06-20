# GetSetPerformance

---

# Value=10

## SetPerformance

```bash
root@e54d97d17810:/data# redis-benchmark -d 10 -t get,set
====== SET ======
  100000 requests completed in 5.42 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.983 milliseconds (cumulative count 1)
50.000% <= 2.247 milliseconds (cumulative count 50059)
75.000% <= 2.671 milliseconds (cumulative count 75050)
87.500% <= 2.975 milliseconds (cumulative count 87722)
93.750% <= 3.135 milliseconds (cumulative count 93738)
96.875% <= 3.255 milliseconds (cumulative count 96885)
98.438% <= 3.375 milliseconds (cumulative count 98447)
99.219% <= 3.503 milliseconds (cumulative count 99203)
99.609% <= 3.639 milliseconds (cumulative count 99569)
99.805% <= 3.799 milliseconds (cumulative count 99758)
99.902% <= 3.991 milliseconds (cumulative count 99859)
99.951% <= 4.223 milliseconds (cumulative count 99905)
99.976% <= 4.463 milliseconds (cumulative count 99930)
99.988% <= 4.671 milliseconds (cumulative count 99941)
99.994% <= 4.799 milliseconds (cumulative count 99948)
99.997% <= 4.903 milliseconds (cumulative count 99950)
99.998% <= 4.999 milliseconds (cumulative count 99952)
99.999% <= 5.063 milliseconds (cumulative count 99953)
100.000% <= 5.063 milliseconds (cumulative count 99953)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.001% <= 1.007 milliseconds (cumulative count 1)
0.008% <= 1.103 milliseconds (cumulative count 8)
0.046% <= 1.207 milliseconds (cumulative count 46)
0.150% <= 1.303 milliseconds (cumulative count 150)
0.360% <= 1.407 milliseconds (cumulative count 360)
0.860% <= 1.503 milliseconds (cumulative count 860)
2.729% <= 1.607 milliseconds (cumulative count 2728)
6.473% <= 1.703 milliseconds (cumulative count 6470)
11.838% <= 1.807 milliseconds (cumulative count 11832)
18.056% <= 1.903 milliseconds (cumulative count 18048)
27.968% <= 2.007 milliseconds (cumulative count 27955)
37.927% <= 2.103 milliseconds (cumulative count 37909)
92.644% <= 3.103 milliseconds (cumulative count 92600)
99.933% <= 4.103 milliseconds (cumulative count 99886)
100.000% <= 5.103 milliseconds (cumulative count 99953)

Summary:
  throughput summary: 18443.38 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.335     0.976     2.247     3.175     3.447     5.063
```

## GetPerformance

```bash
====== GET ======
  100000 requests completed in 5.37 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.007 milliseconds (cumulative count 2)
50.000% <= 2.215 milliseconds (cumulative count 50487)
75.000% <= 2.631 milliseconds (cumulative count 75388)
87.500% <= 2.927 milliseconds (cumulative count 87705)
93.750% <= 3.087 milliseconds (cumulative count 93871)
96.875% <= 3.199 milliseconds (cumulative count 96927)
98.438% <= 3.311 milliseconds (cumulative count 98494)
99.219% <= 3.439 milliseconds (cumulative count 99231)
99.609% <= 3.591 milliseconds (cumulative count 99624)
99.805% <= 3.759 milliseconds (cumulative count 99811)
99.902% <= 3.943 milliseconds (cumulative count 99903)
99.951% <= 4.143 milliseconds (cumulative count 99952)
99.976% <= 4.359 milliseconds (cumulative count 99977)
99.988% <= 4.551 milliseconds (cumulative count 99988)
99.994% <= 4.663 milliseconds (cumulative count 99994)
99.997% <= 4.775 milliseconds (cumulative count 99997)
99.998% <= 4.855 milliseconds (cumulative count 99999)
99.999% <= 4.951 milliseconds (cumulative count 100000)
100.000% <= 4.951 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.002% <= 1.007 milliseconds (cumulative count 2)
0.016% <= 1.103 milliseconds (cumulative count 16)
0.061% <= 1.207 milliseconds (cumulative count 61)
0.184% <= 1.303 milliseconds (cumulative count 184)
0.441% <= 1.407 milliseconds (cumulative count 441)
1.149% <= 1.503 milliseconds (cumulative count 1149)
3.552% <= 1.607 milliseconds (cumulative count 3552)
7.664% <= 1.703 milliseconds (cumulative count 7664)
13.305% <= 1.807 milliseconds (cumulative count 13305)
20.381% <= 1.903 milliseconds (cumulative count 20381)
31.154% <= 2.007 milliseconds (cumulative count 31154)
41.269% <= 2.103 milliseconds (cumulative count 41269)
94.413% <= 3.103 milliseconds (cumulative count 94413)
99.942% <= 4.103 milliseconds (cumulative count 99942)
100.000% <= 5.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 18628.91 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.298     1.000     2.215     3.127     3.383     4.951
```

---

# Value=20

## SetPerformance

```bash
root@e54d97d17810:/data# redis-benchmark -d 20 -t get,set
====== SET ======
  100000 requests completed in 5.46 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.903 milliseconds (cumulative count 1)
50.000% <= 2.247 milliseconds (cumulative count 50088)
75.000% <= 2.671 milliseconds (cumulative count 75105)
87.500% <= 2.975 milliseconds (cumulative count 87685)
93.750% <= 3.143 milliseconds (cumulative count 93996)
96.875% <= 3.263 milliseconds (cumulative count 96925)
98.438% <= 3.391 milliseconds (cumulative count 98495)
99.219% <= 3.535 milliseconds (cumulative count 99229)
99.609% <= 3.735 milliseconds (cumulative count 99613)
99.805% <= 3.975 milliseconds (cumulative count 99806)
99.902% <= 4.295 milliseconds (cumulative count 99903)
99.951% <= 4.671 milliseconds (cumulative count 99952)
99.976% <= 5.079 milliseconds (cumulative count 99976)
99.988% <= 5.607 milliseconds (cumulative count 99988)
99.994% <= 6.023 milliseconds (cumulative count 99994)
99.997% <= 6.503 milliseconds (cumulative count 99997)
99.998% <= 6.791 milliseconds (cumulative count 99999)
99.999% <= 6.863 milliseconds (cumulative count 100000)
100.000% <= 6.863 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.001% <= 0.903 milliseconds (cumulative count 1)
0.008% <= 1.007 milliseconds (cumulative count 8)
0.029% <= 1.103 milliseconds (cumulative count 29)
0.099% <= 1.207 milliseconds (cumulative count 99)
0.243% <= 1.303 milliseconds (cumulative count 243)
0.498% <= 1.407 milliseconds (cumulative count 498)
1.119% <= 1.503 milliseconds (cumulative count 1119)
3.023% <= 1.607 milliseconds (cumulative count 3023)
6.724% <= 1.703 milliseconds (cumulative count 6724)
12.091% <= 1.807 milliseconds (cumulative count 12091)
18.294% <= 1.903 milliseconds (cumulative count 18294)
28.257% <= 2.007 milliseconds (cumulative count 28257)
38.088% <= 2.103 milliseconds (cumulative count 38088)
92.645% <= 3.103 milliseconds (cumulative count 92645)
99.856% <= 4.103 milliseconds (cumulative count 99856)
99.977% <= 5.103 milliseconds (cumulative count 99977)
99.994% <= 6.103 milliseconds (cumulative count 99994)
100.000% <= 7.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 18308.31 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.335     0.896     2.247     3.183     3.479     6.863
```

## GetPerformance

```bash
====== GET ======
  100000 requests completed in 5.37 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.007 milliseconds (cumulative count 2)
50.000% <= 2.215 milliseconds (cumulative count 50460)
75.000% <= 2.631 milliseconds (cumulative count 75301)
87.500% <= 2.927 milliseconds (cumulative count 87689)
93.750% <= 3.087 milliseconds (cumulative count 93831)
96.875% <= 3.199 milliseconds (cumulative count 96933)
98.438% <= 3.311 milliseconds (cumulative count 98458)
99.219% <= 3.447 milliseconds (cumulative count 99253)
99.609% <= 3.591 milliseconds (cumulative count 99614)
99.805% <= 3.759 milliseconds (cumulative count 99811)
99.902% <= 3.919 milliseconds (cumulative count 99904)
99.951% <= 4.111 milliseconds (cumulative count 99952)
99.976% <= 4.263 milliseconds (cumulative count 99976)
99.988% <= 4.455 milliseconds (cumulative count 99989)
99.994% <= 4.583 milliseconds (cumulative count 99994)
99.997% <= 4.655 milliseconds (cumulative count 99997)
99.998% <= 4.767 milliseconds (cumulative count 99999)
99.999% <= 4.863 milliseconds (cumulative count 100000)
100.000% <= 4.863 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.002% <= 1.007 milliseconds (cumulative count 2)
0.019% <= 1.103 milliseconds (cumulative count 19)
0.079% <= 1.207 milliseconds (cumulative count 79)
0.244% <= 1.303 milliseconds (cumulative count 244)
0.589% <= 1.407 milliseconds (cumulative count 589)
1.299% <= 1.503 milliseconds (cumulative count 1299)
3.598% <= 1.607 milliseconds (cumulative count 3598)
7.762% <= 1.703 milliseconds (cumulative count 7762)
13.519% <= 1.807 milliseconds (cumulative count 13519)
20.456% <= 1.903 milliseconds (cumulative count 20456)
31.403% <= 2.007 milliseconds (cumulative count 31403)
41.285% <= 2.103 milliseconds (cumulative count 41285)
94.397% <= 3.103 milliseconds (cumulative count 94397)
99.951% <= 4.103 milliseconds (cumulative count 99951)
100.000% <= 5.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 18621.97 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.297     1.000     2.215     3.127     3.391     4.863
```

---

# Value=50

## SetPerformance

```bash
root@e54d97d17810:/data# redis-benchmark -d 50 -t get,set
====== SET ======
  100000 requests completed in 5.40 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.071 milliseconds (cumulative count 2)
50.000% <= 2.239 milliseconds (cumulative count 50340)
75.000% <= 2.655 milliseconds (cumulative count 75016)
87.500% <= 2.959 milliseconds (cumulative count 87597)
93.750% <= 3.119 milliseconds (cumulative count 93714)
96.875% <= 3.239 milliseconds (cumulative count 96923)
98.438% <= 3.359 milliseconds (cumulative count 98424)
99.219% <= 3.567 milliseconds (cumulative count 99193)
99.609% <= 4.071 milliseconds (cumulative count 99570)
99.805% <= 5.031 milliseconds (cumulative count 99765)
99.902% <= 5.831 milliseconds (cumulative count 99863)
99.951% <= 6.359 milliseconds (cumulative count 99911)
99.976% <= 6.839 milliseconds (cumulative count 99935)
99.988% <= 7.183 milliseconds (cumulative count 99947)
99.994% <= 7.343 milliseconds (cumulative count 99953)
99.997% <= 7.447 milliseconds (cumulative count 99956)
99.998% <= 7.487 milliseconds (cumulative count 99958)
99.999% <= 7.575 milliseconds (cumulative count 99959)
100.000% <= 7.575 milliseconds (cumulative count 99959)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.003% <= 1.103 milliseconds (cumulative count 3)
0.025% <= 1.207 milliseconds (cumulative count 25)
0.091% <= 1.303 milliseconds (cumulative count 91)
0.234% <= 1.407 milliseconds (cumulative count 234)
0.634% <= 1.503 milliseconds (cumulative count 634)
2.254% <= 1.607 milliseconds (cumulative count 2253)
5.914% <= 1.703 milliseconds (cumulative count 5912)
11.473% <= 1.807 milliseconds (cumulative count 11468)
18.031% <= 1.903 milliseconds (cumulative count 18024)
28.524% <= 2.007 milliseconds (cumulative count 28512)
39.002% <= 2.103 milliseconds (cumulative count 38986)
93.149% <= 3.103 milliseconds (cumulative count 93111)
99.623% <= 4.103 milliseconds (cumulative count 99582)
99.816% <= 5.103 milliseconds (cumulative count 99775)
99.935% <= 6.103 milliseconds (cumulative count 99894)
99.984% <= 7.103 milliseconds (cumulative count 99943)
100.000% <= 8.103 milliseconds (cumulative count 99959)

Summary:
  throughput summary: 18511.66 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.334     1.064     2.239     3.167     3.471     7.575
```

## GetPerformance

```bash
====== GET ======
  100000 requests completed in 5.35 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.999 milliseconds (cumulative count 1)
50.000% <= 2.207 milliseconds (cumulative count 50174)
75.000% <= 2.623 milliseconds (cumulative count 75379)
87.500% <= 2.919 milliseconds (cumulative count 87740)
93.750% <= 3.079 milliseconds (cumulative count 94019)
96.875% <= 3.183 milliseconds (cumulative count 96966)
98.438% <= 3.279 milliseconds (cumulative count 98464)
99.219% <= 3.391 milliseconds (cumulative count 99230)
99.609% <= 3.511 milliseconds (cumulative count 99626)
99.805% <= 3.631 milliseconds (cumulative count 99813)
99.902% <= 3.727 milliseconds (cumulative count 99908)
99.951% <= 3.823 milliseconds (cumulative count 99952)
99.976% <= 3.903 milliseconds (cumulative count 99976)
99.988% <= 3.975 milliseconds (cumulative count 99988)
99.994% <= 4.079 milliseconds (cumulative count 99994)
99.997% <= 4.119 milliseconds (cumulative count 99997)
99.998% <= 4.199 milliseconds (cumulative count 99999)
99.999% <= 4.255 milliseconds (cumulative count 100000)
100.000% <= 4.255 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.001% <= 1.007 milliseconds (cumulative count 1)
0.017% <= 1.103 milliseconds (cumulative count 17)
0.086% <= 1.207 milliseconds (cumulative count 86)
0.230% <= 1.303 milliseconds (cumulative count 230)
0.492% <= 1.407 milliseconds (cumulative count 492)
1.108% <= 1.503 milliseconds (cumulative count 1108)
3.411% <= 1.607 milliseconds (cumulative count 3411)
7.801% <= 1.703 milliseconds (cumulative count 7801)
13.579% <= 1.807 milliseconds (cumulative count 13579)
20.480% <= 1.903 milliseconds (cumulative count 20480)
31.561% <= 2.007 milliseconds (cumulative count 31561)
41.584% <= 2.103 milliseconds (cumulative count 41584)
94.809% <= 3.103 milliseconds (cumulative count 94809)
99.995% <= 4.103 milliseconds (cumulative count 99995)
100.000% <= 5.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 18684.60 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.292     0.992     2.207     3.111     3.351     4.255
```

---

# Value=100

## SetPerformance

```bash
root@e54d97d17810:/data# redis-benchmark -d 100 -t get,set
====== SET ======
  100000 requests completed in 5.39 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.959 milliseconds (cumulative count 1)
50.000% <= 2.223 milliseconds (cumulative count 50572)
75.000% <= 2.631 milliseconds (cumulative count 75056)
87.500% <= 2.935 milliseconds (cumulative count 87613)
93.750% <= 3.103 milliseconds (cumulative count 94010)
96.875% <= 3.207 milliseconds (cumulative count 96991)
98.438% <= 3.295 milliseconds (cumulative count 98445)
99.219% <= 3.407 milliseconds (cumulative count 99223)
99.609% <= 3.527 milliseconds (cumulative count 99615)
99.805% <= 3.679 milliseconds (cumulative count 99805)
99.902% <= 3.895 milliseconds (cumulative count 99903)
99.951% <= 4.183 milliseconds (cumulative count 99952)
99.976% <= 4.463 milliseconds (cumulative count 99976)
99.988% <= 4.719 milliseconds (cumulative count 99988)
99.994% <= 4.847 milliseconds (cumulative count 99994)
99.997% <= 4.951 milliseconds (cumulative count 99997)
99.998% <= 5.207 milliseconds (cumulative count 99999)
99.999% <= 5.359 milliseconds (cumulative count 100000)
100.000% <= 5.359 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.004% <= 1.007 milliseconds (cumulative count 4)
0.018% <= 1.103 milliseconds (cumulative count 18)
0.071% <= 1.207 milliseconds (cumulative count 71)
0.186% <= 1.303 milliseconds (cumulative count 186)
0.418% <= 1.407 milliseconds (cumulative count 418)
1.047% <= 1.503 milliseconds (cumulative count 1047)
3.175% <= 1.607 milliseconds (cumulative count 3175)
7.278% <= 1.703 milliseconds (cumulative count 7278)
13.024% <= 1.807 milliseconds (cumulative count 13024)
19.609% <= 1.903 milliseconds (cumulative count 19609)
30.106% <= 2.007 milliseconds (cumulative count 30106)
40.470% <= 2.103 milliseconds (cumulative count 40470)
94.010% <= 3.103 milliseconds (cumulative count 94010)
99.940% <= 4.103 milliseconds (cumulative count 99940)
99.998% <= 5.103 milliseconds (cumulative count 99998)
100.000% <= 6.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 18549.44 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.306     0.952     2.223     3.135     3.367     5.359
```

## GetPerformance

```bash
====== GET ======
  100000 requests completed in 5.33 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.919 milliseconds (cumulative count 1)
50.000% <= 2.207 milliseconds (cumulative count 50175)
75.000% <= 2.623 milliseconds (cumulative count 75273)
87.500% <= 2.919 milliseconds (cumulative count 87621)
93.750% <= 3.079 milliseconds (cumulative count 93816)
96.875% <= 3.191 milliseconds (cumulative count 96938)
98.438% <= 3.295 milliseconds (cumulative count 98450)
99.219% <= 3.423 milliseconds (cumulative count 99188)
99.609% <= 3.559 milliseconds (cumulative count 99576)
99.805% <= 3.735 milliseconds (cumulative count 99768)
99.902% <= 3.919 milliseconds (cumulative count 99864)
99.951% <= 4.087 milliseconds (cumulative count 99910)
99.976% <= 4.231 milliseconds (cumulative count 99934)
99.988% <= 4.431 milliseconds (cumulative count 99946)
99.994% <= 4.599 milliseconds (cumulative count 99952)
99.997% <= 4.759 milliseconds (cumulative count 99955)
99.998% <= 4.879 milliseconds (cumulative count 99957)
99.999% <= 4.935 milliseconds (cumulative count 99958)
100.000% <= 4.935 milliseconds (cumulative count 99958)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.004% <= 1.007 milliseconds (cumulative count 4)
0.018% <= 1.103 milliseconds (cumulative count 18)
0.093% <= 1.207 milliseconds (cumulative count 93)
0.258% <= 1.303 milliseconds (cumulative count 258)
0.536% <= 1.407 milliseconds (cumulative count 536)
1.174% <= 1.503 milliseconds (cumulative count 1174)
3.531% <= 1.607 milliseconds (cumulative count 3530)
7.820% <= 1.703 milliseconds (cumulative count 7817)
13.599% <= 1.807 milliseconds (cumulative count 13593)
20.540% <= 1.903 milliseconds (cumulative count 20531)
31.370% <= 2.007 milliseconds (cumulative count 31357)
41.587% <= 2.103 milliseconds (cumulative count 41570)
94.646% <= 3.103 milliseconds (cumulative count 94606)
99.955% <= 4.103 milliseconds (cumulative count 99913)
100.000% <= 5.103 milliseconds (cumulative count 99958)

Summary:
  throughput summary: 18765.25 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.293     0.912     2.207     3.119     3.375     4.935
```

---

# Value=200

## SetPerformance

```bash
root@e54d97d17810:/data# redis-benchmark -d 200 -t get,set
====== SET ======
  100000 requests completed in 5.63 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.991 milliseconds (cumulative count 1)
50.000% <= 2.335 milliseconds (cumulative count 50128)
75.000% <= 2.783 milliseconds (cumulative count 75324)
87.500% <= 3.079 milliseconds (cumulative count 87491)
93.750% <= 3.287 milliseconds (cumulative count 93745)
96.875% <= 3.487 milliseconds (cumulative count 96869)
98.438% <= 3.663 milliseconds (cumulative count 98405)
99.219% <= 3.831 milliseconds (cumulative count 99200)
99.609% <= 3.975 milliseconds (cumulative count 99565)
99.805% <= 4.087 milliseconds (cumulative count 99758)
99.902% <= 4.231 milliseconds (cumulative count 99855)
99.951% <= 4.399 milliseconds (cumulative count 99903)
99.976% <= 4.575 milliseconds (cumulative count 99927)
99.988% <= 4.711 milliseconds (cumulative count 99940)
99.994% <= 4.775 milliseconds (cumulative count 99945)
99.997% <= 4.879 milliseconds (cumulative count 99948)
99.998% <= 4.975 milliseconds (cumulative count 99950)
99.999% <= 5.319 milliseconds (cumulative count 99951)
100.000% <= 5.319 milliseconds (cumulative count 99951)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.003% <= 1.007 milliseconds (cumulative count 3)
0.011% <= 1.103 milliseconds (cumulative count 11)
0.066% <= 1.207 milliseconds (cumulative count 66)
0.171% <= 1.303 milliseconds (cumulative count 171)
0.410% <= 1.407 milliseconds (cumulative count 410)
0.958% <= 1.503 milliseconds (cumulative count 958)
2.640% <= 1.607 milliseconds (cumulative count 2639)
5.759% <= 1.703 milliseconds (cumulative count 5756)
10.212% <= 1.807 milliseconds (cumulative count 10207)
15.537% <= 1.903 milliseconds (cumulative count 15529)
23.994% <= 2.007 milliseconds (cumulative count 23982)
32.630% <= 2.103 milliseconds (cumulative count 32614)
88.395% <= 3.103 milliseconds (cumulative count 88352)
99.826% <= 4.103 milliseconds (cumulative count 99777)
99.999% <= 5.103 milliseconds (cumulative count 99950)
100.000% <= 6.103 milliseconds (cumulative count 99951)

Summary:
  throughput summary: 17771.46 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.417     0.984     2.335     3.351     3.767     5.319
```

## GetPerformance

```bash
====== GET ======
  100000 requests completed in 5.38 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.847 milliseconds (cumulative count 1)
50.000% <= 2.215 milliseconds (cumulative count 50021)
75.000% <= 2.631 milliseconds (cumulative count 75027)
87.500% <= 2.935 milliseconds (cumulative count 87690)
93.750% <= 3.095 milliseconds (cumulative count 93887)
96.875% <= 3.207 milliseconds (cumulative count 96876)
98.438% <= 3.335 milliseconds (cumulative count 98502)
99.219% <= 3.463 milliseconds (cumulative count 99234)
99.609% <= 3.607 milliseconds (cumulative count 99617)
99.805% <= 3.735 milliseconds (cumulative count 99806)
99.902% <= 3.871 milliseconds (cumulative count 99903)
99.951% <= 4.031 milliseconds (cumulative count 99953)
99.976% <= 4.151 milliseconds (cumulative count 99976)
99.988% <= 4.239 milliseconds (cumulative count 99988)
99.994% <= 4.359 milliseconds (cumulative count 99994)
99.997% <= 4.495 milliseconds (cumulative count 99997)
99.998% <= 4.743 milliseconds (cumulative count 99999)
99.999% <= 4.823 milliseconds (cumulative count 100000)
100.000% <= 4.823 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.002% <= 0.903 milliseconds (cumulative count 2)
0.006% <= 1.007 milliseconds (cumulative count 6)
0.029% <= 1.103 milliseconds (cumulative count 29)
0.101% <= 1.207 milliseconds (cumulative count 101)
0.248% <= 1.303 milliseconds (cumulative count 248)
0.560% <= 1.407 milliseconds (cumulative count 560)
1.276% <= 1.503 milliseconds (cumulative count 1276)
3.535% <= 1.607 milliseconds (cumulative count 3535)
7.740% <= 1.703 milliseconds (cumulative count 7740)
13.338% <= 1.807 milliseconds (cumulative count 13338)
20.110% <= 1.903 milliseconds (cumulative count 20110)
30.854% <= 2.007 milliseconds (cumulative count 30854)
40.695% <= 2.103 milliseconds (cumulative count 40695)
94.128% <= 3.103 milliseconds (cumulative count 94128)
99.969% <= 4.103 milliseconds (cumulative count 99969)
100.000% <= 5.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 18580.45 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.303     0.840     2.215     3.135     3.415     4.823
```

---

# Value=1k

## SetPerformance

```bash
root@e54d97d17810:/data# redis-benchmark -d 1000 -t get,set
====== SET ======
  100000 requests completed in 5.43 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.991 milliseconds (cumulative count 1)
50.000% <= 2.239 milliseconds (cumulative count 50381)
75.000% <= 2.655 milliseconds (cumulative count 75129)
87.500% <= 2.959 milliseconds (cumulative count 87772)
93.750% <= 3.119 milliseconds (cumulative count 93854)
96.875% <= 3.231 milliseconds (cumulative count 96984)
98.438% <= 3.327 milliseconds (cumulative count 98472)
99.219% <= 3.439 milliseconds (cumulative count 99249)
99.609% <= 3.551 milliseconds (cumulative count 99620)
99.805% <= 3.671 milliseconds (cumulative count 99807)
99.902% <= 3.847 milliseconds (cumulative count 99904)
99.951% <= 4.055 milliseconds (cumulative count 99952)
99.976% <= 4.295 milliseconds (cumulative count 99976)
99.988% <= 4.567 milliseconds (cumulative count 99988)
99.994% <= 4.703 milliseconds (cumulative count 99994)
99.997% <= 4.863 milliseconds (cumulative count 99997)
99.998% <= 4.983 milliseconds (cumulative count 99999)
99.999% <= 5.047 milliseconds (cumulative count 100000)
100.000% <= 5.047 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.001% <= 1.007 milliseconds (cumulative count 1)
0.010% <= 1.103 milliseconds (cumulative count 10)
0.051% <= 1.207 milliseconds (cumulative count 51)
0.157% <= 1.303 milliseconds (cumulative count 157)
0.383% <= 1.407 milliseconds (cumulative count 383)
0.907% <= 1.503 milliseconds (cumulative count 907)
2.722% <= 1.607 milliseconds (cumulative count 2722)
6.465% <= 1.703 milliseconds (cumulative count 6465)
12.016% <= 1.807 milliseconds (cumulative count 12016)
18.305% <= 1.903 milliseconds (cumulative count 18305)
28.473% <= 2.007 milliseconds (cumulative count 28473)
38.757% <= 2.103 milliseconds (cumulative count 38757)
93.286% <= 3.103 milliseconds (cumulative count 93286)
99.956% <= 4.103 milliseconds (cumulative count 99956)
100.000% <= 5.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 18426.39 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.324     0.984     2.239     3.159     3.399     5.047
```

## GetPerformance

```bash
====== GET ======
  100000 requests completed in 5.33 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.023 milliseconds (cumulative count 1)
50.000% <= 2.215 milliseconds (cumulative count 50474)
75.000% <= 2.623 milliseconds (cumulative count 74987)
87.500% <= 2.919 milliseconds (cumulative count 87494)
93.750% <= 3.079 milliseconds (cumulative count 93716)
96.875% <= 3.191 milliseconds (cumulative count 96930)
98.438% <= 3.295 milliseconds (cumulative count 98424)
99.219% <= 3.415 milliseconds (cumulative count 99189)
99.609% <= 3.559 milliseconds (cumulative count 99584)
99.805% <= 3.703 milliseconds (cumulative count 99767)
99.902% <= 3.863 milliseconds (cumulative count 99865)
99.951% <= 4.047 milliseconds (cumulative count 99914)
99.976% <= 4.247 milliseconds (cumulative count 99936)
99.988% <= 4.455 milliseconds (cumulative count 99948)
99.994% <= 4.623 milliseconds (cumulative count 99954)
99.997% <= 4.767 milliseconds (cumulative count 99957)
99.998% <= 4.927 milliseconds (cumulative count 99959)
99.999% <= 4.943 milliseconds (cumulative count 99960)
100.000% <= 4.943 milliseconds (cumulative count 99960)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.013% <= 1.103 milliseconds (cumulative count 13)
0.081% <= 1.207 milliseconds (cumulative count 81)
0.220% <= 1.303 milliseconds (cumulative count 220)
0.510% <= 1.407 milliseconds (cumulative count 510)
1.146% <= 1.503 milliseconds (cumulative count 1146)
3.425% <= 1.607 milliseconds (cumulative count 3424)
7.726% <= 1.703 milliseconds (cumulative count 7723)
13.533% <= 1.807 milliseconds (cumulative count 13528)
20.422% <= 1.903 milliseconds (cumulative count 20414)
31.200% <= 2.007 milliseconds (cumulative count 31188)
41.431% <= 2.103 milliseconds (cumulative count 41414)
94.557% <= 3.103 milliseconds (cumulative count 94519)
99.964% <= 4.103 milliseconds (cumulative count 99924)
100.000% <= 5.103 milliseconds (cumulative count 99960)

Summary:
  throughput summary: 18761.73 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.296     1.016     2.215     3.119     3.375     4.943
```

---

# Value=5k

## SetPerformance

```bash
root@e54d97d17810:/data# redis-benchmark -d 5000 -t get,set
====== SET ======
  100000 requests completed in 5.82 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 1.055 milliseconds (cumulative count 1)
50.000% <= 2.407 milliseconds (cumulative count 50275)
75.000% <= 2.855 milliseconds (cumulative count 75029)
87.500% <= 3.167 milliseconds (cumulative count 87583)
93.750% <= 3.383 milliseconds (cumulative count 93759)
96.875% <= 3.599 milliseconds (cumulative count 96944)
98.438% <= 3.791 milliseconds (cumulative count 98484)
99.219% <= 3.959 milliseconds (cumulative count 99224)
99.609% <= 4.135 milliseconds (cumulative count 99622)
99.805% <= 4.303 milliseconds (cumulative count 99807)
99.902% <= 4.471 milliseconds (cumulative count 99904)
99.951% <= 4.631 milliseconds (cumulative count 99953)
99.976% <= 4.783 milliseconds (cumulative count 99976)
99.988% <= 4.967 milliseconds (cumulative count 99988)
99.994% <= 5.111 milliseconds (cumulative count 99994)
99.997% <= 5.191 milliseconds (cumulative count 99997)
99.998% <= 5.223 milliseconds (cumulative count 99999)
99.999% <= 5.359 milliseconds (cumulative count 100000)
100.000% <= 5.359 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.005% <= 1.103 milliseconds (cumulative count 5)
0.036% <= 1.207 milliseconds (cumulative count 36)
0.139% <= 1.303 milliseconds (cumulative count 139)
0.340% <= 1.407 milliseconds (cumulative count 340)
0.705% <= 1.503 milliseconds (cumulative count 705)
1.809% <= 1.607 milliseconds (cumulative count 1809)
3.976% <= 1.703 milliseconds (cumulative count 3976)
7.720% <= 1.807 milliseconds (cumulative count 7720)
12.262% <= 1.903 milliseconds (cumulative count 12262)
19.071% <= 2.007 milliseconds (cumulative count 19071)
27.249% <= 2.103 milliseconds (cumulative count 27249)
85.153% <= 3.103 milliseconds (cumulative count 85153)
99.575% <= 4.103 milliseconds (cumulative count 99575)
99.993% <= 5.103 milliseconds (cumulative count 99993)
100.000% <= 6.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 17176.23 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.489     1.048     2.407     3.455     3.895     5.359
```

## GetPerformance

```bash
====== GET ======
  100000 requests completed in 5.65 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.999 milliseconds (cumulative count 1)
50.000% <= 2.335 milliseconds (cumulative count 50510)
75.000% <= 2.775 milliseconds (cumulative count 75241)
87.500% <= 3.079 milliseconds (cumulative count 87651)
93.750% <= 3.279 milliseconds (cumulative count 93811)
96.875% <= 3.479 milliseconds (cumulative count 96953)
98.438% <= 3.655 milliseconds (cumulative count 98461)
99.219% <= 3.823 milliseconds (cumulative count 99225)
99.609% <= 3.983 milliseconds (cumulative count 99616)
99.805% <= 4.111 milliseconds (cumulative count 99808)
99.902% <= 4.279 milliseconds (cumulative count 99904)
99.951% <= 4.447 milliseconds (cumulative count 99953)
99.976% <= 4.559 milliseconds (cumulative count 99977)
99.988% <= 4.623 milliseconds (cumulative count 99988)
99.994% <= 4.807 milliseconds (cumulative count 99994)
99.997% <= 5.047 milliseconds (cumulative count 99997)
99.998% <= 5.319 milliseconds (cumulative count 99999)
99.999% <= 5.471 milliseconds (cumulative count 100000)
100.000% <= 5.471 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.001% <= 1.007 milliseconds (cumulative count 1)
0.018% <= 1.103 milliseconds (cumulative count 18)
0.106% <= 1.207 milliseconds (cumulative count 106)
0.290% <= 1.303 milliseconds (cumulative count 290)
0.669% <= 1.407 milliseconds (cumulative count 669)
1.461% <= 1.503 milliseconds (cumulative count 1461)
3.548% <= 1.607 milliseconds (cumulative count 3548)
6.764% <= 1.703 milliseconds (cumulative count 6764)
11.331% <= 1.807 milliseconds (cumulative count 11331)
16.616% <= 1.903 milliseconds (cumulative count 16616)
24.765% <= 2.007 milliseconds (cumulative count 24765)
33.288% <= 2.103 milliseconds (cumulative count 33288)
88.568% <= 3.103 milliseconds (cumulative count 88568)
99.802% <= 4.103 milliseconds (cumulative count 99802)
99.997% <= 5.103 milliseconds (cumulative count 99997)
100.000% <= 6.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 17695.98 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        2.407     0.992     2.335     3.343     3.767     5.471
```