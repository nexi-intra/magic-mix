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
 * Create a new item
 * 
 * item - The item to create

 * @returns
 * 
 * Namespace
 */
export default async function call(transactionId: string, item: Namespace) {
  console.log("magic-mix.namespace", "create");

  return run<Namespace>(
    "magic-mix.namespace",
    ["create", JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}
