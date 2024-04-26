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
import { Pods } from "@/services/magic-mix/models/pods";
/**
 * Search for items
 * 
 * query - The search query

 * @returns
 * 
 * Pods
 */
export default async function call(transactionId: string, query: string) {
  console.log("magic-mix.pods", "search");

  return run<Pods>(
    "magic-mix.pods",
    ["search", query],
    transactionId,
    600,
    transactionId
  );
}
