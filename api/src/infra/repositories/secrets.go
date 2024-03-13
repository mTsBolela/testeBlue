package main

// Secrets contém chaves secretas fictícias
var Secrets = map[string]string{
    "GitHub": "gIthuBGitHuB_12345678901234567890123456789012345",
	"Generic API Key": "ApiKey_apiKey1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ-",
	"Generic Secret": "SecretSecrets0123456789abcdefghijklmnopqrstuvwxyz_",
	"Google API Key": "AIzaqpzjqzldvxfpgwutroebymodkdwiwbmBUQa",
	"Google Cloud Platform API Key": "AIzaoplplmkwtkrdoofqwertyuiopqwertyuBNc",
	"Google Cloud Platform OAuth": "1234567890-abcdefghijklmnopqrstuvwxyz_ABCDE.apps.googleusercontent.com",
	"Google Drive API Key": "AIzaelklglzdzdfrpoiyutrewqwertyuiopFgh",
	"Google Drive OAuth": "1234567890-abcdefghijklmnopqrstuvwxyz_ABCDE.apps.googleusercontent.com",
	"Google (GCP) Service-account": "\"type\": \"service_account\"",
	"Google Gmail API Key": "AIzazpqwsXsdfcvghjklqwertyuiopasdfghjKLM",
	"Google Gmail OAuth": "1234567890-abcdefghijklmnopqrstuvwxyz_ABCDE.apps.googleusercontent.com",
	"Google OAuth Access Token": "ya29._zpxcqwoeirutyfghasdzxcvbnmbosJOMKJIHGFEDCBAZXCVBNM",
	"Google YouTube API Key": "AIzaqwertysdfghxcvbnmzxcvbnmlkjhgfdsjkl",
	"Google YouTube OAuth": "1234567890-abcdefghijklmnopqrstuvwxyz_ABCDE.apps.googleusercontent.com",
	"Slack Webhook": "https://hooks.slack.com/services/Tkjsnhdifuuhdbfkhfksbfnfh/BJHkjkjdbdjnfkdbjfksnf/abcdefghijklmnopqrstuvwxyz",
}

func main() {
    // Exemplo de utilização
    for key, value := range Secrets {
        println("Chave:", key, "Valor:", value)
    }
}
