module go-notes

go 1.16

replace (
	github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	contrib.go.opencensus.io/exporter/zipkin v0.1.2
	github.com/ClickHouse/clickhouse-go v1.4.3
	github.com/PuerkitoBio/goquery v1.6.0
	github.com/Shopify/sarama v1.19.0
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.3.3 // indirect
	github.com/aws/aws-sdk-go v1.35.20 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/coreos/bbolt v1.3.4 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dengsgo/math-engine v0.0.0-20200627074419-8918d8f8ea02
	github.com/dlclark/regexp2 v1.4.0
	github.com/eclipse/paho.mqtt.golang v1.3.2
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/gammazero/workerpool v1.1.2 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-redis/redis/v7 v7.4.0 // indirect
	github.com/go-redis/redis/v8 v8.4.0
	github.com/go-redis/redis_rate/v7 v7.0.1
	github.com/go-redis/redis_rate/v8 v8.0.0 // indirect
	github.com/go-redis/redis_rate/v9 v9.1.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gobwas/ws v1.1.0
	github.com/gocolly/colly v1.2.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/btree v1.0.0 // indirect
	github.com/google/gops v0.3.12
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.15.2
	github.com/jinzhu/gorm v1.9.16
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/nsqio/go-nsq v1.0.8
	github.com/olivere/elastic v6.2.35+incompatible
	github.com/opentracing/opentracing-go v1.2.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.5
	github.com/orcaman/concurrent-map v0.0.0-20210501183033-44dafcb38ecc // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.8.0 // indirect
	github.com/qiniu/qmgo v0.8.0
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/rpcxio/rpcx-examples v1.1.6
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/smallnest/rpcx v1.6.2
	github.com/smartystreets/assertions v1.1.1 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/spf13/cast v1.3.1
	github.com/temoto/robotstxt v1.1.1 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	github.com/urfave/cli v1.22.4
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	github.com/xtaci/lossyconn v0.0.0-20200209145036-adba10fffc37 // indirect
	github.com/yuin/gopher-lua v0.0.0-20200816102855-ee81675732da
	go.etcd.io/bbolt v1.3.6 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.mongodb.org/mongo-driver v1.4.0
	go.opencensus.io v0.22.5
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	google.golang.org/api v0.3.1 // indirect
	google.golang.org/genproto v0.0.0-20201026171402-d4b8fe4fd877
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	sigs.k8s.io/yaml v1.2.0 // indirect
)
