package enum

type DNSType string

const (
	DNSTypeDefault  DNSType = "default"
	DNSTypeUDP      DNSType = "udp"
	DNSTypeDoH      DNSType = "doh"
	DNSTypeDoT      DNSType = "dot"
	DNSTypeDoQ      DNSType = "doq"
	DNSTypeDNSCrypt DNSType = "dnscrypt"
)
