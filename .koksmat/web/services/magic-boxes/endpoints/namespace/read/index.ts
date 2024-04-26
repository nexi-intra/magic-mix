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
 * Read a single item
 * 
 * id - The id of the item

 * @returns
 * 
 * Namespace
 */
export default async function call(transactionId: string, id: int) {
  console.log("magic-mix.namespace", "read");

  return run<Namespace>(
    "magic-mix.namespace",
    ["read", id],
    transactionId,
    600,
    transactionId
  );
}
