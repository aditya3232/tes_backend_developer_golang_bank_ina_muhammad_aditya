package log

import (
	"io"
	"os"
	"path/filepath"

	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var New = logrus.New()

type LogArgs struct {
	Endpoint      string `json:"endpoint"`
	Status        string `json:"status"`
	FromIpAddress string `json:"from_ip_address"`
}

func init() {
	log := New

	// Set the log file path
	logFilePath := filepath.Join("..", "tes_backend_developer_golang_bank_ina_muhammad_aditya", "log", "log.log")

	// Open the log file
	file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Warnf("error opening file: %v", err)
	}

	// Set the log output to both stdout and the log file
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

	// Set the log formatter
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "15:04:05 02-01-2006",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
		DisableHTMLEscape: false,
	}

	// Add the config path for viper
	configPath := filepath.Join("..", "tes_backend_developer_golang_bank_ina_muhammad_aditya", "config")
	viper.AddConfigPath(configPath)
}

/*
- dalam menggunakan log gunakan yang ada f diakhirnya, jadi kita bisa memberikan pesan string untuk informasi tambahan terkair errornya
- Contoh Penggunaan: Errorf("Failed to process request: %v", err)
*/

/*
- log info => Digunakan untuk mencatat informasi umum atau langkah-langkah yang dijalankan dengan benar.
- Contoh Penggunaan: Info("A user has logged in")
*/
func Info(args ...interface{}) {
	New.Info(args...)
}

func Infof(format string, args ...interface{}) {
	New.Infof(format, args...)
}

/*
- Log error => Digunakan untuk mencatat pesan kesalahan atau kondisi tidak diharapkan yang terjadi.
- Contoh Penggunaan: Error("Failed to process request")
*/
func Error(args ...interface{}) {
	New.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	New.Errorf(format, args...)
}

/*
- Log fatal => Sama seperti Error, tetapi juga menghentikan eksekusi program setelah mencatat pesan kesalahan.
- Contoh Penggunaan: Fatal("Critical error, shutting down")
*/
func Fatal(args ...interface{}) {
	New.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	New.Fatalf(format, args...)
}

/*
- Log panic => Mencatat pesan kesalahan dan menyebabkan panic, menghentikan eksekusi program.
- Contoh Penggunaan: Panic("Unable to find configuration file")
*/
func Panic(args ...interface{}) {
	New.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	New.Panicf(format, args...)
}

/*
- Log warning => Digunakan untuk mencatat peringatan atau kondisi yang seharusnya mendapatkan perhatian.
- Contoh Penggunaan: Warn("Resource usage is high")
*/
func Warn(args ...interface{}) {
	New.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	New.Warnf(format, args...)
}

/*
- Log debug => Digunakan selama pengembangan untuk mencatat informasi rinci atau langkah-langkah di dalam kode.
- Contoh penggunaan: Debug("Entering function X")
*/
func Debug(args ...interface{}) {
	New.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	New.Debugf(format, args...)
}

/*
- Log trace => Digunakan untuk pencatatan rinci yang sering digunakan dalam debug dan analisis kinerja.
- Contoh Penggunaan: Trace("Function Y execution time:", executionTime)
*/
func Trace(args ...interface{}) {
	New.Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	New.Tracef(format, args...)
}

/*
- Log print => Digunakan untuk mencetak informasi tanpa tingkat log tertentu.
- Contoh penggunaan: Print("Printing status...")
*/
func Print(args ...interface{}) {
	New.Print(args...)
}

func Printf(format string, args ...interface{}) {
	New.Printf(format, args...)
}

/*
- Log(level logrus.Level, args ...interface{}) / Logf(level logrus.Level, format string, args ...interface{}):
- Mencatat pesan log dengan tingkat log tertentu yang ditentukan.
- Contoh Penggunaan: Log(logrus.WarnLevel, "This is a warning message")
*/
func Log(level logrus.Level, args ...interface{}) {
	New.Log(level, args...)
}

func Logf(level logrus.Level, format string, args ...interface{}) {
	New.Logf(level, format, args...)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return New.WithFields(fields)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return New.WithField(key, value)
}

func WithError(err error) *logrus.Entry {
	return New.WithError(err)
}
