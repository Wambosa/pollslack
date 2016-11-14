package pollslack_test

import (
    "os"
    "testing"
    "github.com/wambosa/expect"
    "github.com/wambosa/confman"
    "github.com/wambosa/pollslack"
)

var userId string

func TestMain(m *testing.M){
    
    conf, err := confman.LoadJson("conf.test")
    
    if err != nil {panic(err)}
    
    pollslack.Configure(conf["key"].(string))
    pollslack.ChangeChannel(conf["channelId"].(string))
    userId = conf["user"].(string)
    
    code := m.Run()
    
    os.Exit(code)
}

func Test_GIVEN_slack_is_accesible_with_key_WHEN_GetChannels_is_called_THEN_returns_map(t *testing.T){
    
    channelMap, err := pollslack.GetChannels()
    
    if err != nil {t.Error(err)}
    
    expecting := expect.TestCase {
        T: t,
        Value: channelMap,
    }
    
    expecting.Truthy()
}

func Test_GIVEN_slack_is_accesible_with_key_WHEN_GetChannelIds_is_called_THEN_returns_slice_of_ids(t *testing.T){
    
    idMap, err := pollslack.GetChannelIds()
    
    if err != nil {t.Error(err)}
    
    expecting := expect.TestCase {
        T: t,
        Value: idMap,
    }
    
    expecting.Truthy()
}

func Test_GIVEN_slack_is_accesible_with_key_WHEN_GetMessagesSince_is_called_THEN_returns_map_of_messages(t *testing.T){
    
    messages, err := pollslack.GetMessagesSince("1419143905" + ".00000")
    
    if err != nil {t.Error(err)}
    
    expecting := expect.TestCase {
        T: t,
        Value: messages,
    }
    
    expecting.Truthy()
}

func Test_GIVEN_slack_is_accesible_with_key_WHEN_PostMessagesSince_is_called_THEN_expect_message_to_exist_in_channel(t *testing.T){
    
    ts := pollslack.TimeStamp()
    
    _, err := pollslack.PostMessage("HiGuys!")
    
    if err != nil {t.Error(err)}
    
    messages, err := pollslack.GetMessagesSince(ts)
    
    if err != nil {t.Error(err)}
        
    expecting := expect.TestCase {
        T: t,
        Value: messages,
    }
    
    expecting.Truthy()
}

func Test_GIVEN_slack_is_accesible_with_key_WHEN_GetUserInfo_is_called_THEN_expect_user_map_to_return(t *testing.T){
    
    user, err := pollslack.GetUserInfo(userId)
    
    if err != nil {t.Error(err)}
        
    expecting := expect.TestCase {
        T: t,
        Value: user,
    }
    
    expecting.Property("name").ToBe("imperialsoup")
}