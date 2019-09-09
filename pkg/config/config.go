package config

type Common struct {
	Brokers            []string
	FileDescriptorPath string
	ProtoName          string
	Topic			   string
}

type Producer struct {
	Common
}

type Consumer struct {
	Common
	FromBeginning	bool
	Version			string
}
