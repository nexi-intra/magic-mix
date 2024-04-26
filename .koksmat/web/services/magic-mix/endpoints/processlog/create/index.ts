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
 * Create a new item
 * 
 * item - The item to create

 * @returns
 * 
 * ProcessLog
 */
export default async function call(transactionId: string ,item: ProcessLog) {
  console.log( "magic-mix.processlog", "create");

  return run<ProcessLog>(
    "magic-mix.processlog",
    ["create" , JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}

