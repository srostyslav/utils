package utils

import (
	"encoding/json"
	"fmt"
	"runtime/debug"

	"github.com/srostyslav/adminbot"
	"github.com/srostyslav/log"

	uuid "github.com/satori/go.uuid"
)

func IsEmptyUUID(u uuid.UUID) bool {
	return fmt.Sprint(u) == "00000000-0000-0000-0000-000000000000"
}

func PrettyPrint(v interface{}) (err error) {
	if b, err := json.MarshalIndent(v, "", "  "); err == nil {
		fmt.Println(string(b))
	}
	return err
}

func CatchPanic(title string) {
	if err := recover(); err != nil {
		stack := string(debug.Stack())
		log.ErrorLogger.Println(err, "panic: "+title+"\n"+stack)
		adminbot.AdminBot.SendError(stack, "panic: "+title)
	}
}
