# ollama-chat-with-docs

## usage
First setup ollama and its models
```
1. docker compose up -d
2. docker exec -it ollama_core bash
3. ollama pull llama2
```

Then run go app and add your files.
You can add files and make a query y one line
```
1. cd ./chatgo
2. go run main.go -f path/to/file -q "your query"
```
