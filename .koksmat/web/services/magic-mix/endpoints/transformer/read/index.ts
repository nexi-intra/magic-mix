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
import { Transformer } from "@/services/magic-mix/models/transformer";
/**
 * Read a single item
 * 
 * id - The id of the item

 * @returns
 * 
 * Transformer
 */
export default async function call(transactionId: string ,id: int) {
  console.log( "magic-mix.transformer", "read");

  return run<Transformer>(
    "magic-mix.transformer",
    ["read" , id],
    transactionId,
    600,
    transactionId
  );
}

