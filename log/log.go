package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	New      = logrus.New()
	log      = New
	fileName = "log.log"
)

type LogArgs struct {
	Endpoint      string `json:"endpoint"`
	Status        string `json:"status"`
	FromIpAddress string `json:"from_ip_address"`
}

func init() {
	os.Chdir("../latihan-go-ms-one/log")
	/*
		    - os.O_RDWR: Membuka file untuk membaca dan menulis.
		    - os.O_CREATE: Membuat file jika tidak ada.
		    - os.O_APPEND: Menambahkan data ke akhir file.
			- Mode akses (permission) yang diberikan untuk file yang baru dibuat adalah 0666,
	*/
	writer, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("error write", err)
	}

	defer writer.Close()

	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "15:04:05 02-01-2006",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
		DisableHTMLEscape: false,
	}
}

/*
- log info => Digunakan untuk mencatat informasi umum atau langkah-langkah yang dijalankan dengan benar.
- Contoh Penggunaan: Info("A user has logged in")
*/
func Info(args ...interface{}) {
	New.Info(args...)
}

/*
- Log error => Digunakan untuk mencatat pesan kesalahan atau kondisi tidak diharapkan yang terjadi.
- Contoh Penggunaan: Error("Failed to process request")
*/
func Error(args ...interface{}) {
	New.Error(args...)
}

/*
- Log fatal => Sama seperti Error, tetapi juga menghentikan eksekusi program setelah mencatat pesan kesalahan.
- Contoh Penggunaan: Fatal("Critical error, shutting down")
*/
func Fatal(args ...interface{}) {
	New.Fatal(args...)
}

/*
- Log panic => Mencatat pesan kesalahan dan menyebabkan panic, menghentikan eksekusi program.
- Contoh Penggunaan: Panic("Unable to find configuration file")
*/
func Panic(args ...interface{}) {
	New.Panic(args...)
}

/*
- Log warning => Digunakan untuk mencatat peringatan atau kondisi yang seharusnya mendapatkan perhatian.
- Contoh Penggunaan: Warn("Resource usage is high")
*/
func Warn(args ...interface{}) {
	New.Warn(args...)
}

/*
- Log print => Digunakan untuk mencetak informasi tanpa tingkat log tertentu.
- Contoh penggunaan: Print("Printing status...")
*/
func Print(args ...interface{}) {
	New.Print(args...)
}
