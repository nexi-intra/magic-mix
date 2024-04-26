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
 * Delete an existing item
 * 
 * id - The id of the item

 * @returns
 * 
 * Ingress
 */
export default async function call(transactionId: string, id: int) {
  console.log("magic-mix.ingress", "delete");

  return run<Ingress>(
    "magic-mix.ingress",
    ["delete", id],
    transactionId,
    600,
    transactionId
  );
}
