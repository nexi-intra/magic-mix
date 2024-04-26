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
 * Update an existing item
 * 
 * item - The item to update

 * @returns
 * 
 * Ingress
 */
export default async function call(transactionId: string, item: Ingress) {
  console.log("magic-mix.ingress", "update");

  return run<Ingress>(
    "magic-mix.ingress",
    ["update", JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}
