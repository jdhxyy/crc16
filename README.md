# crc16

## 介绍
查表法实现crc16算法.参数模型是CRC-16/MODBUS 多项式是8005。

## 导入到项目中
```go
import "gitee.com/jdhxyy/crc16"
```

## API
```go
// CheckSum 校验和.参数模型是CRC-16/MODBUS 多项式是8005
// 返回值是高字节在前的16位校验值
func CheckSum(bytes []uint8) uint16
```

## 示例
```go
func TestCheckSum(t *testing.T) {
    arr := []uint8{1, 2, 3}
    fmt.Printf("0x%04x", CheckSum(arr))
}
```

输出:
```
=== RUN   TestCheckSum
0x6161--- PASS: TestCheckSum (0.00s)
PASS
```