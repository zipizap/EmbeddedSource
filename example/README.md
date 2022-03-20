```
paulo@redBox:~/configs/golangs/others/EmbeddedSource/example$ ll
total 48
drwxrwxr-x 3 paulo paulo 4096 mar 20 04:01 ./
drwxrwxr-x 4 paulo paulo 4096 mar 20 03:58 ../
drwxrwxr-x 2 paulo paulo 4096 mar 20 04:00 extracted/
-rw-rw-r-- 1 paulo paulo   74 mar 20 03:42 go.mod
-rw-rw-r-- 1 paulo paulo  185 mar 20 04:01 go.sum
-rw-rw-r-- 1 paulo paulo 1292 mar 20 03:53 main.go


paulo@redBox:~/configs/golangs/others/EmbeddedSource/example$ go run .
Usage: /tmp/go-build454721333/b001/exe/embedMe  --source-files-extract|--source-files-info 


paulo@redBox:~/configs/golangs/others/EmbeddedSource/example$ go run . --source-files-info
2022/03/20 04:02:55 go.mod       74      78410b54873ede734e480c5c446741ab74931efa8cd6cad3c53fe5a069179dd792e96109065ee3dc047ceffe8f3801b9c06d001a4c9563b2e0ed86ae8b43b7dc
2022/03/20 04:02:55 go.sum       185     0979e2d7cbe37131e811d2d11b13df175c1589d239bb4a29a90c763899ed1e3cf85160b46611bdeef9cabf2f2441d21301036c9f59f754b7c888f9ef185fe131
2022/03/20 04:02:55 main.go      1292    51c1398a2060b3e6cb69a9c4d71757df45968db4de0d141fe25394b4b0bb39d9ab7316b5a2c87d42ceb653f60539ff3c3e6f59dc81a8df992e3befc631f5318e


paulo@redBox:~/configs/golangs/others/EmbeddedSource/example$ go run . --source-files-extract
paulo@redBox:~/configs/golangs/others/EmbeddedSource/example$ tree .
.
├── extracted
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── go.mod
├── go.sum
└── main.go
```
