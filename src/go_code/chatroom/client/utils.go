package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

/**
Message的序列化与反序列化
tcp网络传输
*/

func writePkg(conn net.Conn, data []byte) (err error) {

	//先获取到data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes[:]) fail:", err)
		return
	}

	// fmt.Printf("客户端,发送消息的长度=%d,数据内容=%s\n", len(data), string(data))
	//真正的发送数据
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) failed: ", err)
		//此时默认返回错误
		return
	}
	return
}

//读取数据包中的信息，反序列化为对象
func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据......")
	_, err = conn.Read(buf[:4])
	if err != nil {
		// err = errors.New("read pkg header failed")
		return
	}
	// fmt.Println("读到的buf=", buf[:4])
	//根据buf[:4]转换成uint32类型
	pkgLen := binary.BigEndian.Uint32(buf[:4])

	//根据pkgLen读取数据消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//  err = errors.New("read pkg header failed")
		return
	}

	//把buf反序列化->message.Message
	//反序列化 传入地址指针 &mes
	err = json.Unmarshal(buf[:pkgLen], &mes)
	return mes, err
}
