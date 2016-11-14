package pollslack

import (
	"fmt"
	"time"
	"github.com/wambosa/simphttp"
)

const baseUrl = "https://slack.com/api/"

var (
    cachedToken string
    cachedChannelId string
)

type SlackConfig struct {
	Token string
	Channels []string
	LastRunTime string
}

func UnixTime() int64 {
    return time.Now().UTC().Unix()
}

func TimeStamp() string {
    return fmt.Sprintf("%v.000000", UnixTime())
}

func Configure(token string){
	cachedToken = token;
}

func ChangeChannel(channelId string){
    cachedChannelId = channelId
}

func GetChannels() (map[string]map[string]interface{}, error){

	method := fmt.Sprintf("%schannels.list?token=%s", baseUrl, cachedToken)

	response, err := simphttp.GetJson(method)

	if err != nil || response == nil {return nil, err}

	channels := make(map[string]map[string]interface{}, len(response["channels"].([]interface{})))

	for _, channel := range response["channels"].([]interface{}){
	    c := channel.(map[string]interface{})
		channels[c["id"].(string)] = c
	    
	}

	return channels, nil
}

func GetChannelIds()([]string, error){

	chans, err := GetChannels()

	if err != nil {return nil, err}

	keys := make([]string, len(chans))
    
    i := 0
    for k := range chans {
        keys[i] = k
        i++
    }

	return keys, nil
}

func GetMessagesSince(lastRunTime string)([]map[string]interface{}, error){

	method := fmt.Sprintf("%schannels.history?token=%s&channel=%s&oldest=%s", baseUrl, cachedToken, cachedChannelId, lastRunTime)

	response, err := simphttp.GetJson(method)

	if(err != nil || response == nil){return nil, err}

	messages := make([]map[string]interface{}, len(response["messages"].([]interface{})))

	for i, message := range response["messages"].([]interface{}) {
		messages[i] = message.(map[string]interface{})}

	return messages, nil
}

func PostMessage(message string) (response map[string]interface{}, err error){
    return PostMessageTo(message, cachedChannelId)
}

func PostMessageTo(message string, channelId string)(map[string]interface{}, error){

	// todo: do some testing to determine string escaping needs.
	// todo: the name
	name := "Taimur Anwar"
	endpoint := fmt.Sprintf("%schat.postMessage", baseUrl)
	
	return simphttp.Query(endpoint, "token=%s&username=%s&channel=%s&text=%s",
	    cachedToken, 
	    name, 
	    cachedChannelId, 
	    message,
	)
}

func GetUserInfo(userId string)(map[string]interface{}, error){

    endpoint := fmt.Sprintf("%susers.info", baseUrl)

	response, err := simphttp.Query(endpoint, "token=%s&user=%s",
	    cachedToken,
	    userId,
	)

	if err != nil || response == nil {return nil, err}

	userInfo := response["user"].(map[string]interface{})

	return userInfo, nil
}