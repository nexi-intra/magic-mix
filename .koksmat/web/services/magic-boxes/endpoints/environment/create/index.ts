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
import { Environment } from "@/services/magic-mix/models/environment";
/**
 * Create a new item
 * 
 * item - The item to create

 * @returns
 * 
 * Environment
 */
export default async function call(transactionId: string, item: Environment) {
  console.log("magic-mix.environment", "create");

  return run<Environment>(
    "magic-mix.environment",
    ["create", JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}
