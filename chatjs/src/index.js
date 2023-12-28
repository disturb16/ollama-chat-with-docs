import { Ollama } from "langchain/llms/ollama";
import { newMemoryStore } from "./stores/memory.js";
import { getDocs } from "./files_handler.js";
import { Query } from "./prompt.js";

console.time("prompt_1");

const dirPath = process.argv[2];
const prompt = process.argv[3];

if (!dirPath) {
  console.log("Please provide a directory path");
  process.exit(1);
}

const ollama = new Ollama({
  baseUrl: "http://localhost:11434",
  model: "llama2",
  temperature: 0.8,
});

const docs = await getDocs(dirPath);
const store = await newMemoryStore(docs);

const retriever = store.asRetriever();

await Query(ollama, retriever, prompt);
console.timeEnd("prompt_1");
