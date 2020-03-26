# kafka-push

export MY_IP="$(ifconfig en0 | awk '/inet /{print $2}' | cut -f2 -d':')"

docker run -d --name=kafka-push -p 8089:8089 -e KAFKA_BROKERS="$MY_IP:9092" spanglishing/kafka-push: