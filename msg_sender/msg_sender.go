package msg_sender

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"tcResourceAlert/config"
	"time"
)

func MsgSender(data string) {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(len(config.WecomRobot))
	postData := fmt.Sprintf(`{"msgtype": "text", "text": {"content": "%s","mentioned_mobile_list":["%s"]}}`, data, "")
	if _, err := http.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="+config.WecomRobot[randNum], "application/json", strings.NewReader(postData)); err != nil {
		log.Println("群机器人消息推送失败，err：", err)
	}
	log.Println(data)

}
