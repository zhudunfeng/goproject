package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

//这里将这些方法关联到结构体中
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte //这时传输时，使用缓冲
}

//写包
//先发送数据包长度
//再发送数据包本身
func (this *Transfer) WritePkg(data []byte) (err error) {

	//先获取到data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes[:]) fail:", err)
		return
	}

	//发送消息本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) failed: ", err)
		//此时默认返回错误
		return
	}
	return
}

//读取数据包中的信息，反序列化为对象
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据......")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		// err = errors.New("read pkg header failed")
		return
	}
	// fmt.Println("读到的buf=", buf[:4])
	//根据buf[:4]转换成uint32类型
	pkgLen := binary.BigEndian.Uint32(this.Buf[:4])

	//根据pkgLen读取数据消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//  err = errors.New("read pkg header failed")
		return
	}

	//把buf反序列化->message.Message
	//反序列化 传入地址指针 &mes
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	return mes, err
}
