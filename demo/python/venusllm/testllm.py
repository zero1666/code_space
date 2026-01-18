import os
from openai import OpenAI
os.environ['OPENAI_API_KEY'] = "3RgatRUdLTNHAWHMDZByVYy7@3437"
client = OpenAI(base_url="http://v2.open.venus.oa.com/llmproxy")
response = client.chat.completions.create(
  model="deepseek-v3.1-terminus",
  messages=[
    {"role": "user", "content": "你好！你是谁啊"}
  ]
)
print(response)
