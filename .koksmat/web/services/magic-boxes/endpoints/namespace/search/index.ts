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
import { Namespace } from "@/services/magic-mix/models/namespace";
/**
 * Search for items
 * 
 * query - The search query

 * @returns
 * 
 * Namespace
 */
export default async function call(transactionId: string, query: string) {
  console.log("magic-mix.namespace", "search");

  return run<Namespace>(
    "magic-mix.namespace",
    ["search", query],
    transactionId,
    600,
    transactionId
  );
}
