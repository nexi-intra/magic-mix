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
import { Dataset } from "@/services/magic-mix/models/dataset";
/**
 * Search for items
 * 
 * query - The search query

 * @returns
 * 
 * Dataset
 */
export default async function call(transactionId: string ,query: string) {
  console.log( "magic-mix.dataset", "search");

  return run<Dataset>(
    "magic-mix.dataset",
    ["search" , query],
    transactionId,
    600,
    transactionId
  );
}

