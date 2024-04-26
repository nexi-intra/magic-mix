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
 * Create a new item
 * 
 * item - The item to create

 * @returns
 * 
 * Transformer
 */
export default async function call(transactionId: string ,item: Transformer) {
  console.log( "magic-mix.transformer", "create");

  return run<Transformer>(
    "magic-mix.transformer",
    ["create" , JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}

