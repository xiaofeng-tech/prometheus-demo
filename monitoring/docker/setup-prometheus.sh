dir=$(cd ../../ && pwd)
prometheus_config=$dir/alerts-and-dashboard/prometheus/

echo "prometheus config path is $prometheus_config"

docker run \
-p 9090:9090 \
-v $prometheus_config:/etc/prometheus/  \
prom/prometheus