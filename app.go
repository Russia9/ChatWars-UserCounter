package main

import (
	"cw-usercounter/lib"
	"cw-usercounter/messages"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var logger = logrus.New()

func main() {
	// Logger init
	logger.Out = os.Stdout
	logger.Info("Initializing ChatWars Broker")

	// Change logger log level
	switch os.Getenv("CWUC_LOGLEVEL") {
	case "TRACE":
		logger.SetLevel(logrus.TraceLevel)
		break
	case "DEBUG":
		logger.SetLevel(logrus.DebugLevel)
		break
	case "INFO":
		logger.SetLevel(logrus.InfoLevel)
		break
	case "WARN":
		logger.SetLevel(logrus.WarnLevel)
		break
	case "ERROR":
		logger.SetLevel(logrus.ErrorLevel)
		break
	case "FATAL":
		logger.SetLevel(logrus.FatalLevel)
		break
	case "PANIC":
		logger.SetLevel(logrus.PanicLevel)
		break
	default:
		logger.SetLevel(logrus.InfoLevel)
	}
	chatWarsInstance := lib.GetEnv("CWUC_INSTANCE", "cw3-stats")

	// Kafka consumer init
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": lib.GetEnv("CWUC_KAFKA_ADDRESS", "localhost"),
		"group.id":          chatWarsInstance,
		"auto.offset.reset": "latest",
	})

	if err != nil {
		logger.Panic(err)
		return
	}

	consumer.SubscribeTopics([]string{chatWarsInstance + "-offers", chatWarsInstance + "-deals", chatWarsInstance + "-duels"}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			switch *msg.TopicPartition.Topic {
			case chatWarsInstance + "-offers":
				var message messages.OfferMessage
				err = json.Unmarshal(msg.Value, &message)
				if err == nil {
					AddUser("offers", message.SellerID, message.SellerCastle, message.SellerName, "", 0, logger)
				}
			case chatWarsInstance + "-duels":
				var message messages.DuelMessage
				err = json.Unmarshal(msg.Value, &message)
				if err == nil {
					AddUser("duels", message.Winner.ID, message.Winner.Castle, message.Winner.Name, message.Winner.Tag, message.Winner.Level, logger)
					AddUser("duels", message.Loser.ID, message.Loser.Castle, message.Loser.Name, message.Loser.Tag, message.Loser.Level, logger)
				} else {
					logger.Trace(err)
				}
			case chatWarsInstance + "-deals":
				var message messages.DealMessage
				err = json.Unmarshal(msg.Value, &message)
				AddUser("deals", message.SellerID, message.SellerCastle, message.SellerName, "", 0, logger)
				AddUser("deals", message.BuyerID, message.BuyerCastle, message.BuyerName, "", 0, logger)
			}
		} else {
			logger.Error(fmt.Sprintf("Consumer error: %v (%v)\n", err, msg))
		}
	}
}

func AddUser(source string, id string, castle string, name string, tag string, level int, logger *logrus.Logger) {
	DBHost := lib.GetEnv("CWUC_DB_HOST", "localhost")
	DBPort := lib.GetEnv("CWUC_DB_PORT", "3306")
	DBUser := lib.GetEnv("CWUC_DB_USER", "root")
	DBPass := lib.GetEnv("CWUC_DB_PASS", "")
	DBName := lib.GetEnv("CWUC_DB_NAME", "cw3")
	database, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=1&charset=utf8mb4&collation=utf8mb4_unicode_ci", DBUser, DBPass, DBHost, DBPort, DBName)) // ?parseTime=true
	defer database.Close()
	if err != nil {
		logger.Panic("Error creating database: ", err)
		os.Exit(1)
	}
	err = database.Ping()
	if err != nil {
		logger.Panic("Error connecting to the database: ", err)
		os.Exit(1)
	}

	if source == "duels" {
		timestamp := time.Now().Unix()
		_, err := database.Exec("DELETE FROM users WHERE id='" + id + "'")
		if err != nil {
			logger.Debug(err)
		}
		_, _ = database.Exec("INSERT INTO users (id, castle, name, discovered, tag, source, level) VALUES (?,?,?,?,?,?,?)", id, castle, name, timestamp, tag, source, level)
	} else {
		timestamp := time.Now().Unix()
		_, _ = database.Exec("INSERT INTO users (id, castle, name, discovered, source) VALUES (?,?,?,?,?)", id, castle, name, timestamp, source)
	}
}
