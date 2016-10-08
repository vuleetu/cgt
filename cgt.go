package cgt

import "github.com/vuleetu/levelog"

func Log(things ...interface{}) {
    levelog.Error(things...)
}
