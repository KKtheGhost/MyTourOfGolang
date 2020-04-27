package CsvFilter

//本函数目的就是将CSV文件转化为下面的结构体, 并且输出到core功能中, 通过AES解密后, 给Client使用.下面的数据都需要传递.

type AliyunTokenSet struct {
	TokenIndex               int
	AliyunNicknam            string
	AliyunDecryptKey         string
	AliyunTokenKeyEncrypt    string
	AliyunTokenSecretEncrypt string
}
