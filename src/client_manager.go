package main

import (
	"BattleArchive-server/src/packages"
	packageclient "BattleArchive-server/src/packages/client"
	clientevent "BattleArchive-server/src/packages/client_event"
	packageserver "BattleArchive-server/src/packages/server"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

var clients = make(map[byte]net.Conn)

func ClientsSync() {
	for {
		Broadcast(packageserver.EncodeSyncPlayers())
		time.Sleep(time.Millisecond * 20)
	}
}
func Broadcast(data []byte) {
	BroadcastWithIgnore(data, 255)
}

func BroadcastWithIgnore(data []byte, ignoreId byte) {
	for key, client := range clients {
		if key == ignoreId {
			continue
		}
		client.Write(data)
	}
}

func AddClient(conn net.Conn) {
	id := byte(len(clients))
	clients[id] = conn
	go _makeConn(conn, id)
}

func RemoveClient(conn net.Conn) {
	for key, client := range clients {
		if client == conn {
			delete(clients, key)
			break
		}
	}
}

func _makeConn(conn net.Conn, id byte) {
	defer func() {
		conn.Close()
		RemoveClient(conn)
		fmt.Println("客户端断开连接：", id)
	}()

	conn.Write(packageserver.Handshake{Id: id}.Encode())
	for {
		// 先读取2字节的长度
		lengthBytes := make([]byte, 2)
		_, err := conn.Read(lengthBytes)
		if err != nil {
			// 连接已断开
			return
		}
		length := binary.NativeEndian.Uint16(lengthBytes)

		// 根据长度读取完整数据
		data := make([]byte, length)
		readLen := 0
		for readLen < int(length) {
			n, err := conn.Read(data[readLen:])
			if err != nil {
				// 读取数据失败
				return
			}
			readLen += n
		}

		switch data[0] {
		case packageclient.ClientJoinInfoId:
			var clientJoinInfo packageclient.ClientJoinInfo
			clientJoinInfo.Read(data)
			fmt.Println("玩家加入：", clientJoinInfo.Name)
		case packageclient.PlayerMoveId:
			packageclient.ReadPlayerMove(data)
		case clientevent.PlayerFireId:
			var playerFire clientevent.BasePlayerEvent
			playerFire.Read(data)
			BroadcastWithIgnore(packages.EncodePackageWithoutId(data), playerFire.SenderId)
		case clientevent.PlayerDamageId:
			clientevent.ApplyPlayerDamage(data)
			Broadcast(packages.EncodePackageWithoutId(data))
		}
	}
}
