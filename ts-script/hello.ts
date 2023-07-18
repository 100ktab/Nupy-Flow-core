import * as fcl from "@onflow/fcl";
import { getAccountAddress } from "@onflow/flow-js-testing";

const mint = async () => {
  //   fcl.config().put("flow.network", "testnet");
  //   fcl.config().put("accessNode.api", "https://rest-testnet.onflow.org");
  //   fcl.config().put("0xHelloWorld", "0x664896a25fe00e77");
  //   fcl
  //     .config()
  //     .put(
  //       "discovery.authn.endpoint",
  //       "https://fcl-discovery.onflow.org/testnet/authn"
  //     );
  //   //   const result = await fcl.send({
  //   //     cadence: `
  //   //     import HelloWorld from 0x664896a25fe00e77
  //   //           pub fun hello(): String {
  //   //               return self.greeting
  //   //           }
  //   //         `,
  //   //     args: (arg, t) => [
  //   //       //   arg(7, t.Int), // a: Int
  //   //       //   arg(6, t.Int), // b: Int
  //   //       //   arg("0x664896a25fe00e77", t.Address), // addr: Address
  //   //     ],
  //   //   });
  //   const result = fcl
  //     .send([
  //       fcl.script`
  //       import HelloWorld from 0xHelloWorld // will be replaced with 0xf233dc... in config
  //       access(all) fun hello(): String {
  //         return self.greeting
  //     }
  //     `,
  //     ])
  //     .then(fcl.decode);
  //   return result;
};

const main = async () => {
  // await mint();

  fcl.config().put("accessNode.api", "https://rest-testnet.onflow.org");

  const Alice = await getAccountAddress("Alice");
  console.log({ Alice });
};

main();
