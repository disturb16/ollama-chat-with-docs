import { RetrievalQAChain } from "langchain/chains";

/**
 *
 * @param {import("langchain/llms/ollama").Ollama} ollama
 * @param {import("langchain/retrievers/vectorstore").VectorStoreRetriever} retreiver
 * @param {string} query
 */
export async function Query(ollama, retreiver, query) {
  const chain = RetrievalQAChain.fromLLM(ollama, retreiver);
  const response = await chain.call({ query: query });

  console.log("");
  console.log(JSON.stringify(response));
}
