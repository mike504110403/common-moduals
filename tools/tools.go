package tools

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"net"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

var dotNetQueryEscapeMap = map[string]string{
	"%2D": "-",
	"%5F": "_",
	"%2A": "*",
	"%28": "(",
	"%29": ")",
	"%2E": ".",
	"%21": "!",
	"~":   "%7E",
}

// GetMemeryUseKB : 取得目前記憶體使用量(KB)
func GetMemeryUseKB() uint64 {
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)
	return m.Alloc / 1024
}

// DotNetQueryEscape : 以 ASP.NET 方式轉換 QueryEscape
func DotNetQueryEscape(str string) string {
	str = url.QueryEscape(str)
	for old, new := range dotNetQueryEscapeMap {
		str = strings.ReplaceAll(str, old, new)
	}
	return str
}

// OnlyQueryEscapeToLower : 只把 QueryEscape 的字轉為小寫
// 	%2E -> %2e
func OnlyQueryEscapeToLower(str string) string {
	for i, v := range str[:] {
		if string(v) == "%" {
			str = str[:i] + strings.ToLower(str[i:i+3]) + str[i+3:]
		}
	}
	return str
}

// MD5hash : string MD5 hash
func MD5hash(str string) string {
	_hash := md5.Sum([]byte(str))
	return hex.EncodeToString(_hash[:])
}

// CRC32Hash : string CRC32 hash
func CRC32Hash(str string) string {
	_hash := crc32.ChecksumIEEE([]byte(str))
	return fmt.Sprintf("%08x", _hash)
}

// SHA256Hash : string SHA256 hash
func SHA256Hash(s string) string {
	_hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(_hash[:])
}

// SHA1Hash : string SHA-1 hash
func SHA1Hash(s string) string {
	_hash := sha1.Sum([]byte(s))
	return hex.EncodeToString(_hash[:])
}

// GetSign : 產生MD5 簽名檔
func GetSign(signStr string, md5Key string) string {
	result := MD5hash(url.QueryEscape(signStr + md5Key))
	result = base64.StdEncoding.EncodeToString([]byte(result))
	return result
}

// CheckSign : 比對簽名檔
func CheckSign(resData string, targetSign string, md5Key string) bool {
	targetSignByte, err := base64.StdEncoding.DecodeString(targetSign)
	if err != nil {
		return false
	}
	targetSignMD5 := strings.ToLower(string(targetSignByte))
	resSign := strings.ToLower(MD5hash(url.QueryEscape(resData + md5Key)))
	return targetSignMD5 == resSign
}

// CheckInArray : 搜尋字串
func CheckInArray(str string, array []string) bool {
	for _, val := range array {
		if val == str {
			return true
		}
	}
	return false
}

func rsaEncrypt(origData []byte, pubKey []byte) ([]byte, error) {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	partLen := pub.N.BitLen()/8 - 11
	chunks := splitByLimit(origData, partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		bytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(bytes)
	}
	return buffer.Bytes(), nil
}

func PaySign(privateKeyString string, mReq map[string]interface{}) string {
	sorted_keys := make([]string, 0)
	for k := range mReq {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" && k != "remark" && k != "sign" && k != "service_type" {
			sorted_keys = append(sorted_keys, k)
		}
	}
	sort.Strings(sorted_keys)

	var signStrings string
	for i, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if i != (len(sorted_keys) - 1) {
			signStrings = signStrings + k + "=" + value + "&"
		} else {
			signStrings = signStrings + k + "=" + value
		}
	}

	block, _ := base64.StdEncoding.DecodeString(privateKeyString)

	privateKey, err := x509.ParsePKCS8PrivateKey(block)
	if err != nil {
		fmt.Printf("x509.ParsePKCS1PrivateKey-------privateKey----- error : %v\n", err)
		return ""
	} else {
		//fmt.Println("x509.ParsePKCS1PrivateKey-------privateKey-----", privateKey)
	}

	result, err := RsaSign(signStrings, privateKey.(*rsa.PrivateKey))
	return result
}

func PrivateEncryptMD5withRSA(data string, privt *rsa.PrivateKey) (string, error) {
	h := md5.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	sign, err := rsa.SignPKCS1v15(nil, privt, crypto.MD5, hashed)
	if err != nil {
		return "", err
	}

	signData := base64.StdEncoding.EncodeToString(sign)
	return string(signData), nil
}

func RsaSign(origData string, privateKey *rsa.PrivateKey) (string, error) {

	h := md5.New()
	h.Write([]byte(origData))
	digest := h.Sum(nil)

	s, err := rsa.SignPKCS1v15(nil, privateKey, crypto.MD5, digest)
	if err != nil {
		return "", err
	}
	data := base64.StdEncoding.EncodeToString(s)
	return string(data), nil
}

func splitByLimit(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}

