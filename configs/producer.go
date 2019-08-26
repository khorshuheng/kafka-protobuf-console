package configs

type ProducerConfig struct {
	Brokers            []string
	FileDescriptorPath string
	ProtoName          string
	Topic			   string
}
