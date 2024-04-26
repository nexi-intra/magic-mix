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
 * Update an existing item
 * 
 * item - The item to update

 * @returns
 * 
 * Container Image
 */
export default async function call(transactionId: string ,item: Container Image) {
  console.log( "magic-mix.containerimage", "update");

  return run<Container Image>(
    "magic-mix.containerimage",
    ["update" , JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}

