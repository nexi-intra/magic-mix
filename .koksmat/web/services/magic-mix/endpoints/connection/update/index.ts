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
 * Update an existing item
 * 
 * item - The item to update

 * @returns
 * 
 * Connection
 */
export default async function call(transactionId: string ,item: Connection) {
  console.log( "magic-mix.connection", "update");

  return run<Connection>(
    "magic-mix.connection",
    ["update" , JSON.stringify(item)],
    transactionId,
    600,
    transactionId
  );
}

