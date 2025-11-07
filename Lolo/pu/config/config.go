package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"reflect"
	"strconv"
)

type Config struct {
	Log            *Log       `json:"Log"`
	GucooingApiKey string     `json:"GucooingApiKey"`
	Resources      *Resources `json:"Resources"`
	HttpNet        *HttpNet   `json:"HttpNet"`
	GateWay        *GateWay   `json:"GateWay"`
	Game           *Game      `json:"Game"`
	LogServer      *LogServer `json:"LogServer"`
	DB             *DB        `json:"DB"`
}

var DefaultConfig = &Config{
	Log:            defaultLog,
	GucooingApiKey: "123456",
	Resources:      defaultResources,
	HttpNet:        defaultHttpNet,
	GateWay:        defaultGateWay,
	Game:           defaultGame,
	LogServer:      defaultLogServer,
	DB:             defaultDB,
}

var CONF *Config = nil

func SetDefaultConfig() {
	log.Printf("config不存在,使用默认配置\n")
	CONF = DefaultConfig
	CONF.Log.AppName = "App"
}

func GetConfig() *Config {
	if CONF == nil {
		SetDefaultConfig()
	}
	return CONF
}

func GetLog() *Log {
	return GetConfig().Log
}

func GetGucooingApiKey() string {
	return GetConfig().GucooingApiKey
}

var FileNotExist = errors.New("config file not found")

func LoadConfig(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("配置文件读取失败将使用默认配置\n")
		CONF = DefaultConfig
	} else {
		defer func() {
			_ = f.Close()
		}()
		CONF = new(Config)
		d := json.NewDecoder(f)
		if err := d.Decode(CONF); err != nil {
			return err
		}
	}
	overrideWithEnv(reflect.ValueOf(CONF).Elem(), "Config")
	return nil
}

func overrideWithEnv(val reflect.Value, nestKey string) {
	if val.Kind() != reflect.Struct {
		return
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		if !field.CanSet() {
			continue
		}
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}
		envKey := nestKey
		if envKey != "" {
			envKey += "."
		}
		envKey += jsonTag
		if field.Kind() == reflect.Struct {
			overrideWithEnv(field, envKey)
			continue
		}
		if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Struct {
			if field.IsNil() {
				field.Set(reflect.New(field.Type().Elem()))
			}
			overrideWithEnv(field.Elem(), envKey)
			continue
		}
		envValue, exists := os.LookupEnv(envKey)
		if !exists {
			continue
		}
		setFieldValue(field, envValue, envKey)
	}
}

func setFieldValue(field reflect.Value, envValue string, envKey string) {
	target := field
	if field.Kind() == reflect.Ptr {
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		target = field.Elem()
	}
	switch target.Kind() {
	case reflect.String:
		target.SetString(envValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if intVal, err := strconv.ParseInt(envValue, 10, 64); err == nil {
			target.SetInt(intVal)
		} else {
			log.Printf("环境变量 %s 的值 %s 无法转换为整数: %v", envKey, envValue, err)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if uintVal, err := strconv.ParseUint(envValue, 10, 64); err == nil {
			target.SetUint(uintVal)
		} else {
			log.Printf("环境变量 %s 的值 %s 无法转换为无符号整数: %v", envKey, envValue, err)
		}
	case reflect.Bool:
		if boolVal, err := strconv.ParseBool(envValue); err == nil {
			target.SetBool(boolVal)
		} else {
			log.Printf("环境变量 %s 的值 %s 无法转换为布尔值: %v", envKey, envValue, err)
		}
	case reflect.Float32, reflect.Float64:
		if floatVal, err := strconv.ParseFloat(envValue, 64); err == nil {
			target.SetFloat(floatVal)
		} else {
			log.Printf("环境变量 %s 的值 %s 无法转换为浮点数: %v", envKey, envValue, err)
		}
	case reflect.Struct:
		log.Printf("警告: 环境变量 %s 尝试设置结构体字段", envKey)
	default:
		log.Printf("不支持的类型: %s 字段类型 %v", envKey, target.Kind())
	}
}
