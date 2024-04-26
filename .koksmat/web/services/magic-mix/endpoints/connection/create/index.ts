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
import { Connection } from "@/services/magic-mix/models/connection";
/**
 * Create a new item
 * 
 * item - The item to create

 * @returns
 * 
 * Connection
 */
export default async function call(transactionId: string ,item: Connection) {
  console.log( "magic-mix.connection", "create");

  return run<Connection>(
    "magic-mix.connection",
    ["create" , JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}

