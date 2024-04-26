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
 * Update an existing item
 * 
 * item - The item to update

 * @returns
 * 
 * Pods
 */
export default async function call(transactionId: string, item: Pods) {
  console.log("magic-mix.pods", "update");

  return run<Pods>(
    "magic-mix.pods",
    ["update", JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}
