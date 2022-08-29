module go-notes

go 1.15

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
	github.com/alibaba/sentinel-golang v1.0.3
	github.com/allegro/bigcache v1.2.1
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.3.3 // indirect
	github.com/aws/aws-sdk-go v1.35.20 // indirect
	github.com/bluele/gcache v0.0.2
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869
	github.com/boltdb/bolt v1.3.1
	github.com/coreos/bbolt v1.3.4 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/dengsgo/math-engine v0.0.0-20200627074419-8918d8f8ea02
	github.com/dlclark/regexp2 v1.4.0
	github.com/eclipse/paho.mqtt.golang v1.3.2
	github.com/edsrzf/mmap-go v1.0.0
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/gammazero/workerpool v1.1.2
	github.com/getsentry/sentry-go v0.12.0
	github.com/gin-gonic/gin v1.7.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-redis/redis/v8 v8.4.0
	github.com/gocolly/colly v1.2.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/gops v0.3.12
	github.com/google/uuid v1.1.2
	github.com/google/wire v0.5.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.15.2
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/klauspost/cpuid v1.2.3 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/lixd/grpc-go-example v0.0.0-20210731030448-905cb5991806
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/miekg/dns v1.1.27 // indirect
	github.com/nsqio/go-nsq v1.0.8
	github.com/olahol/melody v0.0.0-20180227134253-7bd65910e5ab
	github.com/olivere/elastic v6.2.35+incompatible
	github.com/opentracing/opentracing-go v1.2.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.5
	github.com/orcaman/concurrent-map v0.0.0-20210501183033-44dafcb38ecc
	github.com/panjf2000/ants/v2 v2.5.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/petermattis/goid v0.0.0-20180202154549-b0b1615b78e5
	github.com/pkg/errors v0.9.1
	github.com/qiniu/qmgo v0.8.0
	github.com/rcrowley/go-metrics v0.0.0-20200313005456-10cdbea86bc0
	github.com/rpcxio/rpcx-examples v1.1.6
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/shirou/gopsutil v3.20.11+incompatible // indirect
	github.com/smallnest/rpcx v1.6.2
	github.com/smartystreets/assertions v1.1.1 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.0.0
	github.com/temoto/robotstxt v1.1.1 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	github.com/urfave/cli v1.22.4
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	github.com/xtaci/lossyconn v0.0.0-20200209145036-adba10fffc37 // indirect
	github.com/yuin/gopher-lua v0.0.0-20200816102855-ee81675732da
	go.etcd.io/bbolt v1.3.6 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.mongodb.org/mongo-driver v1.4.0
	go.opencensus.io v0.22.5
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20211008194852-3b03d305991f
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	golang.org/x/sys v0.0.0-20211007075335-d3039528d8ac
	google.golang.org/genproto v0.0.0-20210122163508-8081c04a3579
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/ffmt.v1 v1.5.6
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	sigs.k8s.io/yaml v1.2.0 // indirect
)
