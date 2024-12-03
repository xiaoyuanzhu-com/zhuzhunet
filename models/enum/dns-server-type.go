package enum

type DNSServerType string

const (
	DNSServerTypeDefault  DNSServerType = "default"
	DNSServerTypeUDP      DNSServerType = "udp"
	DNSServerTypeDoH      DNSServerType = "doh"
	DNSServerTypeDoT      DNSServerType = "dot"
	DNSServerTypeDoQ      DNSServerType = "doq"
	DNSServerTypeDNSCrypt DNSServerType = "dnscrypt"
)
