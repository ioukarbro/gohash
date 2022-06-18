package encrypt

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func RSASign(plainText []byte, fileName string) []byte {
	// 1 打开私钥文件
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	buf := make([]byte, fileInfo.Size())
	// 2 将私钥文件内容读取出来
	file.Read(buf)
	file.Close()

	// 3 使用pem对读取的内容解码得到block
	block, _ := pem.Decode(buf)
	// 4 x509将数据解析得到私钥结构体
	privateKey, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
	//转换格式  类型断言
	//privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err != nil {
		panic(err)
	}
	// 5 创建一个hash对象
	//myhash := sha512.New()
	myhash := md5.New()
	// 6 给hash对象添加数据
	myhash.Write(plainText)
	// 7 计算hash值
	hashText := myhash.Sum(nil)
	// 8 使用rsa函数对散列值签名
	signText, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.MD5, hashText)
	if err != nil {
		panic(err)
	}
	return signText
}

func RSAEncrypt(src []byte, path string) (res []byte, err error) {
	//1.获取秘钥（从本地磁盘读取）
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	fileInfo, _ := f.Stat()
	b := make([]byte, fileInfo.Size())
	f.Read(b)
	// 2、将得到的字符串解码
	block, _ := pem.Decode(b)

	// 使用X509将解码之后的数据 解析出来
	//x509.MarshalPKCS1PublicKey(block):解析之后无法用，所以采用以下方法：ParsePKIXPublicKey
	keyInit, err := x509.ParsePKIXPublicKey(block.Bytes) //对应于生成秘钥的x509.MarshalPKIXPublicKey(&publicKey)
	//keyInit1,err:=x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return
	}
	//4.使用公钥加密数据
	pubKey := keyInit.(*rsa.PublicKey)
	res, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, src)
	return
}
