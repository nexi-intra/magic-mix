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
import { ProcessLog } from "@/services/magic-mix/models/processlog";
/**
 * Delete an existing item
 * 
 * id - The id of the item

 * @returns
 * 
 * ProcessLog
 */
export default async function call(transactionId: string ,id: int) {
  console.log( "magic-mix.processlog", "delete");

  return run<ProcessLog>(
    "magic-mix.processlog",
    ["delete" , id],
    transactionId,
    600,
    transactionId
  );
}

