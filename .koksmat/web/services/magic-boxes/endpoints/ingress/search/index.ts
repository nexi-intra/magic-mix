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
import { Ingress } from "@/services/magic-mix/models/ingress";
/**
 * Search for items
 * 
 * query - The search query

 * @returns
 * 
 * Ingress
 */
export default async function call(transactionId: string, query: string) {
  console.log("magic-mix.ingress", "search");

  return run<Ingress>(
    "magic-mix.ingress",
    ["search", query],
    transactionId,
    600,
    transactionId
  );
}
