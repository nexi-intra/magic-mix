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
import { Transformation } from "@/services/magic-mix/models/transformation";
/**
 * Search for items
 * 
 * query - The search query

 * @returns
 * 
 * Transformation
 */
export default async function call(transactionId: string ,query: string) {
  console.log( "magic-mix.transformation", "search");

  return run<Transformation>(
    "magic-mix.transformation",
    ["search" , query],
    transactionId,
    600,
    transactionId
  );
}

