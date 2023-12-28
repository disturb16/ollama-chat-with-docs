import { MemoryVectorStore } from "langchain/vectorstores/memory";
import "@tensorflow/tfjs-node";
import { TensorFlowEmbeddings } from "langchain/embeddings/tensorflow";

/**
 *
 * @param {import("langchain/document").Document<Record<string,any>[]} docs
 * @returns {MemoryVectorStore}
 */
export function newMemoryStore(docs) {
  return MemoryVectorStore.fromDocuments(docs, new TensorFlowEmbeddings());
}
