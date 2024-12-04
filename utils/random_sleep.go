package utils

import (
    "math/rand"
    "time"
    "log"
)

func RandomSleep() {
    duration := time.Duration(100+rand.Intn(300)) * time.Millisecond
    time.Sleep(duration)
    log.Printf("Starting download for sleeping")
}
