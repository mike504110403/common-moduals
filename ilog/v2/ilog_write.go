package ilog

import (
	"fmt"
	"os"
)

// Default : log DEFAULT
func (d *Data) Default(format string, v ...interface{}) *Data { return d.Writef(DEFAULT, format, v...) }

// Debug : log DEBUG
func (d *Data) Debug(format string, v ...interface{}) *Data { return d.Writef(DEBUG, format, v...) }

// Info : log INFO
func (d *Data) Info(format string, v ...interface{}) *Data { return d.Writef(INFO, format, v...) }

// Noti : log NOTICE
func (d *Data) Noti(format string, v ...interface{}) *Data { return d.Writef(NOTICE, format, v...) }

// Warn : log WARNING
func (d *Data) Warn(format string, v ...interface{}) *Data { return d.Writef(WARNING, format, v...) }

// Err : log ERROR
func (d *Data) Err(format string, v ...interface{}) *Data { return d.Writef(ERROR, format, v...) }

// Criti : log CRITICAL
func (d *Data) Criti(format string, v ...interface{}) *Data { return d.Writef(CRITICAL, format, v...) }

// Alert : log ALERT
func (d *Data) Alert(format string, v ...interface{}) *Data { return d.Writef(ALERT, format, v...) }

// Emergency : log EMERGENCY
func (d *Data) Emergency(format string, v ...interface{}) *Data {
	return d.Writef(EMERGENCY, format, v...)
}

func (d *Data) Fatal(format string, v ...interface{}) {
	defer os.Exit(1)
	d.Writef(EMERGENCY, format, v...)
}

func (d *Data) Panic(format string, v ...interface{}) {
	defer panic(fmt.Sprintf(format, v...))
	d.Writef(EMERGENCY, format, v...)
}

// Default : log DEFAULT
func Default(format string, v ...interface{}) *Data { return (&Data{}).Writef(DEFAULT, format, v...) }

// Debug : log DEBUG
func Debug(format string, v ...interface{}) *Data { return (&Data{}).Writef(DEBUG, format, v...) }

// Info : log INFO
func Info(format string, v ...interface{}) *Data { return (&Data{}).Writef(INFO, format, v...) }

// Noti : log NOTICE
func Noti(format string, v ...interface{}) *Data { return (&Data{}).Writef(NOTICE, format, v...) }

// Warn : log WARNING
func Warn(format string, v ...interface{}) *Data { return (&Data{}).Writef(WARNING, format, v...) }

// Err : log ERROR
func Err(format string, v ...interface{}) *Data { return (&Data{}).Writef(ERROR, format, v...) }

// Criti : log CRITICAL
func Criti(format string, v ...interface{}) *Data { return (&Data{}).Writef(CRITICAL, format, v...) }

// Alert : log ALERT
func Alert(format string, v ...interface{}) *Data { return (&Data{}).Writef(ALERT, format, v...) }

// Emergency : log EMERGENCY
func Emergency(format string, v ...interface{}) *Data {
	return (&Data{}).Writef(EMERGENCY, format, v...)
}

func Fatal(format string, v ...interface{}) { (&Data{}).Fatal(format, v...) }
func Panic(format string, v ...interface{}) { (&Data{}).Panic(format, v...) }

// func (d *Data) Def(v ...interface{}) *Data       { return d.Writef(DEFAULT, "%v", v...) }
// func DefF(format string, v ...interface{}) *Data { return (&Data{}).Writef(DEFAULT, format, v...) }
// func Def(v ...interface{}) *Data                 { return (&Data{}).Writef(DEFAULT, "%v", v...) }

// func (d *Data) DebugF(format string, v ...interface{}) *Data { return d.Writef(DEBUG, format, v...) }
// func DebugF(format string, v ...interface{}) *Data           { return (&Data{}).Writef(DEBUG, format, v...) }
// func (d *Data) Debug(v ...interface{}) *Data                 { return d.Writef(DEBUG, "%v", v...) }
// func Debug(v ...interface{}) *Data                           { return (&Data{}).Writef(DEBUG, "%v", v...) }

// func (d *Data) InfF(format string, v ...interface{}) *Data { return d.Writef(INFO, format, v...) }
// func Inf(v ...interface{}) *Data                           { return (&Data{}).Writef(INFO, "%v", v...) }

// func (d *Data) NotiF(format string, v ...interface{}) *Data { return d.Writef(NOTICE, format, v...) }
// func Noti(v ...interface{}) *Data                           { return (&Data{}).Writef(NOTICE, "%v", v...) }

// func (d *Data) WarnF(format string, v ...interface{}) *Data { return d.Writef(WARNING, format, v...) }
// func Warn(v ...interface{}) *Data                           { return (&Data{}).Writef(WARNING, "%v", v...) }

// func (d *Data) ErrF(format string, v ...interface{}) *Data { return d.Writef(ERROR, format, v...) }
// func Err(v ...interface{}) *Data                           { return (&Data{}).Writef(ERROR, "%v", v...) }

// func (d *Data) CRITICALF(format string, v ...interface{}) *Data {
// 	return d.Writef(Default, format, v...)
// }
// func CRITICAL(v ...interface{}) *Data { return (&Data{}).Writef(DEFAULT, "%v", v...) }

// func (d *Data) ALERTF(format string, v ...interface{}) *Data { return d.Writef(DEFAULT, format, v...) }
// func ALERT(v ...interface{}) *Data                           { return (&Data{}).Writef(DEFAULT, "%v", v...) }

// func (d *Data) EMERGENCYF(format string, v ...interface{}) *Data {
// 	return d.Writef(DEFAULT, format, v...)
// }
// func EMERGENCY(v ...interface{}) *Data { return (&Data{}).Writef(DEFAULT, "%v", v...) }
