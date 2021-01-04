package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hamba/avro"
	"github.com/marktsoy/hb_gateway/internal/utils"
	"github.com/streadway/amqp"
)

type Server struct {
	config           *Config
	router           *http.ServeMux
	rabbitConnection *amqp.Connection
	channel          *amqp.Channel
	queue            amqp.Queue
}

func New(c *Config) *Server {
	s := &Server{
		config: c,
		router: http.NewServeMux(),
	}
	s.Init()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) Init() {
	/** Configure routes **/
	s.router.HandleFunc("/", s.createBundle())

	conn, err := amqp.Dial(s.config.RabbitAddr)
	if err != nil {
		panic(err.Error())
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}

	s.channel = ch
	//name string, durable, autoDelete, exclusive, noWait bool
	queue, err := s.channel.QueueDeclare(s.config.MessageQueueName, false, false, false, false, amqp.Table{"x-max-priority": 5})
	if err != nil {
		panic(err.Error())
	}
	s.queue = queue
}

func (s *Server) createBundle() func(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Size     int `json:"size"`
		Priority int `json:"priority"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		req := &request{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		schema := Schema()
		bundleID := utils.StrRandom(14)

		for i := 0; i < req.Size; i++ {
			msg := Message{
				Content:  utils.StrRandom(20) + "; Index " + strconv.Itoa(i) + "Priority" + strconv.Itoa(req.Priority),
				Priority: req.Priority,
				BundleID: bundleID,
			}
			data, err := avro.Marshal(schema, msg)
			if err != nil {
				fmt.Println(err.Error())
			}
			s.pub(data, uint8(msg.Priority))
		}
	}
}

func (s *Server) pub(data []byte, p uint8) error {
	err := s.channel.Publish(
		"",           // exchange
		s.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
			Priority:    p,
		})
	return err
}

func (s *Server) Run() {
	fmt.Printf("Server starting at %v \n", s.config.BindAddr)
	http.ListenAndServe(s.config.BindAddr, s)
}
