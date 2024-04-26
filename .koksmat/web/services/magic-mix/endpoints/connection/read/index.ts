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
 * Read a single item
 * 
 * id - The id of the item

 * @returns
 * 
 * Connection
 */
export default async function call(transactionId: string ,id: int) {
  console.log( "magic-mix.connection", "read");

  return run<Connection>(
    "magic-mix.connection",
    ["read" , id],
    transactionId,
    600,
    transactionId
  );
}