func stringSplitByLimit(buf string, lim int) []string {
	var chunk string
	chunks := make([]string, 0, len(buf)/lim+1)

	for len(buf) >= lim {
		chunk, buf = string([]rune(buf)[:lim]), string([]rune(buf)[lim:])
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}

func HexToCarry(hex string, carry int) string {
	hexSplit := stringSplitByLimit(hex, 5)
	toCarry := ""
	for _, str := range hexSplit {
		i64num, _ := strconv.ParseInt(str, 16, carry)
		str := strconv.FormatInt(i64num, carry)
		toCarry += str
	}
	return toCarry
}

func GetEncrypt(signStr string, rsaPublicKey string) string {
	result := url.QueryEscape(signStr)
	encrypt, err := rsaEncrypt([]byte(result), []byte(rsaPublicKey))
	if err != nil {
		fmt.Println("error:", err)
	}
	result = base64.StdEncoding.EncodeToString(encrypt)
	return result
}

// CreateRandomID_Any : 取得隨機產生的字串
func CreateRandomID_Any(n int) string {
	const alphanum = "0123456789abcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// CreateRandomID_AnySalt : 開頭用hash的隨機數(可指定長度)
func CreateRandomID_AnySalt(salt string, n int) string {
	return CRC32Hash(salt) + CreateRandomID_Any(n)
}

// CreateRandomID_24 : 開頭用hash的隨機數(長度共24字)
func CreateRandomID_24(salt string) string {
	return CRC32Hash(salt) + CreateRandomID_Any(16)
}

// CreateRandomID_32 : 開頭用hash的隨機數(長度共32字)
func CreateRandomID_32(salt string) string {
	var n int64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	text := strconv.FormatInt(n, 10)
	text = base64.StdEncoding.EncodeToString([]byte(text + salt))
	text = MD5hash(text)
	return text
}

// Bool2YN : bool to Y or N string
func Bool2YN(target bool) string {
	if target {
		return "Y"
	}
	return "N"
}

// PrintSelfName : 印出func名稱
func PrintSelfName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

// CallerInfo : 印出func 被誰call
func CallerInfo(callstack int) (string, int) {
	pc, _, line, _ := runtime.Caller(callstack)
	return runtime.FuncForPC(pc).Name(), line
}

func CallerInfo2(callstack int) (file string, funcName string, line int) {
	pc, fileFullPath, l, _ := runtime.Caller(callstack)
	funcFullPath := runtime.FuncForPC(pc).Name()
	functionNameSplit := strings.Split(funcFullPath, "/")
	fileFullPathSplit := strings.Split(fileFullPath, "/")
	if len(fileFullPathSplit) == 1 {
		fileFullPathSplit = strings.Split(fileFullPath, "\\")
	}
	fileName := fileFullPathSplit[len(fileFullPathSplit)-1]
	if len(functionNameSplit) > 2 {
		file = strings.Join(functionNameSplit[1:len(functionNameSplit)-1], "/") + "/" + fileName
	} else {
		file = strings.Join(functionNameSplit, "/") + "/" + fileName
	}

	funcNameSplit := strings.Split(functionNameSplit[len(functionNameSplit)-1], ".")
	funcName = funcNameSplit[len(funcNameSplit)-1]
	return file, funcName, l
}

// StructPrint : 印出資料結構 (自動換行縮排)
func StructPrint(_struct interface{}) {
	prettyJSON, err := json.MarshalIndent(_struct, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}

type temporary interface {
	Temporary() bool // IsTemporary returns true if err is temporary.
}

type timeout interface {
	Timeout() bool
}

// IsTemporary : 確認錯誤是否可重試（全自動）
func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

// IsTimeout : 確認錯誤是否為超時
func IsTimeout(err error) bool {
	te, ok := err.(timeout)
	return ok && te.Timeout()
}

// NowMicroUnixTime : 現在的MicroUnixTime
func NowMicroUnixTime() int64 {
	return time.Now().UnixMilli()
}

// GetMicroUnixTime : 將時間轉為毫秒(ms) unixtime
func GetMicroUnixTime(_t time.Time) int64 {
	return _t.UnixMilli()
}

// IsJSONString : 確認是否為JSON的字串
func IsJSONString(str string) bool {
	return json.Unmarshal([]byte(str), new(interface{})) == nil
}

// IsJSONString : 確認是否為JSON的字串
func IsJSONStringV2(str string) bool {
	i := new(interface{})
	err := json.Unmarshal([]byte(str), &i)
	_, isFloat64 := reflect.ValueOf(i).Elem().Interface().(float64)
	return err == nil && !isFloat64
}

// NumOnly : 將字串只保留數字
func NumOnly(str string) (int64, error) {
	re := regexp.MustCompile("[0-9]+")
	return strconv.ParseInt(strings.Join(re.FindAllString(str, -1), ""), 10, 64)
}

// GetFunctionName : 丟入整個func轉出該func的名字
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// JSONFileUnmarshal : 將JSON檔案轉為Struct
func JSONFileUnmarshal(path string, v interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	_byte, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(_byte, v); err != nil {
		return err
	}
	return nil
}

// JSONAutoUnmarshal : 自動判斷是 JSON 字串或檔案並轉為 Struct
func JSONAutoUnmarshal(str string, v interface{}) error {
	if _, err := os.Stat(str); err == nil {
		return JSONFileUnmarshal(str, v)
	} else {
		return json.Unmarshal([]byte(str), v)
	}
}

// CheckServIsOn : 確認伺服器啟動狀態
func CheckServIsOn(host, port string, wait time.Duration) bool {
	if _, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), wait); err != nil {
		return false
	}
	return true
}
