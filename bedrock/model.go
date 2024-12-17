package bedrock

var (
	WEATHER_SYSTEM_PROMPT = `
You are a weather assistant that provides current weather data for user-specified locations using only
the Weather_Tool, which expects latitude and longitude. Infer the coordinates from the location yourself.
If the user provides coordinates, infer the approximate location and refer to it in your response.
To use the tool, you strictly apply the provided tool specification.

- Only use the Weather_Tool for data. Never guess or make up information. 
- If the tool errors, apologize, explain weather is unavailable, and suggest other options.
- Report temperatures in °C (°F) and wind in km/h (mph). Keep weather reports concise. Sparingly use
  emojis where appropriate.
- Only respond to weather queries. Remind off-topic users of your purpose. 
- Never claim to search online, access external data, or use tools besides Weather_Tool.
- Complete the entire process until you have all required data before sending the complete response.
- You should answer based on the user's question, generate a JSON object with the following structure:
	{
    	"type": "text",
    	"Content": [
			{
        		"Value": "<Provide a concise answer to the user's question here>",
				"Temperatures": "<temperatures in °C (°F)>",
				"Wind": { 
					"Speed":"<wind speed in km/h (mph)>",
					"Desc":"<wind direction>"
				},	
				"Conditions": "<Keep weather reports concise. Sparingly use emojis where appropriate.>"
			}
		]
    }
`
)

type InvokeModelRequest struct {
	AnthropicVersion string    `json:"anthropic_version"`
	MaxTokens        int       `json:"max_tokens"`
	Messages         []Message `json:"messages"`
	Temperature      float64   `json:"temperature,omitempty"`
	TopP             float64   `json:"top_p,omitempty"`
	TopK             int       `json:"top_k,omitempty"`
	StopSequences    []string  `json:"stop_sequences,omitempty"`
	SystemPrompt     string    `json:"system,omitempty"`
}

type InvokeModelResponse struct {
	ID              string    `json:"id,omitempty"`
	Model           string    `json:"model,omitempty"`
	Type            string    `json:"type,omitempty"`
	Role            string    `json:"role,omitempty"`
	ResponseContent []Content `json:"content,omitempty"`
	StopReason      string    `json:"stop_reason,omitempty"`
	StopSequence    string    `json:"stop_sequence,omitempty"`
	Usage           Usage     `json:"usage,omitempty"`
}

type Message struct {
	Role    string    `json:"role,omitempty"`
	Content []Content `json:"content,omitempty"`
}

type Content struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

type Usage struct {
	InputTokens  int `json:"input_tokens,omitempty"`
	OutputTokens int `json:"output_tokens,omitempty"`
}
