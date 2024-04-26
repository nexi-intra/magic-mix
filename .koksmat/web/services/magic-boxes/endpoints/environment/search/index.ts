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
import { Environment } from "@/services/magic-mix/models/environment";
/**
 * Search for items
 * 
 * query - The search query

 * @returns
 * 
 * Environment
 */
export default async function call(transactionId: string, query: string) {
  console.log("magic-mix.environment", "search");

  return run<Environment>(
    "magic-mix.environment",
    ["search", query],
    transactionId,
    600,
    transactionId
  );
}
