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
 * Read a single item
 * 
 * id - The id of the item

 * @returns
 * 
 * Pods
 */
export default async function call(transactionId: string, id: int) {
  console.log("magic-mix.pods", "read");

  return run<Pods>(
    "magic-mix.pods",
    ["read", id],
    transactionId,
    600,
    transactionId
  );
}
