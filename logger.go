package logger

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/fatih/color"
)

// Logger empty struct
type Logger struct {}

func (l *Logger) Write(bytes []byte) (int, error) {
	return fmt.Fprintf(color.Output, "%s", string(bytes))
}

// Init initializes a new logger
func Init() (l *Logger) {
	log.SetFlags(0)
	log.SetOutput(&Logger{})
	return l
}

// Green logs in green color
func (l *Logger) Green(store string, taskNum int, region string, text string) {
	log.Printf("%s %s %s %s %s %s\n", color.HiWhiteString("["+time.Now().Format("2006-01-02 15:04:05.000")+"]"), color.HiWhiteString("::"), color.GreenString("["+store+" "+region+"]"), color.HiWhiteString("[Task-"+strconv.Itoa(taskNum)+"]"), color.HiWhiteString("::"), color.HiGreenString(text))
}

// Black logs in green color
func (l *Logger) Black(store string, taskNum int, region string, text string) {
	log.Printf("%s %s %s %s %s %s\n", color.HiWhiteString("["+time.Now().Format("2006-01-02 15:04:05.000")+"]"), color.HiWhiteString("::"), color.GreenString("["+store+" "+region+"]"), color.HiWhiteString("[Task-"+strconv.Itoa(taskNum)+"]"), color.HiWhiteString("::"), color.HiBlackString(text))
}

// Magenta logs in green color
func (l *Logger) Magenta(store string, taskNum int, region string, text string) {
	log.Printf("%s %s %s %s %s %s\n", color.HiWhiteString("["+time.Now().Format("2006-01-02 15:04:05.000")+"]"), color.HiWhiteString("::"), color.GreenString("["+store+" "+region+"]"), color.HiWhiteString("[Task-"+strconv.Itoa(taskNum)+"]"), color.HiWhiteString("::"), color.HiMagentaString(text))
}

// Red logs in green color
func (l *Logger) Red(store string, taskNum int, region string, text string) {
	log.Printf("%s %s %s %s %s %s\n", color.HiWhiteString("["+time.Now().Format("2006-01-02 15:04:05.000")+"]"), color.HiWhiteString("::"), color.GreenString("["+store+" "+region+"]"), color.HiWhiteString("[Task-"+strconv.Itoa(taskNum)+"]"), color.HiWhiteString("::"), color.HiRedString(text))
}

// White logs in green color
func (l *Logger) White(store string, taskNum int, region string, text string) {
	log.Printf("%s %s %s %s %s %s\n", color.HiWhiteString("["+time.Now().Format("2006-01-02 15:04:05.000")+"]"), color.HiWhiteString("::"), color.GreenString("["+store+" "+region+"]"), color.HiWhiteString("[Task-"+strconv.Itoa(taskNum)+"]"), color.HiWhiteString("::"), color.HiWhiteString(text))
}

// Start logs in green color
func (l *Logger) Start(store string, taskNum int, region string, text string) {
	log.Printf("%s %s %s %s %s %s\n", color.HiWhiteString("["+time.Now().Format("2006-01-02 15:04:05.000")+"]"), color.HiWhiteString("::"), color.HiWhiteString("SYSTEM"), color.HiWhiteString("::"), color.HiRedString(region), color.HiBlackString(text))
}

// LOGGER FORMAT SHAMELESSLY STOLEN FROM dosk#0001 IN BotDev DISCORD :P
