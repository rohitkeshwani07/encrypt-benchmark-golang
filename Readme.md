```
➜  cryptobench git:(main) go test -bench .
goos: darwin
goarch: amd64
pkg: cryptobench
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkEncryptOAEP-8       	   21528	     54400 ns/op
BenchmarkDecryptOAEP-8       	     918	   1279571 ns/op
BenchmarkEncryptPKCS1v15-8   	   23234	     51533 ns/op
BenchmarkDecryptPKCS1v15-8   	     928	   1454803 ns/op
PASS
ok  	cryptobench	7.532s
```

```
➜  cryptobench git:(main) ✗ go test -bench .
goos: darwin
goarch: amd64
pkg: cryptobench
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkEncryptOAEP-8       	   20955	     57318 ns/op
BenchmarkDecryptOAEP-8       	     895	   1307603 ns/op
BenchmarkEncryptPKCS1v15-8   	   21078	     54263 ns/op
BenchmarkDecryptPKCS1v15-8   	     910	   1280667 ns/op
PASS
ok  	cryptobench	6.501s
```

```
➜  cryptobench git:(main) ✗ go test -bench .
goos: darwin
goarch: amd64
pkg: cryptobench
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkEncryptOAEP-8       	   20910	     59892 ns/op
BenchmarkDecryptOAEP-8       	     846	   1282058 ns/op
BenchmarkEncryptPKCS1v15-8   	   22292	     53426 ns/op
BenchmarkDecryptPKCS1v15-8   	     937	   1274007 ns/op
PASS
ok  	cryptobench	6.367s
```

```
➜  cryptobench git:(main) ✗ go test -bench .
goos: darwin
goarch: amd64
pkg: cryptobench
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkEncryptOAEP-8       	   21093	     55512 ns/op
BenchmarkDecryptOAEP-8       	     946	   1266280 ns/op
BenchmarkEncryptPKCS1v15-8   	   22486	     54369 ns/op
BenchmarkDecryptPKCS1v15-8   	     882	   1270490 ns/op
PASS
ok  	cryptobench	6.431s
```

```
➜  cryptobench git:(main) ✗ go test -bench .
goos: darwin
goarch: amd64
pkg: cryptobench
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkEncryptOAEP-8       	   20521	     59381 ns/op
BenchmarkDecryptOAEP-8       	     931	   1275481 ns/op
BenchmarkEncryptPKCS1v15-8   	   21766	     54417 ns/op
BenchmarkDecryptPKCS1v15-8   	     932	   1269098 ns/op
PASS
ok  	cryptobench	6.470s
```