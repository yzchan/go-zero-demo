package common

type ConsulConf struct {
	Address    string
	Token      string
	Datacenter string
}

type ElasticSearch struct {
	Username string
	Password string
	Hosts    []string
}
