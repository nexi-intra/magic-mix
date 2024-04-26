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
 * Search for items
 * 
 * query - The search query

 * @returns
 * 
 * Transformer
 */
export default async function call(transactionId: string ,query: string) {
  console.log( "magic-mix.transformer", "search");

  return run<Transformer>(
    "magic-mix.transformer",
    ["search" , query],
    transactionId,
    600,
    transactionId
  );
}

