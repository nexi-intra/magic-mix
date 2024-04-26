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
import { ProcessLog } from "@/services/magic-mix/models/processlog";
/**
 * Search for items
 * 
 * query - The search query

 * @returns
 * 
 * ProcessLog
 */
export default async function call(transactionId: string ,query: string) {
  console.log( "magic-mix.processlog", "search");

  return run<ProcessLog>(
    "magic-mix.processlog",
    ["search" , query],
    transactionId,
    600,
    transactionId
  );
}

