package config

type AppConfig struct {
	NacosConfig struct {
		NamespaceId string
		IpAddr      string
		Port        int
		DataId      string
		Group       string
	}
}

type Nacos struct {
	Mysql struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Db       string `json:"db"`
	} `json:"mysql"`
	Redis struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		Db       int    `json:"db"`
	} `json:"redis"`
	Aliyun struct {
		AccessKeyID     string `json:"AccessKeyID"`
		AccessKeySecret string `json:"AccessKeySecret"`
		Bucket          string `json:"bucket"`
		URL             string `json:"Url"`
	} `json:"aliyun"`
	Alipay struct {
		AppID      string `json:"AppId"`
		PrivateKey string `json:"PrivateKey"`
		NotifyURL  string `json:"NotifyURL"`
		ReturnURL  string `json:"ReturnURL"`
	} `json:"alipay"`
	Consul struct {
		Name string `json:"name"`
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"consul"`
	QiNiuYun struct {
		Ak     string `json:"AK"`
		Sk     string `json:"Sk"`
		Bucket string `json:"Bucket"`
		URL    string `json:"Url"`
	} `json:"QiNiuYun"`
}
