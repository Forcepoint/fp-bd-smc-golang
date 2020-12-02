package logsocket

import (
	"crypto/tls"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"sync"
)

type Socket struct {
	Conn              *websocket.Conn
	WebsocketDialer   *websocket.Dialer
	Url               string
	ConnectionOptions ConnectionOptions
	RequestHeader     http.Header
	OnConnected       func(socket Socket)
	OnTextMessage     func(message string, socket Socket)
	OnBinaryMessage   func(data []byte, socket Socket)
	OnConnectError    func(err error, socket Socket)
	OnDisconnected    func(err error, socket Socket)
	OnPingReceived    func(data string, socket Socket)
	OnPongReceived    func(data string, socket Socket)
	IsConnected       bool
	sendMu            *sync.Mutex // Prevent "concurrent write to websocket connection"
	receiveMu         *sync.Mutex
}

type ConnectionOptions struct {
	UseCompression bool
	UseSSL         bool
	Proxy          func(*http.Request) (*url.URL, error)
	SubProtocols   []string
}

func New(url string) Socket {
	return Socket{
		Url:           url,
		RequestHeader: http.Header{},
		ConnectionOptions: ConnectionOptions{
			UseCompression: false,
			UseSSL:         true,
		},
		WebsocketDialer: &websocket.Dialer{},
		sendMu:          &sync.Mutex{},
		receiveMu:       &sync.Mutex{},
	}
}

func (s *Socket) setConnectionOptions() {
	s.WebsocketDialer.EnableCompression = s.ConnectionOptions.UseCompression
	s.WebsocketDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: s.ConnectionOptions.UseSSL}
	s.WebsocketDialer.Proxy = s.ConnectionOptions.Proxy
	s.WebsocketDialer.Subprotocols = s.ConnectionOptions.SubProtocols
}

func (s *Socket) Connect() {
	var err error
	s.setConnectionOptions()

	s.Conn, _, err = s.WebsocketDialer.Dial(s.Url, s.RequestHeader)

	if err != nil {
		logrus.Error("Error while connecting to server ", err)
		s.IsConnected = false
		if s.OnConnectError != nil {
			s.OnConnectError(err, *s)
		}
		return
	}

	logrus.Info("Connected to server")

	if s.OnConnected != nil {
		s.IsConnected = true
		s.OnConnected(*s)
	}

	defaultPingHandler := s.Conn.PingHandler()
	s.Conn.SetPingHandler(func(appData string) error {
		logrus.Trace("Received PING from server")
		if s.OnPingReceived != nil {
			s.OnPingReceived(appData, *s)
		}
		return defaultPingHandler(appData)
	})

	defaultPongHandler := s.Conn.PongHandler()
	s.Conn.SetPongHandler(func(appData string) error {
		logrus.Trace("Received PONG from server")
		if s.OnPongReceived != nil {
			s.OnPongReceived(appData, *s)
		}
		return defaultPongHandler(appData)
	})

	defaultCloseHandler := s.Conn.CloseHandler()
	s.Conn.SetCloseHandler(func(code int, text string) error {
		result := defaultCloseHandler(code, text)
		logrus.Warning("Disconnected from server ", result)
		if s.OnDisconnected != nil {
			s.IsConnected = false
			s.OnDisconnected(errors.New(text), *s)
		}
		return result
	})

	go func() {
		for {
			s.receiveMu.Lock()
			messageType, message, err := s.Conn.ReadMessage()
			s.receiveMu.Unlock()
			if err != nil {
				logrus.Error("read:", err)
				return
			}
			switch messageType {
			case websocket.TextMessage:
				if s.OnTextMessage != nil {
					s.OnTextMessage(string(message), *s)
				}
			case websocket.BinaryMessage:
				if s.OnBinaryMessage != nil {
					s.OnBinaryMessage(message, *s)
				}
			}
		}
	}()
}

func (s *Socket) SendText(message string) {
	err := s.send(websocket.TextMessage, []byte(message))
	if err != nil {
		logrus.Error("write:", err)
		return
	}
}

func (s *Socket) SendBinary(data []byte) {
	err := s.send(websocket.BinaryMessage, data)
	if err != nil {
		logrus.Error("write:", err)
		return
	}
}

func (s *Socket) send(messageType int, data []byte) error {
	s.sendMu.Lock()
	err := s.Conn.WriteMessage(messageType, data)
	s.sendMu.Unlock()
	return err
}

func (s *Socket) Close() {
	err := s.send(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		logrus.Error("write close:", err)
	}
	s.Conn.Close()
	if s.OnDisconnected != nil {
		s.IsConnected = false
		s.OnDisconnected(err, *s)
	}
}
