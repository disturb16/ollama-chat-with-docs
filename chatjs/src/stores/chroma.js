import { TensorFlowEmbeddings } from "langchain/embeddings/tensorflow";
import { Chroma } from "langchain/vectorstores/chroma";

/**
 *
 * @param {import("langchain/document").Document<Record<string,any>>[]} docs
 * @param {boolean} collectionExists
 */
export function NewChromaStore(docs, collectionExists) {
  if (collectionExists) {
    return Chroma.fromExistingCollection(new TensorFlowEmbeddings(), {
      url: "http://localhost:8000",
      collectionName: "my-docs",
      collectionMetadata: {
        "hnsw:space": "cosine",
      },
    });
  }

  return Chroma.fromDocuments(docs, new TensorFlowEmbeddings(), {
    collectionName: "my-docs",
    url: "http://localhost:8000",
    collectionMetadata: {
      "hnsw:space": "cosine",
    },
  });
}
