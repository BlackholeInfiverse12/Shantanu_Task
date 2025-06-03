import React, { useEffect, useState } from "react";
// import { Card, CardContent } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"


const BlockchainViewer = () => {
  const [blocks, setBlocks] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/blocks")
      .then((res) => res.text())
      .then((text) => {
        const lines = text.split("\n").filter(line => line.trim() !== "");
        const parsed = lines.map(line => {
          const [index, timestamp, data, prevHash, hash] = line.split(",");
          return {
            Index: index,
            Timestamp: timestamp,
            Data: data,
            PrevHash: prevHash,
            Hash: hash
          };
          });
        setBlocks(parsed);
      });
  }, []);

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-4">Go Blockchain Viewer</h1>
      <div className="grid grid-cols-1 gap-2 ">
        {blocks.map((block, idx) => (
          <Card key={idx} className="bg-white shadow-md p-4 rounded-2xl w-[10vw] md:w-[20vw] lg:w-[30vw] h-[20vh] m-4 flex flex-col justify-between">
            <CardHeader className="flex flex-col items-start space-y-1">
              <CardTitle className="text-lg font-semibold">{block.Index}</CardTitle><br />
              <CardDescription className="text-sm text-gray-500">{block.Timestamp}</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="mb-2"><Badge>{block.Data}</Badge></div>
              <div className="text-xs break-all text-gray-700">{block.PrevHash}</div>
              <div className="text-xs break-all text-gray-700">{block.Hash}</div>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
};

export default BlockchainViewer;
