"use server";
/*
Parameters

*/
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
import { run } from "@/koksmat/magicservices";
import { ShowCodeFragment } from "@/services/ShowCodeFragment";
import { Container Image } from "@/services/magic-mix/models/containerimage";
/**
 * Read a single item
 * 
 * id - The id of the item

 * @returns
 * 
 * Container Image
 */
export default async function call(transactionId: string ,id: int) {
  console.log( "magic-mix.containerimage", "read");

  return run<Container Image>(
    "magic-mix.containerimage",
    ["read" , id],
    transactionId,
    600,
    transactionId
  );
}

