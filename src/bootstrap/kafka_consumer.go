package bootstrap

import (
	"os"
	"strings"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/kafka"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/util"
)

// RegistryKafkaConsumer ...
func RegistryKafkaConsumer() kafka.Consumer {
	return kafka.NewConsumerGroup(&kafka.Config{
		Consumer: kafka.ConsumerConfig{
			// default 8
			SessionTimeoutSecond: 6,
			// default 1000
			HeartbeatInterval: 1000,
			// range, sticky, roundrobin, all of strategy only use for consumer group
			RebalanceStrategy: os.Getenv("KAFKA_CONSUMER_REBALANCE_STRATEGY"),
			// The initial offset to use if no offset was previously committed. Should be -1 = newest or  -2 = oldest
			OffsetInitial: util.StringToInt64(os.Getenv("KAFKA_CONSUMER_OFFSET_INITIAL")),
			// `0 = ReadUncommitted` (default) to consume and return all messages in message channel, 1 = ReadCommitted` to hide messages that are part of an aborted transaction
			IsolationLevel: 1,
		},
		Version:  os.Getenv("KAFKA_VERSION"),
		Brokers:  strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
		ClientID: os.Getenv("KAFKA_CLIENT_ID"),
		SASL: kafka.SASL{
			Enable:    util.StringToBool(os.Getenv("KAFKA_SASL_ENABLE")),
			User:      os.Getenv("KAFKA_USER"),
			Password:  os.Getenv("KAFKA_PASSWORD"),
			Mechanism: os.Getenv("KAFKA_MECHANISME"),
			Version:   int16(util.StringToInt64(os.Getenv("KAFKA_SASL_VERSION"))),
			Handshake: util.StringToBool(os.Getenv("KAFKA_HANDSHAKE")),
		},
		TLS: kafka.TLS{
			Enable:     util.StringToBool(os.Getenv("KAFKA_TLS_ENABLE")),
			CaFile:     os.Getenv("KAFKA_CA_FILE"),
			CertFile:   os.Getenv("KAFKA_CERT_FILE"),
			KeyFile:    os.Getenv("KAFKA_KEY_FILE"),
			SkipVerify: util.StringToBool(os.Getenv("KAFKA_SKIP_VERIFY")),
		},
		// defautl 20
		ChannelBufferSize: 20,
	})
}
