import * as fs from "fs";
import * as path from "path";
import * as t from "@onflow/types";
import * as fcl from "@onflow/fcl";
import { FlowService } from "./services/flow";

const mint = async (flowService: FlowService, recipient: string) => {
  const transaction = fs.readFileSync(
    path.join(__dirname, `../transactions/mint_nft.cdc`),
    "utf8"
  );

  return flowService.sendTx({
    transaction,
    args: [
      fcl.arg(recipient, t.Address),
      fcl.arg(kind, t.UInt8),
      fcl.arg(rarity, t.UInt8),
    ],
  });
};

const result = mint("https://google.com");
