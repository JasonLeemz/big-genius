package openai

import "big-genius/core/config"

func Init() {

	if config.GlobalConfig.OpenAI.ChatGPT.Enable {
		InitChatGPT()
	}
	if config.GlobalConfig.OpenAI.Azure.Enable {
		InitAzure()
	}

}
