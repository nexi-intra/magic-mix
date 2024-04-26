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
 * Update an existing item
 * 
 * item - The item to update

 * @returns
 * 
 * Namespace
 */
export default async function call(transactionId: string, item: Namespace) {
  console.log("magic-mix.namespace", "update");

  return run<Namespace>(
    "magic-mix.namespace",
    ["update", JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}
