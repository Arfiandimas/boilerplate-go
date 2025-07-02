package bootstrap

import (
	"os"
	"strings"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/kafka"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/util"
)

// RegistryKafkaProducer ...
func RegistryKafkaProducer() kafka.Producer {
	return kafka.NewProducer(&kafka.Config{
		Producer: kafka.ProducerConfig{
			TimeoutSecond: util.StringToInt(os.Getenv("KAFKA_PRODUCER_TIMEOUT_SECOND")),
			// -1 wait for all, 0 = NoResponse doesn't send any response, the TCP ACK is all you get,
			// 1 = WaitForLocal waits for only the local commit to succeed before responding
			RequireACK: -1,
			// If enabled, the producer will ensure that exactly one copy of each message is written
			IdemPotent: true,
			// available strategy : hash, roundrobin, manual, random, reference
			PartitionStrategy: os.Getenv("KAFKA_PRODUCER_PARTITION_STRATEGY"),
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
